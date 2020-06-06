package handler

import (
	"net/http"

	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/handler"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/templates"
)

type IndexHandler struct {
	api *api.BildNewsTicker
}

func NewIndexHandler() *IndexHandler {
	h := &IndexHandler{}

	routing.RegisterRoutes(routing.Route{
		Path:    "/bild",
		Handler: h.Get,
	})

	return h
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *IndexHandler) Get(r *http.Request) handler.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "IndexHandler.Get"})
	logger.Info("Handling request")

	// Render view
	markup, err := view.RenderTemplate(templates.IndexTemplate, nil)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return handler.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	return handler.SuccessResponse(markup).WithMaxAge(3600)
}
