package lib

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/mattn/go-runewidth"
)

var gistWidth = 53

func (user User) Render() (string, error) {
	var buf strings.Builder

	emoji, err := tierToEmoji(user.Tier)
	if err != nil {
		return buf.String(), err
	}

	tier := tierMap[user.Tier]
	rating := humanize.Comma(int64(user.Rating))
	rank := humanize.Comma(int64(user.Rank))

	var hl = gistWidth -
		runewidth.StringWidth(
			string(emoji),
		) -
		4 -
		runewidth.StringWidth(rank) -
		runewidth.StringWidth(user.Handle)

	res := fmt.Sprintf("%c %-*s#%s @%s\n", emoji, hl, tier, rank, user.Handle)
	buf.WriteString(res)
	buf.WriteString(user.Bio + "\n")
	pbl := gistWidth - runewidth.StringWidth(rating) - 1
	percentage := ratingToPercentage(user.Rating, user.Tier)
	pb := drawProgressBar(pbl, percentage)
	buf.WriteString(fmt.Sprintf("%s %s\n", pb, rating))
	half := (gistWidth - 1) / 2
	buf.WriteString(
		fmt.Sprint(
			"Solved: %*s Class: %*s\n",
			half-8,
			humanize.Comma(int64(user.SolvedCount)),
			half-7,
			humanize.Comma(int64(user.Class)),
		),
	)
	buf.WriteString(
		fmt.Sprint(
			"Contributions: %*s Rivals: %*s\n",
			half-16,
			humanize.Comma(int64(user.VoteCount)),
			half-8,
			humanize.Comma(int64(user.ReverseRivalCount)),
		),
	)

	return buf.String(), nil
}
