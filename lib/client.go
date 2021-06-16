package lib

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func githubClientContext(pat string) (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client, ctx
}
