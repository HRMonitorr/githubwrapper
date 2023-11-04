package githubwrapper

import (
	"context"
	"github.com/google/go-github/v56/github"
)

func MakeClient(personalToken string) *github.Client {
	client := github.NewClient(nil).WithAuthToken(personalToken)
	return client
}

func ListCommitALL(ctx context.Context, PersonalToken, repoName, ownerName string) ([]*github.RepositoryCommit, error) {
	commits, _, err := MakeClient(PersonalToken).Repositories.ListCommits(ctx, ownerName, repoName, nil)
	if err != nil {
		return nil, err
	}

	return commits, nil
}

func GetCommit(ctx context.Context, personalToken, repoName, ownerName, sha string) (*github.RepositoryCommit, error) {
	commits, _, err := MakeClient(personalToken).Repositories.GetCommit(ctx, ownerName, repoName, sha, nil)
	if err != nil {
		return nil, err
	}
	return commits, nil
}
