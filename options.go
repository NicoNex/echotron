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
	Caption        string `query:"caption"`
	MessageOptions `query:"recursive"`
}
