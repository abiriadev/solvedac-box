package main

import (
	"context"

	"github.com/google/go-github/v60/github"
)

func (client *BoxClient) UpdateGist(id, filename, content string) error {
	ctx := context.Background()

	gist, _, err := client.ghClient.Gists.Get(ctx, id)
	if err != nil {
		return err
	}

	for k := range gist.Files {
		gist.Files[k] = github.GistFile{}
	}

	gist.Files[github.GistFilename(filename)] = github.GistFile{
		Content: &content,
	}

	_, _, err = client.ghClient.Gists.Edit(ctx, id, gist)
	return err
}
