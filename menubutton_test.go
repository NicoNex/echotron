package echotron

import (
	"reflect"
	"testing"
)

var (
	menuBtn = MenuButton{
		Type: MenuButtonTypeCommands,
	}
)

func TestSetChatMenuButton(t *testing.T) {
	_, err := api.SetChatMenuButton(
		&SetChatMenuButtonOptions{
			MenuButton: menuBtn,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetChatMenuButton(t *testing.T) {
	res, err := api.GetChatMenuButton(nil)

	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(*res.Result, menuBtn) {
		t.Logf("expected menu button: %+v", menuBtn)
		t.Logf("got menu button: %+v", res.Result)
		t.Fatal("error: menu buttons mismatch")
	}
}
