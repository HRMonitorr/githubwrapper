package githubwrapper

import (
	"context"
	"fmt"
	"testing"
)

var PersonalToken = "personal"
var Reponame = "namerepo"
var OwnerName = "nameowner"

func TestGetCommitAll(t *testing.T) {
	url, err := ListCommitALL(context.Background(),
		PersonalToken,
		Reponame,
		OwnerName)
	comms, _ := GetCommit(context.Background(),
		PersonalToken,
		Reponame,
		OwnerName,
		url[0].GetSHA())
	//fmt.Printf("%+v\n", url)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", comms)
}

func TestGetBranch2(t *testing.T) {
	branches, err := GetBranch(context.Background(),
		PersonalToken,
		OwnerName,
		Reponame)
	fmt.Printf("%+v\n", branches)
	fmt.Printf("%+v\n", err)
}

func TestGetListRepositories(t *testing.T) {
	list, err := ListRepositoriesOrg(context.Background(),
		PersonalToken,
		OwnerName,
	)
	fmt.Printf("%+v\n", list)
	fmt.Printf("%+v\n", err)
}
