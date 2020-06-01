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
	bildApi := api.NewBildNewsTicker()
	handler.NewNewsTickerHandler(bildApi)
	handler.NewCoronaNewsHandler(bildApi)
	bildIndexHandler := handler.NewIndexHandler()

	// Set the base route
	routing.RegisterRoutes([]routing.Route{
		{
			Path:    "/",
			Handler: bildIndexHandler.Get,
		},
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
