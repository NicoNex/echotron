package echotron

import (
	"reflect"
	"testing"
)

var (
	rights = ChatAdministratorRights{
		CanManageChat:        true,
		CanDeleteMessages:    true,
		CanManageVideo_chats: true,
		CanRestrictMembers:   true,
		CanPromoteMembers:    true,
		CanChangeInfo:        true,
		CanInviteUsers:       true,
		CanPinMessages:       true,
		CanManageTopics:      true,
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
		t.Logf("expected rights: %+v", rights)
		t.Logf("got rights: %+v", res.Result)
		t.Fatal("error: rights mismatch")
	}
}
