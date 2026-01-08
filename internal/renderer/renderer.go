package renderer

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func RenderTemplate(templateName string, data any) (string, error) {
	tplPath := filepath.Join(
		"internal",
		"renderer",
		"html",
		templateName,
	)

	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
