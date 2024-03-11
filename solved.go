package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Handle            string `json:"handle"`
	Bio               string `json:"bio"`
	SolvedCount       int    `json:"solvedCount"`
	VoteCount         int    `json:"voteCount"`
	Tier              int    `json:"tier"`
	Rating            int    `json:"rating"`
	Class             int    `json:"class"`
	ClassDecoration   string `json:"classDecoration"`
	RivalCount        int    `json:"rivalCount"`
	ReverseRivalCount int    `json:"reverseRivalCount"`
	Rank              int    `json:"rank"`
}

var solvedacApiEndpoint = "https://solved.ac/api/v3"
var solvedacUserShowApi = solvedacApiEndpoint + "/user/show"

func (client *BoxClient) FetchUserData(username string) (User, error) {
	var user User

	client.req.R().SetQueryParam("handle", username).Get(solvedacUserShowApi)
	res, err := http.Get("" + username)
	if err != nil {
		return user, err
	}

	err = json.NewDecoder(res.Body).Decode(&user)
	return user, err
}
