package githubwrapper

import (
	"context"
	"github.com/google/go-github/v56/github"
	"golang.org/x/oauth2"
)

func GetCommits(ctx context.Context, personalToken, repoName, ownerName string) ([]*github.RepositoryCommit, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personalToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	commits, _, err := client.Repositories.ListCommits(ctx, ownerName, repoName, nil)
	if err != nil {
		return nil, err
	}

	return commits, nil
}
