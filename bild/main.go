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
	newsTickerHandler := handler.NewNewsTickerHandler(bildApi)
	coronaNewsHandler := handler.NewCoronaNewsHandler(bildApi)

	r := mux.NewRouter()
	r.HandleFunc("/bild/news", newsTickerHandler.Get)
	r.HandleFunc("/bild/corona", coronaNewsHandler.Get)

	log.Info("Starting service")
	http.ListenAndServe(":8080", r)
}
