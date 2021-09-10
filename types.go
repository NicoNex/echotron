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

import "encoding/json"

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

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	URL                  string        `json:"url"`
	HasCustomCertificate bool          `json:"has_custom_certificate"`
	PendingUpdateCount   int           `json:"pending_update_count"`
	IPAddress            string        `json:"ip_address,omitempty"`
	LastErrorDate        int64         `json:"last_error_date,omitempty"`
	LastErrorMessage     string        `json:"last_error_message,omitempty"`
	MaxConnections       int           `json:"max_connections,omitempty"`
	AllowedUpdates       []*UpdateType `json:"allowed_updates,omitempty"`
}

// APIResponseBase is a base type that represents the incoming response from Telegram servers.
// Used by APIResponse* to slim down the implementation.
type APIResponseBase struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

// APIResponseUpdate represents the incoming response from Telegram servers.
// Used by all methods that return an array of Update objects on success.
type APIResponseUpdate struct {
	Result []*Update `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseUser represents the incoming response from Telegram servers.
// Used by all methods that return a User object on success.
type APIResponseUser struct {
	Result *User `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseMessage represents the incoming response from Telegram servers.
// Used by all methods that return a Message object on success.
type APIResponseMessage struct {
	Result *Message `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseMessageID represents the incoming response from Telegram servers.
// Used by all methods that return a MessageID object on success.
type APIResponseMessageID struct {
	Result *MessageID `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseCommands represents the incoming response from Telegram servers.
// Used by all methods that return an array of BotCommand objects on success.
type APIResponseCommands struct {
	Result []*BotCommand `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseBool represents the incoming response from Telegram servers.
// Used by all methods that return True on success.
type APIResponseBool struct {
	Result bool `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseString represents the incoming response from Telegram servers.
// Used by all methods that return a string on success.
type APIResponseString struct {
	Result string `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseChat represents the incoming response from Telegram servers.
// Used by all methods that return a Chat object on success.
type APIResponseChat struct {
	Result *Chat `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseInviteLink represents the incoming response from Telegram servers.
// Used by all methods that return a ChatInviteLink object on success.
type APIResponseInviteLink struct {
	Result *ChatInviteLink `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseStickerSet represents the incoming response from Telegram servers.
// Used by all methods that return a StickerSet object on success.
type APIResponseStickerSet struct {
	Result *StickerSet `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseUserProfile represents the incoming response from Telegram servers.
// Used by all methods that return a UserProfilePhotos object on success.
type APIResponseUserProfile struct {
	Result *UserProfilePhotos `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseFile represents the incoming response from Telegram servers.
// Used by all methods that return a File object on success.
type APIResponseFile struct {
	Result *File `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseAdministrators represents the incoming response from Telegram servers.
// Used by all methods that return an array of ChatMember objects on success.
type APIResponseAdministrators struct {
	Result []*ChatMember `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseChatMember represents the incoming response from Telegram servers.
// Used by all methods that return a ChatMember object on success.
type APIResponseChatMember struct {
	Result *ChatMember `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseInteger represents the incoming response from Telegram servers.
// Used by all methods that return an integer on success.
type APIResponseInteger struct {
	Result int `json:"result,omitempty"`
	APIResponseBase
}

// APIResponsePoll represents the incoming response from Telegram servers.
// Used by all methods that return a Poll object on success.
type APIResponsePoll struct {
	Result *Poll `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseGameHighScore represents the incoming response from Telegram servers.
// Used by all methods that return an array of GameHighScore objects on success.
type APIResponseGameHighScore struct {
	Result []*GameHighScore `json:"result,omitempty"`
	APIResponseBase
}

// APIResponseWebhook represents the incoming response from Telegram servers.
// Used by all methods that return a WebhookInfo object on success.
type APIResponseWebhook struct {
	Result *WebhookInfo `json:"result,omitempty"`
	APIResponseBase
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

// Chat represents a chat.
type Chat struct {
	ID                    int64            `json:"id"`
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
	LinkedChatID          int64            `json:"linked_chat_id,omitempty"`
	Location              *ChatLocation    `json:"location,omitempty"`
}

// Message represents a message.
type Message struct {
	ID                            int                            `json:"message_id"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Date                          int                            `json:"date"`
	Chat                          *Chat                          `json:"chat"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID          int                            `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                         `json:"forward_signature,omitempty"`
	ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                            `json:"forward_date,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int                            `json:"edit_date,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []*MessageEntity               `json:"entities,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int                            `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int                            `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	VoiceChatStarted              *VoiceChatStarted              `json:"voice_chat_started,omitempty"`
	VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

// MessageID represents a unique message identifier.
type MessageID struct {
	MessageID int `json:"message_id"`
}

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
	User   *User  `json:"user,omitempty"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	FileUID  string `json:"file_unique_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileID   string     `json:"file_id"`
	FileUID  string     `json:"file_unique_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string     `json:"file_id"`
	FileUID   string     `json:"file_unique_id"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer,omitempty"`
	Title     string     `json:"title,omitempty"`
	FileName  string     `json:"file_name,omitempty"`
	MimeType  string     `json:"mime_type,omitempty"`
	FileSize  int        `json:"file_size,omitempty"`
	Thumb     *PhotoSize `json:"thumb,omitempty"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string     `json:"file_id"`
	FileUID  string     `json:"file_unique_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`
	FileUID  string     `json:"file_unique_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileID   string     `json:"file_id"`
	FileUID  string     `json:"file_unique_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	FileID   string `json:"file_id"`
	FileUID  string `json:"file_unique_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

// Venue represents a venue.
type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePlaceID   string    `json:"google_place_id,omitempty"`
	GooglePlaceType string    `json:"google_place_type,omitempty"`
}

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// VoiceChatScheduled represents a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VoiceChatStarted represents a service message about a voice chat started in the chat.
type VoiceChatStarted struct{}

// VoiceChatEnded represents a service message about a voice chat ended in the chat.
type VoiceChatEnded struct {
	Duration int `json:"duration"`
}

// VoiceChatParticipantsInvited represents a service message about new members invited to a voice chat.
type VoiceChatParticipantsInvited struct {
	Users []*User `json:"users,omitempty"`
}

// UserProfilePhotos represents a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File represents a file ready to be downloaded.
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
// If the button that originated the query was attached to a message sent by the bot,
// the field message will be present. If the button was attached to a message sent via the bot (in inline mode),
// the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance,omitempty"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID  string `json:"small_file_id"`
	SmallFileUID string `json:"small_file_unique_id"`
	BigFileID    string `json:"big_file_id"`
	BigFileUID   string `json:"big_file_unique_id"`
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

// BotCommand represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// ResponseParameters contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int `json:"retry_after,omitempty"`
}

// InputMediaType is a custom type for the various InputMediaType*'s Type field.
type InputMediaType string

// These are all the possible types for the various InputMediaType*'s Type field.
const (
	MediaTypePhoto     InputMediaType = "photo"
	MediaTypeVideo                    = "video"
	MediaTypeAnimation                = "animation"
	MediaTypeAudio                    = "audio"
	MediaTypeDocument                 = "document"
)

// InputMedia is an interface for the various media types.
type InputMedia interface {
	GetMedia() InputFile
}

// InputMediaGroupable is an interface for the various groupable media types.
type InputMediaGroupable interface {
	GetMedia() InputFile
	ImplementsInputMediaGroupable()
}

// inputMedia is a generic struct for all the various structs under the InputMedia interface.
type inputMedia struct {
	Media string
	InputMedia
}

// MarshalJSON is a custom marshaler for the inputMedia struct.
func (i inputMedia) MarshalJSON() (cnt []byte, err error) {
	var generic map[string]interface{}

	im, err := json.Marshal(i.InputMedia)
	if err != nil {
		return
	}

	err = json.Unmarshal(im, &generic)
	if err != nil {
		return
	}

	generic["media"] = i.Media

	return json.Marshal(generic)
}

// InputMediaPhoto represents a photo to be sent.
// Type MUST BE "photo".
type InputMediaPhoto struct {
	Type            InputMediaType   `json:"type"`
	Media           InputFile        `json:"-"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
}

// ImplementsInputMedia is a dummy method which exists to implement the interface InputMedia.
func (i InputMediaPhoto) GetMedia() InputFile { return i.Media }

// ImplementsInputMediaGroupable is a dummy method which exists to implement the interface InputMediaGroupable.
func (i InputMediaPhoto) ImplementsInputMediaGroupable() { }


// InputMediaVideo represents a video to be sent.
// Type MUST BE "video".
type InputMediaVideo struct {
	Type              InputMediaType   `json:"type"`
	Media             InputFile        `json:"-"`
	Thumb             *InputFile       `json:"thumb,omitempty"`
	Caption           string           `json:"caption,omitempty"`
	ParseMode         ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities   []*MessageEntity `json:"caption_entities,omitempty"`
	Width             int              `json:"width,omitempty"`
	Height            int              `json:"height,omitempty"`
	Duration          int              `json:"duration,omitempty"`
	SupportsStreaming bool             `json:"supports_streaming,omitempty"`
}

// ImplementsInputMedia is a dummy method which exists to implement the interface InputMedia.
func (i InputMediaVideo) GetMedia() InputFile { return i.Media }

// ImplementsInputMediaGroupable is a dummy method which exists to implement the interface InputMediaGroupable.
func (i InputMediaVideo) ImplementsInputMediaGroupable() { }

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
// Type MUST BE "animation".
type InputMediaAnimation struct {
	Type            InputMediaType   `json:"type"`
	Media           InputFile        `json:"-"`
	Thumb           *InputFile       `json:"thumb,omitempty"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	Duration        int              `json:"duration,omitempty"`
}

// ImplementsInputMedia is a dummy method which exists to implement the interface InputMedia.
func (i InputMediaAnimation) GetMedia() InputFile { return i.Media }

// InputMediaAudio represents an audio file to be treated as music to be sent.
// Type MUST BE "audio".
type InputMediaAudio struct {
	Type            InputMediaType   `json:"type"`
	Media           InputFile        `json:"-"`
	Thumb           *InputFile       `json:"thumb,omitempty"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Duration        int              `json:"duration,omitempty"`
	Performer       string           `json:"performer,omitempty"`
	Title           string           `json:"title,omitempty"`
}

// ImplementsInputMedia is a dummy method which exists to implement the interface InputMedia.
func (i InputMediaAudio) GetMedia() InputFile { return i.Media }

// ImplementsInputMediaGroupable is a dummy method which exists to implement the interface InputMediaGroupable.
func (i InputMediaAudio) ImplementsInputMediaGroupable() { }

// InputMediaDocument represents a general file to be sent.
// Type MUST BE "document".
type InputMediaDocument struct {
	Type                        InputMediaType   `json:"type"`
	Media                       InputFile        `json:"-"`
	Thumb                       *InputFile       `json:"thumb,omitempty"`
	Caption                     string           `json:"caption,omitempty"`
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
}

// ImplementsInputMedia is a dummy method which exists to implement the interface InputMedia.
func (i InputMediaDocument) GetMedia() InputFile { return i.Media }

// ImplementsInputMediaGroupable is a dummy method which exists to implement the interface InputMediaGroupable.
func (i InputMediaDocument) ImplementsInputMediaGroupable() { }

// BotCommandScopeType is a custom type for the various bot command scope types.
type BotCommandScopeType string

// These are all the various bot command scope types.
const (
	BCSTDefault               BotCommandScopeType = "default"
	BCSTAllPrivateChats                           = "all_private_chats"
	BCSTAllGroupChats                             = "all_group_chats"
	BCSTAllChatAdministrators                     = "all_chat_administrators"
	BCSTChat                                      = "chat"
	BCSTChatAdministrators                        = "chat_administrators"
	BCSTChatMember                                = "chat_member"
)

// BotCommandScope is an optional parameter used in the SetMyCommands, DeleteMyCommands and GetMyCommands methods.
type BotCommandScope struct {
	Type   BotCommandScopeType `query:"type"`
	ChatID int64               `query:"chat_id"`
	UserID int64               `query:"user_id"`
}

// PermissionOptions is a custom type used to allow proper serialization of ChatPermissions-type parameters in some methods.
type PermissionOptions struct {
	Permissions ChatPermissions `json:"permissions"`
}
