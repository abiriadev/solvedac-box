package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	handle            string
	bio               string
	solvedCount       int
	voteCount         int
	tier              int
	rating            int
	class             int
	classDecoration   string
	rivalCount        int
	reverseRivalCount int
	rank              int
}

func fetchUserData(username string) (User, error) {
	var user User
	res, err := http.Get("https://solved.ac/api/v3/user/show?handle=" + username)
	if err != nil {
		return user, err
	}

	err = json.NewDecoder(res.Body).Decode(&user)
	return user, err
}
