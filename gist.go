package main

import (
	"context"

	"github.com/google/go-github/v60/github"
)

type GistClient struct {
	client github.Client
}

func (client GistClient) UpdateGist(id, filename, content string) error {
	ctx := context.Background()

	gist, _, err := client.client.Gists.Get(ctx, id)
	if err != nil {
		return err
	}

	gFilename := github.GistFilename(filename)
	f := gist.Files[gFilename]
	f.Content = &content
	gist.Files[gFilename] = f

	_, _, err = client.client.Gists.Edit(ctx, id, gist)
	return err
}
