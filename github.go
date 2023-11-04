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

func GetBranch(ctx context.Context, personalToken, repoName, ownerName string) ([]*github.Branch, error) {
	branches, _, err := MakeClient(personalToken).Repositories.ListBranches(ctx, ownerName, repoName, nil)
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func ListRepositoriesOrg(ctx context.Context, personalToken, OrgName string) (dest []*github.Repository, err error) {
	dest, _, err = MakeClient(personalToken).Repositories.ListByOrg(ctx, OrgName, nil)
	if err != nil {
		return nil, err
	}
	return
}

func ListRepositoriesOnlydDetail(ctx context.Context, personalToken, OrgName string) (dest []Repositories, err error) {
	List, err := ListRepositoriesOrg(ctx, personalToken, OrgName)
	if err != nil {
		return nil, err
	}
	dest = make([]Repositories, 0, len(List))
	for _, v := range List {
		data := Repositories{
			Name:     v.Name,
			FullName: v.FullName,
			Homepage: v.Homepage,
		}
		dest = append(dest, data)
	}
	return
}
