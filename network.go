/*
 * Echotron
 * Copyright (C) 2019  Nicolò Santamaria
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
	"strings"
)

func SendGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}

func SendPostRequest(url, filetype string, content []byte) ([]byte, error) {
	var buf = bytes.NewBuffer(content)
	var w = multipart.NewWriter(&buf)
	defer w.Close()

	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Content-Type", writer.FormDataContentType())
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

func SendPostForm(url string, keyVals map[string]string) ([]byte, error) {
	var form url.Values

	for k, v := range keyVals {
		form.Add(k, v)
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
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
