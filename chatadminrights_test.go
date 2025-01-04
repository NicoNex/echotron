package echotron

import (
	"reflect"
	"testing"
)

var (
	rights = ChatAdministratorRights{
		IsAnonymous:         true,
		CanManageChat:       true,
		CanDeleteMessages:   true,
		CanManageVideoChats: true,
		CanRestrictMembers:  true,
		CanPromoteMembers:   true,
		CanChangeInfo:       true,
		CanInviteUsers:      true,
		CanPostStories:      true,
		CanEditStories:      true,
		CanDeleteStories:    true,
	}
)

func TestSetMyDefaultAdministratorRights(t *testing.T) {
	_, err := api.SetMyDefaultAdministratorRights(
		&SetMyDefaultAdministratorRightsOptions{
			Rights: rights,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMyDefaultAdministratorRights(t *testing.T) {
	res, err := api.GetMyDefaultAdministratorRights(nil)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*res.Result, rights) {
		t.Logf("expected: %+v", rights)
		t.Logf("got: %+v", res.Result)
		t.Fatal("error: chat administrator rights mismatch")
	}
}
