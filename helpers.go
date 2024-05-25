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
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

// content contains a file's name, its type and its data.
type content struct {
	fname string
	ftype string
	fdata []byte
}

func check(r APIResponse) error {
	if b := r.Base(); !b.Ok {
		return &APIError{code: b.ErrorCode, desc: b.Description}
	}
	return nil
}

func processMedia(media, thumbnail InputFile) (im mediaEnvelope, cnt []content, err error) {
	switch {
	case media.id != "":
		im = mediaEnvelope{
			media:     media.id,
			thumbnail: "",
		}

	case media.url != "":
		im = mediaEnvelope{
			media:     media.url,
			thumbnail: "",
		}

	case media.path != "" && len(media.content) == 0:
		if media.content, media.path, err = readFile(media); err != nil {
			return
		}
		fallthrough

	case media.path != "" && len(media.content) > 0:
		cnt = append(cnt, content{media.path, media.path, media.content})
		im = mediaEnvelope{
			media:     fmt.Sprintf("attach://%s", media.path),
			thumbnail: "",
		}
	}

	switch {
	case thumbnail.path != "" && len(thumbnail.content) == 0:
		if thumbnail.content, thumbnail.path, err = readFile(thumbnail); err != nil {
			return
		}
		fallthrough

	case thumbnail.path != "" && len(thumbnail.content) > 0:
		cnt = append(cnt, content{thumbnail.path, thumbnail.path, thumbnail.content})
		im.thumbnail = fmt.Sprintf("attach://%s", thumbnail.path)
	}

	return
}

func processSticker(sticker InputFile) (se stickerEnvelope, cnt []content, err error) {
	switch {
	case sticker.id != "":
		se.Sticker = sticker.id

	case sticker.url != "":
		se.Sticker = sticker.url

	case sticker.path != "" && len(sticker.content) == 0:
		if sticker.content, sticker.path, err = readFile(sticker); err != nil {
			return
		}
		fallthrough

	case sticker.path != "" && len(sticker.content) > 0:
		cnt = append(cnt, content{sticker.path, sticker.path, sticker.content})
		se.Sticker = fmt.Sprintf("attach://%s", sticker.path)
	}

	return
}

func readFile(im InputFile) (content []byte, path string, err error) {
	content, err = os.ReadFile(im.path)
	if err != nil {
		return
	}
	path = filepath.Base(im.path)

	return
}

func toContent(ftype string, f InputFile) (content, error) {
	if f.path != "" && len(f.content) == 0 {
		var err error
		if f.content, f.path, err = readFile(f); err != nil {
			return content{}, err
		}
	}

	return content{f.path, ftype, f.content}, nil
}

func toInputMedia(media []GroupableInputMedia) (ret []InputMedia) {
	ret = make([]InputMedia, len(media))

	for i, v := range media {
		ret[i] = v
	}

	return ret
}

func joinURL(base, endpoint string, vals url.Values) (addr string, err error) {
	addr, err = url.JoinPath(base, endpoint)
	if err != nil {
		return
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			addr = fmt.Sprintf("%s?%s", addr, queries)
		}
	}

	return
}

func itoa(i int64) string {
	return strconv.FormatInt(i, 10)
}

func ftoa(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func btoa(b bool) string {
	return strconv.FormatBool(b)
}
