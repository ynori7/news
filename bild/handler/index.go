package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/templates"
	"github.com/ynori7/news/core/log"
	"github.com/ynori7/news/core/view"
)

type IndexHandler struct {
	api *api.BildNewsTicker
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/bild", h.Get)
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := log.WithRequest("IndexHandler.Get", r)
	logger.Info("Handling request")

	markup, err := view.ExecuteHtmlTemplate(templates.IndexTemplate, nil)
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Error("Error rendering view")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	w.Write([]byte(markup))
}
