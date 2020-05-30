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

func (h HtmlTemplate) ExecuteHtmlTemplate() (string, error) {
	t := template.Must(template.New("html").Parse(htmlTemplate))

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := t.Execute(w, h)
	if err != nil {
		return "", err
	}

	w.Flush()
	return b.String(), nil
}

const htmlTemplate = `<html>
<head>
	<meta charset="utf-8" />
	<style>
		h1 {
			text-align: center;
			margin-bottom: 5rem;
		}
		.info {
			font-size: .875rem;
    		line-height: 1.4;
    		color: #787878;
    		text-transform: uppercase;
		}
	</style>
</head>
<body>
<section class="news">
	<h1>Nachrichten</h1>

	{{ range $i, $val := .News }}
	<div class="news-item">
		<h3><a href="{{ $val.Link }}">{{ $val.Title }}</a></h3>
		<div class="info">
			<span class="date">{{ $val.Date }}</span>
			<span class="category">{{ $val.Category }}</span>
		</div>
		<div class="description"><p>{{ $val.Description }}</p></div>
	</div>
	<hr />
	{{ end }}
</section>
</body>
</html>
`
