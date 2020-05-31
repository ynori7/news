package view

import (
	"bufio"
	"bytes"
	"html/template"
)

func ExecuteHtmlTemplate(templateToRender string, data interface{}) (string, error) {
	t := template.Must(template.New("html").Parse(templateToRender))

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := t.Execute(w, data)
	if err != nil {
		return "", err
	}

	w.Flush()
	return b.String(), nil
}

