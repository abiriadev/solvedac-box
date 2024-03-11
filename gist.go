package main

import (
	"context"

	"github.com/google/go-github/v60/github"
)

type GistClient struct {
	client github.Client
}

func NewGistClient(token string) GistClient {
	return GistClient{
		client: *github.NewClient(nil).WithAuthToken(token),
	}
}

func (client GistClient) UpdateGist(id, filename, content string) error {
	ctx := context.Background()

	gist, _, err := client.client.Gists.Get(ctx, id)
	if err != nil {
		return err
	}

	for k := range gist.Files {
		gist.Files[k] = github.GistFile{}
	}

	gist.Files[github.GistFilename(filename)] = github.GistFile{
		Content: &content,
	}

	_, _, err = client.client.Gists.Edit(ctx, id, gist)
	return err
}
