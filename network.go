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
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func SendGetRequest(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	return content
}

func SendPostRequest(url string, filename string, filetype string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(filetype, filepath.Base(file.Name()))
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	io.Copy(part, file)
	writer.Close()
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	return content
}

func SendPostForm(destUrl string, keysArray, valuesArray []string) ([]byte, error) {
	form := url.Values{}
	if len(keysArray) != len(valuesArray) {
		return nil, fmt.Errorf("%s", "Number of keys and values mismatch!")
	}
	for i := 0; i < len(keysArray); i++ {
		form.Add(keysArray[i], valuesArray[i])
	}

	request, err := http.NewRequest("POST", destUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return []byte{}, err
	}
	request.PostForm = form
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return content, nil
}
