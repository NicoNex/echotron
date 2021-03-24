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

// Chat represents a chat.
type Chat struct {
	Id                    int              `json:"id"`
	Type                  string           `json:"type"`
	Title                 string           `json:"title,omitempty"`
	Username              string           `json:"username,omitempty"`
	FirstName             string           `json:"first_name,omitempty"`
	LastName              string           `json:"last_name,omitempty"`
	Photo                 *ChatPhoto       `json:"photo,omitempty"`
	Bio                   string           `json:"bio,omitempty"`
	Description           string           `json:"description,omitempty"`
	InviteLink            string           `json:"invite_link,omitempty"`
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`
	Permissions           *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay         int              `json:"slow_mode_delay,omitempty"`
	MessageAutoDeleteTime int              `json:"message_auto_delete_time,omitempty"`
	StickerSetName        string           `json:"sticker_set_name,omitempty"`
	CanSetStickerSet      bool             `json:"can_set_sticker_set,omitempty"`
	LinkedChatId          int64            `json:"linked_chat_id,omitempty"`
	Location              *ChatLocation    `json:"location,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	InviteLink  string `json:"invite_link"`
	Creator     *User  `json:"creator"`
	IsPrimary   bool   `json:"is_primary"`
	IsRevoked   bool   `json:"is_revoked"`
	ExpireDate  int    `json:"expire_date,omitempty"`
	MemberLimit int    `json:"member_limit,omitempty"`
}

// ChatMember contains information about one member of a chat.
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title,omitempty"`
	IsAnonymous           bool   `json:"is_anonymous,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	CanManageChat         bool   `json:"can_manage_chat,omitempty"`
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`
	CanManageVoiceChats   bool   `json:"can_manage_voice_chats,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	UntilDate             int    `json:"until_date,omitempty"`
}

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	Chat          *Chat           `json:"chat"`
	From          *User           `json:"from"`
	Date          int             `json:"date"`
	OldChatMember *ChatMember     `json:"old_chat_member"`
	NewChatMember *ChatMember     `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// User represents a Telegram user or bot.
type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"Length"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

// Message represents a message.
type Message struct {
	ID             int              `json:"message_id"`
	User           *User            `json:"from"`
	Chat           *Chat            `json:"chat"`
	Date           int64            `json:"date"`
	Text           string           `json:"text"`
	Entities       []*MessageEntity `json:"entities,omitempty"`
	Audio          *Audio           `json:"audio,omitempty"`
	Document       *Document        `json:"document,omitempty"`
	Photo          []*PhotoSize     `json:"photo,omitempty"`
	MediaGroupID   string           `json:"media_group_id,omitempty"`
	Sticker        *Sticker         `json:"sticker,omitempty"`
	Video          *Video           `json:"video,omitempty"`
	VideoNote      *VideoNote       `json:"video_note,omitempty"`
	Voice          *Voice           `json:"voice,omitempty"`
	Caption        string           `json:"caption,omitempty"`
	Contact        *Contact         `json:"contact,omitempty"`
	Location       *Location        `json:"location,omitempty"`
	NewChatMember  []*User          `json:"new_chat_members,omitempty"`
	LeftChatMember *User            `json:"left_chat_member,omitempty"`
	PinnedMessage  *Message         `json:"pinned_message,omitempty"`
}

// Update represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	ID                 int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
}

// APIResponseBase is a base type that represents the incoming response from Telegram servers.
// Used by APIResponse* to slim down the implementation.
type APIResponseBase struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

// APIResponseUpdate represents the incoming response from Telegram servers.
// Used by getUpdates (since it returns an array of Updates).
type APIResponseUpdate struct {
	Result []*Update `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseMessage represents the incoming response from Telegram servers.
// Used by the methods in the api.go module (since they return a Message).
type APIResponseMessage struct {
	Result *Message `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseCommands represents the incoming response from Telegram servers.
// Used by SetMyCommands and GetMyCommands (since they return an array of BotCommands).
type APIResponseCommands struct {
	Result []BotCommand `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseChat represents the incoming response from Telegram servers.
// Used by GetChat (since it returns a Chat).
type APIResponseChat struct {
	Result *Chat `json:"result,omitempty"`
	APIResponseBase
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	Longitude float32
	Latitude  float32
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot,
// the field message will be present. If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID              string   `json:"id"`
	User            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	Data            string   `json:"data,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string     `json:"file_id"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer,omitempty"`
	Title     string     `json:"title,omitempty"`
	MimeType  string     `json:"mime_type,omitempty"`
	FileSize  int        `json:"file_size,omitempty"`
	Thumb     *PhotoSize `json:"thumb,omitempty"`
}

// Video represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// VideoNote represents a video message.
type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// PhotoSize represents one size of a photo or a file/sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"FileSize"`
}

// Voice represents a voice note.
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"FileSize,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// Sticker represents a sticker.
type Sticker struct {
	FileID       string       `json:"file_id"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	Thumb        *PhotoSize   `json:"thumb,omitempty"`
	Emoji        string       `json:"emoji,omitempty"`
	SetName      string       `json:"set_name,omitempty"`
	FileSize     int          `json:"file_size,omitempty"`
	MaskPosition MaskPosition `json:"mask_position"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"sticker"`
}

// Button represents a button in a keyboard.
type Button struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// KbdRow represents a row of buttons in a keyboard.
type KbdRow []Button

// Keyboard represents a keyboard.
type Keyboard struct {
	Keyboard        []KbdRow `json:"keyboard"`
	ResizeKeyboard  bool     `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool     `json:"one_time_keyboard,omitempty"`
	Selective       bool     `json:"selective,omitempty"`
}

// KeyboardRemove represents a keyboard removal request.
type KeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// InlineButton represents a button in an inline keyboard.
type InlineButton struct {
	Text         string `json:"text"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

// InlineKbdRow represents a row of buttons in an inline keyboard.
type InlineKbdRow []InlineButton

// InlineKeyboard represents an inline keyboard.
type InlineKeyboard struct {
	InlineKeyboard []InlineKbdRow `json:"inline_keyboard"`
}

// BotCommand represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
