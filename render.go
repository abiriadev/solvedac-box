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

var tier = map[int]string{
	1:  "Bronze V",
	2:  "Bronze IV",
	3:  "Bronze III",
	4:  "Bronze II",
	5:  "Bronze I",
	6:  "Silver V",
	7:  "Silver IV",
	8:  "Silver III",
	9:  "Silver II",
	10: "Silver I",
	11: "Gold V",
	12: "Gold IV",
	13: "Gold III",
	14: "Gold II",
	15: "Gold I",
	16: "Platinum V",
	17: "Platinum IV",
	18: "Platinum III",
	19: "Platinum II",
	20: "Platinum I",
	21: "Diamond V",
	22: "Diamond IV",
	23: "Diamond III",
	24: "Diamond II",
	25: "Diamond I",
	26: "Ruby V",
	27: "Ruby IV",
	28: "Ruby III",
	29: "Ruby II",
	30: "Ruby I",
	31: "Master",
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
