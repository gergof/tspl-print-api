package main

import (
	"bytes"
	"text/template"
)

func fillTemplate(str string, args map[string]string) (string, error) {
	templ, err := template.New("").Parse(str)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = templ.Execute(&buf, args)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
