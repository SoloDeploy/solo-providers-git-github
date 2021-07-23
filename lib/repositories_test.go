package lib

import (
	"context"
	"testing"

	"github.com/google/go-github/github"
)

func TestGetRepositoryNames(t *testing.T) {
	ctx := context.Background()
	client := github.NewClient(nil)

	client.Repositories.ListByOrg = func(ctx context.Context, org string, opt *github.RepositoryListByOrgOptions) ([]*github.Repository, *github.Response, error) {
		var repos []*github.Repository
		r := "repo1"
		rp := &r
		repo1 := &github.Repository{Name: rp}
		resp := &github.Response{NextPage: 0}
		repos = append(repos, repo1)
		return repos, resp, nil
	}
	names, _ := getRepositoryNames(ctx, client, "org")

	if len(names) == 0 {
		t.Errorf("Expected repositories to be returned")
	}
}
