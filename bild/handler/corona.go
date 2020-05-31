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

type CoronaNewsHandler struct {
	api *api.BildNewsTicker
}

func NewCoronaNewsHandler(a *api.BildNewsTicker) *CoronaNewsHandler {
	return &CoronaNewsHandler{
		api: a,
	}
}

func (h *CoronaNewsHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/bild/corona", h.Get)
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *CoronaNewsHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := log.WithRequest("CoronaNewsHandler.Get", r)
	logger.Info("Handling request")

	news, err := h.api.GetCoronaNews()
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Error("Error getting news")
		w.WriteHeader(500)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	template := view.NewHtmlTemplate(news)
	markup, err := template.ExecuteHtmlTemplate(templates.CoronaNewsTemplate)
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Error("Error rendering view")
		w.WriteHeader(500)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	w.Write([]byte(markup))
}
