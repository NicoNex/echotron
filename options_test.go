package echotron

import "testing"

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
