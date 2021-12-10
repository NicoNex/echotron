package echokit

import (
	"github.com/NicoNex/echotron/v3"
)

// InlineKeyboard allows to quickly cast [][]echotron.InlineKeyboardButton into echotron.ReplyMarkup
func InlineKeyboard(kbd [][]echotron.InlineKeyboardButton) echotron.ReplyMarkup {
	return echotron.InlineKeyboardMarkup{InlineKeyboard: kbd}
}

// GenInlineKeyboard allows to generate an echotron.ReplyMarkup from a list of buttons
func GenInlineKeyboard(column int, fromList ...echotron.InlineKeyboardButton) echotron.ReplyMarkup {
	if column < 1 {
		return echotron.InlineKeyboardMarkup{}
	}

	var (
		row      []echotron.InlineKeyboardButton
		finalKbd [][]echotron.InlineKeyboardButton
	)

	for _, btn := range fromList {
		row = append(row, btn)
		if len(row) >= column {
			finalKbd = append(finalKbd, row)
			row = nil
		}
	}
	if len(row) > 0 {
		finalKbd = append(finalKbd, row)
	}

	return InlineKeyboard(finalKbd)
}

// ForceReply allows to force the user to reply to the current message when sent
func ForceReply(placeholder string, selective bool) echotron.ReplyMarkup {
	return echotron.ForceReply{
		ForceReply:            true,
		InputFieldPlaceholder: placeholder,
		Selective:             selective,
	}
}

// RemoveKeyboard allows to remove a normal (not-inline) keyboard when sent
func RemoveKeyboard(selective bool) echotron.ReplyMarkup {
	return echotron.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// ExtractMessageIDOpt generate and return the MessageIDOptions of a given update using the ID and SenderChat
func ExtractMessageIDOpt(update *echotron.Update) *echotron.MessageIDOptions {
	var (
		message *echotron.Message
		msgID   echotron.MessageIDOptions
		userID  int64
	)

	switch {
	case update.Message != nil:
		message = update.Message
		userID = message.From.ID
	case update.EditedMessage != nil:
		message = update.EditedMessage
		userID = message.From.ID
	case update.ChannelPost != nil:
		message = update.ChannelPost
		userID = message.SenderChat.ID
	case update.EditedChannelPost != nil:
		message = update.EditedChannelPost
		userID = message.SenderChat.ID
	case update.InlineQuery != nil:
		msgID = echotron.NewInlineMessageID(update.InlineQuery.ID)
		return &msgID
	case update.ChosenInlineResult != nil:
		msgID = echotron.NewInlineMessageID(update.ChosenInlineResult.ResultID)
		return &msgID
	case update.CallbackQuery != nil:
		message = update.CallbackQuery.Message
		if message == nil {
			msgID = echotron.NewInlineMessageID(update.CallbackQuery.ID)
			return &msgID
		}
		userID = update.CallbackQuery.From.ID
	}

	if message == nil {
		return nil
	}

	msgID = echotron.NewMessageID(userID, message.ID)
	return &msgID
}
