package view

import (
	"bufio"
	"bytes"
	"html/template"

	"github.com/ynori7/news/bild/model"
)

type HtmlTemplate struct {
	News []model.NewsTickerItem
}

func NewHtmlTemplate(news []model.NewsTickerItem) HtmlTemplate {
	return HtmlTemplate{
		News: news,
	}
}

func (h HtmlTemplate) ExecuteHtmlTemplate(templateToRender string) (string, error) {
	t := template.Must(template.New("html").Parse(templateToRender))

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := t.Execute(w, h)
	if err != nil {
		return "", err
	}

	w.Flush()
	return b.String(), nil
}

