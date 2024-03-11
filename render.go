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
	TierEmoji rune
	Rating    string
	Rank      string
}

func NewGistPropFromUser(user User) (GistProp, error) {
	emoji, err := TierToEmoji(user.Tier)
	if err != nil {
		return GistProp{}, err
	}

	return GistProp{
		Handle:    user.Handle,
		Bio:       user.Bio,
		Tier:      TierMap[user.Tier],
		TierEmoji: emoji,
		Rating:    string(user.Rating),
		Rank:      string(user.Rank),
	}, nil
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
