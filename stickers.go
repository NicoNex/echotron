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
	FileID       string       `json:"file_id"`
	FileUniqueID string       `json:"file_unique_id"`
	Width        int          `json:"width"`
	Height       int          `json:"height"`
	IsAnimated   bool         `json:"is_animated"`
	Thumb        *PhotoSize   `json:"thumb,omitempty"`
	Emoji        string       `json:"emoji,omitempty"`
	SetName      string       `json:"set_name,omitempty"`
	MaskPosition MaskPosition `json:"mask_position"`
	FileSize     int          `json:"file_size,omitempty"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"sticker"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// SendSticker is used to send static .WEBP or animated .TGS stickers.
func (a API) SendSticker(stickerID string, chatID int64, opts *BaseOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendSticker?chat_id=%d&sticker=%s&%s",
		string(a),
		chatID,
		encode(stickerID),
		querify(opts),
	)

	content, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}

// GetStickerSet is used to get a sticker set.
func (a API) GetStickerSet(name string) (APIResponseStickerSet, error) {
	var res APIResponseStickerSet
	var url = fmt.Sprintf("%sgetStickerSet?name=%s", string(a), encode(name))

	content, err := sendGetRequest(url)
	if err != nil {
		return APIResponseStickerSet{}, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}
