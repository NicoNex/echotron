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

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// API is the object that contains all the functions that wrap those of the Telegram Bot API.
type API string

// Option is a custom type for the various frequent options used by some methods of the API.
type Option string

// These are all the possible options that can be used by some methods.
const (
	ParseMarkdown         Option = "&parse_mode=markdown"
	ParseHTML                    = "&parse_mode=html"
	DisableWebPagePreview        = "&disable_web_page_preview=true"
	DisableNotification          = "&disable_notification=true"
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

func encode(s string) string {
	return url.QueryEscape(s)
}

func parseOpts(opts ...Option) string {
	var buf strings.Builder

	for _, o := range opts {
		buf.WriteString(string(o))
	}
	return buf.String()
}

func makeInlineKeyboard(rows ...InlineKbdRow) InlineKeyboard {
	return InlineKeyboard{rows}
}

// NewAPI returns a new API object.
func NewAPI(token string) API {
	return API(fmt.Sprintf("https://api.telegram.org/bot%s/", token))
}

// DeleteWebhook is used to remove webhook integration if you decide to switch back to getUpdates.
func (a API) DeleteWebhook() (APIResponseUpdate, error) {
	var res APIResponseUpdate

	content, err := SendGetRequest(fmt.Sprintf("%sdeleteWebhook", string(a)))
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SetWebhook is used to specify a url and receive incoming updates via an outgoing webhook.
func (a API) SetWebhook(url string) (APIResponseUpdate, error) {
	var res APIResponseUpdate

	keyVal := map[string]string{"url": url}
	content, err := SendPostForm(fmt.Sprintf("%ssetWebhook", string(a)), keyVal)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// GetUpdates is used to receive incoming updates using long polling.
func (a API) GetUpdates(offset, timeout int) (APIResponseUpdate, error) {
	var res APIResponseUpdate
	var url = fmt.Sprintf("%sgetUpdates?timeout=%d", string(a), timeout)

	if offset != 0 {
		url = fmt.Sprintf("%s&offset=%d", url, offset)
	}

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// GetChat is used to get up to date information about the chat.
// (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.)
func (a API) GetChat(chatID int64) (Chat, error) {
	var res Chat
	var url = fmt.Sprintf("%sgetChat?chat_id=%d", string(a), chatID)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// GetStickerSet is used to get a sticker set.
func (a API) GetStickerSet(name string) (StickerSet, error) {
	var res StickerSet
	var url = fmt.Sprintf("%sgetStickerSet?name=%s", string(a), encode(name))

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendMessage is used to send text messages.
func (a API) SendMessage(text string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d%s",
		string(a),
		encode(text),
		chatID,
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendMessageReply is used to send a message as a reply to a previously received one.
func (a API) SendMessageReply(text string, chatID int64, messageID int, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&reply_to_message_id=%d%s",
		string(a),
		encode(text),
		chatID,
		messageID,
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}

	json.Unmarshal(content, &res)
	return res, nil
}

// SendMessageWithKeyboard is used to send a message with a keyboard.
func (a API) SendMessageWithKeyboard(text string, chatID int64, keyboard []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&reply_markup=%s%s",
		string(a),
		encode(text),
		chatID,
		keyboard,
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// DeleteMessage is used to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
func (a API) DeleteMessage(chatID int64, messageID int) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%sdeleteMessage?chat_id=%d&message_id=%d",
		string(a),
		chatID,
		messageID,
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendPhoto is used to send photos.
func (a API) SendPhoto(filepath, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseMessage{}, err
	}
	return a.SendPhotoBytes(filepath, caption, chatID, b, opts...)
}

// SendPhotoBytes is used to send photos as a slice of bytes.
func (a API) SendPhotoBytes(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendPhoto?chat_id=%d&caption=%s%s",
		string(a),
		chatID,
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "photo", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendPhotoByID is used to send photos through an ID of a photo that already exists on the Telegram servers.
func (a API) SendPhotoByID(photoID, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendPhoto?chat_id=%d&photo=%s&caption=%s%s",
		string(a),
		chatID,
		encode(photoID),
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendPhotoWithKeyboard is used to send photos with a keyboard.
func (a API) SendPhotoWithKeyboard(filepath, caption string, chatID int64, keyboard []byte, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseMessage{}, err
	}
	return a.SendPhotoWithKeyboardBytes(filepath, caption, chatID, b, keyboard, opts...)
}

// SendPhotoWithKeyboardBytes is used to send photos as a slice of bytes with a keyboard.
func (a API) SendPhotoWithKeyboardBytes(filepath, caption string, chatID int64, data []byte, keyboard []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendPhoto?chat_id=%d&caption=%s&reply_markup=%s%s",
		string(a),
		chatID,
		encode(caption),
		keyboard,
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "photo", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendAudio is used to send audio files,
// if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
func (a API) SendAudio(filepath, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseMessage{}, err
	}
	return a.SendAudioBytes(filepath, caption, chatID, b, opts...)
}

// SendAudioBytes is used to send audio files as a slice of bytes,
// if you want Telegram clients to display them in the music player.
func (a API) SendAudioBytes(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendAudio?chat_id=%d&caption=%s%s",
		string(a),
		chatID,
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "audio", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendAudioByID is used to send audio files that already exist on the Telegram servers,
// if you want Telegram clients to display them in the music player.
func (a API) SendAudioByID(audioID, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendAudio?chat_id=%d&audio=%s&caption=%s%s",
		string(a),
		chatID,
		encode(audioID),
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendDocument is used to send general files.
func (a API) SendDocument(filepath, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseMessage{}, err
	}
	return a.SendDocumentBytes(filepath, caption, chatID, b, opts...)
}

// SendDocumentBytes is used to send general files as a slice of bytes.
func (a API) SendDocumentBytes(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendDocument?chat_id=%d&caption=%s%s",
		string(a),
		chatID,
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "document", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendDocumentByID is used to send general files that already exist on the Telegram servers.
func (a API) SendDocumentByID(documentID, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendDocument?chat_id=%d&document=%s&caption=%s%s",
		string(a),
		chatID,
		encode(documentID),
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendVideo is used to send video files.
// Telegram clients support mp4 videos (other formats may be sent with SendDocument).
func (a API) SendVideo(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseMessage{}, err
	}
	return a.SendVideoBytes(filepath, caption, chatID, b, opts...)
}

// SendVideoBytes is used to send video files as a slice of bytes.
func (a API) SendVideoBytes(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVideo?chat_id=%d&caption=%s%s",
		string(a),
		chatID,
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "video", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendVideoByID is used to send video files that already exist on the Telegram servers.
func (a API) SendVideoByID(videoID, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVideo?chat_id=%d&video=%s&caption=%s%s",
		string(a),
		chatID,
		encode(videoID),
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendVideoNote is used to send video messages.
func (a API) SendVideoNote(videoID string, chatID int64) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVideoNote?chat_id=%d&video_note=%s",
		string(a),
		chatID,
		encode(videoID),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendVoice is used to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document).
func (a API) SendVoice(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseMessage{}, err
	}
	return a.SendVoiceBytes(filepath, caption, chatID, b, opts...)
}

// SendVoiceBytes is used to send audio files as a slice of bytes,
// if you want Telegram clients to display the file as a playable voice message.
func (a API) SendVoiceBytes(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVoice?chat_id=%d&caption=%s%s",
		string(a),
		chatID,
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "voice", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendVoiceByID is used to send audio files that already exists on the Telegram servers,
// if you want Telegram clients to display the file as a playable voice message.
func (a API) SendVoiceByID(voiceID, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendVoice?chat_id=%d&voice=%s%s",
		string(a),
		chatID,
		encode(voiceID),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendContact is used to send phone contacts.
func (a API) SendContact(phoneNumber, firstName, lastName string, chatID int64) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendContact?chat_id=%d&phone_number=%s&first_name=%s&last_name=%s",
		string(a),
		chatID,
		encode(phoneNumber),
		encode(firstName),
		encode(lastName),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendSticker is used to send static .WEBP or animated .TGS stickers.
func (a API) SendSticker(stickerID string, chatID int64) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendSticker?chat_id=%d&sticker=%s",
		string(a),
		chatID,
		encode(stickerID),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendChatAction is used to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
func (a API) SendChatAction(action ChatAction, chatID int64) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendChatAction?chat_id=%d&action=%s",
		string(a),
		chatID,
		action,
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// KeyboardButton is a wrapper method for the Button type.
// It's used to create a new keyboard button.
func (a API) KeyboardButton(text string, requestContact, requestLocation bool) Button {
	return Button{text, requestContact, requestLocation}
}

// KeyboardRow is a wrapper method for the KbdRow type.
// It's used to create a row of keyboard buttons.
func (a API) KeyboardRow(buttons ...Button) (kbdRow KbdRow) {
	for _, button := range buttons {
		kbdRow = append(kbdRow, button)
	}

	return
}

// KeyboardMarkup represents a custom keyboard with reply options.
// This method generates the actual JSON that will be sent to Telegram to make the desired keyboard show up in a message.
func (a API) KeyboardMarkup(resizeKeyboard, oneTimeKeyboard, selective bool, keyboardRows ...KbdRow) (kbd []byte) {
	kbd, _ = json.Marshal(Keyboard{
		keyboardRows,
		resizeKeyboard,
		oneTimeKeyboard,
		selective,
	})
	return
}

// KeyboardRemove generates the object to send in a message to remove
// the current custom keyboard and display the default letter-keyboard.
func (a API) KeyboardRemove(selective bool) (kbdrmv []byte) {
	kbdrmv, _ = json.Marshal(KeyboardRemove{true, selective})
	return
}

// InlineKbdBtn is a wrapper method for the InlineButton type.
// It's used to create a new inline keyboard button.
func (a API) InlineKbdBtn(text, url, callbackData string) InlineButton {
	return InlineButton{
		encode(text),
		url,
		callbackData,
	}
}

// InlineKbdBtnURL is a wrapper method for InlineKbdBtn, but only with url.
func (a API) InlineKbdBtnURL(text, url string) InlineButton {
	return a.InlineKbdBtn(text, url, "")
}

// InlineKbdBtnCbd is a wrapper method for InlineKbdBtn, but only with callbackData.
func (a API) InlineKbdBtnCbd(text, callbackData string) InlineButton {
	return a.InlineKbdBtn(text, "", callbackData)
}

// InlineKbdRow is a wrapper method for the InlineKbdRow type.
// It's used to create a row of inline keyboard buttons.
func (a API) InlineKbdRow(inlineButtons ...InlineButton) InlineKbdRow {
	return inlineButtons
}

// InlineKbdMarkup represents an inline keyboard that appears right next to the message it belongs to.
// This method generates the actual JSON that will be sent to Telegram to make the desired inline keyboard show up in a message.
func (a API) InlineKbdMarkup(inlineKbdRows ...InlineKbdRow) (jsn []byte) {
	jsn, _ = json.Marshal(makeInlineKeyboard(inlineKbdRows...))
	return
}

// EditMessageReplyMarkup is used to edit only the reply markup of messages.
func (a API) EditMessageReplyMarkup(chatID int64, messageID int, keyboard []byte) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageReplyMarkup?chat_id=%d&message_id=%d&reply_markup=%s",
		string(a),
		chatID,
		messageID,
		keyboard,
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// EditMessageText is used to edit text and game messages.
func (a API) EditMessageText(chatID int64, messageID int, text string, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageText?chat_id=%d&message_id=%d&text=%s%s",
		string(a),
		chatID,
		messageID,
		encode(text),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// EditMessageTextWithKeyboard is the same as EditMessageText, but allows to send a custom keyboard.
func (a API) EditMessageTextWithKeyboard(chatID int64, messageID int, text string, keyboard []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%seditMessageText?chat_id=%d&message_id=%d&text=%s&reply_markup=%s%s",
		string(a),
		chatID,
		messageID,
		encode(text),
		keyboard,
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// AnswerCallbackQuery is used to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
func (a API) AnswerCallbackQuery(id, text string, showAlert bool) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%sanswerCallbackQuery?callback_query_id=%s&text=%s&show_alert=%s",
		string(a),
		id,
		text,
		strconv.FormatBool(showAlert),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// GetMyCommands is used to get the current list of the bot's commands.
func (a API) GetMyCommands() (APIResponseCommands, error) {
	var res APIResponseCommands
	var url = fmt.Sprintf(
		"%sgetMyCommands",
		string(a),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SetMyCommands is used to change the list of the bot's commands.
func (a API) SetMyCommands(commands ...BotCommand) (APIResponseCommands, error) {
	var res APIResponseCommands
	jsn, _ := json.Marshal(commands)

	var url = fmt.Sprintf(
		"%ssetMyCommands?commands=%s",
		string(a),
		jsn,
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// Command is a wrapper method for the BotCommand type.
// It's used to create a new command for the bot.
func (a API) Command(command, description string) BotCommand {
	return BotCommand{command, description}
}

func (a Api) SendAnimation(filepath, caption string, chatId int64, opts ...Option) (APIResponseCommands, error) {
// SendAnimation is used to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
func (a API) SendAnimation(filepath, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return APIResponseCommands{}, err
	}
	return a.SendAnimationBytes(filepath, caption, chatID, b, opts...)
}

// SendAnimationBytes is used to send animation files (GIF or H.264/MPEG-4 AVC video without sound) as a slice of bytes.
func (a API) SendAnimationBytes(filepath, caption string, chatID int64, data []byte, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendAnimation?chat_id=%d&caption=%s%s",
		string(a),
		chatID,
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendPostRequest(url, filepath, "animation", data)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// SendAnimationByID is used to send animation files (GIF or H.264/MPEG-4 AVC video without sound) that already exist on the Telegram servers.
func (a API) SendAnimationByID(animationID, caption string, chatID int64, opts ...Option) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendAnimation?chat_id=%d&animation=%s&caption=%s%s",
		string(a),
		chatID,
		encode(animationID),
		encode(caption),
		parseOpts(opts...),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// AnswerInlineQuery is a wrapper method for AnswerInlineQueryOptions.
func (a API) AnswerInlineQuery(inlineQueryID string, results []InlineQueryResult) (APIResponseBase, error) {
	return a.AnswerInlineQueryOptions(inlineQueryID, results, InlineQueryOptions{CacheTime: 300})
}

// AnswerInlineQueryOptions is used to send answers to an inline query.
func (a API) AnswerInlineQueryOptions(inlineQueryID string, results []InlineQueryResult, opts InlineQueryOptions) (APIResponseBase, error) {
	var res APIResponseBase
	jsn, _ := json.Marshal(results)

	var url = fmt.Sprintf(
		"%sanswerInlineQuery?inline_query_id=%s&results=%s%s",
		string(a),
		inlineQueryID,
		jsn,
		parseInlineQueryOpts(opts),
	)

	content, err := SendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// parseInlineQueryOpts is an helper method to properly format InlineQueryOptions as a string to add to the request.
func parseInlineQueryOpts(opts InlineQueryOptions) string {
	return fmt.Sprintf(
		"&cache_time=%d&is_personal=%t&next_offset=%s&switch_pm_text=%s&switch_pm_parameter=%s",
		opts.CacheTime,
		opts.IsPersonal,
		opts.NextOffset,
		opts.SwitchPmText,
		opts.SwitchPmParameter,
	)
}
