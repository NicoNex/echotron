/*
 * Echotron
 * Copyright (C) 2018-2021  The Echotron Devs
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

import (
	"net/url"
	"strconv"
)

// ParseMode is a custom type for the various frequent options used by some methods of the API.
type ParseMode string

// These are all the possible options that can be used by some methods.
const (
	Markdown   ParseMode = "Markdown"
	MarkdownV2           = "MarkdownV2"
	HTML                 = "HTML"
)

type PollType string

const (
	Quiz    PollType = "quiz"
	Regular          = "regular"
	Any              = ""
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

type ReplyMarkup interface {
	ImplementsReplyMarkup()
}

// KeyboardButton represents a button in a keyboard.
type KeyboardButton struct {
	Text            string   `json:"text"`
	RequestContact  bool     `json:"request_contact,omitempty"`
	RequestLocation bool     `json:"request_location,omitempty"`
	RequestPoll     PollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type PollType `json:"type"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

func (i ReplyKeyboardMarkup) ImplementsReplyMarkup() {}

// ReplyKeyboardRemove is used to remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
// RemoveKeyboard MUST BE true.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

func (r ReplyKeyboardRemove) ImplementsReplyMarkup() {}

// InlineKeyboardButton represents a button in an inline keyboard.
type InlineKeyboardButton struct {
	Text                         string   `json:"text"`
	URL                          string   `json:"url,omitempty"`
	LoginURL                     LoginURL `json:"login_url,omitempty"`
	CallbackData                 string   `json:"callback_data,omitempty"`
	SwitchInlineQuery            string   `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string   `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 CallbackGame `json:"callback_game,omitempty"`
	Pay bool `json:"pay,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard" query:"inline_keyboard"`
}

func (i InlineKeyboardMarkup) ImplementsReplyMarkup() {}

// ForceReply is used to display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"`
}

func (f ForceReply) ImplementsReplyMarkup() {}

type BaseOptions struct {
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}

// MessageOptions contains the optional parameters used in some Telegram API methods.
type MessageOptions struct {
	BaseOptions           `query:"recursive"`
	ParseMode             ParseMode       `query:"parse_mode"`
	Entities              []MessageEntity `query:"entities"`
	DisableWebPagePreview bool            `query:"disable_web_page_preview"`
}

type InputFile struct {
	id      string
	path    string
	content []byte
}

func NewInputFileID(ID string) InputFile {
	return InputFile{id: ID}
}

func NewInputFilePath(filePath string) InputFile {
	return InputFile{path: filePath}
}

func NewInputFileBytes(fileName string, content []byte) InputFile {
	return InputFile{path: fileName, content: content}
}

// PhotoOptions contains the optional parameters used in API.SendPhoto method.
type PhotoOptions struct {
	BaseOptions     `query:"recursive"`
	ParseMode       ParseMode       `query:"parse_mode"`
	Caption         string          `query:"caption"`
	CaptionEntities []MessageEntity `query:"caption_entities"`
}

// AudioOptions contains the optional parameters used in API.SendAudio method.
type AudioOptions struct {
	BaseOptions     `query:"recursive"`
	ParseMode       ParseMode       `query:"parse_mode"`
	Caption         string          `query:"caption"`
	CaptionEntities []MessageEntity `query:"caption_entities"`
	Duration        int             `query:"duration"`
	Performer       string          `query:"performer"`
	Title           string          `query:"title"`
	Thumb           InputFile       `query:"thumb"`
}

// DocumentOptions contains the optional parameters used in API.SendDocument method.
type DocumentOptions struct {
	BaseOptions                 `query:"recursive"`
	ParseMode                   ParseMode       `query:"parse_mode"`
	Caption                     string          `query:"caption"`
	CaptionEntities             []MessageEntity `query:"caption_entities"`
	DisableContentTypeDetection bool            `query:"disable_content_type_detection"`
	Thumb                       InputFile       `query:"thumb"`
}

// VideoOptions contains the optional parameters used in API.SendVideo method.
type VideoOptions struct {
	BaseOptions       `query:"recursive"`
	ParseMode         ParseMode       `query:"parse_mode"`
	Caption           string          `query:"caption"`
	CaptionEntities   []MessageEntity `query:"caption_entities"`
	Duration          int             `query:"duration"`
	Width             int             `query:"width"`
	Height            int             `query:"height"`
	Thumb             InputFile       `query:"thumb"`
	SupportsStreaming bool            `query:"supports_streaming"`
}

// AnimationOptions contains the optional parameters used in API.SendAnimation method.
type AnimationOptions struct {
	BaseOptions     `query:"recursive"`
	ParseMode       ParseMode       `query:"parse_mode"`
	Caption         string          `query:"caption"`
	CaptionEntities []MessageEntity `query:"caption_entities"`
	Duration        int             `query:"duration"`
	Width           int             `query:"width"`
	Height          int             `query:"height"`
	Thumb           InputFile       `query:"thumb"`
}

// VoiceOptions contains the optional parameters used in API.SendVoice method.
type VoiceOptions struct {
	BaseOptions     `query:"recursive"`
	ParseMode       ParseMode       `query:"parse_mode"`
	Caption         string          `query:"caption"`
	CaptionEntities []MessageEntity `query:"caption_entities"`
	Duration        int             `query:"duration"`
}

// VideoNoteOptions contains the optional parameters used in API.SendVideoNote method.
type VideoNoteOptions struct {
	BaseOptions `query:"recursive"`
	Duration    int       `query:"duration"`
	Length      int       `query:"length"`
	Thumb       InputFile `query:"thumb"`
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

// ContactOptions contains the optional parameters used in API.SendContact method.
type ContactOptions struct {
	BaseOptions `query:"recursive"`
	LastName    string `query:"last_name"`
	VCard       string `query:"vcard"`
}

// CallbackQueryOptions contains the optional parameters used in API.AnswerCallbackQuery method.
type CallbackQueryOptions struct {
	Text      string `query:"text"`
	ShowAlert bool   `query:"show_alert"`
	URL       string `query:"url"`
	CacheTime int    `query:"cache_time"`
}

type MessageIDOptions struct {
	chatID          int64  `query:"chat_id"`
	messageID       int    `query:"message_id"`
	inlineMessageID string `query:"inline_message_id"`
}

func NewMessageID(chatID int64, messageID int) *MessageIDOptions {
	return &MessageIDOptions{chatID: chatID, messageID: messageID}
}

func NewInlineMessageID(ID string) *MessageIDOptions {
	return &MessageIDOptions{inlineMessageID: ID}
}

func (m MessageIDOptions) EncodeValues(key string, v *url.Values) error {
	if m.chatID != 0 {
		v.Add("chat_id", strconv.FormatInt(m.chatID, 10))
	}
	if m.messageID != 0 {
		v.Add("message_id", strconv.FormatInt(int64(m.messageID), 10))
	}
	if m.inlineMessageID != "" {
		v.Add("inline_message_id", m.inlineMessageID)
	}
	return nil
}

type MessageTextOptions struct {
	ParseMode             string               `query:"parse_mode"`
	Entities              []MessageEntity      `query:"entities"`
	DisableWebPagePreview bool                 `query:"disable_web_page_preview"`
	ReplyMarkup           InlineKeyboardMarkup `query:"reply_markup"`
}

type MessageCaptionOptions struct {
	Caption         string               `query:"caption"`
	ParseMode       string               `query:"parse_mode"`
	CaptionEntities []MessageEntity      `query:"caption_entities"`
	ReplyMarkup     InlineKeyboardMarkup `query:"reply_markup"`
}

type MessageMediaOptions struct {
	Media       InputMedia           `query:"media"`
	ReplyMarkup InlineKeyboardMarkup `query:"reply_markup"`
}

type MessageReplyMarkup struct {
	ReplyMarkup InlineKeyboardMarkup `query:"reply_markup"`
}
