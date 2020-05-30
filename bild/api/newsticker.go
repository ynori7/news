package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ynori7/news/bild/model"
)

const (
	baseUrl         = "https://m.bild.de"
	newsTickerUrl   = "https://m.bild.de/news/alle-news-home-home/nachricht/alle-meldungen-49391716.bildMobile.html"
	coronaTickerUrl = "https://www.bild.de/news/inland/news-inland/coronavirus-live-ticker-70411946.bild.html"
)

type BildNewsTicker struct {
	httpClient      *http.Client
	newsTickerUrl   string
	coronaTickerUrl string
}

func NewBildNewsTicker() *BildNewsTicker {
	return &BildNewsTicker{
		httpClient:      &http.Client{},
		newsTickerUrl:   newsTickerUrl,
		coronaTickerUrl: coronaTickerUrl,
	}
}

func (b *BildNewsTicker) GetNews() ([]model.NewsTickerItem, error) {
	// Load the HTML document
	doc, err := b.getDocument(b.newsTickerUrl)
	if err != nil {
		return nil, err
	}

	newsItems := make([]model.NewsTickerItem, 0)

	// Find the latest news
	doc.Find(".module:not(.module-alt) .hentry").Each(func(i int, s *goquery.Selection) {
		newsItem := model.NewsTickerItem{}

		linkNode := s.Find("a")
		newsItem.Link = baseUrl + linkNode.AttrOr("href", "")
		newsItem.Title = strings.TrimSpace(linkNode.Find(".entry-title .headline").First().Text())
		newsItem.ImageUrl = linkNode.Find("img").AttrOr("data-src", "")

		if newsItem.Title == "" {
			return //this one must be some garbage node
		}

		newsItem.Description = strings.TrimSpace(s.Find(".entry-content").First().Text())

		infoNode := s.Find(".info")
		newsItem.Date = infoNode.Find("time").AttrOr("datetime", "")
		newsItem.Category = infoNode.Nodes[0].LastChild.Data

		newsItems = append(newsItems, newsItem)
	})

	return newsItems, nil
}

func (b *BildNewsTicker) GetCoronaNews() ([]model.NewsTickerItem, error) {
	// Load the HTML document
	doc, err := b.getDocument(b.coronaTickerUrl)
	if err != nil {
		return nil, err
	}

	newsItems := make([]model.NewsTickerItem, 0)

	// Find the latest news
	doc.Find(".stream-list .stream-item:not(.stream-datetime) .stream-body").Each(func(i int, s *goquery.Selection) {
		newsItem := model.NewsTickerItem{}

		timeNode := s.Find("time").First()
		newsItem.Date = timeNode.AttrOr("datetime", "") + " " + timeNode.Text()

		newsItem.Title = s.Find(".entry-title .headline").Text()

		newsItem.Description = s.Find("p").Text()

		newsItems = append(newsItems, newsItem)
	})

	return newsItems, nil
}

func (b *BildNewsTicker) getDocument(url string) (*goquery.Document, error) {
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