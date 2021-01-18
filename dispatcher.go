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
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Bot is the interface that must be implemented by your definition of
// the struct thus it represent each open session with a user on Telegram.
type Bot interface {
	// Update will be called upon receiving any update from Telegram.
	Update(*Update)
}

type NewBotFn func(chatId int64) Bot

type Dispatcher struct {
	api        Api
	sessionMap map[int64]Bot
	newBot     NewBotFn
}

// NewDispatcher returns a new instance of the Dispatcher object;
// useful for polling telegram and dispatch every update to the
// corresponding Bot instance.
func NewDispatcher(token string, newBot NewBotFn) Dispatcher {
	return Dispatcher{
		NewApi(token),
		make(map[int64]Bot),
		newBot,
	}
}

// DelSession deletes the Bot instance, seen as a session, from the
// map with all of them.
func (d *Dispatcher) DelSession(chatId int64) {
	delete(d.sessionMap, chatId)
}

// AddSession allows to arbitrarily create a new Bot instance.
func (d *Dispatcher) AddSession(chatId int64) {
	if _, isIn := d.sessionMap[chatId]; !isIn {
		d.sessionMap[chatId] = d.newBot(chatId)
	}
}

// Run starts the polling loop and calls the function Update
// upon receiving any update from Telegram.
func (d *Dispatcher) Run() {
	var timeout int
	var chatId int64
	var response APIResponseUpdate

	var firstRun = true
	var lastUpdateId = -1

	// deletes webhook if present to run in long polling mode
	response = d.api.DeleteWebhook()
	if !response.Ok {
		log.Fatalln("Could not disable webhook, running in long polling mode is not possible.")
	}

	for {
		response = d.api.GetUpdates(lastUpdateId+1, timeout)
		if response.Ok {
			for _, update := range response.Result {
				lastUpdateId = update.ID

				if update.Message != nil {
					chatId = update.Message.Chat.ID
				} else if update.EditedMessage != nil {
					chatId = update.EditedMessage.Chat.ID
				} else if update.ChannelPost != nil {
					chatId = update.ChannelPost.Chat.ID
				} else if update.EditedChannelPost != nil {
					chatId = update.EditedChannelPost.Chat.ID
				} else {
					continue
				}

				if _, isIn := d.sessionMap[chatId]; !isIn {
					d.sessionMap[chatId] = d.newBot(chatId)
				}

				if !firstRun {
					if bot, ok := d.sessionMap[chatId]; ok {
						go bot.Update(update)
					}
				}
			}
		}

		if firstRun {
			firstRun = false
			timeout = 120
		}
	}
}

// RunWithWebhook sets webhook and starts a server listening for incoming updates
func (d *Dispatcher) RunWithWebhook(url string, internalPort int) {
	var chatId int64
	var response APIResponseUpdate

	response = d.api.SetWebhook(url)
	if response.Ok {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			var update Update

			var reader io.ReadCloser
			var err error
			switch request.Header.Get("Content-Encoding") {
			case "gzip":
				reader, err = gzip.NewReader(request.Body)
				if err != nil {
					log.Println(err)
				}
				defer reader.Close()
			default:
				reader = request.Body
			}

			err = json.NewDecoder(reader).Decode(&update)
			if err != nil {
				log.Println(err)
			}

			if update.Message != nil {
				chatId = update.Message.Chat.ID
			} else if update.EditedMessage != nil {
				chatId = update.EditedMessage.Chat.ID
			} else if update.ChannelPost != nil {
				chatId = update.ChannelPost.Chat.ID
			} else if update.EditedChannelPost != nil {
				chatId = update.EditedChannelPost.Chat.ID
			}

			if _, isIn := d.sessionMap[chatId]; !isIn {
				d.sessionMap[chatId] = d.newBot(chatId)
			}

			if bot, ok := d.sessionMap[chatId]; ok {
				go bot.Update(&update)
			}

		})
		err := http.ListenAndServe(":"+strconv.Itoa(internalPort), nil)
		log.Fatalln(err)
	}
}
