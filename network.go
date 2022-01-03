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
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

// content is a struct which contains a file's name, its type and its data.
type content struct {
	fname string
	ftype string
	fdata []byte
}

// sendGetRequest is used to send an HTTP GET request.
func sendGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// sendPostRequest is used to send an HTTP POST request.
func sendPostRequest(url string, files ...content) ([]byte, error) {
	var buf = new(bytes.Buffer)
	var w = multipart.NewWriter(buf)

	for _, f := range files {
		part, err := w.CreateFormFile(f.ftype, filepath.Base(f.fname))
		if err != nil {
			return []byte{}, err
		}
		part.Write(f.fdata)
	}

	w.Close()

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Content-Type", w.FormDataContentType())

	var client = new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	cnt, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return cnt, nil
}

// sendPostForm is used to send an "application/x-www-form-urlencoded" through an HTTP POST request.
func sendPostForm(reqURL string, keyVals map[string]string) ([]byte, error) {
	var form = make(url.Values)

	for k, v := range keyVals {
		form.Add(k, v)
	}

	request, err := http.NewRequest("POST", reqURL, strings.NewReader(form.Encode()))
	if err != nil {
		return []byte{}, err
	}
	request.PostForm = form
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	var client http.Client

	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}
