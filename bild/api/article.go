package api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ynori7/news/bild/model"
	"net/http"
)

const articlePathFormat = "https://www.bild.de/article-%s.bild.html"

type BildArticleApi struct {
	httpClient        *http.Client
	articlePathFormat string
}

func NewBildArticleApi() *BildArticleApi {
	return &BildArticleApi{
		httpClient:        &http.Client{},
		articlePathFormat: articlePathFormat,
	}
}

func (b *BildArticleApi) GetNewsArticle(articleId string) (*model.Article, error) {
	link := fmt.Sprintf(b.articlePathFormat, articleId)

	// Load the HTML document
	doc, err := b.getDocument(link)
	if err != nil {
		return nil, err
	}

	newsArticle := &model.Article{
		Id:   articleId,
		Link: link,
	}

	articleRoot := doc.Find("article").First()
	var articleBody *goquery.Selection

	rawTitle, err := articleRoot.Find("#cover .headline").First().Html()
	if err != nil {
		return nil, err
	}
	if rawTitle != "" {
		newsArticle.Title = rawTitle

		authorData := articleRoot.Find(".authors").First()
		newsArticle.Author = authorData.Find(".authors__name").First().Text()
		newsArticle.DatePublished = authorData.Find(".authors__pubdate").First().AttrOr("datetime", "")

		articleBody = articleRoot.Find(".txt").First()
	} else {
		rawTitle, err := articleRoot.Find(".article-header__headline").First().Html()
		if err != nil {
			return nil, err
		}
		newsArticle.Title = rawTitle

		authorData := articleRoot.Find(".author").First()
		newsArticle.Author = authorData.Find(".author__name").First().Text()
		newsArticle.DatePublished = articleRoot.Find(".datetime--article").First().AttrOr("datetime", "")

		articleBody = articleRoot.Find(".article-body").First()
	}

	articleBody.Children().Each(func(i int, s *goquery.Selection) {
		if s.Is("p") {
			h, err := s.Html()
			if err != nil {
				return
			}
			newsArticle.BodyLines = append(newsArticle.BodyLines, fmt.Sprintf("<p>%s</p>", h))
		}
		if s.Is("h2") {
			h, err := s.Html()
			if err != nil {
				return
			}
			newsArticle.BodyLines = append(newsArticle.BodyLines, fmt.Sprintf("<h2>%s</h2>", h))
		}
	})

	newsArticle.OriginalLink = doc.Find("link[rel=canonical]").First().AttrOr("href", "")

	return newsArticle, nil
}

func (b *BildArticleApi) getDocument(url string) (*goquery.Document, error) {
	// Request the HTML page.
	res, err := b.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
