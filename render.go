package main

import (
	"strings"
	"text/template"
)

var gistTemplate = "gist.tmpl"

func renderGist(user User) (string, error) {
	var buf strings.Builder
	if tmpl, err := template.New("gist").ParseFiles(gistTemplate); err != nil {
		return "", err
	} else if err := tmpl.Execute(&buf, user); err != nil {
		return "", err
	}
	return buf.String(), nil
}
