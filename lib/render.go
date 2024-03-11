package lib

import (
	"fmt"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/mattn/go-runewidth"
)

var gistWidth = 53

func renderField(field string, num int, width int) string {
	fl := runewidth.StringWidth(field)
	return fmt.Sprintf(
		"%s%*s",
		field,
		width-fl,
		humanize.Comma(int64(num)),
	)
}

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

	half := (gistWidth-1)/2 - 2
	buf.WriteString(
		fmt.Sprintf("%s     %s\n%s     %s\n",
			renderField("✅ Solved: ", user.SolvedCount, half),
			renderField("💠 Class: ", user.Class, half),
			renderField("💡 Contributions: ", user.VoteCount, half),
			renderField("🔥 Rivals: ", user.ReverseRivalCount, half),
		),
	)

	return buf.String(), nil
}
