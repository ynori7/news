package core

import (
	"html/template"
	"strings"
)

const ErrorTemplate = `<html>
<head></head>
<body>
<h1>{{ .Status }}</h1>
<p>{{ .Message }}</p>
</body>
</html>
`
var CoreTemplateFuncs = template.FuncMap{
	"br2space": func(s string) string {
		s = strings.ReplaceAll(s, "<br/>", " ")
		s = strings.ReplaceAll(s, "<br>", " ")
		return s
	},
}
