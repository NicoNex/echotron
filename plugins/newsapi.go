/*
*	 Echotron-GO
*    Copyright (C) 2018  Nicolò Santamaria
*
*    This program is free software: you can redistribute it and/or modify
*    it under the terms of the GNU General Public License as published by
*    the Free Software Foundation, either version 3 of the License, or
*    (at your option) any later version.
*
*    This program is distributed in the hope that it will be useful,
*    but WITHOUT ANY WARRANTY; without even the implied warranty of
*    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
*    GNU General Public License for more details.
*
*    You should have received a copy of the GNU General Public License
*    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package plugins

import (
		"fmt"
		"strings"
		"encoding/json"
		"../core"
		)

type article struct {
	Author string `json:"author"`
	Title string `json:"title"`
	Description string `json:"description"`
	Url string `json:"url"`
	UrlToImage string `json:"string"`
	PublishedAt string `json:"publishedAt"`
	Content string `json:"content"`
}


type NewsApiResponse struct {
	Status string `json:"status"`
	TotalResult uint `json:"totalResult"`
	Articles []article `json:"articles"`
}


type source struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	URL string `json:"url"`
	Category string `json:"category"`
	Language string `json:"language"`
	Country string `json:"country"`
}


type NewsApiSources struct {
	Status string `json:"status"`
	Sources []source `json:"sources"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}


type NewsApi struct {
	key string
	baseUrl string
	lastUpdate string
}


func NewNewsApi (key string) *NewsApi {
	api := new(NewsApi)
	api.key = key
	api.baseUrl = "https://newsapi.org/v2/"

	return api
}


func (api *NewsApi) GetUpdates (sources []string, oldest string) NewsApiResponse {
	var url string

	if len(sources) > 0 {
		url = fmt.Sprintf("%severything?apiKey=%s&sources=%s&from=%s", api.baseUrl, api.key, strings.Join(sources, ","), oldest)
	} else {
		url = fmt.Sprintf("%severything?apiKey=%s&from=%s", api.baseUrl, api.key, oldest)
	}

	content := core.SendGetRequest(url)
	var response NewsApiResponse

	json.Unmarshal(content, &response)

	return response
}


func (api *NewsApi) GetSources (language string) NewsApiSources {
	var url = fmt.Sprintf("%ssources?apiKey=%s&language=%s", api.baseUrl, api.key, language)

	content := core.SendGetRequest(url)
	var response NewsApiSources

	json.Unmarshal(content, &response)

	return response
}
