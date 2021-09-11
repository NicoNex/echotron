/*
 * Echotron
 * Copyright (C) 2021  The Echotron Devs
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
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []Sticker  `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// NewStickerSetOptions contains the optional parameters used in the CreateNewStickerSet method.
type NewStickerSetOptions struct {
	ContainsMasks bool         `query:"contains_masks"`
	MaskPosition  MaskPosition `query:"mask_position"`
}

// StickerType is a custom type for the various sticker types.
type StickerType string

// These are all the possible sticker types.
const (
	PNGSticker StickerType = "png_sticker"
	TGSSticker             = "tgs_sticker"
)

// StickerFile is a struct which contains info about sticker files.
type StickerFile struct {
	File InputFile
	Type StickerType
}

// SendSticker is used to send static .WEBP or animated .TGS stickers.
func (a API) SendSticker(stickerID string, chatID int64, opts *BaseOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendSticker?chat_id=%d&sticker=%s&%s",
		a.base,
		chatID,
		encode(stickerID),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetStickerSet is used to get a sticker set.
func (a API) GetStickerSet(name string) (APIResponseStickerSet, error) {
	var res APIResponseStickerSet
	var url = fmt.Sprintf(
		"%sgetStickerSet?name=%s",
		a.base,
		encode(name),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// UploadStickerFile is used to upload a .PNG file with a sticker for later use in
// CreateNewStickerSet and AddStickerToSet methods (can be used multiple times).
func (a API) UploadStickerFile(userID int64, sticker StickerFile) (APIResponseFile, error) {
	var res APIResponseFile
	var url = fmt.Sprintf(
		"%suploadStickerFile?user_id=%d",
		a.base,
		userID,
	)

	cnt, err := sendFile(sticker.File, InputFile{}, url, string(sticker.Type))
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// CreateNewStickerSet is used to create a new sticker set owned by a user.
func (a API) CreateNewStickerSet(userID int64, name, title, emojis string, sticker StickerFile, opts *NewStickerSetOptions) (APIResponseBase, error) {
	var res APIResponseBase
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
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// AddStickerToSet is used to add a new sticker to a set created by the bot.
func (a API) AddStickerToSet(userID int64, name, emojis string, sticker StickerFile, opts *MaskPosition) (APIResponseBase, error) {
	var res APIResponseBase
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
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetStickerPositionInSet is used to move a sticker in a set created by the bot to a specific position.
func (a API) SetStickerPositionInSet(sticker string, position int) (APIResponseBase, error) {
	var res APIResponseBase
	var url = fmt.Sprintf(
		"%ssetStickerPositionInSet?sticker=%s&position=%d",
		a.base,
		encode(sticker),
		position,
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// DeleteStickerFromSet is used to delete a sticker from a set created by the bot.
func (a API) DeleteStickerFromSet(sticker string) (APIResponseBase, error) {
	var res APIResponseBase
	var url = fmt.Sprintf(
		"%sdeleteStickerFromSet?sticker=%s",
		a.base,
		encode(sticker),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetStickerSetThumb is used to set the thumbnail of a sticker set.
func (a API) SetStickerSetThumb(name string, userID int64, thumb InputFile) (APIResponseBase, error) {
	var res APIResponseBase
	var url = fmt.Sprintf(
		"%ssetStickerSetThumb?name=%s&user_id=%d",
		a.base,
		encode(name),
		userID,
	)

	cnt, err := sendFile(thumb, InputFile{}, url, "thumb")
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}
