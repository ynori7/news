package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetNewsArticle(t *testing.T) {
	//given
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("testdata/spacex-article-70990862.html")
		require.NoError(t, err, "There was an error reading the test data file")
		rw.Write(dat)
	}))
	defer server.Close()

	newsClient := NewBildArticleApi()
	newsClient.httpClient = server.Client()
	newsClient.articlePathFormat = server.URL + "?%s"

	//when
	article, err := newsClient.GetNewsArticle("70990862")

	//then
	require.NoError(t, err, "There was an error getting the article")
	assert.Equal(t, "Wie sieht der Alltag auf der ISS aus?", article.Title)
	assert.Equal(t, "https://www.bild.de/ratgeber/2020/ratgeber/spacex-astronauten-an-bord-wie-sieht-der-alltag-auf-der-iss-aus-70990862.bild.html", article.OriginalLink)
	assert.Equal(t, "2020-06-02T10:49:06+02:00", article.DatePublished)
	assert.Equal(t, "Claudia Mende", article.Author)
	assert.Equal(t, "70990862", article.Id)
	assert.Equal(t, 43, len(article.BodyLines))
}

func Test_GetNewsArticle_NewFormat(t *testing.T) {
	//given
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("testdata/instagram-account-gekapert-hacker-angriff-auf-bayerischen-landtag-77753956.bild.html")
		require.NoError(t, err, "There was an error reading the test data file")
		rw.Write(dat)
	}))
	defer server.Close()

	newsClient := NewBildArticleApi()
	newsClient.httpClient = server.Client()
	newsClient.articlePathFormat = server.URL + "?%s"

	//when
	article, err := newsClient.GetNewsArticle("70990862")

	//then
	require.NoError(t, err, "There was an error getting the article")
	assert.Equal(t, "Hacker-Angriff auf Bayerischen Landtag", article.Title)
	assert.Equal(t, "https://www.bild.de/regional/muenchen/muenchen-aktuell/instagram-account-gekapert-hacker-angriff-auf-bayerischen-landtag-77753956.bild.html", article.OriginalLink)
	assert.Equal(t, "2021-09-22T15:11:55.678Z", article.DatePublished)
	assert.Equal(t, "Thomas GAUTIER", article.Author)
	assert.Equal(t, "70990862", article.Id)
	assert.Equal(t, 6, len(article.BodyLines))
}
