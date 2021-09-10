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
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

// API is the object that contains all the functions that wrap those of the Telegram Bot API.
type API struct {
	token string
	base  string
}

func encode(s string) string {
	return url.QueryEscape(s)
}

func sendFile(file InputFile, url, fileType string) (res []byte, err error) {
	switch {
	case file.id != "":
		res, err = sendGetRequest(fmt.Sprintf("%s&%s=%s", url, fileType, file.id))

	case file.path != "" && len(file.content) == 0:
		file.content, err = os.ReadFile(file.path)
		if err != nil {
			return
		}
		file.path = filepath.Base(file.path)
		fallthrough

	case file.path != "" && len(file.content) > 0:
		res, err = sendPostRequest(url, content{file.path, fileType, file.content})
	}

	if err != nil {
		return
	}

	return res, nil
}

func sendMediaFiles(url string, files ...InputMedia) (res []byte, err error) {
	var jsn []inputMedia
	var cnt []content
	var dat []byte

	for _, file := range files {
		media := file.GetMedia()

		switch {
		case media.id != "":
			jsn = append(jsn, inputMedia{media.id, file})
		
		case media.path != "" && len(media.content) == 0:
			media.content, err = os.ReadFile(media.path)
			if err != nil {
				return
			}
			media.path = filepath.Base(media.path)
			fallthrough

		case media.path != "" && len(media.content) > 0:
			cnt = append(cnt, content{media.path, media.path, media.content})
			jsn = append(jsn, inputMedia{fmt.Sprintf("attach://%s", media.path), file})
		}
	}

	if len(jsn) == 1 {
		dat, err = json.Marshal(jsn[0])
	} else {
		dat, err = json.Marshal(jsn)
	}

	if err != nil {
		return
	}

	url = fmt.Sprintf("%s&media=%s", url, dat)

	if len(cnt) != 0 {
		return sendPostRequest(url, cnt...)
	} else {
		return sendGetRequest(url)
	}
}

func serializePerms(permissions ChatPermissions) (string, error) {
	perm, err := json.Marshal(PermissionOptions{permissions})
	if err != nil {
		return "", err
	}

	return string(perm), nil
}

// NewAPI returns a new API object.
func NewAPI(token string) API {
	return API{
		token: token,
		base:  fmt.Sprintf("https://api.telegram.org/bot%s/", token),
	}
}

// GetUpdates is used to receive incoming updates using long polling.
func (a API) GetUpdates(opts *UpdateOptions) (APIResponseUpdate, error) {
	var res APIResponseUpdate
	var url = fmt.Sprintf(
		"%sgetUpdates?%s",
		a.base,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetWebhook is used to specify a url and receive incoming updates via an outgoing webhook.
func (a API) SetWebhook(webhookURL string, dropPendingUpdates bool, opts *WebhookOptions) (APIResponseBase, error) {
	var res APIResponseBase
	var url = fmt.Sprintf(
		"%ssetWebhook?drop_pending_updates=%t&%s",
		a.base,
		dropPendingUpdates,
		querify(opts),
	)

	keyVal := map[string]string{"url": webhookURL}
	cnt, err := sendPostForm(url, keyVal)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DeleteWebhook is used to remove webhook integration if you decide to switch back to GetUpdates.
func (a API) DeleteWebhook(dropPendingUpdates bool) (APIResponseBase, error) {
	var res APIResponseBase
	var url = fmt.Sprintf(
		"%sdeleteWebhook?drop_pending_updates=%t",
		a.base,
		dropPendingUpdates,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetWebhookInfo is used to get current webhook status.
func (a API) GetWebhookInfo() (APIResponseWebhook, error) {
	var res APIResponseWebhook
	var url = fmt.Sprintf(
		"%sgetWebhookInfo",
		a.base,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetMe is a simple method for testing your bot's auth token.
func (a API) GetMe() (APIResponseUser, error) {
	var res APIResponseUser
	var url = fmt.Sprintf(
		"%sgetMe",
		a.base,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// LogOut is used to log out from the cloud Bot API server before launching the bot locally.
// You MUST log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates.
// After a successful call, you can immediately log in on a local server,
// but will not be able to log in back to the cloud Bot API server for 10 minutes.
func (a API) LogOut() (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%slogOut",
		a.base,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// Close is used to close the bot instance before moving it from one local server to another.
// You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart.
// The method will return error 429 in the first 10 minutes after the bot is launched.
func (a API) Close() (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sclose",
		a.base,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendMessage is used to send text messages.
func (a API) SendMessage(text string, chatID int64, opts *MessageOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&%s",
		a.base,
		encode(text),
		chatID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// ForwardMessage is used to forward messages of any kind.
// Service messages can't be forwarded.
func (a API) ForwardMessage(chatID, fromChatID int64, messageID int, opts *DisableNotificationOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%sforwardMessage?chat_id=%d&from_chat_id=%d&message_id=%d&%s",
		a.base,
		chatID,
		fromChatID,
		messageID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// CopyMessage is used to copy messages of any kind.
// Service messages and invoice messages can't be copied.
// The method is analogous to the method ForwardMessage,
// but the copied message doesn't have a link to the original message.
func (a API) CopyMessage(chatID, fromChatID int64, messageID int, opts *CopyOptions) (APIResponseMessageID, error) {
	var res APIResponseMessageID
	var url = fmt.Sprintf(
		"%scopyMessage?chat_id=%d&from_chat_id=%d&message_id=%d&%s",
		a.base,
		chatID,
		fromChatID,
		messageID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendPhoto is used to send photos.
func (a API) SendPhoto(file InputFile, chatID int64, opts *PhotoOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendPhoto?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "photo")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendAudio is used to send audio files,
// if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
func (a API) SendAudio(file InputFile, chatID int64, opts *AudioOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendAudio?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "audio")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendDocument is used to send general files.
func (a API) SendDocument(file InputFile, chatID int64, opts *DocumentOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendDocument?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "document")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendVideo is used to send video files.
// Telegram clients support mp4 videos (other formats may be sent with SendDocument).
func (a API) SendVideo(file InputFile, chatID int64, opts *VideoOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVideo?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "video")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendAnimation is used to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
func (a API) SendAnimation(file InputFile, chatID int64, opts *AnimationOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendAnimation?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "animation")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendVoice is used to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document).
func (a API) SendVoice(file InputFile, chatID int64, opts *VoiceOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVoice?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "voice")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendVideoNote is used to send video messages.
func (a API) SendVideoNote(file InputFile, chatID int64, opts *VideoNoteOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVideoNote?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendFile(file, url, "video_note")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendLocation is used to send point on the map.
func (a API) SendLocation(chatID int64, latitude, longitude float64, opts *LocationOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendLocation?chat_id=%d&latitude=%f&longitude=%f&%s",
		a.base,
		chatID,
		latitude,
		longitude,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// EditMessageLiveLocation is used to edit live location messages.
// A location can be edited until its `LivePeriod` expires or editing is explicitly disabled by a call to `StopMessageLiveLocation`.
func (a API) EditMessageLiveLocation(msg MessageIDOptions, latitude, longitude float64, opts *EditLocationOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageLiveLocation?latitude=%f&longitude=%f&%s&%s",
		a.base,
		latitude,
		longitude,
		querify(msg),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// StopMessageLiveLocation is used to stop updating a live location message before `LivePeriod` expires.
func (a API) StopMessageLiveLocation(msg MessageIDOptions, opts *MessageReplyMarkup) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%sstopMessageLiveLocation?%s&%s",
		a.base,
		querify(msg),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendVenue is used to send information about a venue.
func (a API) SendVenue(chatID int64, latitude, longitude float64, title, address string, opts *VenueOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVenue?chat_id=%d&latitude=%f&longitude=%f&title=%s&address=%s&%s",
		a.base,
		chatID,
		latitude,
		longitude,
		encode(title),
		encode(address),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendContact is used to send phone contacts.
func (a API) SendContact(phoneNumber, firstName string, chatID int64, opts *ContactOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendContact?chat_id=%d&phone_number=%s&first_name=%s&%s",
		a.base,
		chatID,
		encode(phoneNumber),
		encode(firstName),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendPoll is used to send a native poll.
func (a API) SendPoll(chatID int64, question string, options []string, opts *PollOptions) (APIResponseMessage, error) {
	var res APIResponseMessage

	pollOpts, err := json.Marshal(options)
	if err != nil {
		return res, err
	}

	var url = fmt.Sprintf(
		"%ssendPoll?chat_id=%d&question=%s&options=%s&%s",
		a.base,
		chatID,
		question,
		encode(string(pollOpts)),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendDice is used to send an animated emoji that will display a random value.
func (a API) SendDice(chatID int64, emoji DiceEmoji, opts *BaseOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendDice?chat_id=%d&emoji=%s&%s",
		a.base,
		chatID,
		encode(string(emoji)),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SendChatAction is used to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
func (a API) SendChatAction(action ChatAction, chatID int64) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%ssendChatAction?chat_id=%d&action=%s",
		a.base,
		chatID,
		action,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetUserProfilePhotos is used to get a list of profile pictures for a user.
func (a API) GetUserProfilePhotos(userID int64, opts *UserProfileOptions) (APIResponseUserProfile, error) {
	var res APIResponseUserProfile
	var url = fmt.Sprintf(
		"%sgetUserProfilePhotos?user_id=%d&%s",
		a.base,
		userID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetFile returns the basic info about a file and prepares it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// The file can then be downloaded with DownloadFile where filePath is taken from the response.
// It is guaranteed that the file will be downloadable for at least 1 hour.
// When the download file expires, a new one can be requested by calling GetFile again.
func (a API) GetFile(fileID string) (APIResponseFile, error) {
	var res APIResponseFile
	var url = fmt.Sprintf(
		"%sgetFile?file_id=%s",
		a.base,
		fileID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DownloadFile returns the bytes of the file corresponding to the given filePath.
// This function is callable for at least 1 hour since the call to GetFile.
// When the download expires a new one can be requested by calling GetFile again.
func (a API) DownloadFile(filePath string) ([]byte, error) {
	return sendGetRequest(fmt.Sprintf(
		"https://api.telegram.org/file/bot%s/%s",
		a.token,
		filePath,
	))
}

// BanChatMember is used to ban a user in a group, a supergroup or a channel.
// In the case of supergroups or channels, the user will not be able to return to the chat
// on their own using invite links, etc., unless unbanned first (through the UnbanChatMember method).
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
func (a API) BanChatMember(chatID, userID int64, opts *BanOptions) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sbanChatMember?chat_id=%d&user_id=%d&%s",
		a.base,
		chatID,
		userID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// UnbanChatMember is used to unban a previously banned user in a supergroup or channel.
// The user will NOT return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work.
// By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it.
// So if the user is a member of the chat they will also be REMOVED from the chat.
// If you don't want this, use the parameter `OnlyIfBanned`.
func (a API) UnbanChatMember(chatID, userID int64, opts *UnbanOptions) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sunbanChatMember?chat_id=%d&user_id=%d&%s",
		a.base,
		chatID,
		userID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// RestrictChatMember is used to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
func (a API) RestrictChatMember(chatID, userID int64, permissions ChatPermissions, opts *RestrictOptions) (APIResponseBool, error) {
	var res APIResponseBool

	perm, err := serializePerms(permissions)
	if err != nil {
		return res, err
	}

	var url = fmt.Sprintf(
		"%srestrictChatMember?chat_id=%d&user_id=%d&permissions=%s&%s",
		a.base,
		chatID,
		userID,
		encode(perm),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// PromoteChatMember is used to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
func (a API) PromoteChatMember(chatID, userID int64, opts *PromoteOptions) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%spromoteChatMember?chat_id=%d&user_id=%d&%s",
		a.base,
		chatID,
		userID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetChatAdministratorCustomTitle is used to set a custom title for an administrator in a supergroup promoted by the bot.
func (a API) SetChatAdministratorCustomTitle(chatID, userID int64, customTitle string) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%ssetChatAdministratorCustomTitle?chat_id=%d&user_id=%d&custom_title=%s",
		a.base,
		chatID,
		userID,
		encode(customTitle),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetChatPermissions is used to set default chat permissions for all members.
// The bot must be an administrator in the supergroup for this to work and must have the can_restrict_members admin rights.
func (a API) SetChatPermissions(chatID int64, permissions ChatPermissions) (APIResponseBool, error) {
	var res APIResponseBool

	perm, err := serializePerms(permissions)
	if err != nil {
		return res, err
	}

	var url = fmt.Sprintf(
		"%ssetChatPermissions?chat_id=%d&permissions=%s",
		a.base,
		chatID,
		encode(perm),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// ExportChatInviteLink is used to generate a new primary invite link for a chat;
// any previously generated primary link is revoked.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
func (a API) ExportChatInviteLink(chatID int64) (APIResponseString, error) {
	var res APIResponseString
	var url = fmt.Sprintf(
		"%sexportChatInviteLink?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// CreateChatInviteLink is used to create an additional invite link for a chat.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
// The link can be revoked using the method RevokeChatInviteLink.
func (a API) CreateChatInviteLink(chatID int64, opts *InviteLinkOptions) (APIResponseInviteLink, error) {
	var res APIResponseInviteLink
	var url = fmt.Sprintf(
		"%screateChatInviteLink?chat_id=%d&%s",
		a.base,
		chatID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// EditChatInviteLink is used to edit a non-primary invite link created by the bot.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
func (a API) EditChatInviteLink(chatID int64, inviteLink string, opts *InviteLinkOptions) (APIResponseInviteLink, error) {
	var res APIResponseInviteLink
	var url = fmt.Sprintf(
		"%seditChatInviteLink?chat_id=%d&invite_link=%s&%s",
		a.base,
		chatID,
		encode(inviteLink),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// RevokeChatInviteLink is used to revoke an invite link created by the bot.
// If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights.
func (a API) RevokeChatInviteLink(chatID int64, inviteLink string) (APIResponseInviteLink, error) {
	var res APIResponseInviteLink
	var url = fmt.Sprintf(
		"%srevokeChatInviteLink?chat_id=%d&invite_link=%s",
		a.base,
		chatID,
		encode(inviteLink),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetChatPhoto is used to set a new profile photo for the chat.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
func (a API) SetChatPhoto(file InputFile, chatID int64) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%ssetChatPhoto?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendFile(file, url, "photo")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DeleteChatPhoto is used to delete a chat photo.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
func (a API) DeleteChatPhoto(chatID int64) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sdeleteChatPhoto?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetChatTitle is used to change the title of a chat.
// Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
func (a API) SetChatTitle(chatID int64, title string) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%ssetChatTitle?chat_id=%d&title=%s",
		a.base,
		chatID,
		encode(title),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetChatDescription is used to change the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
func (a API) SetChatDescription(chatID int64, description string) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%ssetChatDescription?chat_id=%d&description=%s",
		a.base,
		chatID,
		encode(description),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// PinChatMessage is used to add a message to the list of pinned messages in the chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work
// and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel.
func (a API) PinChatMessage(chatID int64, messageID int, opts *DisableNotificationOptions) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%spinChatMessage?chat_id=%d&message_id=%d&%s",
		a.base,
		chatID,
		messageID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// UnpinChatMessage is used to remove a message from the list of pinned messages in the chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work
// and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel.
func (a API) UnpinChatMessage(chatID int64, messageID int) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sunpinChatMessage?chat_id=%d&message_id=%d",
		a.base,
		chatID,
		messageID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// UnpinAllChatMessages is used to clear the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work
// and must have the 'can_pin_messages' admin right in a supergroup or 'can_edit_messages' admin right in a channel.
func (a API) UnpinAllChatMessages(chatID int64) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sunpinAllChatMessages?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// LeaveChat is used to make the bot leave a group, supergroup or channel.
func (a API) LeaveChat(chatID int64) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sleaveChat?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetChat is used to get up to date information about the chat.
// (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.)
func (a API) GetChat(chatID int64) (APIResponseChat, error) {
	var res APIResponseChat
	var url = fmt.Sprintf(
		"%sgetChat?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetChatAdministrators is used to get a list of administrators in a chat.
func (a API) GetChatAdministrators(chatID int64) (APIResponseAdministrators, error) {
	var res APIResponseAdministrators
	var url = fmt.Sprintf(
		"%sgetChatAdministrators?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetChatMemberCount is used to get the number of members in a chat.
func (a API) GetChatMemberCount(chatID int64) (APIResponseInteger, error) {
	var res APIResponseInteger
	var url = fmt.Sprintf(
		"%sgetChatMemberCount?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetChatMember is used to get information about a member of a chat.
func (a API) GetChatMember(chatID, userID int64) (APIResponseChatMember, error) {
	var res APIResponseChatMember
	var url = fmt.Sprintf(
		"%sgetChatMember?chat_id=%d&user_id=%d",
		a.base,
		chatID,
		userID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetChatStickerSet is used to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Use the field `CanSetStickerSet` optionally returned in GetChat requests to check if the bot can use this method.
func (a API) SetChatStickerSet(chatID int64, stickerSetName string) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%ssetChatStickerSet?chat_id=%d&sticker_set_name=%s",
		a.base,
		chatID,
		encode(stickerSetName),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DeleteChatStickerSet is used to delete a group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Use the field `CanSetStickerSet` optionally returned in GetChat requests to check if the bot can use this method.
func (a API) DeleteChatStickerSet(chatID int64) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sdeleteChatStickerSet?chat_id=%d",
		a.base,
		chatID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// AnswerCallbackQuery is used to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
func (a API) AnswerCallbackQuery(callbackID string, opts *CallbackQueryOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%sanswerCallbackQuery?callback_query_id=%s&%s",
		a.base,
		callbackID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetMyCommands is used to change the list of the bot's commands for the given scope and user language.
func (a API) SetMyCommands(opts *CommandOptions, commands ...BotCommand) (APIResponseBool, error) {
	var res APIResponseBool
	jsn, _ := json.Marshal(commands)

	var url = fmt.Sprintf(
		"%ssetMyCommands?commands=%s&%s",
		a.base,
		jsn,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DeleteMyCommands is used to delete the list of the bot's commands for the given scope and user language.
func (a API) DeleteMyCommands(opts *CommandOptions) (APIResponseBool, error) {
	var res APIResponseBool
	var url = fmt.Sprintf(
		"%sdeleteMyCommands?%s",
		a.base,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetMyCommands is used to get the current list of the bot's commands for the given scope and user language.
func (a API) GetMyCommands(opts *CommandOptions) (APIResponseCommands, error) {
	var res APIResponseCommands
	var url = fmt.Sprintf(
		"%sgetMyCommands?%s",
		a.base,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// EditMessageText is used to edit text and game messages.
func (a API) EditMessageText(text string, msg MessageIDOptions, opts *MessageTextOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageText?text=%s&%s&%s",
		a.base,
		encode(text),
		querify(msg),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// EditMessageCaption is used to edit captions of messages.
func (a API) EditMessageCaption(msg MessageIDOptions, opts *MessageCaptionOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageCaption?%s&%s",
		a.base,
		querify(msg),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// EditMessageMedia is used to edit animation, audio, document, photo or video messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise.
// When an inline message is edited, a new file can't be uploaded.
// Use a previously uploaded file via its file_id or specify a URL.
func (a API) EditMessageMedia(msg MessageIDOptions, media InputMedia, opts *MessageReplyMarkup) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageMedia?%s&%s",
		a.base,
		querify(msg),
		querify(opts),
	)

	cnt, err := sendMediaFiles(url, media)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// EditMessageReplyMarkup is used to edit only the reply markup of messages.
func (a API) EditMessageReplyMarkup(msg MessageIDOptions, opts *MessageReplyMarkup) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageReplyMarkup?%s&%s",
		a.base,
		querify(msg),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// StopPoll is used to stop a poll which was sent by the bot.
func (a API) StopPoll(chatID int64, messageID int, opts *MessageReplyMarkup) (APIResponsePoll, error) {
	var res APIResponsePoll
	var url = fmt.Sprintf(
		"%sstopPoll?chat_id=%d&message_id=%d&%s",
		a.base,
		chatID,
		messageID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DeleteMessage is used to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
func (a API) DeleteMessage(chatID int64, messageID int) (APIResponseBase, error) {
	var res APIResponseBase
	var url = fmt.Sprintf(
		"%sdeleteMessage?chat_id=%d&message_id=%d",
		a.base,
		chatID,
		messageID,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}
