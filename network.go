/*
 * Echotron
 * Copyright (C) 2019  Nicolò Santamaria
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package echotron

import (
	"io"
	"os"
	"log"
	"bytes"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"mime/multipart"
)

func SendGetRequest (url string) []byte {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	return content
}


func SendPostRequest (url string, filename string, filetype string) []byte {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Println(err)
		return []byte{}
	}


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
	defer response.Body.Close()
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}

	return content
}
