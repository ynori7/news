package main

import (
	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/handler"
	"github.com/ynori7/news/core"
)

func main() {
	// Bild
	bildNewsTickerApi := api.NewBildNewsTicker()
	bildArticleApi := api.NewBildArticleApi()
	handler.NewNewsTickerHandler(bildNewsTickerApi)
	handler.NewNewsArticleHandler(bildArticleApi)
	bildIndexHandler := handler.NewIndexHandler()

	// Set the base route
	routing.RegisterRoutes(routing.Route{
		Path:    "/",
		Handler: bildIndexHandler.Get,
	})

	view.RegisterGlobalTemplateFuncs(core.CoreTemplateFuncs)
	errors.UseMarkupErrors(core.ErrorTemplate)

	log.Info("Starting service")
	routing.ServeHttp(":8080")
}
