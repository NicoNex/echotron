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
	"encoding/json"
	"fmt"
	"strings"
)

type Api struct {
	url string
}

const (
	PARSE_MARKDOWN           = 1 << iota
	PARSE_HTML               = 1 << iota
	DISABLE_WEB_PAGE_PREVIEW = 1 << iota
	DISABLE_NOTIFICATION     = 1 << iota
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

// NewApi returns a new Api object.
func NewApi(token string) Api {
	return Api{
		url: fmt.Sprintf("https://api.telegram.org/bot%s/", token),
	}
}

// GetResponse returns the incoming updates from telegram.
func (a Api) GetUpdates(offset int, timeout int) (response APIResponseUpdate) {
	var url = fmt.Sprintf("%sgetUpdates?timeout=%d", a.url, timeout)

	if offset != 0 {
		url = fmt.Sprintf("%s&offset=%d", url, offset)
	}
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) GetChat(chatId int64) (response Chat) {
	var url = fmt.Sprintf("%sgetChat?chat_id=%d", a.url, chatId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) GetStickerSet(name string) (response StickerSet) {
	var url = fmt.Sprintf("%sgetStickerSet?name=%s", a.url, name)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)

	return
}

func (a Api) SendMessage(text string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendMessage?text=%s&chat_id=%d", a.url, strings.Replace(text, "\n", "%0A", -1), chatId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendMessageOptions(text string, chatId int64, options int) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendMessage?text=%s&chat_id=%d", a.url, strings.Replace(text, "\n", "%0A", -1), chatId)

	if options&PARSE_MARKDOWN != 0 {
		url += "&parse_mode=markdown"
	}

	if options&PARSE_HTML != 0 {
		url += "&parse_mode=html"
	}

	if options&DISABLE_WEB_PAGE_PREVIEW != 0 {
		url += "&disable_web_page_preview=true"
	}

	if options&DISABLE_NOTIFICATION != 0 {
		url += "&disable_notification=true"
	}

	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendMessageReply(text string, chatId int64, messageId int) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendMessage?text=%s&chat_id=%d&reply_to_message_id=%d", a.url, strings.Replace(text, "\n", "%0A", -1), chatId, messageId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendMessageReplyOptions(text string, chatId int64, messageId int, options int) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendMessage?text=%s&chat_id=%d&reply_to_message_id=%d", a.url, strings.Replace(text, "\n", "%0A", -1), chatId, messageId)

	if options&PARSE_MARKDOWN != 0 {
		url += "&parse_mode=markdown"
	}

	if options&PARSE_HTML != 0 {
		url += "&parse_mode=html"
	}

	if options&DISABLE_WEB_PAGE_PREVIEW != 0 {
		url += "&disable_web_page_preview=true"
	}

	if options&DISABLE_NOTIFICATION != 0 {
		url += "&disable_notification=true"
	}

	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendMessageWithKeyboard(text string, chatId int64, keyboard []byte) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendMessage?text=%s&chat_id=%d&parse_mode=markdown&reply_markup=%s", a.url, strings.Replace(text, "\n", "%0A", -1), chatId, keyboard)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) DeleteMessage(chatId int64, messageId int) (response APIResponseMessage) {
	var url = fmt.Sprintf("%sdeleteMessage?chat_id=%d&message_id=%d", a.url, chatId, messageId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendPhoto(filename string, chatId int64, caption string) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendPhoto?chat_id=%d&caption=%s", a.url, chatId, caption)
	var content = SendPostRequest(url, filename, "photo")

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendPhotoByID(photoId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendPhoto?chat_id=%d&photo=%s", a.url, chatId, photoId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendAudio(filename string, chatId int64, caption string) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendAudio?chat_id=%d&caption=%s", a.url, chatId, caption)
	var content = SendPostRequest(url, filename, "audio")

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendAudioByID(audioId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendAudio?chat_id=%d&audio=%s", a.url, chatId, audioId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendDocument(filename string, caption string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendDocument?chat_id=%d&caption=%s", a.url, chatId, caption)
	var content = SendPostRequest(url, filename, "document")

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendDocumentByID(documentId string, caption string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendDocument?chat_id=%d&document=%s&caption=%s", a.url, chatId, documentId, caption)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVideo(filename string, chatId int64, caption string) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVideo?chat_id=%d&caption=%s", a.url, chatId, caption)
	var content = SendPostRequest(url, filename, "video")

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVideoByID(videoId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVideo?chat_id=%d&video=%s", a.url, chatId, videoId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVideoNoteByID(videoId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVideoNote?chat_id=%d&video_note=%s", a.url, chatId, videoId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVoice(filename string, chatId int64, caption string) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVoice?chat_id=%d&caption=%s", a.url, chatId, caption)
	var content = SendPostRequest(url, filename, "voice")

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendVoiceByID(voiceId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendVoice?chat_id=%d&voice=%s", a.url, chatId, voiceId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendContact(phoneNumber string, firstName string, lastName string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendContact?chat_id=%d&phone_number=%s&first_name=%s&last_name=%s", a.url, chatId, phoneNumber, firstName, lastName)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendStickerByID(stickerId string, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendSticker?chat_id=%d&sticker=%s", a.url, chatId, stickerId)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) SendChatAction(action ChatAction, chatId int64) (response APIResponseMessage) {
	var url = fmt.Sprintf("%ssendChatAction?chat_id=%d&action=%s", a.url, chatId, action)
	var content = SendGetRequest(url)

	json.Unmarshal(content, &response)
	return
}

func (a Api) KeyboardButton(text string, requestContact bool, requestLocation bool) Button {
	return Button{
		text,
		requestContact,
		requestLocation,
	}
}

func (a Api) KeyboardRow(buttons ...Button) (kbdRow KbdRow) {
	for _, button := range buttons {
		kbdRow = append(kbdRow, button)
	}

	return
}

func (a Api) KeyboardMarkup(resizeKeyboard bool, oneTimeKeyboard bool, selective bool, keyboardRows ...KbdRow) (kbd []byte) {
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
	kbdrmv, _ = json.Marshal(KeyboardRemove{
		true,
		selective,
	})

	return
}

// Returns a new inline keyboard button with the provided data.
func (a Api) InlineKbdBtn(text string, url string, callbackData string) InlineButton {
	return InlineButton{
		text,
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
