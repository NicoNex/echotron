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

// ParseMode is a custom type for the various frequent options used by some methods of the API.
type ParseMode string

// These are all the possible options that can be used by some methods.
const (
	Markdown   ParseMode = "Markdown"
	MarkdownV2           = "MarkdownV2"
	HTML                 = "HTML"
)

// PollType is a custom type for the various types of poll that can be sent.
type PollType string

// These are all the possible poll types.
const (
	Quiz    PollType = "quiz"
	Regular          = "regular"
	Any              = ""
)

// DiceEmoji is a custom type for the various emojis that can be sent through the SendDice method.
type DiceEmoji string

// These are all the possible emojis that can be sent through the SendDice method.
const (
	Die     DiceEmoji = "🎲"
	Darts             = "🎯"
	Basket            = "🏀"
	Goal              = "⚽️"
	Bowling           = "🎳"
	Slot              = "🎰"
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

// UpdateType is a custom type for the various update types that a bot can be subscribed to.
type UpdateType string

// These are all the possible types that a bot can be subscribed to.
const (
	MessageUpdate            UpdateType = "message"
	EditedMessageUpdate                 = "edited_message"
	ChannelPostUpdate                   = "channel_post"
	EditedChannelPostUpdate             = "edited_channel_post"
	InlineQueryUpdate                   = "inline_query"
	ChosenInlineResultUpdate            = "chosen_inline_result"
	CallbackQueryUpdate                 = "callback_query"
	ShippingQueryUpdate                 = "shipping_query"
	PreCheckoutQueryUpdate              = "pre_checkout_query"
	PollUpdate                          = "poll"
	PollAnswerUpdate                    = "poll_answer"
	MyChatMemberUpdate                  = "my_chat_member"
	ChatMemberUpdate                    = "chat_member"
)

// ReplyMarkup is an interface for the various keyboard types.
type ReplyMarkup interface {
	ImplementsReplyMarkup()
}

// KeyboardButton represents a button in a keyboard.
type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`
	RequestLocation bool                    `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type PollType `json:"type"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Selective             bool               `json:"selective,omitempty"`
}

// ImplementsReplyMarkup is a dummy method which exists to implement the interface ReplyMarkup.
func (i ReplyKeyboardMarkup) ImplementsReplyMarkup() {}

// ReplyKeyboardRemove is used to remove the current custom keyboard and display the default letter-keyboard.
// By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
// RemoveKeyboard MUST BE true.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// ImplementsReplyMarkup is a dummy method which exists to implement the interface ReplyMarkup.
func (r ReplyKeyboardRemove) ImplementsReplyMarkup() {}

// InlineKeyboardButton represents a button in an inline keyboard.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          string        `json:"url,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	CallbackData                 string        `json:"callback_data,omitempty"`
	SwitchInlineQuery            string        `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string        `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard" query:"inline_keyboard"`
}

// ImplementsReplyMarkup is a dummy method which exists to implement the interface ReplyMarkup.
func (i InlineKeyboardMarkup) ImplementsReplyMarkup() {}

// ForceReply is used to display a reply interface to the user (act as if the user has selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective"`
}

// ImplementsReplyMarkup is a dummy method which exists to implement the interface ReplyMarkup.
func (f ForceReply) ImplementsReplyMarkup() {}

// UpdateOptions contains the optional parameters used by the GetUpdates method.
type UpdateOptions struct {
	Offset         int          `query:"offset"`
	Limit          int          `query:"limit"`
	Timeout        int          `query:"timeout"`
	AllowedUpdates []UpdateType `query:"allowed_updates"`
}

// WebhookOptions contains the optional parameters used by the SetWebhook method.
type WebhookOptions struct {
	Certificate    InputFile
	IPAddress      string       `query:"ip_address"`
	MaxConnections int          `query:"max_connections"`
	AllowedUpdates []UpdateType `query:"allowed_updates"`
}

// BaseOptions contains the optional parameters used frequently in some Telegram API methods.
type BaseOptions struct {
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}

// MessageOptions contains the optional parameters used by some Telegram API methods.
type MessageOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Entities                 []MessageEntity `query:"entities"`
	DisableWebPagePreview    bool            `query:"disable_web_page_preview"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// DisableNotificationOptions contains the optional parameters used by some Telegram API methods.
type DisableNotificationOptions struct {
	DisableNotification bool `query:"disable_notification"`
}

// CopyOptions contains the optional parameters used by the CopyMessage method.
type CopyOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Caption                  string          `query:"caption"`
	CaptionEntities          []MessageEntity `query:"caption_entities"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// InputFile is a struct which contains data about a file to be sent.
type InputFile struct {
	id      string
	path    string
	ftype   string
	content []byte
}

// NewInputFileID is a wrapper for InputFile which only fills the id field.
func NewInputFileID(ID string) InputFile {
	return InputFile{id: ID}
}

// NewInputFilePath is a wrapper for InputFile which only fills the path field.
func NewInputFilePath(filePath string) InputFile {
	return InputFile{path: filePath}
}

// NewInputFileBytes is a wrapper for InputFile which only fills the path and content fields.
func NewInputFileBytes(fileName string, content []byte) InputFile {
	return InputFile{path: fileName, content: content}
}

// PhotoOptions contains the optional parameters used by the SendPhoto method.
type PhotoOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Caption                  string          `query:"caption"`
	CaptionEntities          []MessageEntity `query:"caption_entities"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// AudioOptions contains the optional parameters used by the SendAudio method.
type AudioOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Caption                  string          `query:"caption"`
	CaptionEntities          []MessageEntity `query:"caption_entities"`
	Duration                 int             `query:"duration"`
	Performer                string          `query:"performer"`
	Title                    string          `query:"title"`
	Thumb                    InputFile
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// DocumentOptions contains the optional parameters used by the SendDocument method.
type DocumentOptions struct {
	ParseMode                   ParseMode       `query:"parse_mode"`
	Caption                     string          `query:"caption"`
	CaptionEntities             []MessageEntity `query:"caption_entities"`
	DisableContentTypeDetection bool            `query:"disable_content_type_detection"`
	Thumb                       InputFile
	DisableNotification         bool            `query:"disable_notification"`
	ReplyToMessageID            int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply    bool            `query:"allow_sending_without_reply"`
	ReplyMarkup                 ReplyMarkup     `query:"reply_markup"`
}

// VideoOptions contains the optional parameters used by the SendVideo method.
type VideoOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Caption                  string          `query:"caption"`
	CaptionEntities          []MessageEntity `query:"caption_entities"`
	Duration                 int             `query:"duration"`
	Width                    int             `query:"width"`
	Height                   int             `query:"height"`
	Thumb                    InputFile
	SupportsStreaming        bool            `query:"supports_streaming"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// AnimationOptions contains the optional parameters used by the SendAnimation method.
type AnimationOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Caption                  string          `query:"caption"`
	CaptionEntities          []MessageEntity `query:"caption_entities"`
	Duration                 int             `query:"duration"`
	Width                    int             `query:"width"`
	Height                   int             `query:"height"`
	Thumb                    InputFile
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// VoiceOptions contains the optional parameters used by the SendVoice method.
type VoiceOptions struct {
	ParseMode                ParseMode       `query:"parse_mode"`
	Caption                  string          `query:"caption"`
	CaptionEntities          []MessageEntity `query:"caption_entities"`
	Duration                 int             `query:"duration"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// VideoNoteOptions contains the optional parameters used by the SendVideoNote method.
type VideoNoteOptions struct {
	Duration                 int         `query:"duration"`
	Length                   int         `query:"length"`
	Thumb                    InputFile
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}

// MediaGroupOptions contains the optional parameters used by the SendMediaGroup method.
type MediaGroupOptions struct {
	DisableNotification      bool `query:"disable_notification"`
	ReplyToMessageID         int  `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool `query:"allow_sending_without_reply"`
}

// LocationOptions contains the optional parameters used by the SendLocation method.
type LocationOptions struct {
	HorizontalAccuracy       float64     `query:"horizontal_accuracy"`
	LivePeriod               int         `query:"live_period"`
	Heading                  int         `query:"heading"`
	ProximityAlertRadius     int         `query:"proximity_alert_radius"`
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}

// EditLocationOptions contains the optional parameters used by the EditMessageLiveLocation method.
type EditLocationOptions struct {
	HorizontalAccuracy   float64              `query:"horizontal_accuracy"`
	Heading              int                  `query:"heading"`
	ProximityAlertRadius int                  `query:"proximity_alert_radius"`
	ReplyMarkup          InlineKeyboardMarkup `query:"reply_markup"`
}

// VenueOptions contains the optional parameters used by the SendVenue method.
type VenueOptions struct {
	FoursquareID             string      `query:"foursquare_id"`
	FoursquareType           string      `query:"foursquare_type"`
	GooglePlaceID            string      `query:"google_place_id"`
	GooglePlaceType          string      `query:"google_place_type"`
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}

// ContactOptions contains the optional parameters used by the SendContact method.
type ContactOptions struct {
	LastName                 string      `query:"last_name"`
	VCard                    string      `query:"vcard"`
	DisableNotification      bool        `query:"disable_notification"`
	ReplyToMessageID         int         `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool        `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup `query:"reply_markup"`
}

// PollOptions contains the optional parameters used by the SendPoll method.
type PollOptions struct {
	IsAnonymous              bool            `query:"is_anonymous"`
	Type                     PollType        `query:"type"`
	AllowsMultipleAnswers    bool            `query:"allows_multiple_answers"`
	CorrectOptionID          int             `query:"correct_option_id"`
	Explanation              string          `query:"explanation"`
	ExplanationParseMode     ParseMode       `query:"explanation_parse_mode"`
	ExplanationEntities      []MessageEntity `query:"explanation_entities"`
	OpenPeriod               int             `query:"open_period"`
	CloseDate                int             `query:"close_date"`
	IsClosed                 bool            `query:"is_closed"`
	DisableNotification      bool            `query:"disable_notification"`
	ReplyToMessageID         int             `query:"reply_to_message_id"`
	AllowSendingWithoutReply bool            `query:"allow_sending_without_reply"`
	ReplyMarkup              ReplyMarkup     `query:"reply_markup"`
}

// BanOptions contains the optional parameters used by the BanChatMember method.
type BanOptions struct {
	UntilDate      int  `query:"until_date"`
	RevokeMessages bool `query:"revoke_messages"`
}

// UnbanOptions contains the optional parameters used by the UnbanChatMember method.
type UnbanOptions struct {
	OnlyIfBanned bool `query:"only_if_banned"`
}

// RestrictOptions contains the optional parameters used by the RestrictChatMember method.
type RestrictOptions struct {
	UntilDate int `query:"until_date"`
}

// PromoteOptions contains the optional parameters used by the PromoteChatMember method.
type PromoteOptions struct {
	IsAnonymous         bool `query:"is_anonymous"`
	CanManageChat       bool `query:"can_manage_chat"`
	CanPostMessages     bool `query:"can_post_messages"`
	CanEditMessages     bool `query:"can_edit_messages"`
	CanDeleteMessages   bool `query:"can_delete_messages"`
	CanManageVoiceChats bool `query:"can_manage_voice_chats"`
	CanRestrictMembers  bool `query:"can_restrict_members"`
	CanPromoteMembers   bool `query:"can_promote_members"`
	CanChangeInfo       bool `query:"can_change_info"`
	CanInviteUsers      bool `query:"can_invite_users"`
	CanPinMessages      bool `query:"can_pin_messages"`
}

// UserProfileOptions contains the optional parameters used by the GetUserProfilePhotos method.
type UserProfileOptions struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

// InviteLinkOptions contains the optional parameters used by the CreateChatInviteLink and EditChatInviteLink methods.
type InviteLinkOptions struct {
	ExpireDate  int64 `query:"expire_date"`
	MemberLimit int   `query:"member_limit"`
}

// CallbackQueryOptions contains the optional parameters used by the AnswerCallbackQuery method.
type CallbackQueryOptions struct {
	Text      string `query:"text"`
	ShowAlert bool   `query:"show_alert"`
	URL       string `query:"url"`
	CacheTime int    `query:"cache_time"`
}

// MessageIDOptions is a struct which contains data about a message to edit.
type MessageIDOptions struct {
	chatID          int64  `query:"chat_id"`
	messageID       int    `query:"message_id"`
	inlineMessageID string `query:"inline_message_id"`
}

// NewMessageID is a wrapper for MessageIDOptions which only fills the chatID and messageID fields.
func NewMessageID(chatID int64, messageID int) MessageIDOptions {
	return MessageIDOptions{chatID: chatID, messageID: messageID}
}

// NewInlineMessageID is a wrapper for MessageIDOptions which only fills the inlineMessageID fields.
func NewInlineMessageID(ID string) MessageIDOptions {
	return MessageIDOptions{inlineMessageID: ID}
}

// MessageTextOptions contains the optional parameters used by the EditMessageText method.
type MessageTextOptions struct {
	ParseMode             ParseMode            `query:"parse_mode"`
	Entities              []MessageEntity      `query:"entities"`
	DisableWebPagePreview bool                 `query:"disable_web_page_preview"`
	ReplyMarkup           InlineKeyboardMarkup `query:"reply_markup"`
}

// MessageCaptionOptions contains the optional parameters used by the EditMessageCaption method.
type MessageCaptionOptions struct {
	Caption         string               `query:"caption"`
	ParseMode       ParseMode            `query:"parse_mode"`
	CaptionEntities []MessageEntity      `query:"caption_entities"`
	ReplyMarkup     InlineKeyboardMarkup `query:"reply_markup"`
}

// MessageReplyMarkup contains the optional parameters used by the method which only require reply_markup.
type MessageReplyMarkup struct {
	ReplyMarkup InlineKeyboardMarkup `query:"reply_markup"`
}

// CommandOptions contains the optional parameters used by the SetMyCommands, DeleteMyCommands and GetMyCommands methods.
type CommandOptions struct {
	Scope        BotCommandScope `query:"scope"`
	LanguageCode string          `query:"language_code"`
}
