package main

import (
	"github.com/abiriadev/solvedac-box/lib"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GhToken  string `envconfig:"GH_TOKEN" required:"true"`
	Username string `envconfig:"USERNAME" required:"true"`
	GistId   string `envconfig:"GIST_ID"  required:"true"`
}

func main() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}

	client := lib.NewBoxClient(
		config.GhToken,
	)

	res, err := client.FetchUserData(config.Username)
	if err != nil {
		panic(err)
	}

	rendered, err := res.Render()
	if err != nil {
		panic(err)
	}

	err = client.UpdateGist(config.GistId, "💯 My Solved.ac Profile", rendered)
	if err != nil {
		panic(err)
	}
}
