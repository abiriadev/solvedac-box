package main

import "errors"

var TierMap = map[int]string{
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

var TierEmojis = []rune("🟫⬜🟨🟩🟦🟥🟪")

func TierToEmoji(tier int) (rune, error) {
	if 0 < tier && tier <= 5 {
		return TierEmojis[0], nil
	} else if 5 < tier && tier <= 10 {
		return TierEmojis[1], nil
	} else if 10 < tier && tier <= 15 {
		return TierEmojis[2], nil
	} else if 15 < tier && tier <= 20 {
		return TierEmojis[3], nil
	} else if 20 < tier && tier <= 25 {
		return TierEmojis[4], nil
	} else if 25 < tier && tier <= 30 {
		return TierEmojis[5], nil
	} else if tier == 31 {
		return TierEmojis[6], nil
	} else {
		return 0, errors.New("tier out of range")
	}
}
