package view_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ynori7/news/bild/model"
	"github.com/ynori7/news/bild/templates"
	"github.com/ynori7/news/core/view"
)

func Test_ExecuteHtmlTemplate(t *testing.T) {
	// given
	news := make([]model.NewsTickerItem, 1)
	news[0] = model.NewsTickerItem{
		Title:       "Some news happened",
		Description: "Lots of really bad stuff happened. Oh no!",
		Link:        "http://www.blah.com/blop",
		Date:        "2020-05-22 22:05",
		Category:    "News",
	}

	// when
	_, err := view.ExecuteHtmlTemplate(templates.NewsTickerTemplate, struct {
		News []model.NewsTickerItem
	}{News: news})

	// then
	require.NoError(t, err)
}
