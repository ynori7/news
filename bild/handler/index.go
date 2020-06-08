package handler

import (
	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/http"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
)

type IndexHandler struct {
	api *api.BildNewsTicker
}

func NewIndexHandler() *IndexHandler {
	h := &IndexHandler{}

	http.RegisterRoutes(http.Route{
		Path:    "/bild",
		Handler: h.Get,
	})

	return h
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *IndexHandler) Get(r http.Request) http.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "IndexHandler.Get"})
	logger.Info("Handling request")

	// Render view
	markup, err := view.New("layout", "bild/templates/index.gohtml").Render(nil)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return http.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	return http.SuccessResponse(markup).WithMaxAge(3600)
}
