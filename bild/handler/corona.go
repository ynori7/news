package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/view"
	"github.com/ynori7/news/bild/view/templates"
)

type CoronaNewsHandler struct {
	api *api.BildNewsTicker
}

func NewCoronaNewsHandler(a *api.BildNewsTicker) *CoronaNewsHandler {
	return &CoronaNewsHandler{
		api: a,
	}
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *CoronaNewsHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"Logger": "CoronaNewsHandler.Get"})
	logger.Info("Handling request")

	news, err := h.api.GetCoronaNews()
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting news")
		w.WriteHeader(500)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	template := view.NewHtmlTemplate(news)
	markup, err := template.ExecuteHtmlTemplate(templates.CoronaNewsTemplate)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		w.WriteHeader(500)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	w.Write([]byte(markup))
}
