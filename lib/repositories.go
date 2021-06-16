package lib

import (
	"context"

	"github.com/google/go-github/github"
)

func GetRepositoryNames(org string, pat string) ([]string, error) {
	client, ctx := githubClientContext(pat)

	return getRepositoryNames(ctx, client, org)
}

func getRepositoryNames(ctx context.Context, client *github.Client, org string) ([]string, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	var repoNames []string
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			repoNames = append(repoNames, repo.GetName())
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return repoNames, nil
}
