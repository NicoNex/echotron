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
)

// Sticker represents a sticker.
type Sticker struct {
	Thumb            *PhotoSize     `json:"thumb,omitempty"`
	MaskPosition     *MaskPosition  `json:"mask_position,omitempty"`
	FileUniqueID     string         `json:"file_unique_id"`
	SetName          string         `json:"set_name,omitempty"`
	FileID           string         `json:"file_id"`
	Emoji            string         `json:"emoji,omitempty"`
	PremiumAnimation File           `json:"premium_animation,omitempty"`
	FileSize         int            `json:"file_size,omitempty"`
	Width            int            `json:"width"`
	Height           int            `json:"height"`
	IsVideo          bool           `json:"is_video"`
	IsAnimated       bool           `json:"is_animated"`
	CustomEmojiID    string         `json:"custom_emoji_id,omitempty"`
	Type             StickerSetType `json:"type"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Thumb       *PhotoSize     `json:"thumb,omitempty"`
	Title       string         `json:"title"`
	Name        string         `json:"name"`
	Stickers    []Sticker      `json:"stickers"`
	IsAnimated  bool           `json:"is_animated"`
	IsVideo     bool           `json:"is_video"`
	StickerType StickerSetType `json:"sticker_type"`
}

// StickerType represents the type of a sticker or of the entire set
type StickerSetType string

const (
	RegularStickerType     StickerType = "regular"
	MaskStickerType                    = "mask"
	CustomEmojiStickerType             = "custom_emoji"
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
	var url = fmt.Sprintf(
		"%ssendSticker?chat_id=%d&sticker=%s&%s",
		a.base,
		chatID,
		encode(stickerID),
		querify(opts),
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

// GetStickerSet is used to get a sticker set.
func (a API) GetStickerSet(name string) (res APIResponseStickerSet, err error) {
	var url = fmt.Sprintf(
		"%sgetStickerSet?name=%s",
		a.base,
		encode(name),
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
	var url = fmt.Sprintf(
		"%suploadStickerFile?user_id=%d",
		a.base,
		userID,
	)

	cnt, err := sendFile(sticker.File, InputFile{}, url, string(sticker.Type))
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

// CreateNewStickerSet is used to create a new sticker set owned by a user.
func (a API) CreateNewStickerSet(userID int64, name, title, emojis string, sticker StickerFile, opts *NewStickerSetOptions) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%screateNewStickerSet?user_id=%d&name=%s&title=%s&emojis=%s&%s",
		a.base,
		userID,
		encode(name),
		encode(title),
		encode(emojis),
		querify(opts),
	)

	cnt, err := sendFile(sticker.File, InputFile{}, url, string(sticker.Type))
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

// AddStickerToSet is used to add a new sticker to a set created by the bot.
func (a API) AddStickerToSet(userID int64, name, emojis string, sticker StickerFile, opts *MaskPosition) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%saddStickerToSet?user_id=%d&name=%s&emojis=%s&%s",
		a.base,
		userID,
		encode(name),
		encode(emojis),
		querify(opts),
	)

	cnt, err := sendFile(sticker.File, InputFile{}, url, string(sticker.Type))
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

// SetStickerPositionInSet is used to move a sticker in a set created by the bot to a specific position.
func (a API) SetStickerPositionInSet(sticker string, position int) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%ssetStickerPositionInSet?sticker=%s&position=%d",
		a.base,
		encode(sticker),
		position,
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

// DeleteStickerFromSet is used to delete a sticker from a set created by the bot.
func (a API) DeleteStickerFromSet(sticker string) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%sdeleteStickerFromSet?sticker=%s",
		a.base,
		encode(sticker),
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

// SetStickerSetThumb is used to set the thumbnail of a sticker set.
func (a API) SetStickerSetThumb(name string, userID int64, thumb InputFile) (res APIResponseBase, err error) {
	var url = fmt.Sprintf(
		"%ssetStickerSetThumb?name=%s&user_id=%d",
		a.base,
		encode(name),
		userID,
	)

	cnt, err := sendFile(thumb, InputFile{}, url, "thumb")
	if err != nil {
		return
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}
