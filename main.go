package main

import (
	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/handler"
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

	errors.UseMarkupErrors(errorTemplate)

	log.Info("Starting service")
	routing.ServeHttp(":8080")
}

const errorTemplate = `<html>
<head></head>
<body>
<h1>{{ .Status }}</h1>
<p>{{ .Message }}</p>
</body>
</html>
`
