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
	"strconv"
)

// Sticker represents a sticker.
type Sticker struct {
	Thumb            *PhotoSize     `json:"thumb,omitempty"`
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
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Thumb       *PhotoSize     `json:"thumb,omitempty"`
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

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// NewStickerSetOptions contains the optional parameters used in the CreateNewStickerSet method.
type NewStickerSetOptions struct {
	StickerType  StickerSetType `query:"sticker_type"`
	MaskPosition MaskPosition   `query:"mask_position"`
}

// StickerType is a custom type for the various sticker types.
type StickerType string

// These are all the possible sticker types.
const (
	PNGSticker  StickerType = "png_sticker"
	TGSSticker              = "tgs_sticker"
	WEBMSticker             = "webm_sticker"
)

// StickerFile is a struct which contains info about sticker files.
type StickerFile struct {
	Type StickerType
	File InputFile
}

// SendSticker is used to send static .WEBP or animated .TGS stickers.
func (a API) SendSticker(stickerID string, chatID int64, opts *BaseOptions) (res APIResponseMessage, err error) {
	var vals = make(url.Values)

	vals.Set("sticker", stickerID)
	vals.Set("chat_id", strconv.FormatInt(chatID, 10))
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
func (a API) UploadStickerFile(userID int64, sticker StickerFile) (res APIResponseFile, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", strconv.FormatInt(userID, 10))
	return postFile[APIResponseFile](a.base, "uploadStickerFile", string(sticker.Type), sticker.File, InputFile{}, vals)
}

// CreateNewStickerSet is used to create a new sticker set owned by a user.
func (a API) CreateNewStickerSet(userID int64, name, title, emojis string, sticker StickerFile, opts *NewStickerSetOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", strconv.FormatInt(userID, 10))
	vals.Set("name", name)
	vals.Set("title", title)
	vals.Set("emojis", emojis)
	return postFile[APIResponseBase](a.base, "createNewStickerSet", string(sticker.Type), sticker.File, InputFile{}, addValues(vals, opts))
}

// AddStickerToSet is used to add a new sticker to a set created by the bot.
func (a API) AddStickerToSet(userID int64, name, emojis string, sticker StickerFile, opts *MaskPosition) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("user_id", strconv.FormatInt(userID, 10))
	vals.Set("name", name)
	vals.Set("emojis", emojis)
	return postFile[APIResponseBase](a.base, "addStickerToSet", string(sticker.Type), sticker.File, InputFile{}, addValues(vals, opts))
}

// SetStickerPositionInSet is used to move a sticker in a set created by the bot to a specific position.
func (a API) SetStickerPositionInSet(sticker string, position int) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("sticker", sticker)
	vals.Set("position", strconv.FormatInt(int64(position), 10))
	return get[APIResponseBase](a.base, "setStickerPositionInSet", vals)
}

// DeleteStickerFromSet is used to delete a sticker from a set created by the bot.
func (a API) DeleteStickerFromSet(sticker string) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("sticker", sticker)
	return get[APIResponseBase](a.base, "deleteStickerFromSet", vals)
}

// SetStickerSetThumb is used to set the thumbnail of a sticker set.
func (a API) SetStickerSetThumb(name string, userID int64, thumb InputFile) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	vals.Set("name", name)
	vals.Set("user_id", strconv.FormatInt(userID, 10))
	return postFile[APIResponseBase](a.base, "setStickerSetThumb", "thumb", thumb, InputFile{}, vals)
}
