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

import "encoding/json"

// Update represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ID                      int                          `json:"update_id"`
}

// ChatID returns the ID of the chat the update is coming from.
func (u Update) ChatID() int64 {
	switch {
	case u.ChatJoinRequest != nil:
		return u.ChatJoinRequest.Chat.ID
	case u.ChatBoost != nil:
		return u.ChatBoost.Chat.ID
	case u.RemovedChatBoost != nil:
		return u.RemovedChatBoost.Chat.ID
	case u.Message != nil:
		return u.Message.Chat.ID
	case u.EditedMessage != nil:
		return u.EditedMessage.Chat.ID
	case u.ChannelPost != nil:
		return u.ChannelPost.Chat.ID
	case u.EditedChannelPost != nil:
		return u.EditedChannelPost.Chat.ID
	case u.BusinessConnection != nil:
		return u.BusinessConnection.User.ID
	case u.BusinessMessage != nil:
		return u.BusinessMessage.Chat.ID
	case u.EditedBusinessMessage != nil:
		return u.EditedBusinessMessage.Chat.ID
	case u.DeletedBusinessMessages != nil:
		return u.DeletedBusinessMessages.Chat.ID
	case u.MessageReaction != nil:
		return u.MessageReaction.Chat.ID
	case u.MessageReactionCount != nil:
		return u.MessageReactionCount.Chat.ID
	case u.InlineQuery != nil:
		return u.InlineQuery.From.ID
	case u.ChosenInlineResult != nil:
		return u.ChosenInlineResult.From.ID
	case u.CallbackQuery != nil:
		return u.CallbackQuery.Message.Chat.ID
	case u.ShippingQuery != nil:
		return u.ShippingQuery.From.ID
	case u.PreCheckoutQuery != nil:
		return u.PreCheckoutQuery.From.ID
	case u.PollAnswer != nil:
		return u.PollAnswer.User.ID
	case u.MyChatMember != nil:
		return u.MyChatMember.Chat.ID
	case u.ChatMember != nil:
		return u.ChatMember.Chat.ID
	default:
		return 0
	}
}

// WebhookInfo contains information about the current status of a webhook.
type WebhookInfo struct {
	URL                          string        `json:"url"`
	IPAddress                    string        `json:"ip_address,omitempty"`
	LastErrorMessage             string        `json:"last_error_message,omitempty"`
	AllowedUpdates               []*UpdateType `json:"allowed_updates,omitempty"`
	MaxConnections               int           `json:"max_connections,omitempty"`
	LastErrorDate                int64         `json:"last_error_date,omitempty"`
	LastSynchronizationErrorDate int64         `json:"last_synchronization_error_date,omitempty"`
	PendingUpdateCount           int           `json:"pending_update_count"`
	HasCustomCertificate         bool          `json:"has_custom_certificate"`
}

// APIResponse is implemented by all the APIResponse* types.
type APIResponse interface {
	// Base returns the object of type APIResponseBase contained in each implemented type.
	Base() APIResponseBase
}

// APIResponseBase is a base type that represents the incoming response from Telegram servers.
// Used by APIResponse* to slim down the implementation.
type APIResponseBase struct {
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Ok          bool   `json:"ok"`
}

// Base returns the APIResponseBase itself.
func (a APIResponseBase) Base() APIResponseBase {
	return a
}

// APIResponseUpdate represents the incoming response from Telegram servers.
// Used by all methods that return an array of Update objects on success.
type APIResponseUpdate struct {
	Result []*Update `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseUpdate) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseUser represents the incoming response from Telegram servers.
// Used by all methods that return a User object on success.
type APIResponseUser struct {
	Result *User `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseUser) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseMessage represents the incoming response from Telegram servers.
// Used by all methods that return a Message object on success.
type APIResponseMessage struct {
	Result *Message `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseMessage) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseMessageArray represents the incoming response from Telegram servers.
// Used by all methods that return an array of Message objects on success.
type APIResponseMessageArray struct {
	Result []*Message `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseMessageArray) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseMessageID represents the incoming response from Telegram servers.
// Used by all methods that return a MessageID object on success.
type APIResponseMessageID struct {
	Result *MessageID `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseMessageID) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseMessageIDs represents the incoming response from Telegram servers.
// Used by all methods that return a MessageID object on success.
type APIResponseMessageIDs struct {
	Result []*MessageID `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseMessageIDs) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseCommands represents the incoming response from Telegram servers.
// Used by all methods that return an array of BotCommand objects on success.
type APIResponseCommands struct {
	Result []*BotCommand `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseCommands) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseBool represents the incoming response from Telegram servers.
// Used by all methods that return True on success.
type APIResponseBool struct {
	APIResponseBase
	Result bool `json:"result,omitempty"`
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBool) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseString represents the incoming response from Telegram servers.
// Used by all methods that return a string on success.
type APIResponseString struct {
	Result string `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseString) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseChat represents the incoming response from Telegram servers.
// Used by all methods that return a ChatFullInfo object on success.
type APIResponseChat struct {
	Result *ChatFullInfo `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseChat) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseInviteLink represents the incoming response from Telegram servers.
// Used by all methods that return a ChatInviteLink object on success.
type APIResponseInviteLink struct {
	Result *ChatInviteLink `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseInviteLink) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseStickers represents the incoming response from Telegram servers.
// Used by all methods that return an array of Stickers on success.
type APIResponseStickers struct {
	Result []*Sticker `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseStickers) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseStickerSet represents the incoming response from Telegram servers.
// Used by all methods that return a StickerSet object on success.
type APIResponseStickerSet struct {
	Result *StickerSet `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseStickerSet) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseUserProfile represents the incoming response from Telegram servers.
// Used by all methods that return a UserProfilePhotos object on success.
type APIResponseUserProfile struct {
	Result *UserProfilePhotos `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseUserProfile) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseFile represents the incoming response from Telegram servers.
// Used by all methods that return a File object on success.
type APIResponseFile struct {
	Result *File `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseFile) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseAdministrators represents the incoming response from Telegram servers.
// Used by all methods that return an array of ChatMember objects on success.
type APIResponseAdministrators struct {
	Result []*ChatMember `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseAdministrators) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseChatMember represents the incoming response from Telegram servers.
// Used by all methods that return a ChatMember object on success.
type APIResponseChatMember struct {
	Result *ChatMember `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseChatMember) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseInteger represents the incoming response from Telegram servers.
// Used by all methods that return an integer on success.
type APIResponseInteger struct {
	APIResponseBase
	Result int `json:"result,omitempty"`
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseInteger) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponsePoll represents the incoming response from Telegram servers.
// Used by all methods that return a Poll object on success.
type APIResponsePoll struct {
	Result *Poll `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponsePoll) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseGameHighScore represents the incoming response from Telegram servers.
// Used by all methods that return an array of GameHighScore objects on success.
type APIResponseGameHighScore struct {
	Result []*GameHighScore `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseGameHighScore) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseWebhook represents the incoming response from Telegram servers.
// Used by all methods that return a WebhookInfo object on success.
type APIResponseWebhook struct {
	Result *WebhookInfo `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseWebhook) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseSentWebAppMessage represents the incoming response from Telegram servers.
// Used by all methods that return a SentWebAppMessage object on success.
type APIResponseSentWebAppMessage struct {
	Result *SentWebAppMessage `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseSentWebAppMessage) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseMenuButton represents the incoming response from Telegram servers.
// Used by all methods that return a MenuButton object on success.
type APIResponseMenuButton struct {
	Result *MenuButton `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseMenuButton) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseChatAdministratorRights represents the incoming response from Telegram servers.
// Used by all methods that return a ChatAdministratorRights object on success.
type APIResponseChatAdministratorRights struct {
	Result *ChatAdministratorRights `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseChatAdministratorRights) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseForumTopic represents the incoming response from Telegram servers.
// Used by all methods that return a ForumTopic object on success.
type APIResponseForumTopic struct {
	Result *ForumTopic `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseForumTopic) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseBotDescription represents the incoming response from Telegram servers.
// Used by all methods that return a BotDescription object on success.
type APIResponseBotDescription struct {
	Result *BotDescription `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBotDescription) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseBotShortDescription represents the incoming response from Telegram servers.
// Used by all methods that return a BotShortDescription object on success.
type APIResponseBotShortDescription struct {
	Result *BotShortDescription `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBotShortDescription) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseBotName represents the incoming response from Telegram servers.
// Used by all methods that return a BotName object on success.
type APIResponseBotName struct {
	Result *BotName `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBotName) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseUserChatBoosts represents the incoming response from Telegram servers.
// Used by all methods that return a UserChatBoosts object on success.
type APIResponseUserChatBoosts struct {
	Result *UserChatBoosts `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseUserChatBoosts) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseBusinessConnection represents the incoming response from Telegram servers.
// Used by all methods that return a BusinessConnection object on success.
type APIResponseBusinessConnection struct {
	Result *BusinessConnection `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBusinessConnection) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseStarTransactions represents the incoming response from Telegram servers.
// Used by all methods that return a StarTransactions object on success.
type APIResponseStarTransactions struct {
	Result *StarTransactions `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseStarTransactions) Base() APIResponseBase {
	return a.APIResponseBase
}

// User represents a Telegram user or bot.
type User struct {
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
}

// Chat represents a chat.
type Chat struct {
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	ID        int64  `json:"id"`
	IsForum   bool   `json:"is_forum,omitempty"`
}

// ChatFullInfo contains full information about a chat.
type ChatFullInfo struct {
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    *[]string             `json:"active_usernames,omitempty"`
	AvailableReactions                 *[]ReactionType       `json:"available_reactions,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id,omitempty"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	Username                           string                `json:"username,omitempty"`
	Title                              string                `json:"title,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	Description                        string                `json:"description,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	Type                               string                `json:"type"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id,omitempty"`
	MaxReactionCount                   int                   `json:"max_reaction_count,omitempty"`
	ProfileAccentColorID               int                   `json:"profile_accent_color_id,omitempty"`
	EmojiStatusExpirationDate          int                   `json:"emoji_status_expiration_date,omitempty"`
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`
	ID                                 int64                 `json:"id"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
}

// Message represents a message.
type Message struct {
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	BusinessConnectionID          string                         `json:"business_connection_id,omitempty"`
	EffectID                      string                         `json:"effect_id,omitempty"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Entities                      []*MessageEntity               `json:"entities,omitempty"`
	Chat                          Chat                           `json:"chat"`
	ID                            int                            `json:"message_id"`
	ThreadID                      int                            `json:"message_thread_id,omitempty"`
	MigrateFromChatID             int                            `json:"migrate_from_chat_id,omitempty"`
	Date                          int                            `json:"date"`
	MigrateToChatID               int                            `json:"migrate_to_chat_id,omitempty"`
	EditDate                      int                            `json:"edit_date,omitempty"`
	SenderBoostCount              int                            `json:"sender_boost_count,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`
}

// MessageID represents a unique message identifier.
type MessageID struct {
	MessageID int `json:"message_id"`
}

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	User          *User             `json:"user,omitempty"`
	Type          MessageEntityType `json:"type"`
	URL           string            `json:"url,omitempty"`
	Language      string            `json:"language,omitempty"`
	CustomEmojiID string            `json:"custom_emoji_id,omitempty"`
	Offset        int               `json:"offset"`
	Length        int               `json:"length"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Duration     int        `json:"duration"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	MimeType     string `json:"mime_type,omitempty"`
	Duration     int    `json:"duration"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int              `json:"voter_count"`
}

// InputPollOption contains information about one answer option in a poll to send.
type InputPollOption struct {
	Text          string           `json:"text"`
	TextParseMode ParseMode        `json:"text_parse_mode,omitempty"`
	TextEntities  []*MessageEntity `json:"text_entities,omitempty"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	Type                  string           `json:"type"`
	Question              string           `json:"question"`
	Explanation           string           `json:"explanation,omitempty"`
	ID                    string           `json:"id"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	QuestionEntities      []*MessageEntity `json:"question_entities,omitempty"`
	Options               []*PollOption    `json:"options"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	TotalVoterCount       int              `json:"total_voter_count"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
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

// VideoChatScheduled represents a service message about a voice chat scheduled in the chat.
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VideoChatStarted represents a service message about a voice chat started in the chat.
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a voice chat ended in the chat.
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members invited to a voice chat.
type VideoChatParticipantsInvited struct {
	Users []*User `json:"users,omitempty"`
}

// UserProfilePhotos represents a user's profile pictures.
type UserProfilePhotos struct {
	Photos     [][]PhotoSize `json:"photos"`
	TotalCount int           `json:"total_count"`
}

// File represents a file ready to be downloaded.
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FilePath     string `json:"file_path,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

// SwitchInlineQueryChosenChat represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
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
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	Creator                 *User  `json:"creator"`
	InviteLink              string `json:"invite_link"`
	Name                    string `json:"name,omitempty"`
	PendingJoinRequestCount int    `json:"pending_join_request_count,omitempty"`
	ExpireDate              int    `json:"expire_date,omitempty"`
	MemberLimit             int    `json:"member_limit,omitempty"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
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
	CanManageVideoChats   bool   `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendAudios         bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool   `json:"can_send_photos,omitempty"`
	CanSendVideos         bool   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	CanManageTopics       bool   `json:"can_manage_topics,omitempty"`
	CanPostStories        bool   `json:"can_post_stories,omitempty"`
	CanEditStories        bool   `json:"can_edit_stories,omitempty"`
	CanDeleteStories      bool   `json:"can_delete_stories,omitempty"`
	UntilDate             int    `json:"until_date,omitempty"`
}

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	Date                    int             `json:"date"`
	ViaChatFolderInviteLink bool            `json:"via_chat_folder_invite_link,omitempty"`
	ViaJoinRequest          bool            `json:"via_join_request,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

// Birthdate
type Birthdate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// BusinessIntro
type BusinessIntro struct {
	Sticker *Sticker `json:"sticker,omitempty"`
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
}

// BusinessLocation
type BusinessLocation struct {
	Location *Location `json:"location,omitempty"`
	Address  string    `json:"address"`
}

// BusinessOpeningHoursInterval
type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

// BusinessOpeningHours
type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
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

// InputMediaType is a custom type for the various InputMedia*'s Type field.
type InputMediaType string

// These are all the possible types for the various InputMedia*'s Type field.
const (
	MediaTypePhoto     InputMediaType = "photo"
	MediaTypeVideo                    = "video"
	MediaTypeAnimation                = "animation"
	MediaTypeAudio                    = "audio"
	MediaTypeDocument                 = "document"
)

// InputMedia is an interface for the various media types.
type InputMedia interface {
	media() InputFile
	thumbnail() InputFile
}

// GroupableInputMedia is an interface for the various groupable media types.
type GroupableInputMedia interface {
	InputMedia
	groupable()
}

// mediaEnvelope is a generic struct for all the various structs under the InputMedia interface.
type mediaEnvelope struct {
	InputMedia
	media     string
	thumbnail string
}

// MarshalJSON is a custom marshaler for the mediaEnvelope struct.
func (i mediaEnvelope) MarshalJSON() (cnt []byte, err error) {
	var tmp any

	switch o := i.InputMedia.(type) {
	case InputMediaPhoto:
		tmp = struct {
			Media string `json:"media"`
			InputMediaPhoto
		}{
			InputMediaPhoto: o,
			Media:           i.media,
		}

	case InputMediaVideo:
		tmp = struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaVideo
		}{
			InputMediaVideo: o,
			Media:           i.media,
			Thumbnail:       i.thumbnail,
		}

	case InputMediaAnimation:
		tmp = struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaAnimation
		}{
			InputMediaAnimation: o,
			Media:               i.media,
			Thumbnail:           i.thumbnail,
		}

	case InputMediaAudio:
		tmp = struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaAudio
		}{
			InputMediaAudio: o,
			Media:           i.media,
			Thumbnail:       i.thumbnail,
		}

	case InputMediaDocument:
		tmp = struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaDocument
		}{
			InputMediaDocument: o,
			Media:              i.media,
			Thumbnail:          i.thumbnail,
		}
	}

	return json.Marshal(tmp)
}

// InputMediaPhoto represents a photo to be sent.
// Type MUST BE "photo".
type InputMediaPhoto struct {
	Type                  InputMediaType   `json:"type"`
	Media                 InputFile        `json:"-"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaPhoto) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaPhoto) thumbnail() InputFile { return InputFile{} }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputMediaPhoto) groupable() {}

// InputMediaVideo represents a video to be sent.
// Type MUST BE "video".
type InputMediaVideo struct {
	Type                  InputMediaType   `json:"type"`
	Media                 InputFile        `json:"-"`
	Thumbnail             InputFile        `json:"-"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	Width                 int              `json:"width,omitempty"`
	Height                int              `json:"height,omitempty"`
	Duration              int              `json:"duration,omitempty"`
	SupportsStreaming     bool             `json:"supports_streaming,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaVideo) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaVideo) thumbnail() InputFile { return i.Thumbnail }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputMediaVideo) groupable() {}

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
// Type MUST BE "animation".
type InputMediaAnimation struct {
	Type                  InputMediaType   `json:"type"`
	Media                 InputFile        `json:"-"`
	Thumbnail             InputFile        `json:"-"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	Width                 int              `json:"width,omitempty"`
	Height                int              `json:"height,omitempty"`
	Duration              int              `json:"duration,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaAnimation) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaAnimation) thumbnail() InputFile { return i.Thumbnail }

// InputMediaAudio represents an audio file to be treated as music to be sent.
// Type MUST BE "audio".
type InputMediaAudio struct {
	Type            InputMediaType   `json:"type"`
	Performer       string           `json:"performer,omitempty"`
	Title           string           `json:"title,omitempty"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`
	Media           InputFile        `json:"-"`
	Thumbnail       InputFile        `json:"-"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Duration        int              `json:"duration,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaAudio) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaAudio) thumbnail() InputFile { return i.Thumbnail }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputMediaAudio) groupable() {}

// InputMediaDocument represents a general file to be sent.
// Type MUST BE "document".
type InputMediaDocument struct {
	Type                        InputMediaType   `json:"type"`
	Media                       InputFile        `json:"-"`
	Thumbnail                   InputFile        `json:"-"`
	Caption                     string           `json:"caption,omitempty"`
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities             []*MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaDocument) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaDocument) thumbnail() InputFile { return i.Thumbnail }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputMediaDocument) groupable() {}

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
	Type   BotCommandScopeType `json:"type"`
	ChatID int64               `json:"chat_id"`
	UserID int64               `json:"user_id"`
}

// BotDescription represents the bot's description.
type BotDescription struct {
	Description string `json:"description"`
}

// BotShortDescription represents the bot's short description.
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// BotName represents the bot's name.
type BotName struct {
	Name string `json:"name"`
}

// ChatJoinRequest represents a join request sent to a chat.
type ChatJoinRequest struct {
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	Bio        string          `json:"bio,omitempty"`
	Chat       Chat            `json:"chat"`
	From       User            `json:"user"`
	Date       int             `json:"date"`
	UserChatID int64           `json:"user_chat_id"`
}

// ChatBoostAdded represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

// BackgroundFill describes the way a background is filled based on the selected colors.
type BackgroundFill interface {
	ImplementsBackgroundFill()
}

// BackgroundFillSolid is a background filled using the selected color.
// Type MUST be "solid".
type BackgroundFillSolid struct {
	Type  string `json:"type"`
	Color int    `json:"color"`
}

func (b BackgroundFillSolid) ImplementsBackgroundFill() {}

// BackgroundFillGradient is a background with a gradient fill.
// Type MUST be "gradient".
type BackgroundFillGradient struct {
	Type          string `json:"type"`
	TopColor      int    `json:"top_color"`
	BottomColor   int    `json:"bottom_color"`
	RotationAngle int    `json:"rotation_angle"`
}

func (b BackgroundFillGradient) ImplementsBackgroundFill() {}

// BackgroundFillFreeformGradient is a background with a freeform gradient that rotates after every message in the chat.
// Type MUST be "freeform_gradient".
type BackgroundFillFreeformGradient struct {
	Type   string `json:"type"`
	Colors []int  `json:"colors"`
}

func (b BackgroundFillFreeformGradient) ImplementsBackgroundFill() {}

// BackgroundType describes the type of a background.
type BackgroundType interface {
	ImplementsBackgroundType()
}

// BackgroundTypeFill is a background which is automatically filled based on the selected colors.
// Type MUST be "fill".
type BackgroundTypeFill struct {
	Fill             BackgroundFill `json:"fill"`
	Type             string         `json:"type"`
	DarkThemeDimming int            `json:"dark_theme_dimming"`
}

func (b BackgroundTypeFill) ImplementsBackgroundType() {}

// BackgroundTypeWallpaper is a background which is a wallpaper in the JPEG format.
// Type MUST be "wallpaper".
type BackgroundTypeWallpaper struct {
	Type             string   `json:"type"`
	Document         Document `json:"document"`
	DarkThemeDimming int      `json:"dark_theme_dimming"`
	IsBlurred        bool     `json:"is_blurred,omitempty"`
	IsMoving         bool     `json:"is_moving,omitempty"`
}

func (b BackgroundTypeWallpaper) ImplementsBackgroundType() {}

// BackgroundTypePattern is a PNG or TGV (gzipped subset of SVG with MIME type “application/x-tgwallpattern”) pattern
// to be combined with the background fill chosen by the user.
// Type MUST be "pattern".
type BackgroundTypePattern struct {
	Fill       BackgroundFill `json:"fill"`
	Type       string         `json:"type"`
	Document   Document       `json:"document"`
	Intensity  int            `json:"intensity"`
	IsInverted bool           `json:"is_inverted,omitempty"`
	IsMoving   bool           `json:"is_moving,omitempty"`
}

func (b BackgroundTypePattern) ImplementsBackgroundType() {}

// BackgroundTypeChatTheme is taken directly from a built-in chat theme.
// Type MUST be "chat_theme".
type BackgroundTypeChatTheme struct {
	Type      string `json:"type"`
	ThemeName string `json:"theme_name"`
}

func (b BackgroundTypeChatTheme) ImplementsBackgroundType() {}

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	IconColor         int    `json:"icon_color"`
}

// ChatBackground represents a chat background.
type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

// ForumTopicClosed represents a service message about a forum topic closed in the chat.
type ForumTopicClosed struct{}

// ForumTopicEdited represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	Name              string `json:"name"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
}

// ForumTopicReopened represents a service message about a forum topic reopened in the chat.
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about General forum topic hidden in the chat.
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about General forum topic unhidden in the chat.
type GeneralForumTopicUnhidden struct{}

// WriteAccessAllowed represents a service message about a user allowing a bot added to the attachment menu to write messages.
type WriteAccessAllowed struct {
	WebAppName         string `json:"web_app_name,omitempty"`
	FromRequest        bool   `json:"from_request,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

// IconColor represents a forum topic icon in RGB format.
type IconColor int

// These are all the various icon colors.
const (
	LightBlue IconColor = 0x6FB9F0
	Yellow              = 0xFFD67E
	Purple              = 0xCB86DB
	Green               = 0x8EEE98
	Pink                = 0xFF93B2
	Red                 = 0xFB6F5F
)

// ForumTopic represents a forum topic.
type ForumTopic struct {
	Name              string    `json:"name"`
	IconCustomEmojiID string    `json:"icon_custom_emoji_id"`
	IconColor         IconColor `json:"icon_color"`
	MessageThreadID   int64     `json:"message_thread_id"`
}

// UserShared contains information about the user whose identifier was shared with the bot using a KeyboardButtonRequestUser button.
type UserShared struct {
	RequestID int   `json:"request_id"`
	UserID    int64 `json:"user_id"`
}

// ChatShared contains information about the chat whose identifier was shared with the bot using a KeyboardButtonRequestChat button.
type ChatShared struct {
	Photo     *[]PhotoSize `json:"photo,omitempty"`
	Title     string       `json:"title,omitempty"`
	Username  string       `json:"username,omitempty"`
	RequestID int          `json:"request_id"`
	ChatID    int64        `json:"chat_id"`
}

// Story represents a story.
type Story struct {
	Chat Chat  `json:"chat"`
	ID   int64 `json:"id"`
}

type ReactionType struct {
	Type        string `json:"type"`
	Emoji       string `json:"emoji"`
	CustomEmoji string `json:"custom_emoji"`
}

// ReactionCount represents a reaction added to a message along with the number of times it was added.
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

// MessageReactionUpdated represents a change of a reaction on a message performed by a user.
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	ActorChat   Chat           `json:"actor_chat,omitempty"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
	User        User           `json:"user,omitempty"`
	MessageID   int            `json:"message_id"`
	Date        int            `json:"date"`
}

// MessageReactionCountUpdated represents reaction changes on a message with anonymous reactions.
type MessageReactionCountUpdated struct {
	Reactions []ReactionCount `json:"reactions"`
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int             `json:"date"`
}

// TextQuote contains information about the quoted part of a message that is replied to by the given message.
type TextQuote struct {
	Entities *[]MessageEntity `json:"entities,omitempty"`
	Text     string           `json:"text"`
	Position int              `json:"position"`
	IsManual bool             `json:"is_manual,omitempty"`
}

// ExternalReplyInfo contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	Venue              Venue              `json:"venue,omitempty"`
	Chat               Chat               `json:"chat,omitempty"`
	Document           Document           `json:"document,omitempty"`
	Origin             MessageOrigin      `json:"origin"`
	Contact            Contact            `json:"contact,omitempty"`
	Invoice            Invoice            `json:"invoice,omitempty"`
	Dice               Dice               `json:"dice,omitempty"`
	LinkPreviewOptions LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Photo              []PhotoSize        `json:"photo,omitempty"`
	Audio              Audio              `json:"audio,omitempty"`
	Story              Story              `json:"story,omitempty"`
	Voice              Voice              `json:"voice,omitempty"`
	VideoNote          VideoNote          `json:"video_note,omitempty"`
	Game               Game               `json:"game,omitempty"`
	Video              Video              `json:"video,omitempty"`
	Animation          Animation          `json:"animation,omitempty"`
	Sticker            Sticker            `json:"sticker,omitempty"`
	Giveaway           Giveaway           `json:"giveaway,omitempty"`
	Poll               Poll               `json:"poll,omitempty"`
	GiveawayWinners    GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Location           Location           `json:"location,omitempty"`
	MessageID          int                `json:"message_id,omitempty"`
	HasMediaSpoiler    bool               `json:"has_media_spoiler,omitempty"`
}

// MessageOrigin describes the origin of a message.
type MessageOrigin struct {
	SenderChat      *Chat  `json:"sender_chat,omitempty"`
	SenderUser      *User  `json:"sender_user,omitempty"`
	Type            string `json:"type"`
	SenderUserName  string `json:"sender_user_name,omitempty"`
	AuthorSignature string `json:"author_signature,omitempty"`
	Date            int    `json:"date"`
}

// LinkPreviewOptions describes the options used for link preview generation.
type LinkPreviewOptions struct {
	URL              string `json:"url,omitempty"`
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

// ReplyParameters describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	Quote                    string          `json:"quote,omitempty"`
	QuoteParseMode           string          `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entitites,omitempty"`
	MessageID                int             `json:"message_id"`
	ChatID                   int64           `json:"chat_id,omitempty"`
	QuotePosition            int             `json:"quote_position,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
}

// SharedUser contains information about a user that was shared with the bot using a KeyboardButtonRequestUser button.
type SharedUser struct {
	Photo     *[]PhotoSize `json:"photo,omitempty"`
	FirstName string       `json:"firstname,omitempty"`
	LastName  string       `json:"lastname,omitempty"`
	Username  string       `json:"username,omitempty"`
	UserID    int64        `json:"user_id"`
}

// UsersShared contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	Users     []SharedUser `json:"users"`
	RequestID int          `json:"request_id"`
}

// ChatBoost contains information about a chat boost.
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	Source         ChatBoostSource `json:"source"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
}

// ChatBoostSourceType is a custom type for the various chat boost sources.
type ChatBoostSourceType string

// These are all the possible chat boost types.
const (
	ChatBoostSourcePremium  ChatBoostSourceType = "premium"
	ChatBoostSourceGiftCode                     = "gift_code"
	ChatBoostSourceGiveaway                     = "giveaway"
)

// ChatBoostSource describes the source of a chat boost.
type ChatBoostSource struct {
	User              *User               `json:"user,omitempty"`
	Source            ChatBoostSourceType `json:"source"`
	GiveawayMessageID int                 `json:"giveaway_message_id,omitempty"`
	IsUnclaimed       bool                `json:"is_unclaimed,omitempty"`
}

// ChatBoostUpdated represents a boost added to a chat or changed.
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved represents a boost removed from a chat.
type ChatBoostRemoved struct {
	BoostID    string          `json:"boost_id"`
	Chat       Chat            `json:"chat"`
	Source     ChatBoostSource `json:"source"`
	RemoveDate int             `json:"remove_date"`
}

// UserChatBoosts represents a list of boosts added to a chat by a user.
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

// BusinessConnection describes the connection of the bot with a business account.
type BusinessConnection struct {
	ID         string `json:"id"`
	User       User   `json:"user"`
	UserChatID int64  `json:"user_chat_id"`
	Date       int64  `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

// BusinessMessagesDeleted is received when messages are deleted from a connected business account.
type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	MessageIDs           []int  `json:"message_ids"`
	Chat                 Chat   `json:"chat"`
}

// Giveaway represents a message about a scheduled giveaway.
type Giveaway struct {
	CountryCodes                  *[]string `json:"country_codes,omitempty"`
	PrizeDescription              string    `json:"prize_description,omitempty"`
	Chats                         []Chat    `json:"chats"`
	WinnersSelectionDate          int       `json:"winners_selection_date"`
	WinnerCount                   int       `json:"winner_count"`
	PremiumSubscriptionMonthCount int       `json:"premium_subscription_month_count,omitempty"`
	OnlyNewMembers                bool      `json:"only_new_members,omitempty"`
	HasPublicWinners              bool      `json:"has_public_winners,omitempty"`
}

// GiveawayCreated represents a service message about the creation of a scheduled giveaway.
// Currently holds no information.
type GiveawayCreated struct{}

// GiveawayWinners represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	PrizeDescription              string `json:"prize_description,omitempty"`
	Chats                         []Chat `json:"chats"`
	Winners                       []User `json:"winners"`
	GiveawayMessageID             int    `json:"giveaway_message_id"`
	WinnersSelectionDate          int    `json:"winners_selection_date"`
	WinnerCount                   int    `json:"winner_count"`
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`
	WasRefunded                   bool   `json:"was_refunded,omitempty"`
}

// GiveawayCompleted represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
}
