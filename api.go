/*
 * Echotron
 * Copyright (C) 2019  Nicolò Santamaria, Michele Dimaggio, Alessandro Ianne
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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type Api string

type Option int

const (
	PARSE_MARKDOWN Option = iota
	PARSE_HTML
	DISABLE_WEB_PAGE_PREVIEW
	DISABLE_NOTIFICATION
)

type ChatAction string

const (
	TYPING            ChatAction = "typing"
	UPLOAD_PHOTO                 = "upload_photo"
	RECORD_VIDEO                 = "record_video"
	UPLOAD_VIDEO                 = "upload_video"
	RECORD_AUDIO                 = "record_audio"
	UPLOAD_AUDIO                 = "upload_audio"
	UPLOAD_DOCUMENT              = "upload_document"
	FIND_LOCATION                = "find_location"
	RECORD_VIDEO_NOTE            = "record_video_note"
	UPLOAD_VIDEO_NOTE            = "upload_video_note"
)

func encode(s string) string {
	return url.QueryEscape(s)
}

func parseOpts(opts ...Option) string {
	var buf bytes.Buffer

	for _, o := range opts {
		switch o {
		case PARSE_MARKDOWN:
			buf.WriteString("&parse_mode=markdown")

		case PARSE_HTML:
			buf.WriteString("&parse_mode=html")

		case DISABLE_WEB_PAGE_PREVIEW:
			buf.WriteString("&disable_web_page_preview=true")

		case DISABLE_NOTIFICATION:
			buf.WriteString("&disable_notification=true")
		}
	}
	return buf.String()
}

// NewApi returns a new Api object.
func NewApi(token string) Api {
	return Api(fmt.Sprintf("https://api.telegram.org/bot%s/", token))
}

// DeleteWebhook deletes webhook
func (a Api) DeleteWebhook() (response APIResponseUpdate) {
	content := SendGetRequest(string(a) + "deleteWebhook")
	json.Unmarshal(content, &response)
	return
}

// SetWebhook sets the webhook to bot on Telegram servers
func (a Api) SetWebhook(url string) (response APIResponseUpdate) {
	keyVal := map[string]string{"url": url}
	content, err := SendPostForm(fmt.Sprintf("%ssetWebhook", string(a)), keyVal)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal(content, &response)
	return
}

// GetResponse returns the incoming updates from telegram.
func (a Api) GetUpdates(offset, timeout int) (response APIResponseUpdate) {
	var url = fmt.Sprintf("%sgetUpdates?timeout=%d", string(a), timeout)

	if offset != 0 {
		url = fmt.Sprintf("%s&offset=%d", url, offset)
	}
	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

// Returns the current chat in use.
func (a Api) GetChat(chatId int64) (response Chat) {
	var url = fmt.Sprintf("%sgetChat?chat_id=%d", string(a), chatId)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) GetStickerSet(name string) (response StickerSet) {
	var url = fmt.Sprintf("%sgetStickerSet?name=%s", string(a), encode(name))

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendMessage(text string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d%s",
		string(a),
		encode(text),
		chatId,
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

// Sends a message as a reply to a previously received one.
func (a Api) SendMessageReply(text string, chatId int64, messageId int, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&reply_to_message_id=%d%s",
		string(a),
		encode(text),
		chatId,
		messageId,
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendMessageWithKeyboard(text string, chatId int64, keyboard []byte) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendMessage?text=%s&chat_id=%d&parse_mode=markdown&reply_markup=%s",
		string(a),
		encode(text),
		chatId,
		keyboard,
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) DeleteMessage(chatId int64, messageId int) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%sdeleteMessage?chat_id=%d&message_id=%d",
		string(a),
		chatId,
		messageId,
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendPhoto(filename, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendPhoto?chat_id=%d&caption=%s%s",
		string(a),
		chatId,
		encode(caption),
		parseOpts(opts...),
	)

	content := SendPostRequest(url, filename, "photo")
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendPhotoByID(photoId, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendPhoto?chat_id=%d&photo=%s&caption=%s%s",
		string(a),
		chatId,
		encode(photoId),
		encode(caption),
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendAudio(filename, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendAudio?chat_id=%d&caption=%s%s",
		string(a),
		chatId,
		encode(caption),
		parseOpts(opts...),
	)

	content := SendPostRequest(url, filename, "audio")
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendAudioByID(audioId, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendAudio?chat_id=%d&audio=%s&caption=%s%s",
		string(a),
		chatId,
		encode(audioId),
		encode(caption),
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendDocument(filename, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendDocument?chat_id=%d&caption=%s%s",
		string(a),
		chatId,
		encode(caption),
		parseOpts(opts...),
	)

	content := SendPostRequest(url, filename, "document")
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendDocumentByID(documentId, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendDocument?chat_id=%d&document=%s&caption=%s%s",
		string(a),
		chatId,
		encode(documentId),
		encode(caption),
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVideo(filename, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendVideo?chat_id=%d&caption=%s%s",
		string(a),
		chatId,
		encode(caption),
		parseOpts(opts...),
	)

	content := SendPostRequest(url, filename, "video")
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVideoByID(videoId, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf(
		"%ssendVideo?chat_id=%d&video=%s&caption=%s%s",
		string(a),
		chatId,
		encode(videoId),
		encode(caption),
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVideoNoteByID(videoId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVideoNote?chat_id=%d&video_note=%s",
		string(a),
		chatId,
		encode(videoId),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVoice(filename, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVoice?chat_id=%d&caption=%s%s",
		string(a),
		chatId,
		encode(caption),
		parseOpts(opts...),
	)

	content := SendPostRequest(url, filename, "voice")
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVoiceByID(voiceId, caption string, chatId int64, opts ...Option) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVoice?chat_id=%d&voice=%s%s",
		string(a),
		chatId,
		encode(voiceId),
		parseOpts(opts...),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendContact(phoneNumber, firstName, lastName string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendContact?chat_id=%d&phone_number=%s&first_name=%s&last_name=%s",
		string(a),
		chatId,
		encode(phoneNumber),
		encode(firstName),
		encode(lastName),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendStickerByID(stickerId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendSticker?chat_id=%d&sticker=%s",
		string(a),
		chatId,
		encode(stickerId),
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) SendChatAction(action ChatAction, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendChatAction?chat_id=%d&action=%s",
		string(a),
		chatId,
		action,
	)

	content := SendGetRequest(url)
	json.Unmarshal(content, &response)
	return
}

func (a Api) KeyboardButton(text string, requestContact, requestLocation bool) Button {
	return Button{text, requestContact, requestLocation}
}

func (a Api) KeyboardRow(buttons ...Button) (kbdRow KbdRow) {
	for _, button := range buttons {
		kbdRow = append(kbdRow, button)
	}

	return
}

func (a Api) KeyboardMarkup(resizeKeyboard, oneTimeKeyboard, selective bool, keyboardRows ...KbdRow) (kbd []byte) {
	keyboard := Keyboard{
		nil,
		resizeKeyboard,
		oneTimeKeyboard,
		selective,
	}

	for _, row := range keyboardRows {
		keyboard.Keyboard = append(keyboard.Keyboard, row)
	}

	kbd, _ = json.Marshal(keyboard)
	return
}

func (a Api) KeyboardRemove(selective bool) (kbdrmv []byte) {
	kbdrmv, _ = json.Marshal(KeyboardRemove{true, selective})
	return
}

// Returns a new inline keyboard button with the provided data.
func (a Api) InlineKbdBtn(text, url, callbackData string) InlineButton {
	return InlineButton{
		encode(text),
		url,
		callbackData,
	}
}

// Returns a new inline keyboard row with the given buttons.
func (a Api) InlineKbdRow(inlineButtons ...InlineButton) (inlineKbdRow InlineKbdRow) {
	return append(inlineKbdRow, inlineButtons...)
}

// Returns an inline keyboard object with the specified rows.
func (a Api) NewInlineKeyboard(rows ...InlineKbdRow) (ret InlineKeyboard) {
	ret.InlineKeyboard = append(ret.InlineKeyboard, rows...)
	return
}

// Returns a byte slice containing the inline keyboard json data.
func (a Api) InlineKbdMarkup(inlineKbdRows ...InlineKbdRow) (jsn []byte) {
	keyboard := a.NewInlineKeyboard(inlineKbdRows...)
	jsn, _ = json.Marshal(keyboard)
	return
}
