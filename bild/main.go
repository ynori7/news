package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/handler"
)

func main() {
	log.SetLevel(log.DebugLevel)

	bildApi := api.NewBildNewsTicker()
	h := handler.NewNewsTickerHandler(bildApi)

	r := mux.NewRouter()
	r.HandleFunc("/news/bild", h.Get)

	log.Info("Starting service")
	http.ListenAndServe(":8080", r)
}
