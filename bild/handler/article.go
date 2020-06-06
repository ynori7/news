package handler

import (
	"net/http"

	"github.com/ynori7/lilypad/errors"
	"github.com/ynori7/lilypad/handler"
	"github.com/ynori7/lilypad/log"
	"github.com/ynori7/lilypad/routing"
	"github.com/ynori7/lilypad/view"
	"github.com/ynori7/news/bild/api"
	"github.com/ynori7/news/bild/templates"
)

type NewsArticleHandler struct {
	api *api.BildArticleApi
}

func NewNewsArticleHandler(a *api.BildArticleApi) *NewsArticleHandler {
	h := &NewsArticleHandler{
		api: a,
	}

	routing.RegisterRoutes(routing.Route{
		Path:    "/bild/news/{id}",
		Handler: h.Get,
	})

	return h
}

// Get fetches the HTML markup for the Bild Newsticker
func (h *NewsArticleHandler) Get(r *http.Request) handler.Response {
	logger := log.WithRequest(r).WithFields(log.Fields{"Logger": "NewsArticleHandler.Get"})
	logger.Info("Handling request")

	id := routing.GetVar(r, "id")
	if id == "" {
		logger.Debug("Missing id from request")
		return handler.ErrorResponse(errors.BadRequestError("missing article id from request"))
	}

	// Get the data
	article, err := h.api.GetNewsArticle(id)
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error getting article")
		return handler.ErrorResponse(errors.InternalServerError("error getting article"))
	}

	// Render view
	markup, err := view.RenderTemplate(templates.NewsArticleTemplate, templates.ArticleData{Article: article})
	if err != nil {
		logger.WithFields(log.Fields{"error": err}).Error("Error rendering view")
		return handler.ErrorResponse(errors.InternalServerError("error getting news"))
	}

	return handler.SuccessResponse(markup).WithMaxAge(3600)
}
