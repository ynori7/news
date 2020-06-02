package templates

import "github.com/ynori7/news/bild/model"

type ArticleData struct {
	Article *model.Article
}

const NewsArticleTemplate = `<html>
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<style>
		h1 {
			text-align: center;
			margin-bottom: 3rem;
		}
		.info {
			font-size: .875rem;
    		line-height: 1.4;
    		color: #787878;
		}
		.back {
			margin-bottom: 2rem;
		}
	</style>
	<title>Bild News - {{ br2space .Article.Title }}</title>
	<meta name="description" content="Bild News - {{ br2space .Article.Title }}" property="og:description">
</head>
<body>
<article>
	<div class="back"><a href="/bild/news">Züruck</a></div>

	<h1><a href="{{ .Article.OriginalLink }}" target="_blank">{{ br2space .Article.Title }}</a></h1>
	
	<div class="info">
		{{ if .Article.Author }}<span class="author">Von {{ .Article.Author }}</span>{{ end }}
		<span class="date">{{ .Article.DatePublished }}</span>
	</div>

	<div class="body">
	{{ .Article.HtmlBody }}
	</div>
</article>
</body>
</html>
`

