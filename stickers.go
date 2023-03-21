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

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// Sticker represents a sticker.
type Sticker struct {
	Thumbnail        *PhotoSize     `json:"thumbnail,omitempty"`
	MaskPosition     *MaskPosition  `json:"mask_position,omitempty"`
	Type             StickerSetType `json:"type"`
	FileUniqueID     string         `json:"file_unique_id"`
	SetName          string         `json:"set_name,omitempty"`
	FileID           string         `json:"file_id"`
	Emoji            string         `json:"emoji,omitempty"`
	CustomEmojiID    string         `json:"custom_emoji_id,omitempty"`
	PremiumAnimation File           `json:"premium_animation,omitempty"`
	FileSize         int            `json:"file_size,omitempty"`
	Width            int            `json:"width"`
	Height           int            `json:"height"`
	IsVideo          bool           `json:"is_video"`
	IsAnimated       bool           `json:"is_animated"`
	NeedsRepainting  bool           `json:"needs_repainting,omitempty"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Thumbnail   *PhotoSize     `json:"thumbnail,omitempty"`
	Title       string         `json:"title"`
	Name        string         `json:"name"`
	StickerType StickerSetType `json:"sticker_type"`
	Stickers    []Sticker      `json:"stickers"`
	IsAnimated  bool           `json:"is_animated"`
	IsVideo     bool           `json:"is_video"`
}

// StickerSetType represents the type of a sticker or of the entire set
type StickerSetType string

const (
	RegularStickerSet     StickerSetType = "regular"
	MaskStickerSet                       = "mask"
	CustomEmojiStickerSet                = "custom_emoji"
)

// StickerFormat is a custom type for the various sticker formats.
type StickerFormat string

// These are all the possible sticker formats.
const (
	StaticFormat   StickerFormat = "static"
	AnimatedFormat               = "animated"
	VideoFormat                  = "video"
)

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  MaskPoint `json:"point"`
	XShift float32   `json:"x_shift"`
	YShift float32   `json:"y_shift"`
	Scale  float32   `json:"scale"`
}

// MaskPoint is a custom type for the various part of face where a mask should be placed.
type MaskPoint string

// These are all the possible parts of the face for a mask.
const (
	ForeheadPoint MaskPoint = "forehead"
	EyesPoint               = "eyes"
	MouthPoint              = "mouth"
	ChinPoint               = "chin"
)

// NewStickerSetOptions contains the optional parameters used in the CreateNewStickerSet method.
type NewStickerSetOptions struct {
	StickerType     StickerSetType `query:"sticker_type"`
	NeedsRepainting bool           `query:"needs_repainting"`
}

// InputSticker is a struct which describes a sticker to be added to a sticker set.
type InputSticker struct {
	Sticker      InputFile     `json:"-"`
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     *[]string     `json:"keywords,omitempty"`
}

// stickerEnvelope is a generic struct for all the various structs under the InputSticker interface.
type stickerEnvelope struct {
	InputSticker
	Sticker string `json:"sticker"`
}

// SendSticker is used to send static .WEBP or animated .TGS stickers.
func (a API) SendSticker(stickerID string, chatID int64, opts *StickerOptions) (res APIResponseMessage, err error) {
	var vals = make(url.Values)

	vals.Set("sticker", stickerID)
	vals.Set("chat_id", itoa(chatID))
	return get[APIResponseMessage](a.base, "sendSticker", addValues(vals, opts))
}

// GetStickerSet is used to get a sticker set.
func (a API) GetStickerSet(name string) (res APIResponseStickerSet, err error) {
	var vals = make(url.Values)

	vals.Set("name", name)
	return get[APIResponseStickerSet](a.base, "getStickerSet", vals)
}

// GetCustomEmojiStickers is used to get information about custom emoji stickers by their identifiers.
func (a API) GetCustomEmojiStickers(customEmojiIDs ...string) (res APIResponseStickers, err error) {
	jsn, _ := json.Marshal(customEmojiIDs)

	var url = fmt.Sprintf(
		"%sgetCustomEmojiStickers?custom_emoji_ids=%s",
		a.base,
		jsn,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

// UploadStickerFile is used to upload a .PNG file with a sticker for later use in
// CreateNewStickerSet and AddStickerToSet methods (can be used multiple times).
func (a API) UploadStickerFile(userID int64, sticker InputFile, format StickerFormat) (res APIResponseFile, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", itoa(userID))
	vals.Set("sticker_format", string(format))
	return postFile[APIResponseFile](a.base, "uploadStickerFile", "sticker", sticker, InputFile{}, vals)
}

// CreateNewStickerSet is used to create a new sticker set owned by a user.
func (a API) CreateNewStickerSet(userID int64, name, title string, stickers []InputSticker, format StickerFormat, opts *NewStickerSetOptions) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", itoa(userID))
	vals.Set("name", name)
	vals.Set("title", title)
	vals.Set("sticker_format", string(format))
	return postStickers[APIResponseBool](a.base, "createNewStickerSet", addValues(vals, opts), stickers...)
}

// AddStickerToSet is used to add a new sticker to a set created by the bot.
func (a API) AddStickerToSet(userID int64, name string, sticker InputSticker) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", itoa(userID))
	vals.Set("name", name)
	return postStickers[APIResponseBool](a.base, "addStickerToSet", vals, sticker)
}

// SetStickerPositionInSet is used to move a sticker in a set created by the bot to a specific position.
func (a API) SetStickerPositionInSet(sticker string, position int) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("sticker", sticker)
	vals.Set("position", itoa(int64(position)))
	return get[APIResponseBase](a.base, "setStickerPositionInSet", vals)
}

// DeleteStickerFromSet is used to delete a sticker from a set created by the bot.
func (a API) DeleteStickerFromSet(sticker string) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("sticker", sticker)
	return get[APIResponseBase](a.base, "deleteStickerFromSet", vals)
}

// SetStickerEmojiList is used to change the list of emoji assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot.
func (a API) SetStickerEmojiList(sticker string, emojis []string) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	jsn, _ := json.Marshal(emojis)

	vals.Set("sticker", sticker)
	vals.Set("emoji_list", string(jsn))
	return get[APIResponseBool](a.base, "setStickerEmojiList", vals)
}

// SetStickerKeywords is used to change search keywords assigned to a regular or custom emoji sticker.
// The sticker must belong to a sticker set created by the bot.
func (a API) SetStickerKeywords(sticker string, keywords []string) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	jsn, _ := json.Marshal(keywords)

	vals.Set("sticker", sticker)
	vals.Set("keywords", string(jsn))
	return get[APIResponseBool](a.base, "setStickerKeywords", vals)
}

// SetStickerMaskPosition is used to change the mask position of a mask sticker.
// The sticker must belong to a sticker set that was created by the bot.
func (a API) SetStickerMaskPosition(sticker string, mask MaskPosition) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	jsn, _ := json.Marshal(mask)

	vals.Set("sticker", sticker)
	vals.Set("mask_position", string(jsn))
	return get[APIResponseBool](a.base, "setStickerMaskPosition", vals)
}

// SetStickerSetTitle is used to set the title of a created sticker set.
func (a API) SetStickerSetTitle(name, title string) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	vals.Set("name", name)
	vals.Set("title", title)
	return get[APIResponseBool](a.base, "setStickerSetTitle", vals)
}

// SetStickerSetThumbnail is used to set the thumbnail of a sticker set.
func (a API) SetStickerSetThumbnail(name string, userID int64, thumbnail InputFile) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("name", name)
	vals.Set("user_id", itoa(userID))
	return postFile[APIResponseBase](a.base, "setStickerSetThumbnail", "thumbnail", thumbnail, InputFile{}, vals)
}

// SetCustomEmojiStickerSetThumbnail is used to set the thumbnail of a custom emoji sticker set.
func (a API) SetCustomEmojiStickerSetThumbnail(name, emojiID string) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	vals.Set("name", name)
	vals.Set("custom_emoji_id", emojiID)
	return get[APIResponseBool](a.base, "setCustomEmojiStickerSetThumbnail", vals)
}

// DeleteStickerSet is used to delete a sticker set that was created by the bot.
func (a API) DeleteStickerSet(name string) (res APIResponseBool, err error) {
	var vals = make(url.Values)

	vals.Set("name", name)
	return get[APIResponseBool](a.base, "DeleteStickerSet", vals)
}

// GetForumTopicIconStickers is used to get custom emoji stickers, which can be used as a forum topic icon by any user.
func (a API) GetForumTopicIconStickers() (res APIResponseStickers, err error) {
	return get[APIResponseStickers](a.base, "getForumTopicIconStickers", nil)
}
