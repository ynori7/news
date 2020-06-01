package handler

import (
	"net/http"

	"github.com/ynori7/lilypad/handler"
	"github.com/ynori7/lilypad/routing"

	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/filter"
	"github.com/ynori7/news/bild/model"
	"github.com/ynori7/news/bild/templates"
)

type NewsTickerHandler struct {
	api *api.BildNewsTicker
}

func NewNewsTickerHandler(a *api.BildNewsTicker) *NewsTickerHandler {
	h := &NewsTickerHandler{
		api: a,
	}

	routing.RegisterRoutes([]routing.Route{
		{
			Path:    "/bild/news",
			Handler: h.Get,
		},
	})

	return h
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *NewsTickerHandler) Get(r *http.Request) handler.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "NewsTickerHandler.Get"})
	logger.Info("Handling request")

	// Get the data
	news, err := h.api.GetNews()
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting news")
		return handler.ErrorResponse(http.StatusInternalServerError, ErrInternalError)
	}

	// Filter results
	news = filter.FilterNewsTickerItems(news)

	// Render view
	markup, err := view.RenderTemplate(templates.NewsTickerTemplate, struct {
		News []model.NewsTickerItem
	}{News: news})
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return handler.ErrorResponse(http.StatusInternalServerError, ErrInternalError)
	}

	return handler.SuccessResponse(markup)
}
