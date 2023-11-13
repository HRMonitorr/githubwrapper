package githubwrapper

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/go-github/v56/github"
	"io"
	"log"
	"os"
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

func UploadFileToRepository(val PushRepositories) (response *github.RepositoryContentResponse, err error) {
	file, err := os.Open(val.Path)
	if err != nil {
		log.Fatalf("Gagal open file %s", err.Error())
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String(val.Message),
		Content:   bs,
		Branch:    github.String(val.Branch),
		Committer: &github.CommitAuthor{Name: github.String(val.Username), Email: github.String(val.Email)},
	}
	response, _, err = MakeClient(val.PersonalToken).Repositories.CreateFile(val.Context, val.OwnerName, val.Reponame, val.Path, opts)
	if err != nil {
		fmt.Printf("%+v", err.Error())
		return
	}
	return
}
