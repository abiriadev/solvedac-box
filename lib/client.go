package lib

import (
	"github.com/google/go-github/v60/github"
	"github.com/imroc/req/v3"
)

type BoxClient struct {
	ghClient github.Client
	req      req.Client
}

func NewBoxClient(token string) BoxClient {
	return BoxClient{
		ghClient: *github.NewClient(nil).WithAuthToken(token),
		req:      *req.C(),
	}
}
