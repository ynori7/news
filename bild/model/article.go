package model

import (
	"html/template"
	"strings"
)

type Article struct {
	Id            string
	Title         string
	Link          string
	OriginalLink  string
	Author        string
	DatePublished string
	BodyLines     []string
}

func (a Article) HtmlTitle() template.HTML {
	return template.HTML(a.Title)
}

func (a Article) HtmlBody() template.HTML {
	body := strings.Join(a.BodyLines, "\n")
	return template.HTML(body)
}
