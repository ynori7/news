package handler

import (
	"github.com/ynori7/news/bild/filter"
	"github.com/ynori7/news/bild/model"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/templates"
	"github.com/ynori7/news/core/log"
	"github.com/ynori7/news/core/view"
)

type NewsTickerHandler struct {
	api *api.BildNewsTicker
}

func NewNewsTickerHandler(a *api.BildNewsTicker) *NewsTickerHandler {
	return &NewsTickerHandler{
		api: a,
	}
}

func (h *NewsTickerHandler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/bild/news", h.Get)
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *NewsTickerHandler) Get(w http.ResponseWriter, r *http.Request) {
	logger := log.WithRequest("NewsTickerHandler.Get", r)
	logger.Info("Handling request")

	// Get the data
	news, err := h.api.GetNews()
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Error("Error getting news")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrInternalError.Error()))
		return
	}

	// Filter results
	news = filter.FilterNewsTickerItems(news)

	// Render view
	markup, err := view.ExecuteHtmlTemplate(templates.NewsTickerTemplate, struct {
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
