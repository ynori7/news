package handler

import (
	"github.com/ynori7/news/bild/model"
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	markup, err := view.ExecuteHtmlTemplate(templates.CoronaNewsTemplate, struct {
		News []model.NewsTickerItem
	}{News: news})
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Error("Error rendering view")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	w.Write([]byte(markup))
}
