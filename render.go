package main

import (
	"strings"
	"text/template"
)

var gistTemplate = "gist.tmpl"

type GistProp struct {
	Handle    string
	Bio       string
	Tier      string
	TierEmoji string
	Rating    string
	Rank      string
}

func NewGistPropFromUser(user User) GistProp {
	return GistProp{
		Handle:    user.Handle,
		Bio:       user.Bio,
		Tier:      string(user.Tier),
		TierEmoji: string(user.Tier),
		Rating:    string(user.Rating),
		Rank:      string(user.Rank),
	}
}

func (prop GistProp) Render() (string, error) {
	var buf strings.Builder
	if tmpl, err := template.New("gist").ParseFiles(gistTemplate); err != nil {
		return "", err
	} else if err := tmpl.Execute(&buf, prop); err != nil {
		return "", err
	}
	return buf.String(), nil
}
