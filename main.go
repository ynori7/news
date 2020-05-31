package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/handler"
	_ "github.com/ynori7/news/core/log"
)

func main() {
	r := mux.NewRouter()

	// Bild
	bildApi := api.NewBildNewsTicker()
	handler.NewNewsTickerHandler(bildApi).AddRoutes(r)
	handler.NewCoronaNewsHandler(bildApi).AddRoutes(r)

	bildIndexHandler := handler.NewIndexHandler()
	bildIndexHandler.AddRoutes(r)
	r.HandleFunc("/", bildIndexHandler.Get)

	log.Info("Starting service")
	http.ListenAndServe(":8080", r)
}
