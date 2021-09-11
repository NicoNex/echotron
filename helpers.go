/*
 * Echotron
 * Copyright (C) 2018-2021  The Echotron Devs
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
)

func encode(s string) string {
	return url.QueryEscape(s)
}

func parseInputMedia(media, thumb InputFile) (im mediaEnvelope, cnt []content, err error) {
	switch {
	case media.id != "":
		im = mediaEnvelope{media.id, "", nil}

	case media.path != "" && len(media.content) == 0:
		media.content, media.path, err = readFile(media)
		if err != nil {
			return
		}
		fallthrough

	case media.path != "" && len(media.content) > 0:
		cnt = append(cnt, content{media.path, media.path, media.content})
		im = mediaEnvelope{fmt.Sprintf("attach://%s", media.path), "", nil}
	}

	switch {
	case thumb.id != "":
		im.thumb = thumb.id

	case thumb.path != "" && len(thumb.content) == 0:
		thumb.content, thumb.path, err = readFile(thumb)
		if err != nil {
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

func sendFile(file InputFile, url, fileType string) (res []byte, err error) {
	switch {
	case file.id != "":
		res, err = sendGetRequest(fmt.Sprintf("%s&%s=%s", url, fileType, file.id))

	case file.path != "" && len(file.content) == 0:
		file.content, file.path, err = readFile(file)
		if err != nil {
			return
		}
		fallthrough

	case file.path != "" && len(file.content) > 0:
		res, err = sendPostRequest(url, content{file.path, fileType, file.content})
	}

	if err != nil {
		return
	}

	return res, nil
}

func sendMediaFiles(url string, isSingleFile bool, files ...InputMedia) (res []byte, err error) {
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

		im, cntArr, err = parseInputMedia(media, thumb)
		if err != nil {
			return
		}

		im.InputMedia = file

		med = append(med, im)
		cnt = append(cnt, cntArr...)
	}

	if isSingleFile {
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
	} else {
		return sendGetRequest(url)
	}
}

func serializePerms(permissions ChatPermissions) (string, error) {
	perm, err := json.Marshal(PermissionOptions{permissions})
	if err != nil {
		return "", err
	}

	return string(perm), nil
}

func toInputMedia(media []GroupableInputMedia) (ret []InputMedia) {
	ret = make([]InputMedia, len(media))

	for i, v := range media {
		ret[i] = v
	}

	return ret
}
