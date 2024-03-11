package main

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/mattn/go-runewidth"
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
		Rating:    humanize.Comma(int64(user.Rating)),
		Rank:      humanize.Comma(int64(user.Rank)),
	}, nil
}

func (prop GistProp) Render() (string, error) {
	var buf strings.Builder

	var fl = 53
	var hl = fl - runewidth.StringWidth(string(prop.TierEmoji)) - 4 - runewidth.StringWidth(prop.Rank) - runewidth.StringWidth(prop.Handle)

	res := fmt.Sprintf("%c %-*s#%s @%s", prop.TierEmoji, hl, prop.Tier, prop.Rank, prop.Handle)
	buf.WriteString(res)

	return buf.String(), nil
}
