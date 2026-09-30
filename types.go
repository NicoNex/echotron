/*
 * Echotron
 * Copyright (C) 2018 The Echotron Contributors
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
	MessageReactionCount     *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	ChatMember               *ChatMemberUpdated           `json:"chat_member,omitempty"`
	RemovedChatBoost         *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	Message                  *Message                     `json:"message,omitempty"`
	EditedMessage            *Message                     `json:"edited_message,omitempty"`
	ChannelPost              *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost        *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection       *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage          *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage    *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages  *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction          *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	ChatBoost                *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	ChosenInlineResult       *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	ChatJoinRequest          *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	CallbackQuery            *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery            *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery         *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	Poll                     *Poll                        `json:"poll,omitempty"`
	PollAnswer               *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember             *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	InlineQuery              *InlineQuery                 `json:"inline_query,omitempty"`
	PurchasedPaidMedia       *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	GuestMessage             *Message                     `json:"guest_message,omitempty"`
	ManagedBot               *ManagedBotUpdated           `json:"managed_bot,omitempty"`
	Subscription             *BotSubscriptionUpdated      `json:"subscription,omitempty"`
	StoppedMessageGeneration *MessageGenerationStopped    `json:"stopped_message_generation,omitempty"`
	ID                       int                          `json:"update_id"`
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
	case u.GuestMessage != nil:
		return u.GuestMessage.Chat.ID
	case u.ManagedBot != nil:
		return u.ManagedBot.User.ID
	case u.Subscription != nil:
		return u.Subscription.User.ID
	case u.StoppedMessageGeneration != nil:
		return u.StoppedMessageGeneration.Chat.ID
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

// APIResponsePreparedInlineMessage represents the incoming response from Telegram servers.
// Used by all methods that return a PreparedInlineMessage object on success.
type APIResponsePreparedInlineMessage struct {
	Result *PreparedInlineMessage `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponsePreparedInlineMessage) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseGifts represents the incoming response from Telegram servers.
// Used by all methods that return a Gifts object on success.
type APIResponseGifts struct {
	Result *Gifts `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseGifts) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseUserProfileAudios represents the incoming response from Telegram servers.
// Used by all methods that return a UserProfileAudios object on success.
type APIResponseUserProfileAudios struct {
	Result *UserProfileAudios `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseUserProfileAudios) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseOwnedGifts represents the incoming response from Telegram servers.
// Used by all methods that return an OwnedGifts object on success.
type APIResponseOwnedGifts struct {
	Result *OwnedGifts `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseOwnedGifts) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseStarAmount represents the incoming response from Telegram servers.
// Used by all methods that return a StarAmount object on success.
type APIResponseStarAmount struct {
	Result *StarAmount `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseStarAmount) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseSentGuestMessage represents the incoming response from Telegram servers.
// Used by all methods that return a SentGuestMessage object on success.
type APIResponseSentGuestMessage struct {
	Result *SentGuestMessage `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseSentGuestMessage) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponseBotAccessSettings represents the incoming response from Telegram servers.
// Used by all methods that return a BotAccessSettings object on success.
type APIResponseBotAccessSettings struct {
	Result *BotAccessSettings `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponseBotAccessSettings) Base() APIResponseBase {
	return a.APIResponseBase
}

// APIResponsePreparedKeyboardButton represents the incoming response from Telegram servers.
// Used by all methods that return a PreparedKeyboardButton object on success.
type APIResponsePreparedKeyboardButton struct {
	Result *PreparedKeyboardButton `json:"result,omitempty"`
	APIResponseBase
}

// Base returns the contained object of type APIResponseBase.
func (a APIResponsePreparedKeyboardButton) Base() APIResponseBase {
	return a.APIResponseBase
}

// User represents a Telegram user or bot.
type User struct {
	FirstName                  string `json:"first_name"`
	LastName                   string `json:"last_name,omitempty"`
	Username                   string `json:"username,omitempty"`
	LanguageCode               string `json:"language_code,omitempty"`
	ID                         int64  `json:"id"`
	IsBot                      bool   `json:"is_bot"`
	IsPremium                  bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu      bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups              bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages    bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries      bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness       bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp              bool   `json:"has_main_web_app,omitempty"`
	HasTopicsEnabled           bool   `json:"has_topics_enabled,omitempty"`
	AllowsUsersToCreateTopics  bool   `json:"allows_users_to_create_topics,omitempty"`
	SupportsGuestQueries       bool   `json:"supports_guest_queries,omitempty"`
	CanManageBots              bool   `json:"can_manage_bots,omitempty"`
	SupportsJoinRequestQueries bool   `json:"supports_join_request_queries,omitempty"`
}

// Chat represents a chat.
type Chat struct {
	Type             string `json:"type"`
	Title            string `json:"title,omitempty"`
	Username         string `json:"username,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	ID               int64  `json:"id"`
	IsForum          bool   `json:"is_forum,omitempty"`
	IsDirectMessages bool   `json:"is_direct_messages,omitempty"`
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
	ParentChat                         *Chat                 `json:"parent_chat,omitempty"`
	Rating                             *UserRating           `json:"rating,omitempty"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors,omitempty"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio,omitempty"`
	GuardBot                           *User                 `json:"guard_bot,omitempty"`
	Community                          *Community            `json:"community,omitempty"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	Username                           string                `json:"username,omitempty"`
	Title                              string                `json:"title,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	Description                        string                `json:"description,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`
	Type                               string                `json:"type"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	BackgroundCustomEmojiID            string                `json:"background_custom_emoji_id,omitempty"`
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`
	ProfileAccentColorID               int                   `json:"profile_accent_color_id,omitempty"`
	EmojiStatusExpirationDate          int                   `json:"emoji_status_expiration_date,omitempty"`
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`
	MaxReactionCount                   int                   `json:"max_reaction_count,omitempty"`
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`
	LinkedChatID                       int64                 `json:"linked_chat_id,omitempty"`
	ID                                 int64                 `json:"id"`
	PaidMessageStarCount               int                   `json:"paid_message_star_count,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id,omitempty"`
	AcceptedGiftTypes                  AcceptedGiftTypes     `json:"accepted_gift_types,omitempty"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
	IsDirectMessages                   bool                  `json:"is_direct_messages,omitempty"`
}

// AcceptedGiftTypes describes the types of gifts that can be gifted to a user or a chat.
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	GiftsFromChannels   bool `json:"gifts_from_channels"`
}

// Message represents a message.
type Message struct {
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
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
	Contact                       *Contact                       `json:"contact,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	Checklist                     *Checklist                     `json:"checklist,omitempty"`
	ChecklistTasksDone            *ChecklistTasksDone            `json:"checklist_tasks_done,omitempty"`
	ChecklistTasksAdded           *ChecklistTasksAdded           `json:"checklist_tasks_added,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	DirectMessagePriceChanged     *DirectMessagePriceChanged     `json:"direct_message_price_changed,omitempty"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic,omitempty"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info,omitempty"`
	SuggestedPostRefunded         *SuggestedPostRefunded         `json:"suggested_post_refunded,omitempty"`
	SuggestedPostPaid             *SuggestedPostPaid             `json:"suggested_post_paid,omitempty"`
	SuggestedPostDeclined         *SuggestedPostDeclined         `json:"suggested_post_declined,omitempty"`
	SuggestedPostApprovalFailed   *SuggestedPostApprovalFailed   `json:"suggested_post_approval_failed,omitempty"`
	SuggestedPostApproved         *SuggestedPostApproved         `json:"suggested_post_approved,omitempty"`
	GiftUpgradeSent               *GiftInfo                      `json:"gift_upgrade_sent,omitempty"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	ChatOwnerChanged              *ChatOwnerChanged              `json:"chat_owner_changed,omitempty"`
	ChatOwnerLeft                 *ChatOwnerLeft                 `json:"chat_owner_left,omitempty"`
	GuestBotCallerUser            *User                          `json:"guest_bot_caller_user,omitempty"`
	GuestBotCallerChat            *Chat                          `json:"guest_bot_caller_chat,omitempty"`
	RichMessage                   *RichMessage                   `json:"rich_message,omitempty"`
	LivePhoto                     *LivePhoto                     `json:"live_photo,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ManagedBotCreated             *ManagedBotCreated             `json:"managed_bot_created,omitempty"`
	PollOptionAdded               *PollOptionAdded               `json:"poll_option_added,omitempty"`
	PollOptionDeleted             *PollOptionDeleted             `json:"poll_option_deleted,omitempty"`
	ReceiverUser                  *User                          `json:"receiver_user,omitempty"`
	CommunityChatAdded            *CommunityChatAdded            `json:"community_chat_added,omitempty"`
	CommunityChatJoined           *CommunityChatJoined           `json:"community_chat_joined,omitempty"`
	CommunityChatRemoved          *CommunityChatRemoved          `json:"community_chat_removed,omitempty"`
	GuestQueryID                  string                         `json:"guest_query_id,omitempty"`
	ReplyToPollOptionID           string                         `json:"reply_to_poll_option_id,omitempty"`
	SenderTag                     string                         `json:"sender_tag,omitempty"`
	EffectID                      string                         `json:"effect_id,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	BusinessConnectionID          string                         `json:"business_connection_id,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	Entities                      []*MessageEntity               `json:"entities,omitempty"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Chat                          Chat                           `json:"chat"`
	Date                          int                            `json:"date"`
	MigrateFromChatID             int                            `json:"migrate_from_chat_id,omitempty"`
	ThreadID                      int                            `json:"message_thread_id,omitempty"`
	ID                            int                            `json:"message_id"`
	MigrateToChatID               int                            `json:"migrate_to_chat_id,omitempty"`
	EditDate                      int                            `json:"edit_date,omitempty"`
	SenderBoostCount              int                            `json:"sender_boost_count,omitempty"`
	ReplyToChecklistTaskID        int                            `json:"reply_to_checklist_task_id,omitempty"`
	PaidStarCount                 int                            `json:"paid_star_count,omitempty"`
	EphemeralMessageID            int                            `json:"ephemeral_message_id,omitempty"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`
	IsPaidPost                    bool                           `json:"is_paid_post,omitempty"`
}

// MessageID represents a unique message identifier.
type MessageID struct {
	MessageID int `json:"message_id"`
}

// MessageEntity represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	User           *User             `json:"user,omitempty"`
	Type           MessageEntityType `json:"type"`
	URL            string            `json:"url,omitempty"`
	Language       string            `json:"language,omitempty"`
	CustomEmojiID  string            `json:"custom_emoji_id,omitempty"`
	DateTimeFormat string            `json:"date_time_format,omitempty"`
	Offset         int               `json:"offset"`
	Length         int               `json:"length"`
	UnixTime       int               `json:"unix_time,omitempty"`
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

// VideoQuality represents a video file of a specific quality.
type VideoQuality struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Codec        string `json:"codec"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	Thumbnail      *PhotoSize     `json:"thumbnail,omitempty"`
	FileID         string         `json:"file_id"`
	FileUniqueID   string         `json:"file_unique_id"`
	FileName       string         `json:"file_name,omitempty"`
	MimeType       string         `json:"mime_type,omitempty"`
	Cover          []PhotoSize    `json:"cover,omitempty"`
	Qualities      []VideoQuality `json:"qualities,omitempty"`
	Width          int            `json:"width"`
	Height         int            `json:"height"`
	Duration       int            `json:"duration"`
	FileSize       int64          `json:"file_size,omitempty"`
	StartTimestamp int            `json:"start_timestamp,omitempty"`
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

// PaidMediaInfo describes the paid media added to a message.
type PaidMediaInfo struct {
	PaidMedia []PaidMedia `json:"paid_media"`
	StarCount int         `json:"star_count"`
}

// PaidMedia describes paid media.
type PaidMedia struct {
	Photo     *[]PhotoSize `json:"photo,omitempty"`
	Video     *Video       `json:"video,omitempty"`
	LivePhoto *LivePhoto   `json:"live_photo,omitempty"`
	Type      string       `json:"type"`
	Width     int          `json:"width,omitempty"`
	Height    int          `json:"height,omitempty"`
	Duration  int          `json:"duration,omitempty"`
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
	Media        *PollMedia       `json:"media,omitempty"`
	AddedByUser  *User            `json:"added_by_user,omitempty"`
	AddedByChat  *Chat            `json:"added_by_chat,omitempty"`
	PersistentID string           `json:"persistent_id"`
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int              `json:"voter_count"`
	AdditionDate int              `json:"addition_date,omitempty"`
}

// InputPollOption contains information about one answer option in a poll to send.
// Media can be an InputMediaAnimation, InputMediaLink, InputMediaLivePhoto, InputMediaLocation,
// InputMediaPhoto, InputMediaSticker, InputMediaVenue or InputMediaVideo.
type InputPollOption struct {
	Media         InputMedia       `json:"-"`
	Text          string           `json:"text"`
	TextParseMode ParseMode        `json:"text_parse_mode,omitempty"`
	TextEntities  []*MessageEntity `json:"text_entities,omitempty"`
}

// MarshalJSON is a custom marshaler for the InputPollOption struct.
func (i InputPollOption) MarshalJSON() ([]byte, error) {
	type alias InputPollOption
	return json.Marshal(struct {
		alias
		Media *mediaEnvelope `json:"media,omitempty"`
	}{alias(i), inlineMediaPtr(i.Media)})
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID              string   `json:"poll_id"`
	VoterChat           *Chat    `json:"voter_chat,omitempty"`
	User                *User    `json:"user,omitempty"`
	OptionIDs           []int    `json:"option_ids"`
	OptionPersistentIDs []string `json:"option_persistent_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	ExplanationMedia      *PollMedia       `json:"explanation_media,omitempty"`
	Media                 *PollMedia       `json:"media,omitempty"`
	Type                  string           `json:"type"`
	Question              string           `json:"question"`
	Explanation           string           `json:"explanation,omitempty"`
	Description           string           `json:"description,omitempty"`
	ID                    string           `json:"id"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	QuestionEntities      []*MessageEntity `json:"question_entities,omitempty"`
	DescriptionEntities   []*MessageEntity `json:"description_entities,omitempty"`
	Options               []*PollOption    `json:"options"`
	CountryCodes          []string         `json:"country_codes,omitempty"`
	CorrectOptionIDs      []int            `json:"correct_option_ids,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	TotalVoterCount       int              `json:"total_voter_count"`
	CloseDate             int              `json:"close_date,omitempty"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	AllowsRevoting        bool             `json:"allows_revoting"`
	MembersOnly           bool             `json:"members_only"`
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
	SubscriptionPeriod      int    `json:"subscription_period,omitempty"`
	SubscriptionPrice       int    `json:"subscription_price,omitempty"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
}

// ChatMember contains information about one member of a chat.
type ChatMember struct {
	User                    *User  `json:"user"`
	Status                  string `json:"status"`
	CustomTitle             string `json:"custom_title,omitempty"`
	Tag                     string `json:"tag,omitempty"`
	UntilDate               int    `json:"until_date,omitempty"`
	CanSendMessages         bool   `json:"can_send_messages,omitempty"`
	CanSendVideos           bool   `json:"can_send_videos,omitempty"`
	CanEditMessages         bool   `json:"can_edit_messages,omitempty"`
	CanDeleteMessages       bool   `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     bool   `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      bool   `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       bool   `json:"can_promote_members,omitempty"`
	CanChangeInfo           bool   `json:"can_change_info,omitempty"`
	CanInviteUsers          bool   `json:"can_invite_users,omitempty"`
	CanPinMessages          bool   `json:"can_pin_messages,omitempty"`
	IsMember                bool   `json:"is_member,omitempty"`
	CanManageChat           bool   `json:"can_manage_chat,omitempty"`
	CanSendAudios           bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments        bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos           bool   `json:"can_send_photos,omitempty"`
	CanPostMessages         bool   `json:"can_post_messages,omitempty"`
	CanSendVideoNotes       bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes       bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls            bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages    bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews   bool   `json:"can_add_web_page_previews,omitempty"`
	CanManageTopics         bool   `json:"can_manage_topics,omitempty"`
	CanPostStories          bool   `json:"can_post_stories,omitempty"`
	CanEditStories          bool   `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool   `json:"can_delete_stories,omitempty"`
	CanManageDirectMessages bool   `json:"can_manage_direct_messages,omitempty"`
	CanBeEdited             bool   `json:"can_be_edited,omitempty"`
	IsAnonymous             bool   `json:"is_anonymous,omitempty"`
	CanEditTag              bool   `json:"can_edit_tag,omitempty"`
	CanManageTags           bool   `json:"can_manage_tags,omitempty"`
	CanReactToMessages      bool   `json:"can_react_to_messages,omitempty"`
	CanSendWelcomeMessages  bool   `json:"can_send_welcome_messages,omitempty"`
}

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
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
	CanManageTags         bool `json:"can_manage_tags,omitempty"`
	CanReactToMessages    bool `json:"can_react_to_messages,omitempty"`
	CanEditTag            bool `json:"can_edit_tag,omitempty"`
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
	IsEphemeral bool   `json:"is_ephemeral,omitempty"`
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
	MediaTypeLivePhoto                = "live_photo"
	MediaTypeSticker                  = "sticker"
	MediaTypeLocation                 = "location"
	MediaTypeVenue                    = "venue"
	MediaTypeLink                     = "link"
	MediaTypeVoiceNote                = "voice_note"
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
	switch o := i.InputMedia.(type) {
	case InputMediaPhoto:
		return json.Marshal(struct {
			Media string `json:"media"`
			InputMediaPhoto
		}{
			InputMediaPhoto: o,
			Media:           i.media,
		})

	case InputMediaVideo:
		return json.Marshal(struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaVideo
		}{
			InputMediaVideo: o,
			Media:           i.media,
			Thumbnail:       i.thumbnail,
		})

	case InputMediaAnimation:
		return json.Marshal(struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaAnimation
		}{
			InputMediaAnimation: o,
			Media:               i.media,
			Thumbnail:           i.thumbnail,
		})

	case InputMediaAudio:
		return json.Marshal(struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaAudio
		}{
			InputMediaAudio: o,
			Media:           i.media,
			Thumbnail:       i.thumbnail,
		})

	case InputMediaDocument:
		return json.Marshal(struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputMediaDocument
		}{
			InputMediaDocument: o,
			Media:              i.media,
			Thumbnail:          i.thumbnail,
		})

	case InputPaidMediaPhoto:
		return json.Marshal(struct {
			Media string `json:"media"`
			InputPaidMediaPhoto
		}{
			InputPaidMediaPhoto: o,
			Media:               i.media,
		})

	case InputPaidMediaVideo:
		return json.Marshal(struct {
			Media     string `json:"media"`
			Thumbnail string `json:"thumbnail,omitempty"`
			InputPaidMediaVideo
		}{
			InputPaidMediaVideo: o,
			Media:               i.media,
			Thumbnail:           i.thumbnail,
		})

	case InputMediaLivePhoto:
		return json.Marshal(struct {
			Media string `json:"media"`
			Photo string `json:"photo"`
			InputMediaLivePhoto
		}{
			InputMediaLivePhoto: o,
			Media:               i.media,
			Photo:               i.thumbnail,
		})

	case InputPaidMediaLivePhoto:
		return json.Marshal(struct {
			Media string `json:"media"`
			Photo string `json:"photo"`
			InputPaidMediaLivePhoto
		}{
			InputPaidMediaLivePhoto: o,
			Media:                   i.media,
			Photo:                   i.thumbnail,
		})

	case InputMediaSticker:
		return json.Marshal(struct {
			Media string `json:"media"`
			InputMediaSticker
		}{
			InputMediaSticker: o,
			Media:             i.media,
		})

	case InputMediaVoiceNote:
		return json.Marshal(struct {
			Media string `json:"media"`
			InputMediaVoiceNote
		}{
			InputMediaVoiceNote: o,
			Media:               i.media,
		})

	case InputMediaLocation, InputMediaVenue, InputMediaLink:
		return json.Marshal(o)

	default:
		return []byte("null"), nil
	}
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
	Cover                 string           `json:"cover,omitempty"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`
	Type                  InputMediaType   `json:"type"`
	Media                 InputFile        `json:"-"`
	Thumbnail             InputFile        `json:"-"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	Width                 int              `json:"width,omitempty"`
	Duration              int              `json:"duration,omitempty"`
	Height                int              `json:"height,omitempty"`
	StartTimestamp        int              `json:"start_timestamp,omitempty"`
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

// InputMediaLivePhoto represents a live photo to be sent.
// Type MUST BE "live_photo".
type InputMediaLivePhoto struct {
	Type                  InputMediaType   `json:"type"`
	Media                 InputFile        `json:"-"`
	Photo                 InputFile        `json:"-"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []*MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool             `json:"has_spoiler,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaLivePhoto) media() InputFile { return i.Media }

// thumbnail returns the static photo, which is uploaded alongside the video like a thumbnail.
func (i InputMediaLivePhoto) thumbnail() InputFile { return i.Photo }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputMediaLivePhoto) groupable() {}

// InputMediaSticker represents a sticker file to be sent.
// Type MUST BE "sticker".
type InputMediaSticker struct {
	Type  InputMediaType `json:"type"`
	Media InputFile      `json:"-"`
	Emoji string         `json:"emoji,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaSticker) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaSticker) thumbnail() InputFile { return InputFile{} }

// InputMediaVoiceNote represents a voice message file to be sent.
// Type MUST BE "voice_note".
type InputMediaVoiceNote struct {
	Type            InputMediaType   `json:"type"`
	Media           InputFile        `json:"-"`
	Caption         string           `json:"caption,omitempty"`
	ParseMode       ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	Duration        int              `json:"duration,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaVoiceNote) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaVoiceNote) thumbnail() InputFile { return InputFile{} }

// InputMediaLocation represents a location to be sent.
// Type MUST BE "location".
type InputMediaLocation struct {
	Type               InputMediaType `json:"type"`
	Latitude           float64        `json:"latitude"`
	Longitude          float64        `json:"longitude"`
	HorizontalAccuracy float64        `json:"horizontal_accuracy,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaLocation) media() InputFile { return InputFile{} }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaLocation) thumbnail() InputFile { return InputFile{} }

// InputMediaVenue represents a venue to be sent.
// Type MUST BE "venue".
type InputMediaVenue struct {
	Type            InputMediaType `json:"type"`
	Title           string         `json:"title"`
	Address         string         `json:"address"`
	FoursquareID    string         `json:"foursquare_id,omitempty"`
	FoursquareType  string         `json:"foursquare_type,omitempty"`
	GooglePlaceID   string         `json:"google_place_id,omitempty"`
	GooglePlaceType string         `json:"google_place_type,omitempty"`
	Latitude        float64        `json:"latitude"`
	Longitude       float64        `json:"longitude"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaVenue) media() InputFile { return InputFile{} }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaVenue) thumbnail() InputFile { return InputFile{} }

// InputMediaLink represents an HTTP link to be sent.
// Type MUST BE "link".
type InputMediaLink struct {
	Type InputMediaType `json:"type"`
	URL  string         `json:"url"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputMedia* struct.
func (i InputMediaLink) media() InputFile { return InputFile{} }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputMedia* struct.
func (i InputMediaLink) thumbnail() InputFile { return InputFile{} }

// InputPaidMediaType represents the various InputPaidMedia types.
type InputPaidMediaType string

// These are the various InputPaidMediaType values.
const (
	InputPaidMediaTypePhoto     InputPaidMediaType = "photo"
	InputPaidMediaTypeVideo                        = "video"
	InputPaidMediaTypeLivePhoto                    = "live_photo"
)

// InputPaidMediaPhoto represents a paid photo to send.
type InputPaidMediaPhoto struct {
	Type  InputPaidMediaType `json:"type"`
	Media InputFile          `json:"-"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputPaidMedia* struct.
func (i InputPaidMediaPhoto) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputPaidMedia* struct.
func (i InputPaidMediaPhoto) thumbnail() InputFile { return InputFile{} }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputPaidMediaPhoto) groupable() {}

// InputPaidMediaVideo represents a paid video to send.
type InputPaidMediaVideo struct {
	Type              InputPaidMediaType `json:"type"`
	Cover             string             `json:"cover,omitempty"`
	Media             InputFile          `json:"-"`
	Thumbnail         InputFile          `json:"-"`
	Width             int                `json:"width,omitempty"`
	Height            int                `json:"height,omitempty"`
	Duration          int                `json:"duration,omitempty"`
	StartTimestamp    int                `json:"start_timestamp,omitempty"`
	SupportsStreaming bool               `json:"supports_streaming,omitempty"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputPaidMedia* struct.
func (i InputPaidMediaVideo) media() InputFile { return i.Media }

// thumbnail is a method which allows to obtain the Thumbnail (type InputFile) field from the InputPaidMedia* struct.
func (i InputPaidMediaVideo) thumbnail() InputFile { return i.Thumbnail }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputPaidMediaVideo) groupable() {}

// InputPaidMediaLivePhoto represents a paid live photo to send.
type InputPaidMediaLivePhoto struct {
	Type  InputPaidMediaType `json:"type"`
	Media InputFile          `json:"-"`
	Photo InputFile          `json:"-"`
}

// media is a method which allows to obtain the Media (type InputFile) field from the InputPaidMedia* struct.
func (i InputPaidMediaLivePhoto) media() InputFile { return i.Media }

// thumbnail returns the static photo, which is uploaded alongside the video like a thumbnail.
func (i InputPaidMediaLivePhoto) thumbnail() InputFile { return i.Photo }

// groupable is a dummy method which exists to implement the interface GroupableInputMedia.
func (i InputPaidMediaLivePhoto) groupable() {}

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
	QueryID    string          `json:"query_id,omitempty"`
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
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
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
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
	IsNameImplicit    bool      `json:"is_name_implicit,omitempty"`
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
	PaidMedia          PaidMediaInfo      `json:"paid_media,omitempty"`
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
	LivePhoto          LivePhoto          `json:"live_photo,omitempty"`
	Checklist          Checklist          `json:"checklist,omitempty"`
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
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	PollOptionID             string          `json:"poll_option_id,omitempty"`
	MessageID                int             `json:"message_id,omitempty"`
	EphemeralMessageID       int             `json:"ephemeral_message_id,omitempty"`
	ChatID                   int64           `json:"chat_id,omitempty"`
	QuotePosition            int             `json:"quote_position,omitempty"`
	ChecklistTaskID          int             `json:"checklist_task_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
}

// SharedUser contains information about a user that was shared with the bot using a KeyboardButtonRequestUser button.
type SharedUser struct {
	Photo     *[]PhotoSize `json:"photo,omitempty"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
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
	PrizeStarCount    int                 `json:"prize_star_count,omitempty"`
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
	Rights     *BusinessBotRights `json:"rights,omitempty"`
	ID         string             `json:"id"`
	User       User               `json:"user"`
	UserChatID int64              `json:"user_chat_id"`
	Date       int64              `json:"date"`
	IsEnabled  bool               `json:"is_enabled"`
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
	PrizeStarCount                int       `json:"prize_star_count,omitempty"`
	WinnersSelectionDate          int       `json:"winners_selection_date"`
	WinnerCount                   int       `json:"winner_count"`
	PremiumSubscriptionMonthCount int       `json:"premium_subscription_month_count,omitempty"`
	OnlyNewMembers                bool      `json:"only_new_members,omitempty"`
	HasPublicWinners              bool      `json:"has_public_winners,omitempty"`
}

// GiveawayCreated represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	PrizeStarCount int `json:"prize_star_count,omitempty"`
}

// GiveawayWinners represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	PrizeDescription              string `json:"prize_description,omitempty"`
	Chat                          Chat   `json:"chat"`
	Winners                       []User `json:"winners"`
	PrizeStarCount                int    `json:"prize_star_count,omitempty"`
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
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
}

// Gift represents a gift that can be sent by the bot.
type Gift struct {
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	ID                     string          `json:"id"`
	Sticker                Sticker         `json:"sticker"`
	PersonalRemainingCount int             `json:"personal_remaining_count,omitempty"`
	RemainingCount         int             `json:"remaining_count,omitempty"`
	PersonalTotalCount     int             `json:"personal_total_count,omitempty"`
	TotalCount             int             `json:"total_count,omitempty"`
	UpgradeStarCount       int             `json:"upgrade_star_count,omitempty"`
	UniqueGiftVariantCount int             `json:"unique_gift_variant_count,omitempty"`
	StarCount              int             `json:"star_count"`
	IsPremium              bool            `json:"is_premium,omitempty"`
	HasColors              bool            `json:"has_colors,omitempty"`
}

// Gifts represents a list of gifts.
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

// StarAmount describes an amount of Telegram Stars.
type StarAmount struct {
	Amount         int `json:"amount"`
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}

// GiftBackground describes the background of a gift.
type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

// UniqueGiftModel describes the model of a unique gift.
type UniqueGiftModel struct {
	Name           string  `json:"name"`
	Rarity         string  `json:"rarity,omitempty"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// UniqueGiftSymbol describes the symbol shown on the pattern of a unique gift.
type UniqueGiftSymbol struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors describes the colors of the backdrop of a unique gift.
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

// UniqueGiftBackdrop describes the backdrop of a unique gift.
type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

// UniqueGiftColors contains information about the color scheme for a user's name, message replies and link previews based on a unique gift.
type UniqueGiftColors struct {
	ModelCustomEmojiID    string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string `json:"symbol_custom_emoji_id"`
	LightThemeOtherColors []int  `json:"light_theme_other_colors"`
	DarkThemeOtherColors  []int  `json:"dark_theme_other_colors"`
	LightThemeMainColor   int    `json:"light_theme_main_color"`
	DarkThemeMainColor    int    `json:"dark_theme_main_color"`
}

// UniqueGift describes a unique gift that was upgraded from a regular gift.
type UniqueGift struct {
	PublisherChat    *Chat              `json:"publisher_chat,omitempty"`
	Colors           *UniqueGiftColors  `json:"colors,omitempty"`
	GiftID           string             `json:"gift_id"`
	BaseName         string             `json:"base_name"`
	Name             string             `json:"name"`
	Model            UniqueGiftModel    `json:"model"`
	Symbol           UniqueGiftSymbol   `json:"symbol"`
	Backdrop         UniqueGiftBackdrop `json:"backdrop"`
	Number           int                `json:"number"`
	IsPremium        bool               `json:"is_premium,omitempty"`
	IsBurned         bool               `json:"is_burned,omitempty"`
	IsFromBlockchain bool               `json:"is_from_blockchain,omitempty"`
}

// GiftInfo describes a service message about a regular gift that was sent or received.
type GiftInfo struct {
	Gift                    Gift             `json:"gift"`
	OwnedGiftID             string           `json:"owned_gift_id,omitempty"`
	Text                    string           `json:"text,omitempty"`
	Entities                []*MessageEntity `json:"entities,omitempty"`
	ConvertStarCount        int              `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int              `json:"prepaid_upgrade_star_count,omitempty"`
	UniqueGiftNumber        int              `json:"unique_gift_number,omitempty"`
	IsUpgradeSeparate       bool             `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           bool             `json:"can_be_upgraded,omitempty"`
	IsPrivate               bool             `json:"is_private,omitempty"`
}

// UniqueGiftInfo describes a service message about a unique gift that was sent or received.
type UniqueGiftInfo struct {
	OwnedGiftID        string           `json:"owned_gift_id,omitempty"`
	LastResaleCurrency string           `json:"last_resale_currency,omitempty"`
	Origin             string           `json:"origin"`
	Text               string           `json:"text,omitempty"`
	Entities           []*MessageEntity `json:"entities,omitempty"`
	Gift               UniqueGift       `json:"gift"`
	LastResaleAmount   int              `json:"last_resale_amount,omitempty"`
	TransferStarCount  int              `json:"transfer_star_count,omitempty"`
	NextTransferDate   int              `json:"next_transfer_date,omitempty"`
	IsPrivate          bool             `json:"is_private,omitempty"`
}

// OwnedGiftRegular describes a regular gift owned by a user or a chat.
type OwnedGiftRegular struct {
	SenderUser              *User            `json:"sender_user,omitempty"`
	Gift                    Gift             `json:"gift"`
	OwnedGiftID             string           `json:"owned_gift_id,omitempty"`
	Text                    string           `json:"text,omitempty"`
	Type                    string           `json:"type"`
	Entities                []*MessageEntity `json:"entities,omitempty"`
	SendDate                int              `json:"send_date"`
	ConvertStarCount        int              `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int              `json:"prepaid_upgrade_star_count,omitempty"`
	UniqueGiftNumber        int              `json:"unique_gift_number,omitempty"`
	IsPrivate               bool             `json:"is_private,omitempty"`
	IsSaved                 bool             `json:"is_saved,omitempty"`
	CanBeUpgraded           bool             `json:"can_be_upgraded,omitempty"`
	WasRefunded             bool             `json:"was_refunded,omitempty"`
	IsUpgradeSeparate       bool             `json:"is_upgrade_separate,omitempty"`
}

// OwnedGiftUnique describes a unique gift received and owned by a user or a chat.
type OwnedGiftUnique struct {
	SenderUser        *User      `json:"sender_user,omitempty"`
	OwnedGiftID       string     `json:"owned_gift_id,omitempty"`
	Type              string     `json:"type"`
	Gift              UniqueGift `json:"gift"`
	SendDate          int        `json:"send_date"`
	TransferStarCount int        `json:"transfer_star_count,omitempty"`
	NextTransferDate  int        `json:"next_transfer_date,omitempty"`
	IsSaved           bool       `json:"is_saved,omitempty"`
	CanBeTransferred  bool       `json:"can_be_transferred,omitempty"`
}

// OwnedGifts contains the list of gifts received and owned by a user or a chat.
type OwnedGifts struct {
	NextOffset string            `json:"next_offset,omitempty"`
	Gifts      []json.RawMessage `json:"gifts"`
	TotalCount int               `json:"total_count"`
}

// UserProfileAudios represents the audios displayed on a user's profile.
type UserProfileAudios struct {
	Audios     []Audio `json:"audios"`
	TotalCount int     `json:"total_count"`
}

// UserRating describes the rating of a user based on their Telegram Star spendings.
type UserRating struct {
	Level              int `json:"level"`
	Rating             int `json:"rating"`
	CurrentLevelRating int `json:"current_level_rating"`
	NextLevelRating    int `json:"next_level_rating,omitempty"`
}

// ChecklistTask describes a task in a checklist.
type ChecklistTask struct {
	CompletedByUser *User            `json:"completed_by_user,omitempty"`
	CompletedByChat *Chat            `json:"completed_by_chat,omitempty"`
	Text            string           `json:"text"`
	TextEntities    []*MessageEntity `json:"text_entities,omitempty"`
	ID              int              `json:"id"`
	CompletionDate  int              `json:"completion_date,omitempty"`
}

// Checklist describes a checklist.
type Checklist struct {
	Title                    string           `json:"title"`
	TitleEntities            []*MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []ChecklistTask  `json:"tasks"`
	OthersCanAddTasks        bool             `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool             `json:"others_can_mark_tasks_as_done,omitempty"`
}

// InputChecklistTask describes a task to add to a checklist.
type InputChecklistTask struct {
	Text         string           `json:"text"`
	ParseMode    ParseMode        `json:"parse_mode,omitempty"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	ID           int              `json:"id"`
}

// InputChecklist describes a checklist to create.
type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                ParseMode            `json:"parse_mode,omitempty"`
	TitleEntities            []*MessageEntity     `json:"title_entities,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool                 `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool                 `json:"others_can_mark_tasks_as_done,omitempty"`
}

// ChecklistTasksDone describes a service message about checklist tasks marked as done or not done.
type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message,omitempty"`
	MarkedAsDoneTaskIDs    []int    `json:"marked_as_done_task_ids,omitempty"`
	MarkedAsNotDoneTaskIDs []int    `json:"marked_as_not_done_task_ids,omitempty"`
}

// ChecklistTasksAdded describes a service message about tasks added to a checklist.
type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message,omitempty"`
	Tasks            []ChecklistTask `json:"tasks"`
}

// PaidMessagePriceChanged describes a service message about a change in the price of paid messages within a chat.
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

// DirectMessagePriceChanged describes a service message about a change in the price of direct messages sent to a channel chat.
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int  `json:"direct_message_star_count,omitempty"`
}

// SuggestedPostPrice describes the price of a suggested post.
type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// SuggestedPostInfo contains information about a suggested post.
type SuggestedPostInfo struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	State    string              `json:"state"`
	SendDate int                 `json:"send_date,omitempty"`
}

// SuggestedPostParameters contains parameters of a post that is being suggested by the bot.
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int                 `json:"send_date,omitempty"`
}

// SuggestedPostApproved describes a service message about the approval of a suggested post.
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int                 `json:"send_date"`
}

// SuggestedPostApprovalFailed describes a service message about the failed approval of a suggested post.
type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message           `json:"suggested_post_message,omitempty"`
	Price                SuggestedPostPrice `json:"price"`
}

// SuggestedPostDeclined describes a service message about the rejection of a suggested post.
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              string   `json:"comment,omitempty"`
}

// SuggestedPostPaid describes a service message about a successful payment for a suggested post.
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
	Currency             string      `json:"currency"`
	Amount               int         `json:"amount,omitempty"`
}

// SuggestedPostRefunded describes a service message about a payment refund for a suggested post.
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}

// DirectMessagesTopic describes a topic of a direct messages chat.
type DirectMessagesTopic struct {
	User    *User `json:"user,omitempty"`
	TopicID int64 `json:"topic_id"`
}

// ChatOwnerLeft describes a service message about the chat owner leaving the chat.
type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner,omitempty"`
}

// ChatOwnerChanged describes a service message about an ownership change in the chat.
type ChatOwnerChanged struct {
	NewOwner User `json:"new_owner"`
}

// BusinessBotRights represents the rights of a business bot.
type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply,omitempty"`
	CanReadMessages            bool `json:"can_read_messages,omitempty"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages,omitempty"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages,omitempty"`
	CanEditName                bool `json:"can_edit_name,omitempty"`
	CanEditBio                 bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            bool `json:"can_edit_username,omitempty"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings,omitempty"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars,omitempty"`
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars,omitempty"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           bool `json:"can_manage_stories,omitempty"`
}

// APIResponseStory represents the incoming response from Telegram servers.
// Used by the PostStory, RepostStory, and EditStory methods.
type APIResponseStory struct {
	APIResponseBase
	Result Story `json:"result"`
}

// Base is used to get the APIResponseBase of the APIResponseStory type.
func (a APIResponseStory) Base() APIResponseBase {
	return a.APIResponseBase
}

// LocationAddress describes the physical address of a location.
type LocationAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state,omitempty"`
	City        string `json:"city,omitempty"`
	Street      string `json:"street,omitempty"`
}

// StoryAreaType describes the type of a clickable area on a story.
type StoryAreaType interface {
	ImplementsStoryAreaType()
}

// StoryAreaTypeLocation describes a story area pointing to a location.
type StoryAreaTypeLocation struct {
	Address   *LocationAddress `json:"address,omitempty"`
	Type      string           `json:"type"`
	Latitude  float64          `json:"latitude"`
	Longitude float64          `json:"longitude"`
}

// ImplementsStoryAreaType is used to implement the StoryAreaType interface.
func (s StoryAreaTypeLocation) ImplementsStoryAreaType() {}

// StoryAreaTypeSuggestedReaction describes a story area pointing to a suggested reaction.
type StoryAreaTypeSuggestedReaction struct {
	ReactionType ReactionType `json:"reaction_type"`
	Type         string       `json:"type"`
	IsDark       bool         `json:"is_dark,omitempty"`
	IsFlipped    bool         `json:"is_flipped,omitempty"`
}

// ImplementsStoryAreaType is used to implement the StoryAreaType interface.
func (s StoryAreaTypeSuggestedReaction) ImplementsStoryAreaType() {}

// StoryAreaTypeLink describes a story area pointing to an HTTP or tg:// link.
type StoryAreaTypeLink struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// ImplementsStoryAreaType is used to implement the StoryAreaType interface.
func (s StoryAreaTypeLink) ImplementsStoryAreaType() {}

// StoryAreaTypeWeather describes a story area containing weather information.
type StoryAreaTypeWeather struct {
	Type            string  `json:"type"`
	Emoji           string  `json:"emoji"`
	Temperature     float64 `json:"temperature"`
	BackgroundColor int     `json:"background_color"`
}

// ImplementsStoryAreaType is used to implement the StoryAreaType interface.
func (s StoryAreaTypeWeather) ImplementsStoryAreaType() {}

// StoryAreaTypeUniqueGift describes a story area pointing to a unique gift.
type StoryAreaTypeUniqueGift struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// ImplementsStoryAreaType is used to implement the StoryAreaType interface.
func (s StoryAreaTypeUniqueGift) ImplementsStoryAreaType() {}

// StoryAreaPosition describes the position of a clickable area within a story.
type StoryAreaPosition struct {
	XPercentage            float64 `json:"x_percentage"`
	YPercentage            float64 `json:"y_percentage"`
	WidthPercentage        float64 `json:"width_percentage"`
	HeightPercentage       float64 `json:"height_percentage"`
	RotationAngle          float64 `json:"rotation_angle"`
	CornerRadiusPercentage float64 `json:"corner_radius_percentage"`
}

// StoryArea describes a clickable area on a story media.
type StoryArea struct {
	Type     StoryAreaType     `json:"type"`
	Position StoryAreaPosition `json:"position"`
}

// InputStoryContent describes the content of a story to post.
type InputStoryContent interface {
	ImplementsInputStoryContent()
	storyContentFile() InputFile
}

// InputStoryContentPhoto describes a photo to post as a story.
type InputStoryContentPhoto struct {
	Photo InputFile
}

// ImplementsInputStoryContent is used to implement the InputStoryContent interface.
func (i InputStoryContentPhoto) ImplementsInputStoryContent() {}

func (i InputStoryContentPhoto) storyContentFile() InputFile { return i.Photo }

// InputStoryContentVideo describes a video to post as a story.
type InputStoryContentVideo struct {
	Video               InputFile
	Duration            float64
	CoverFrameTimestamp float64
	IsAnimation         bool
}

// ImplementsInputStoryContent is used to implement the InputStoryContent interface.
func (i InputStoryContentVideo) ImplementsInputStoryContent() {}

func (i InputStoryContentVideo) storyContentFile() InputFile { return i.Video }

// storyContentEnvelope is used internally to marshal InputStoryContent to JSON.
type storyContentEnvelope struct {
	content InputStoryContent
	ref     string // "attach://filename" or file_id or url
}

// MarshalJSON is a custom marshaler for the storyContentEnvelope struct.
func (s storyContentEnvelope) MarshalJSON() ([]byte, error) {
	switch o := s.content.(type) {
	case InputStoryContentPhoto:
		return json.Marshal(struct {
			Type  string `json:"type"`
			Photo string `json:"photo"`
		}{
			Type:  "photo",
			Photo: s.ref,
		})

	case InputStoryContentVideo:
		return json.Marshal(struct {
			Type                string  `json:"type"`
			Video               string  `json:"video"`
			Duration            float64 `json:"duration,omitempty"`
			CoverFrameTimestamp float64 `json:"cover_frame_timestamp,omitempty"`
			IsAnimation         bool    `json:"is_animation,omitempty"`
		}{
			Type:                "video",
			Video:               s.ref,
			Duration:            o.Duration,
			CoverFrameTimestamp: o.CoverFrameTimestamp,
			IsAnimation:         o.IsAnimation,
		})

	default:
		return []byte("null"), nil
	}
}

// InputProfilePhoto describes a profile photo to set.
type InputProfilePhoto interface {
	ImplementsInputProfilePhoto()
	profilePhotoFile() InputFile
}

// SentGuestMessage contains information about a message sent by a guest bot.
type SentGuestMessage struct {
	InlineMessageID string `json:"inline_message_id"`
}

// BotAccessSettings contains information about the access settings of a managed bot.
type BotAccessSettings struct {
	AddedUsers         []User `json:"added_users,omitempty"`
	IsAccessRestricted bool   `json:"is_access_restricted"`
}

// PreparedKeyboardButton describes a prepared keyboard button.
type PreparedKeyboardButton struct {
	ID string `json:"id"`
}

// InputRichMessage describes a rich message to be sent.
// Exactly one of the fields HTML, Markdown or Blocks must be used.
type InputRichMessage struct {
	HTML                string                  `json:"html,omitempty"`
	Markdown            string                  `json:"markdown,omitempty"`
	Blocks              []InputRichBlock        `json:"blocks,omitempty"`
	Media               []InputRichMessageMedia `json:"media,omitempty"`
	IsRTL               bool                    `json:"is_rtl,omitempty"`
	SkipEntityDetection bool                    `json:"skip_entity_detection,omitempty"`
}

// InputRichMessageMedia describes a media element embedded in an outgoing rich message
// through tg://photo?id=, tg://video?id=, tg://document?id= or tg://audio?id= links.
// Media can be an InputMediaAnimation, InputMediaAudio, InputMediaDocument,
// InputMediaPhoto, InputMediaVideo or InputMediaVoiceNote.
type InputRichMessageMedia struct {
	Media InputMedia `json:"-"`
	ID    string     `json:"id"`
}

// MarshalJSON is a custom marshaler for the InputRichMessageMedia struct.
func (i InputRichMessageMedia) MarshalJSON() ([]byte, error) {
	type alias InputRichMessageMedia
	return json.Marshal(struct {
		alias
		Media *mediaEnvelope `json:"media,omitempty"`
	}{alias(i), inlineMediaPtr(i.Media)})
}

// InputProfilePhotoStatic describes a static profile photo in the .JPG format.
type InputProfilePhotoStatic struct {
	Photo InputFile
}

// ImplementsInputProfilePhoto is used to implement the InputProfilePhoto interface.
func (i InputProfilePhotoStatic) ImplementsInputProfilePhoto() {}

func (i InputProfilePhotoStatic) profilePhotoFile() InputFile { return i.Photo }

// InputProfilePhotoAnimated describes an animated profile photo in the MPEG4 format.
type InputProfilePhotoAnimated struct {
	Animation          InputFile
	MainFrameTimestamp float64
}

// ImplementsInputProfilePhoto is used to implement the InputProfilePhoto interface.
func (i InputProfilePhotoAnimated) ImplementsInputProfilePhoto() {}

func (i InputProfilePhotoAnimated) profilePhotoFile() InputFile { return i.Animation }

// profilePhotoEnvelope is used internally to marshal InputProfilePhoto to JSON.
type profilePhotoEnvelope struct {
	content InputProfilePhoto
	ref     string // "attach://filename" or file_id or url
}

// MarshalJSON is a custom marshaler for the profilePhotoEnvelope struct.
func (p profilePhotoEnvelope) MarshalJSON() ([]byte, error) {
	switch o := p.content.(type) {
	case InputProfilePhotoStatic:
		return json.Marshal(struct {
			Type  string `json:"type"`
			Photo string `json:"photo"`
		}{
			Type:  "static",
			Photo: p.ref,
		})

	case InputProfilePhotoAnimated:
		return json.Marshal(struct {
			Type               string  `json:"type"`
			Animation          string  `json:"animation"`
			MainFrameTimestamp float64 `json:"main_frame_timestamp,omitempty"`
		}{
			Type:               "animated",
			Animation:          p.ref,
			MainFrameTimestamp: o.MainFrameTimestamp,
		})

	default:
		return []byte("null"), nil
	}
}

// LivePhoto represents a live photo.
type LivePhoto struct {
	Photo        []PhotoSize `json:"photo,omitempty"`
	FileID       string      `json:"file_id"`
	FileUniqueID string      `json:"file_unique_id"`
	MimeType     string      `json:"mime_type,omitempty"`
	Width        int         `json:"width"`
	Height       int         `json:"height"`
	Duration     int         `json:"duration"`
	FileSize     int64       `json:"file_size,omitempty"`
}

// Link represents an HTTP link.
type Link struct {
	URL string `json:"url"`
}

// PollMedia represents the media added to a poll, a poll option or a quiz explanation.
// At most one of the fields can be present.
type PollMedia struct {
	Animation *Animation  `json:"animation,omitempty"`
	Audio     *Audio      `json:"audio,omitempty"`
	Document  *Document   `json:"document,omitempty"`
	Link      *Link       `json:"link,omitempty"`
	LivePhoto *LivePhoto  `json:"live_photo,omitempty"`
	Location  *Location   `json:"location,omitempty"`
	Sticker   *Sticker    `json:"sticker,omitempty"`
	Venue     *Venue      `json:"venue,omitempty"`
	Video     *Video      `json:"video,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// PollOptionAdded describes a service message about an option added to a poll.
type PollOptionAdded struct {
	PollMessage        *Message         `json:"poll_message,omitempty"`
	OptionPersistentID string           `json:"option_persistent_id"`
	OptionText         string           `json:"option_text"`
	OptionTextEntities []*MessageEntity `json:"option_text_entities,omitempty"`
}

// PollOptionDeleted describes a service message about an option deleted from a poll.
type PollOptionDeleted PollOptionAdded

// ManagedBotCreated contains information about the bot that was created to be managed by the current bot.
type ManagedBotCreated struct {
	Bot User `json:"bot"`
}

// ManagedBotUpdated contains information about the creation, token update, or owner update of a bot managed by the current bot.
type ManagedBotUpdated struct {
	User User `json:"user"`
	Bot  User `json:"bot"`
}

// RichText represents rich formatted text.
// It is either plain text (Plain), a concatenation of rich texts (Texts)
// or a formatted entity identified by Type, in which case the other fields
// relevant to that type are set.
type RichText struct {
	Text            *RichText          `json:"text,omitempty"`
	User            *User              `json:"user,omitempty"`
	Button          *RichMessageButton `json:"button,omitempty"`
	Type            string             `json:"type"`
	Plain           string             `json:"-"`
	DateTimeFormat  string             `json:"date_time_format,omitempty"`
	CustomEmojiID   string             `json:"custom_emoji_id,omitempty"`
	AlternativeText string             `json:"alternative_text,omitempty"`
	Expression      string             `json:"expression,omitempty"`
	URL             string             `json:"url,omitempty"`
	EmailAddress    string             `json:"email_address,omitempty"`
	PhoneNumber     string             `json:"phone_number,omitempty"`
	BankCardNumber  string             `json:"bank_card_number,omitempty"`
	Username        string             `json:"username,omitempty"`
	Hashtag         string             `json:"hashtag,omitempty"`
	Cashtag         string             `json:"cashtag,omitempty"`
	BotCommand      string             `json:"bot_command,omitempty"`
	Name            string             `json:"name,omitempty"`
	AnchorName      string             `json:"anchor_name,omitempty"`
	ReferenceName   string             `json:"reference_name,omitempty"`
	Texts           []RichText         `json:"-"`
	UnixTime        int                `json:"unix_time,omitempty"`
}

// MarshalJSON is a custom marshaler for the RichText struct.
func (r RichText) MarshalJSON() ([]byte, error) {
	type alias RichText

	switch {
	case r.Type != "":
		return json.Marshal(alias(r))
	case r.Texts != nil:
		return json.Marshal(r.Texts)
	default:
		return json.Marshal(r.Plain)
	}
}

// UnmarshalJSON is a custom unmarshaler for the RichText struct.
func (r *RichText) UnmarshalJSON(b []byte) error {
	type alias RichText

	switch {
	case len(b) > 0 && b[0] == '"':
		return json.Unmarshal(b, &r.Plain)
	case len(b) > 0 && b[0] == '[':
		return json.Unmarshal(b, &r.Texts)
	default:
		return json.Unmarshal(b, (*alias)(r))
	}
}

// RichBlockCaption represents the caption of a rich formatted block.
type RichBlockCaption struct {
	Credit *RichText `json:"credit,omitempty"`
	Text   RichText  `json:"text"`
}

// RichBlockTableCell represents a cell in a table.
type RichBlockTableCell struct {
	Text     *RichText `json:"text,omitempty"`
	Align    string    `json:"align"`
	Valign   string    `json:"valign"`
	Colspan  int       `json:"colspan,omitempty"`
	Rowspan  int       `json:"rowspan,omitempty"`
	IsHeader bool      `json:"is_header,omitempty"`
}

// RichBlockListItem represents an item of a list.
type RichBlockListItem struct {
	Label       string      `json:"label"`
	Type        string      `json:"type,omitempty"`
	Blocks      []RichBlock `json:"blocks"`
	Value       int         `json:"value,omitempty"`
	HasCheckbox bool        `json:"has_checkbox,omitempty"`
	IsChecked   bool        `json:"is_checked,omitempty"`
}

// RichBlock represents a block in a rich formatted message.
// Type identifies the kind of block and which of the other fields are set.
// Caption is set for media, collage, slideshow and map blocks, while
// TableCaption is set for table blocks.
type RichBlock struct {
	Text         *RichText              `json:"text,omitempty"`
	Credit       *RichText              `json:"credit,omitempty"`
	Summary      *RichText              `json:"summary,omitempty"`
	Caption      *RichBlockCaption      `json:"-"`
	TableCaption *RichText              `json:"-"`
	Location     *Location              `json:"location,omitempty"`
	Animation    *Animation             `json:"animation,omitempty"`
	Audio        *Audio                 `json:"audio,omitempty"`
	Document     *Document              `json:"document,omitempty"`
	Video        *Video                 `json:"video,omitempty"`
	VoiceNote    *Voice                 `json:"voice_note,omitempty"`
	Type         string                 `json:"type"`
	Language     string                 `json:"language,omitempty"`
	Expression   string                 `json:"expression,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Align        string                 `json:"align,omitempty"`
	Photo        []PhotoSize            `json:"photo,omitempty"`
	Items        []RichBlockListItem    `json:"items,omitempty"`
	Blocks       []RichBlock            `json:"blocks,omitempty"`
	Cells        [][]RichBlockTableCell `json:"cells,omitempty"`
	Buttons      []RichMessageButton    `json:"buttons,omitempty"`
	Size         int                    `json:"size,omitempty"`
	Zoom         int                    `json:"zoom,omitempty"`
	Width        int                    `json:"width,omitempty"`
	Height       int                    `json:"height,omitempty"`
	HasSpoiler   bool                   `json:"has_spoiler,omitempty"`
	IsBordered   bool                   `json:"is_bordered,omitempty"`
	IsStriped    bool                   `json:"is_striped,omitempty"`
	IsCompact    bool                   `json:"is_compact,omitempty"`
	IsOpen       bool                   `json:"is_open,omitempty"`
}

// UnmarshalJSON is a custom unmarshaler for the RichBlock struct.
// The caption of table blocks is a RichText, while it's a RichBlockCaption for the others.
func (r *RichBlock) UnmarshalJSON(b []byte) error {
	type alias RichBlock
	aux := struct {
		*alias
		Caption json.RawMessage `json:"caption"`
	}{alias: (*alias)(r)}

	if err := json.Unmarshal(b, &aux); err != nil || aux.Caption == nil {
		return err
	}

	if r.Type == "table" {
		return json.Unmarshal(aux.Caption, &r.TableCaption)
	}
	return json.Unmarshal(aux.Caption, &r.Caption)
}

// RichMessage represents a rich formatted message.
type RichMessage struct {
	Blocks []RichBlock `json:"blocks"`
	IsRTL  bool        `json:"is_rtl,omitempty"`
}

// RichMessageButton represents a button in a rich message.
// Exactly one of the fields other than Text and Style must be used to specify the type of the button.
type RichMessageButton struct {
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	Disabled                     *DisabledButton              `json:"disabled,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	Style                        string                       `json:"style,omitempty"`
	URL                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	Text                         RichText                     `json:"text"`
}

// InputRichBlockListItem represents an item of a list to be sent.
type InputRichBlockListItem struct {
	Type        string           `json:"type,omitempty"`
	Blocks      []InputRichBlock `json:"blocks"`
	Value       int              `json:"value,omitempty"`
	HasCheckbox bool             `json:"has_checkbox,omitempty"`
	IsChecked   bool             `json:"is_checked,omitempty"`
}

// InputRichBlock represents a block of a rich message to be sent.
// Type identifies the kind of block and which of the other fields are used.
// Media is used by the "animation", "audio", "document", "photo", "video"
// and "voice_note" blocks and must be of the matching InputMedia* type.
// Caption is used by media, collage, slideshow and map blocks, while
// TableCaption is used by table blocks.
type InputRichBlock struct {
	Text         *RichText                `json:"text,omitempty"`
	Credit       *RichText                `json:"credit,omitempty"`
	Summary      *RichText                `json:"summary,omitempty"`
	Caption      *RichBlockCaption        `json:"-"`
	TableCaption *RichText                `json:"-"`
	Location     *Location                `json:"location,omitempty"`
	Media        InputMedia               `json:"-"`
	Type         string                   `json:"type"`
	Language     string                   `json:"language,omitempty"`
	Expression   string                   `json:"expression,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Align        string                   `json:"align,omitempty"`
	Items        []InputRichBlockListItem `json:"items,omitempty"`
	Blocks       []InputRichBlock         `json:"blocks,omitempty"`
	Cells        [][]RichBlockTableCell   `json:"cells,omitempty"`
	Buttons      []RichMessageButton      `json:"buttons,omitempty"`
	Size         int                      `json:"size,omitempty"`
	Zoom         int                      `json:"zoom,omitempty"`
	Width        int                      `json:"width,omitempty"`
	Height       int                      `json:"height,omitempty"`
	IsBordered   bool                     `json:"is_bordered,omitempty"`
	IsStriped    bool                     `json:"is_striped,omitempty"`
	IsCompact    bool                     `json:"is_compact,omitempty"`
	IsOpen       bool                     `json:"is_open,omitempty"`
}

// MarshalJSON is a custom marshaler for the InputRichBlock struct.
func (i InputRichBlock) MarshalJSON() ([]byte, error) {
	type alias InputRichBlock
	var caption any

	if i.TableCaption != nil {
		caption = i.TableCaption
	} else if i.Caption != nil {
		caption = i.Caption
	}

	b, err := json.Marshal(struct {
		alias
		Caption any `json:"caption,omitempty"`
	}{alias(i), caption})
	if err != nil || i.Media == nil {
		return b, err
	}

	// The media is serialised under the key named after the block type.
	var obj map[string]json.RawMessage
	if err = json.Unmarshal(b, &obj); err != nil {
		return nil, err
	}
	if obj[i.Type], err = json.Marshal(inlineMedia(i.Media)); err != nil {
		return nil, err
	}
	return json.Marshal(obj)
}

// Community represents a community, a group of chats linked together around a shared topic or audience.
type Community struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

// CommunityChatAdded describes a service message about a chat or a bot being added to a community.
type CommunityChatAdded struct {
	Community Community `json:"community"`
}

// CommunityChatJoined describes a service message about a chat being joined by a user from a community.
type CommunityChatJoined struct {
	Community Community `json:"community"`
}

// CommunityChatRemoved describes a service message about a chat or a bot being removed from a community.
type CommunityChatRemoved struct{}

// BotSubscriptionUpdated contains information about changes to a user payment subscription toward the bot.
type BotSubscriptionUpdated struct {
	InvoicePayload string `json:"invoice_payload"`
	State          string `json:"state"`
	User           User   `json:"user"`
}

// MessageGenerationStopped describes an update about a user stopping message generation.
type MessageGenerationStopped struct {
	Chat            Chat `json:"chat"`
	MessageThreadID int  `json:"message_thread_id,omitempty"`
	DraftID         int  `json:"draft_id"`
}
