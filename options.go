/*
 * Echotron
 * Copyright (C) 2018-2022 The Echotron Devs
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
	ChooseSticker              = "choose_sticker"
)

// MessageEntityType is a custom type for the various MessageEntity types used in various methods.
type MessageEntityType string

// These are all the possible types for MessageEntityType.
const (
	MentionEntity       MessageEntityType = "mention"
	HashtagEntity                         = "hashtag"
	CashtagEntity                         = "cashtag"
	BotCommandEntity                      = "bot_command"
	UrlEntity                             = "url"
	EmailEntity                           = "email"
	PhoneNumberEntity                     = "phone_number"
	BoldEntity                            = "bold"
	ItalicEntity                          = "italic"
	UnderlineEntity                       = "underline"
	StrikethroughEntity                   = "strikethrough"
	SpoilerEntity                         = "spoiler"
	CodeEntity                            = "code"
	PreEntity                             = "pre"
	TextLinkEntity                        = "text_link"
	TextMentionEntity                     = "text_mention"
	CustomEmojiEntity                     = "custom_emoji"
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
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	Text            string                      `json:"text"`
	RequestContact  bool                        `json:"request_contact,omitempty"`
	RequestLocation bool                        `json:"request_location,omitempty"`
}

// KeyboardButtonPollType represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type PollType `json:"type"`
}

// KeyboardButtonRequestUsers defines the criteria used to request suitable users.
// The identifiers of the selected users will be shared with the bot when the corresponding button is pressed.
type KeyboardButtonRequestUsers struct {
	RequestID     int  `json:"request_id"`
	MaxQuantity   int  `json:"max_quantity,omitempty"`
	UserIsBot     bool `json:"user_is_bot,omitempty"`
	UserIsPremium bool `json:"user_is_premium,omitempty"`
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable chat.
// The identifier of the selected chat will be shared with the bot when the corresponding button is pressed.
type KeyboardButtonRequestChat struct {
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	RequestID               int                      `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel,omitempty"`
	ChatIsForum             bool                     `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool                     `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool                     `json:"chat_is_created,omitempty"`
	BotIsMember             bool                     `json:"bot_is_member,omitempty"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	InputFieldPlaceholder string             `json:"input_field_placeholder,omitempty"`
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          bool               `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool               `json:"one_time_keyboard,omitempty"`
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
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	Text                         string                       `json:"text"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	URL                          string                       `json:"url,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
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
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	ForceReply            bool   `json:"force_reply"`
	Selective             bool   `json:"selective"`
}

// ImplementsReplyMarkup is a dummy method which exists to implement the interface ReplyMarkup.
func (f ForceReply) ImplementsReplyMarkup() {}

// UpdateOptions contains the optional parameters used by the GetUpdates method.
type UpdateOptions struct {
	AllowedUpdates []UpdateType `query:"allowed_updates"`
	Offset         int          `query:"offset"`
	Limit          int          `query:"limit"`
	Timeout        int          `query:"timeout"`
}

// WebhookOptions contains the optional parameters used by the SetWebhook method.
type WebhookOptions struct {
	IPAddress      string `query:"ip_address"`
	SecretToken    string `query:"secret_token"`
	Certificate    InputFile
	AllowedUpdates []UpdateType `query:"allowed_updates"`
	MaxConnections int          `query:"max_connections"`
}

// BaseOptions contains the optional parameters used frequently in some Telegram API methods.
type BaseOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	MessageThreadID     int             `query:"message_thread_id"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// MessageOptions contains the optional parameters used by some Telegram API methods.
type MessageOptions struct {
	ReplyMarkup         ReplyMarkup        `query:"reply_markup"`
	ReplyParameters     ReplyParameters    `query:"reply_parameters"`
	ParseMode           ParseMode          `query:"parse_mode"`
	LinkPreviewOptions  LinkPreviewOptions `query:"link_preview_options"`
	Entities            []MessageEntity    `query:"entities"`
	MessageThreadID     int64              `query:"message_thread_id"`
	DisableNotification bool               `query:"disable_notification"`
	ProtectContent      bool               `query:"protect_content"`
}

// PinMessageOptions contains the optional parameters used by the PinChatMember method.
type PinMessageOptions struct {
	DisableNotification bool `query:"disable_notification"`
}

// ForwardOptions contains the optional parameters used by the ForwardMessage method.
type ForwardOptions struct {
	MessageThreadID     int  `query:"message_thread_id"`
	DisableNotification bool `query:"disable_notification"`
	ProtectContent      bool `query:"protect_content"`
}

// CopyOptions contains the optional parameters used by the CopyMessage method.
type CopyOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	ParseMode           ParseMode       `query:"parse_mode"`
	Caption             string          `query:"caption"`
	CaptionEntities     []MessageEntity `query:"caption_entities"`
	MessageThreadID     int             `query:"message_thread_id"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// CopyMessagesOptions contains the optional parameters used by the CopyMessages methods.
type CopyMessagesOptions struct {
	MessageThreadID     int  `query:"message_thread_id"`
	DisableNotification bool `query:"disable_notification"`
	ProtectContent      bool `query:"protect_content"`
	RemoveCaption       bool `query:"remove_caption"`
}

// StickerOptions contains the optional parameters used by the SendSticker method.
type StickerOptions struct {
	Emoji               string          `query:"emoji"`
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	MessageThreadID     int             `query:"message_thread_id"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// InputFile is a struct which contains data about a file to be sent.
type InputFile struct {
	id      string
	path    string
	url     string
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

// NewInputFileURL is a wrapper for InputFile which only fills the url field.
func NewInputFileURL(url string) InputFile {
	return InputFile{url: url}
}

// NewInputFileBytes is a wrapper for InputFile which only fills the path and content fields.
func NewInputFileBytes(fileName string, content []byte) InputFile {
	return InputFile{path: fileName, content: content}
}

// PhotoOptions contains the optional parameters used by the SendPhoto method.
type PhotoOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	ParseMode           ParseMode       `query:"parse_mode"`
	Caption             string          `query:"caption"`
	CaptionEntities     []MessageEntity `query:"caption_entities"`
	MessageThreadID     int             `query:"message_thread_id"`
	HasSpoiler          bool            `query:"has_spoiler"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// AudioOptions contains the optional parameters used by the SendAudio method.
type AudioOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	ParseMode           ParseMode       `query:"parse_mode"`
	Caption             string          `query:"caption"`
	Performer           string          `query:"performer"`
	Title               string          `query:"title"`
	Thumbnail           InputFile
	CaptionEntities     []MessageEntity `query:"caption_entities"`
	MessageThreadID     int             `query:"message_thread_id"`
	Duration            int             `query:"duration"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// DocumentOptions contains the optional parameters used by the SendDocument method.
type DocumentOptions struct {
	ReplyMarkup                 ReplyMarkup     `query:"reply_markup"`
	ReplyParameters             ReplyParameters `query:"reply_parameters"`
	ParseMode                   ParseMode       `query:"parse_mode"`
	Caption                     string          `query:"caption"`
	Thumbnail                   InputFile
	CaptionEntities             []MessageEntity `query:"caption_entities"`
	MessageThreadID             int             `query:"message_thread_id"`
	DisableNotification         bool            `query:"disable_notification"`
	ProtectContent              bool            `query:"protect_content"`
	DisableContentTypeDetection bool            `query:"disable_content_type_detection"`
}

// VideoOptions contains the optional parameters used by the SendVideo method.
type VideoOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	Caption             string          `query:"caption"`
	ParseMode           ParseMode       `query:"parse_mode"`
	Thumbnail           InputFile
	CaptionEntities     []MessageEntity `query:"caption_entities"`
	MessageThreadID     int             `query:"message_thread_id"`
	Duration            int             `query:"duration"`
	Width               int             `query:"width"`
	Height              int             `query:"height"`
	HasSpoiler          bool            `query:"has_spoiler"`
	SupportsStreaming   bool            `query:"supports_streaming"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// AnimationOptions contains the optional parameters used by the SendAnimation method.
type AnimationOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	ParseMode           ParseMode       `query:"parse_mode"`
	Caption             string          `query:"caption"`
	Thumbnail           InputFile
	CaptionEntities     []MessageEntity `query:"caption_entities"`
	MessageThreadID     int             `query:"message_thread_id"`
	Duration            int             `query:"duration"`
	Width               int             `query:"width"`
	Height              int             `query:"height"`
	HasSpoiler          bool            `query:"has_spoiler"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// VoiceOptions contains the optional parameters used by the SendVoice method.
type VoiceOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	ParseMode           ParseMode       `query:"parse_mode"`
	Caption             string          `query:"caption"`
	CaptionEntities     []MessageEntity `query:"caption_entities"`
	MessageThreadID     int             `query:"message_thread_id"`
	Duration            int             `query:"duration"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// VideoNoteOptions contains the optional parameters used by the SendVideoNote method.
type VideoNoteOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	Thumbnail           InputFile
	MessageThreadID     int  `query:"message_thread_id"`
	Duration            int  `query:"duration"`
	Length              int  `query:"length"`
	DisableNotification bool `query:"disable_notification"`
	ProtectContent      bool `query:"protect_content"`
}

// MediaGroupOptions contains the optional parameters used by the SendMediaGroup method.
type MediaGroupOptions struct {
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	MessageThreadID     int             `query:"message_thread_id"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// LocationOptions contains the optional parameters used by the SendLocation method.
type LocationOptions struct {
	ReplyMarkup          ReplyMarkup     `query:"reply_markup"`
	ReplyParameters      ReplyParameters `query:"reply_parameters"`
	HorizontalAccuracy   float64         `query:"horizontal_accuracy"`
	MessageThreadID      int             `query:"message_thread_id"`
	LivePeriod           int             `query:"live_period"`
	ProximityAlertRadius int             `query:"proximity_alert_radius"`
	Heading              int             `query:"heading"`
	DisableNotification  bool            `query:"disable_notification"`
	ProtectContent       bool            `query:"protect_content"`
}

// EditLocationOptions contains the optional parameters used by the EditMessageLiveLocation method.
type EditLocationOptions struct {
	ReplyMarkup          InlineKeyboardMarkup `query:"reply_markup"`
	HorizontalAccuracy   float64              `query:"horizontal_accuracy"`
	Heading              int                  `query:"heading"`
	ProximityAlertRadius int                  `query:"proximity_alert_radius"`
}

// VenueOptions contains the optional parameters used by the SendVenue method.
type VenueOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	FoursquareID        string          `query:"foursquare_id"`
	FoursquareType      string          `query:"foursquare_type"`
	GooglePlaceType     string          `query:"google_place_type"`
	GooglePlaceID       string          `query:"google_place_id"`
	MessageThreadID     int             `query:"message_thread_id"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// ContactOptions contains the optional parameters used by the SendContact method.
type ContactOptions struct {
	ReplyMarkup         ReplyMarkup     `query:"reply_markup"`
	ReplyParameters     ReplyParameters `query:"reply_parameters"`
	VCard               string          `query:"vcard"`
	LastName            string          `query:"last_name"`
	MessageThreadID     int             `query:"message_thread_id"`
	DisableNotification bool            `query:"disable_notification"`
	ProtectContent      bool            `query:"protect_content"`
}

// PollOptions contains the optional parameters used by the SendPoll method.
type PollOptions struct {
	ReplyMarkup           ReplyMarkup     `query:"reply_markup"`
	ReplyParameters       ReplyParameters `query:"reply_parameters"`
	Explanation           string          `query:"explanation"`
	ExplanationParseMode  ParseMode       `query:"explanation_parse_mode"`
	Type                  PollType        `query:"type"`
	ExplanationEntities   []MessageEntity `query:"explanation_entities"`
	MessageThreadID       int             `query:"message_thread_id"`
	CorrectOptionID       int             `query:"correct_option_id"`
	CloseDate             int             `query:"close_date"`
	OpenPeriod            int             `query:"open_period"`
	IsClosed              bool            `query:"is_closed"`
	DisableNotification   bool            `query:"disable_notification"`
	ProtectContent        bool            `query:"protect_content"`
	AllowsMultipleAnswers bool            `query:"allows_multiple_answers"`
	IsAnonymous           bool            `query:"is_anonymous"`
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
	UseIndependentChatPermissions bool `query:"use_independent_chat_permissions"`
	UntilDate                     int  `query:"until_date"`
}

// PromoteOptions contains the optional parameters used by the PromoteChatMember method.
type PromoteOptions struct {
	IsAnonymous         bool `query:"is_anonymous,omitempty"`
	CanManageChat       bool `query:"can_manage_chat,omitempty"`
	CanPostMessages     bool `query:"can_post_messages,omitempty"`
	CanEditMessages     bool `query:"can_edit_messages,omitempty"`
	CanDeleteMessages   bool `query:"can_delete_messages,omitempty"`
	CanManageVideoChats bool `query:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool `query:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool `query:"can_promote_members,omitempty"`
	CanChangeInfo       bool `query:"can_change_info,omitempty"`
	CanInviteUsers      bool `query:"can_invite_users,omitempty"`
	CanPinMessages      bool `query:"can_pin_messages,omitempty"`
	CanPostStories      bool `json:"can_post_stories,omitempty"`
	CanEditStories      bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`
	CanManageTopics     bool `query:"can_manage_topics,omitempty"`
}

// UserProfileOptions contains the optional parameters used by the GetUserProfilePhotos method.
type UserProfileOptions struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
}

// ChatPermissionsOptions contains the optional parameters used by the SetChatPermissions method.
type ChatPermissionsOptions struct {
	UseIndependentChatPermissions bool `query:"use_independent_chat_permissions"`
}

// InviteLinkOptions contains the optional parameters used by the CreateChatInviteLink and EditChatInviteLink methods.
type InviteLinkOptions struct {
	Name               string `query:"name"`
	ExpireDate         int64  `query:"expire_date"`
	MemberLimit        int    `query:"member_limit"`
	CreatesJoinRequest bool   `query:"creates_join_request"`
}

// CallbackQueryOptions contains the optional parameters used by the AnswerCallbackQuery method.
type CallbackQueryOptions struct {
	Text      string `query:"text"`
	URL       string `query:"url"`
	CacheTime int    `query:"cache_time"`
	ShowAlert bool   `query:"show_alert"`
}

// MessageIDOptions is a struct which contains data about a message to edit.
type MessageIDOptions struct {
	inlineMessageID string `query:"inline_message_id"`
	chatID          int64  `query:"chat_id"`
	messageID       int    `query:"message_id"`
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
	ParseMode          ParseMode            `query:"parse_mode"`
	Entities           []MessageEntity      `query:"entities"`
	ReplyMarkup        InlineKeyboardMarkup `query:"reply_markup"`
	LinkPreviewOptions LinkPreviewOptions   `query:"link_preview_options"`
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
	LanguageCode string          `query:"language_code"`
	Scope        BotCommandScope `query:"scope"`
}

// InvoiceOptions contains the optional parameters used by the SendInvoice API method.
type InvoiceOptions struct {
	StartParameter            string               `query:"start_parameter"`
	ProviderData              string               `query:"provider_data"`
	PhotoURL                  string               `query:"photo_url"`
	ReplyMarkup               InlineKeyboardMarkup `query:"reply_markup"`
	ReplyParameters           ReplyParameters      `query:"reply_parameters"`
	SuggestedTipAmount        []int                `query:"suggested_tip_amounts"`
	MessageThreadID           int                  `query:"message_thread_id"`
	PhotoSize                 int                  `query:"photo_size"`
	PhotoWidth                int                  `query:"photo_width"`
	PhotoHeight               int                  `query:"photo_height"`
	MaxTipAmount              int                  `query:"max_tip_amount"`
	NeedPhoneNumber           bool                 `query:"need_phone_number"`
	NeepShippingAddress       bool                 `query:"need_shipping_address"`
	SendPhoneNumberToProvider bool                 `query:"send_phone_number_to_provider"`
	SendEmailToProvider       bool                 `query:"send_email_to_provider"`
	IsFlexible                bool                 `query:"is_flexible"`
	DisableNotification       bool                 `query:"disable_notification"`
	ProtectContent            bool                 `query:"protect_content"`
	NeedName                  bool                 `query:"need_name"`
	NeedEmail                 bool                 `query:"need_email"`
}

// CreateInvoiceLinkOptions contains the optional parameters used by the CreateInvoiceLink API method.
type CreateInvoiceLinkOptions struct {
	ProviderData              string `query:"provider_data"`
	PhotoURL                  string `query:"photo_url"`
	SuggestedTipAmounts       []int  `query:"suggested_tip_amounts"`
	PhotoSize                 int    `query:"photo_size"`
	PhotoWidth                int    `query:"photo_width"`
	PhotoHeight               int    `query:"photo_height"`
	MaxTipAmount              int    `query:"max_tip_amount"`
	NeedPhoneNumber           bool   `query:"need_phone_number"`
	NeepShippingAddress       bool   `query:"need_shipping_address"`
	SendPhoneNumberToProvider bool   `query:"send_phone_number_to_provider"`
	SendEmailToProvider       bool   `query:"send_email_to_provider"`
	IsFlexible                bool   `query:"is_flexible"`
	NeedName                  bool   `query:"need_name"`
	NeedEmail                 bool   `query:"need_email"`
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	ID     string         `query:"id"`
	Title  string         `query:"title"`
	Prices []LabeledPrice `query:"prices"`
}

// ShippingQueryOptions contains the optional parameters used by the AnswerShippingQuery API method.
type ShippingQueryOptions struct {
	ErrorMessage    string           `query:"error_message"`
	ShippingOptions []ShippingOption `query:"shipping_options"`
}

// PreCheckoutOptions contains the optional parameters used by the AnswerPreCheckoutQuery API method.
type PreCheckoutOptions struct {
	ErrorMessage string `query:"error_message"`
}

// CreateTopicOptions contains the optional parameters used by the CreateForumTopic API method.
type CreateTopicOptions struct {
	IconCustomEmojiID string    `query:"icon_custom_emoji_id"`
	IconColor         IconColor `query:"icon_color"`
}

// EditTopicOptions contains the optional parameters used by the EditForumTopic API method.
type EditTopicOptions struct {
	Name              string `query:"name"`
	IconCustomEmojiID string `query:"icon_custom_emoji_id"`
}

// ChatActionOptions contains the optional parameters used by the SendChatAction API method.
type ChatActionOptions struct {
	MessageThreadID int `query:"message_thread_id"`
}

// MessageReactionOptions contains the optional parameters used by the SetMessageReaction API method.
type MessageReactionOptions struct {
	Reaction []ReactionType `query:"reaction"`
	IsBig    bool           `query:"is_big"`
}
