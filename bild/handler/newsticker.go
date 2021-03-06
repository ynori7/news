package handler

import (
	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/http"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/filter"
	"github.com/ynori7/news/bild/templates"
)

type NewsTickerHandler struct {
	api *api.BildNewsTicker
}

func NewNewsTickerHandler(a *api.BildNewsTicker) *NewsTickerHandler {
	h := &NewsTickerHandler{
		api: a,
	}

	http.RegisterRoutes([]http.Route{
		{
			Path:    "/bild/news",
			Handler: h.Get,
		},
		{
			Path:    "/bild/corona",
			Handler: h.Corona,
		},
	}...)

	return h
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *NewsTickerHandler) Get(r http.Request) http.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "NewsTickerHandler.Get"})
	logger.Info("Handling request")

	// Get the data
	news, err := h.api.GetNews()
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting news")
		return http.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	// Filter results
	news = filter.FilterNewsTickerItems(news)

	// Render view
	markup, err := view.New("layout", "bild/templates/newsticker.gohtml").Render(templates.NewsTickerData{News: news})
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return http.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	return http.SuccessResponse(markup).WithMaxAge(300)
}

// Get fetches the HTML markup for the Bild Corona Newsticker
func (h *NewsTickerHandler) Corona(r http.Request) http.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "NewsTickerHandler.Corona"})
	logger.Info("Handling request")

	// Fetch news
	news, err := h.api.GetCoronaNews()
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting news")
		return http.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	// Filter results
	news = filter.FilterCoronaNewsItems(news)

	// Render view
	markup, err := view.New("layout", "bild/templates/corona.gohtml").Render(templates.CoronaNewsData{News: news})
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return http.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	return http.SuccessResponse(markup).WithMaxAge(300)
}