package main

import (
	"errors"
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

func (client BoxClient) FetchUserData(username string) (User, error) {
	var user User

	res, err := client.req.R().SetQueryParam("handle", username).
		SetSuccessResult(&user).
		Get(solvedacUserShowApi)
	if err != nil {
		return user, err
	}

	if res.IsSuccessState() {
		return user, nil
	} else {
		return user, errors.New("http response was not successful")
	}
}
