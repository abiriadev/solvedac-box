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

var GistWidth = 53

func (user User) Render() (string, error) {
	var buf strings.Builder

	emoji, err := TierToEmoji(user.Tier)
	if err != nil {
		return buf.String(), err
	}

	tier := TierMap[user.Tier]
	rating := humanize.Comma(int64(user.Rating))
	rank := humanize.Comma(int64(user.Rank))

	var hl = GistWidth - runewidth.StringWidth(string(emoji)) - 4 - runewidth.StringWidth(rank) - runewidth.StringWidth(user.Handle)

	res := fmt.Sprintf("%c %-*s#%s @%s\n", emoji, hl, tier, rank, user.Handle)
	buf.WriteString(res)
	buf.WriteString(user.Bio + "\n")
	pbl := GistWidth - runewidth.StringWidth(rating) - 1
	pb := drawProgressBar(pbl, 0.8)
	buf.WriteString("%s %s", pb, rating)

	return buf.String(), nil
}
