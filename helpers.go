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
	"os"
	"path/filepath"
	"strconv"
)

func check(r APIResponse) error {
	if b := r.Base(); !b.Ok {
		return &APIError{code: b.ErrorCode, desc: b.Description}
	}
	return nil
}

func processMedia(media, thumb InputFile) (im mediaEnvelope, cnt []content, err error) {
	switch {
	case media.id != "":
		im = mediaEnvelope{
			InputMedia: nil,
			media:      media.id,
			thumb:      "",
		}

	case media.path != "" && len(media.content) == 0:
		if media.content, media.path, err = readFile(media); err != nil {
			return
		}
		fallthrough

	case media.path != "" && len(media.content) > 0:
		cnt = append(cnt, content{media.path, media.path, media.content})
		im = mediaEnvelope{
			InputMedia: nil,
			media:      fmt.Sprintf("attach://%s", media.path),
			thumb:      "",
		}
	}

	switch {
	case thumb.path != "" && len(thumb.content) == 0:
		if thumb.content, thumb.path, err = readFile(thumb); err != nil {
			return
		}
		fallthrough

	case thumb.path != "" && len(thumb.content) > 0:
		cnt = append(cnt, content{thumb.path, thumb.path, thumb.content})
		im.thumb = fmt.Sprintf("attach://%s", thumb.path)
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

func sendFile(file, thumb InputFile, url, fileType string) (res []byte, err error) {
	var cnt []content

	if file.id != "" {
		url = fmt.Sprintf("%s&%s=%s", url, fileType, file.id)
	} else if c, e := toContent(fileType, file); e == nil {
		cnt = append(cnt, c)
	} else {
		err = e
	}

	if c, e := toContent("thumb", thumb); e == nil {
		cnt = append(cnt, c)
	} else {
		err = e
	}

	if len(cnt) > 0 {
		res, err = sendPostRequest(url, cnt...)
	} else {
		res, err = sendGetRequest(url)
	}
	return
}

func sendMediaFiles(url string, editSingle bool, files ...InputMedia) (res []byte, err error) {
	var (
		med []mediaEnvelope
		cnt []content
		jsn []byte
	)

	for _, file := range files {
		var im mediaEnvelope
		var cntArr []content

		media := file.media()
		thumb := file.thumb()

		im, cntArr, err = processMedia(media, thumb)
		if err != nil {
			return
		}

		im.InputMedia = file

		med = append(med, im)
		cnt = append(cnt, cntArr...)
	}

	if editSingle {
		jsn, err = json.Marshal(med[0])
	} else {
		jsn, err = json.Marshal(med)
	}

	if err != nil {
		return
	}

	url = fmt.Sprintf("%s&media=%s", url, jsn)

	if len(cnt) > 0 {
		return sendPostRequest(url, cnt...)
	}

	return sendGetRequest(url)
}

func serializePerms(permissions ChatPermissions) (string, error) {
	perm, err := json.Marshal(PermissionOptions{permissions})
	if err != nil {
		return "", err
	}

	return string(perm), nil
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

func get[T APIResponse](base, endpoint string, vals url.Values) (res T, err error) {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return res, err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}
	err = check(res)
	return
}

func postFile[T APIResponse](base, endpoint, fileType string, file, thumb InputFile, vals url.Values) (res T, err error) {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return res, err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}

	cnt, err := sendFile(file, thumb, url, fileType)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
	return
}

func postMedia[T APIResponse](base, endpoint string, editSingle bool, vals url.Values, files ...InputMedia) (res T, err error) {
	url, err := url.JoinPath(base, endpoint)
	if err != nil {
		return res, err
	}

	if vals != nil {
		if queries := vals.Encode(); queries != "" {
			url = fmt.Sprintf("%s?%s", url, queries)
		}
	}

	cnt, err := sendMediaFiles(url, editSingle, files...)
	if err != nil {
		return res, err
	}

	if err = json.Unmarshal(cnt, &res); err != nil {
		return
	}

	err = check(res)
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
