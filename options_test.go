package echotron

import (
	"reflect"
	"testing"
)

var (
	msgIDOpts = MessageIDOptions{
		chatID:    1,
		messageID: 2,
	}

	inlineMsgIDOpts = MessageIDOptions{
		inlineMessageID: "inline",
	}
)

func TestNewMessageID(t *testing.T) {
	new := NewMessageID(1, 2)

	if !reflect.DeepEqual(new, msgIDOpts) {
		t.Logf("expected MessageIDOptions: %+v", msgIDOpts)
		t.Logf("got MessageIDOptions: %+v", new)
		t.Fatal("error: MessageIDOptions mismatch")
	}
}

func TestNewInlineMessageID(t *testing.T) {
	new := NewInlineMessageID("inline")

	if !reflect.DeepEqual(new, inlineMsgIDOpts) {
		t.Logf("expected MessageIDOptions: %+v", inlineMsgIDOpts)
		t.Logf("got MessageIDOptions: %+v", new)
		t.Fatal("error: MessageIDOptions mismatch")
	}
}

func TestReplyKeyboardMarkupImplementsReplyMarkup(_ *testing.T) {
	i := ReplyKeyboardMarkup{}
	i.ImplementsReplyMarkup()
}

func TestReplyKeyboardRemoveImplementsReplyMarkup(_ *testing.T) {
	i := ReplyKeyboardRemove{}
	i.ImplementsReplyMarkup()
}

func TestInlineKeyboardMarkupImplementsReplyMarkup(_ *testing.T) {
	i := InlineKeyboardMarkup{}
	i.ImplementsReplyMarkup()
}

func TestForceReplyImplementsReplyMarkup(_ *testing.T) {
	i := ForceReply{}
	i.ImplementsReplyMarkup()
}
