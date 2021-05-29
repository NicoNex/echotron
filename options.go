/*
 * Echotron
 * Copyright (C) 2018-2021  Nicolò Santamaria, Michele Dimaggio
 *
 * Echotron is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Echotron is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron

// ParseMode is a custom type for the various frequent options used by some methods of the API.
type ParseMode string

// These are all the possible options that can be used by some methods.
const (
	Markdown   ParseMode = "markdown"
	MarkdownV2           = "markdownv2"
	HTML                 = "html"
)

// ChatAction is a custom type for the various actions that can be sent through the SendChatAction method.
type ChatAction string

// These are all the possible actions that can be sent through the SendChatAction method.
const (
	Typing          ChatAction = "typing"
	UploadPhoto                = "upload_photo"
	RecordVideo                = "record_video"
	UploadVideo                = "upload_video"
	RecordAudio                = "record_audio"
	UploadAudio                = "upload_audio"
	UploadDocument             = "upload_document"
	FindLocation               = "find_location"
	RecordVideoNote            = "record_video_note"
	UploadVideoNote            = "upload_video_note"
)

// InlineQueryOptions is a custom type which contains the various options required by the AnswerInlineQueryOptions method.
type InlineQueryOptions struct {
	CacheTime         int
	IsPersonal        bool
	NextOffset        string
	SwitchPmText      string
	SwitchPmParameter string
}

type ReplyMarkup interface {
	ImplementsReplyMarkup()
}

// InlineKeyboardButton represents a button in an inline keyboard.
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard" query:"inline_keyboard"`
}

func (i InlineKeyboardMarkup) ImplementsReplyMarkup() {}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	InlineKeyboardMarkup
	ResizeKeyboard  bool `json:"resize_keyboard"`
	OneTimeKeyboard bool `json:"one_time_keyboard"`
	Selective       bool `json:"selective"`
}

// ReplyKeyboardRemove is used to remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

func (r ReplyKeyboardRemove) ImplementsReplyMarkup() {}

// ForceReply is used to display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

func (f ForceReply) ImplementsReplyMarkup() {}

// MessageOptions contains the optional parameters used in some Telegram API methods.
type MessageOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Entities                 []MessageEntity `query:"entities"`
	DisableWebPagePreview    bool            `query:"disable_web_page_preview"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// PhotoOptions contains the optional parameters used in API.SendPhoto method.
type PhotoOptions struct {
	MessageOptions `query:"recursive"`
	Caption        string `query:"caption"`
}

type InputFile struct {
	ID      string
	Path    string
	Content []byte
}

func NewInputFileID(ID string) InputFile {
	return InputFile{ID: ID}
}

func NewInputFilePath(filePath string) InputFile {
	return InputFile{Path: filePath}
}

func NewInputFileBytes(filePath string, content []byte) InputFile {
	return InputFile{Path: filePath, Content: content}
}

// AudioOptions contains the optional parameters used in API.SendAudio method.
type AudioOptions struct {
	MessageOptions `query:"recursive"`
	Caption        string    `query:"caption"`
	Duration       int       `query:"duration"`
	Performer      string    `query:"performer"`
	Title          string    `query:"title"`
	Thumb          InputFile `query:"thumb"`
}

// DocumentOptions contains the optional parameters used in API.SendDocument method.
type DocumentOptions struct {
	MessageOptions `query:"recursive"`
	Thumb          InputFile `query:"thumb"`
}

// VideoOptions contains the optional parameters used in API.SendVideo method.
type VideoOptions struct {
	MessageOptions    `query:"recursive"`
	Duration          int       `query:"duration"`
	Width             int       `query:"width"`
	Height            int       `query:"height"`
	Thumb             InputFile `query:"thumb"`
	SupportsStreaming bool      `query:"supports_streaming"`
}

// AnimationOptions contains the optional parameters used in API.SendAnimation method.
type AnimationOptions struct {
	MessageOptions `query:"recursive"`
	Duration       int       `query:"duration"`
	Width          int       `query:"width"`
	Height         int       `query:"height"`
	Thumb          InputFile `query:"thumb"`
}

// VoiceOptions contains the optional parameters used in API.SendVoice method.
type VoiceOptions MessageOptions

// VideoNoteOptions contains the optional parameters used in API.SendVideoNote method.
type VideoNoteOptions struct {
	MessageOptions `query:"recursive"`
	Duration       int       `query:"duration"`
	Length         int       `query:"length"`
	Thumb          InputFile `query:"thumb"`
}

// MediaGroupOptions contains the optional parameters used in API.SendMediaGroup method.
type MediaGroupOptions struct {
	DisableNotification      bool `query:"disable_notification"`
	ReplyToMessageID         int  `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool `query:"allow_sending_without_reply"`
}

// LocationOptions contains the optional parameters used in API.SendLocation method.
type LocationOptions struct {
	HorizontalAccuracy       float64     `query:"horizontal_accuracy"`
	LivePeriod               int         `query:"live_period"`
	Heading                  int         `query:"heading"`
	ProximityAlertRadius     int         `query:"ProximityAlertRadius"`
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}
