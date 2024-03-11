package lib

import (
	"errors"
	"strings"
)

var tierMap = []string{
	"Unrated",
	"Bronze V",
	"Bronze IV",
	"Bronze III",
	"Bronze II",
	"Bronze I",
	"Silver V",
	"Silver IV",
	"Silver III",
	"Silver II",
	"Silver I",
	"Gold V",
	"Gold IV",
	"Gold III",
	"Gold II",
	"Gold I",
	"Platinum V",
	"Platinum IV",
	"Platinum III",
	"Platinum II",
	"Platinum I",
	"Diamond V",
	"Diamond IV",
	"Diamond III",
	"Diamond II",
	"Diamond I",
	"Ruby V",
	"Ruby IV",
	"Ruby III",
	"Ruby II",
	"Ruby I",
	"Master",
}

var tierPercentageMap = []int{
	0,                    // unrated
	30, 60, 90, 120, 150, // bronze
	200, 300, 400, 500, 650, // silver
	800, 950, 1100, 1250, 1400, // gold
	1600, 1750, 1900, 2000, 2100, // platinum
	2200, 2300, 2400, 2500, 2600, // diamond
	2700, 2800, 2850, 2900, 2950, // ruby
	3000, // master
}

var tierEmojis = []rune("🟫⬜🟨🟩🟦🟥🟪")

func tierToEmoji(tier int) (rune, error) {
	if 0 < tier && tier <= 5 {
		return tierEmojis[0], nil
	} else if 5 < tier && tier <= 10 {
		return tierEmojis[1], nil
	} else if 10 < tier && tier <= 15 {
		return tierEmojis[2], nil
	} else if 15 < tier && tier <= 20 {
		return tierEmojis[3], nil
	} else if 20 < tier && tier <= 25 {
		return tierEmojis[4], nil
	} else if 25 < tier && tier <= 30 {
		return tierEmojis[5], nil
	} else if tier == 31 {
		return tierEmojis[6], nil
	} else {
		return 0, errors.New("tier out of range")
	}
}

func ratingToPercentage(rating int, tier int) float64 {
	// master progress is always 100%
	if tier == 31 {
		return 1
	}

	base := tierPercentageMap[tier]
	next := tierPercentageMap[tier+1]

	delta, curr := next-base, rating-base

	return float64(curr) / float64(delta)
}

var progressBarChars = []rune("░▏▎▍▌▋▊▉█")
var semi = 8

func drawProgressBar(size int, frac float64) string {
	l := int(float64(semi) * frac * float64(size))
	fl := l / semi

	var buf strings.Builder
	buf.WriteString(strings.Repeat(string(progressBarChars[semi]), min(fl, size)))

	if fl >= size {
		return buf.String()
	}

	buf.WriteRune(progressBarChars[l%semi])
	buf.WriteString(strings.Repeat(string(progressBarChars[0]), size-fl-1))

	return buf.String()
}
