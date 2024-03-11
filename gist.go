package main

import "github.com/google/go-github/v60/github"

type GistClient struct {
	client github.Client
}

func (client GistClient) UpdateGist(id, content string) {
	// todo!()
}
