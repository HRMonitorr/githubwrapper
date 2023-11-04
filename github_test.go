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
