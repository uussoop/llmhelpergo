package utils

import (
	"bytes"
	"html/template"
)

func TemplateRender(data interface{}, tmpl string) (string, error) {

	t, err := template.New("t").Parse(tmpl)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil

}
