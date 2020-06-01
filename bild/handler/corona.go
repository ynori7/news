package handler

import (
	"net/http"

	"github.com/ynori7/lilypad/handler"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/templates"
)

type CoronaNewsHandler struct {
	api *api.BildNewsTicker
}

func NewCoronaNewsHandler(a *api.BildNewsTicker) *CoronaNewsHandler {
	h := &CoronaNewsHandler{
		api: a,
	}

	routing.RegisterRoutes(routing.Route{
		Path:    "/bild/corona",
		Handler: h.Get,
	})

	return h
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *CoronaNewsHandler) Get(r *http.Request) handler.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "CoronaNewsHandler.Get"})
	logger.Info("Handling request")

	// Fetch news
	news, err := h.api.GetCoronaNews()
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting news")
		return handler.ErrorResponse(http.StatusInternalServerError, ErrInternalError)
	}

	// Render view
	markup, err := view.RenderTemplate(templates.CoronaNewsTemplate, templates.CoronaNewsData{News: news})
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return handler.ErrorResponse(http.StatusInternalServerError, ErrInternalError)
	}

	return handler.SuccessResponse(markup)
}
