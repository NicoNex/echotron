/*
 * Echotron
 * Copyright (C) 2019  Nicolò Santamaria, Michele Dimaggio
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron


// This object represents a chat.
type Chat struct {
	ID int64 `json:"id"`
	Type string `json:"type"`
	Title string `json:"title,omitempty"`
	Username string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	AllMembersAreAdmin bool `json:"all_members_are_administrators,omitempty"`
	Description string `json:"description,omitempty"`
	InviteLink string `json:"invite_link,omitempty"`
	PinnedMessage *Message `json:"pinned_message,omitempty"`
}


// This object represents a Telegram user or bot.
type User struct {
	ID int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	Username string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}


// This object represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type string `json:"type"`
	Offset int `json:"offset"`
	Length int `json:"Length"`
	Url string `json:"url,omitempty"`
	User *User `json:"user,omitempty"`
}


// WIP: this object represents a message.
type Message struct {
	ID int `json:"message_id"`
	User *User `json:"from"`
	Chat *Chat `json:"chat"`
	Date int64 `json:"date"`
	Text string `json:"text"`
	Entities []*MessageEntity `json:"entities,omitempty"`
	Audio *Audio `json:"audio,omitempty"`
	Document *Document `json:"document,omitempty"`
	Photo []*PhotoSize `json:"photo,omitempty"`
	MediaGroupId string `json:"media_group_id,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
	Video *Video `json:"video,omitempty"`
	VideoNote *VideoNote `json:"video_note,omitempty"`
	Voice *Voice `json:"voice,omitempty"`
	Caption string `json:"caption,omitempty"`
	Contact *Contact `json:"contact,omitempty"`
	Location *Location `json:"location,omitempty"`
	NewChatMember []*User `json:"new_chat_members,omitempty"`
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	PinnedMessage *Message `json:"pinned_message,omitempty"`
}


// This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	ID int `json:"update_id"`
	Message *Message `json:"message,omitempty"`
	EditedMessage *Message `json:"edited_message,omitempty"`
	ChannelPost *Message `json:"channel_post,omitempty"`
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
}


// This object represents the incoming response from Telegram servers.
type APIResponse struct {
	Ok bool `json:"ok"`
	Result	[]*Update `json:"result,omitempty"`
	ErrorCode int `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}


// This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name,omitempty"`
	UserID int `json:"user_id,omitempty"`
	Vcard string `json:"vcard,omitempty"`
}



// This object represents a point on the map.
type Location struct {
	Longitude float32
	Latitude float32
}


// This object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID string `json:"id"`
	User *User `json:"user"`
	Query string `json:"query"`
	Offset string `json:"offset"`
}


// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ID string `json:"result_id"`
	User *User `json:"user"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Query string `json:"query,omitempty"`
}


// This object represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot,
// the field message will be present. If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID string `json:"id"`
	User *User `json:"user"`
	Message string `json:"message,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
}


// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileId string `json:"file_id"`
	Duration int `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title string `json:"title,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
}


// This object represents a video file.
type Video struct {
	FileId string `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Duration int `json:"duration"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int `json:"file_size,omitempty"`
}


// This object represents a video message.
type VideoNote struct {
	FileId string `json:"file_id"`
	Length int `json:"length"`
	Duration int `json:"duration"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	FileSize int  `json:"file_size,omitempty"`
}


// This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileId string `json:"file_id"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	FileName string `json:"file_name,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int `json:"file_size,omitempty"`
}


// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileId string `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	FileSize int `json:"FileSize"`
}


// This object represents a voice note.
type Voice struct {
	FileId string `json:"file_id"`
	Duration int `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int `json:"FileSize,omitempty"`
}


// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point string `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale float32 `json:"scale"`
}


// This object represents a sticker.
type Sticker struct {
	FileId string `json:"file_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Thumb *PhotoSize `json:"thumb,omitempty"`
	Emoji string `json:"emoji,omitempty"`
	SetName string `json:"set_name,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	MaskPosition MaskPosition `json:"mask_position"`
}


// This object represents a sticker set.
type StickerSet struct {
	Name string `json:"name"`
	Title string `json:"title"`
	ContainsMasks bool `json:"contains_masks"`
	Stickers []*Sticker `json:"sticker"`
}


// N/A
type Button struct {
	Text string `json:"text"`
	RequestContact bool `json:"request_contact,omitempty"`
	RequestLocation bool `json:"request_location,omitempty"`
}


// N/A
type KbdRow []Button


// N/A
type Keyboard struct {
	Keyboard []KbdRow `json:"keyboard"`
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	Selective bool `json:"selective,omitempty"`
}


// N/A
type KeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective bool `json:"selective,omitempty"`
}


// N/A
type InlineButton struct {
	Text string `json:"text"`
	URL string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}


// N/A
type InlineKbdRow []InlineButton


// N/A
type InlineKeyboard struct{
	InlineKeyboard []InlineKbdRow `json:"inline_keyboard"`
}
