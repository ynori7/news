package api

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetNewsTicker(t *testing.T) {
	//given
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("testdata/alle-meldungen-49391716.bildMobile.html")
		require.NoError(t, err, "There was an error reading the test data file")
		rw.Write(dat)
	}))
	defer server.Close()

	newsClient := &BildNewsTicker{httpClient: server.Client(), newsTickerUrl: server.URL}

	//when
	news, err := newsClient.GetNews()

	//then
	require.NoError(t, err, "There was an error getting the premieres")
	assert.Equal(t, 128, len(news), "There should have been the expected number of news items")
}

func Test_GetCoronaNews(t *testing.T) {
	//given
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		dat, err := ioutil.ReadFile("testdata/coronavirus-live-ticker.html")
		require.NoError(t, err, "There was an error reading the test data file")
		rw.Write(dat)
	}))
	defer server.Close()

	newsClient := &BildNewsTicker{httpClient: server.Client(), coronaTickerUrl: server.URL}

	//when
	news, err := newsClient.GetCoronaNews()

	//then
	require.NoError(t, err, "There was an error getting the premieres")
	assert.Equal(t, 40, len(news), "There should have been the expected number of news items")
}