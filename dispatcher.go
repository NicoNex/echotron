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
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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
	updates    chan *Update
}

// NewDispatcher returns a new instance of the Dispatcher object;
// useful for polling telegram and dispatch every update to the
// corresponding Bot instance.
func NewDispatcher(token string, newBot NewBotFn) Dispatcher {
	d := Dispatcher{
		NewApi(token),
		make(map[int64]Bot),
		newBot,
		make(chan *Update),
	}
	go d.listen()
	return d
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

// Poll starts the polling loop so that the dispatcher calls the function Update
// upon receiving any update from Telegram.
func (d *Dispatcher) Poll() error {
	var timeout int
	var firstRun = true
	var lastUpdateId = -1

	// deletes webhook if present to run in long polling mode
	response, err := d.api.DeleteWebhook()
	if err != nil {
		return err
	} else if !response.Ok {
		return errors.New("could not disable webhook, running in long polling mode is not possible.")
	}

	for {
		response, err := d.api.GetUpdates(lastUpdateId+1, timeout)
		if err != nil {
			return err
		} else if response.Ok {
			if !firstRun {
				for _, u := range response.Result {
					d.updates <- u
				}
			}

			if l := len(response.Result); l > 0 {
				lastUpdateId = response.Result[l-1].ID
			}
		}

		if firstRun {
			firstRun = false
			timeout = 120
		}
	}

	return nil
}

func (d *Dispatcher) listen() {
	for update := range d.updates {
		var chatId int64

		if update.Message != nil {
			chatId = update.Message.Chat.ID
		} else if update.EditedMessage != nil {
			chatId = update.EditedMessage.Chat.ID
		} else if update.ChannelPost != nil {
			chatId = update.ChannelPost.Chat.ID
		} else if update.EditedChannelPost != nil {
			chatId = update.EditedChannelPost.Chat.ID
		} else if update.CallbackQuery != nil {
			chatId = update.CallbackQuery.Message.Chat.ID
		} else {
			continue
		}

		if _, isIn := d.sessionMap[chatId]; !isIn {
			d.sessionMap[chatId] = d.newBot(chatId)
		}

		if bot, ok := d.sessionMap[chatId]; ok {
			go bot.Update(update)
		}
	}
}

// ListenWebhook sets a webhook and listens for incoming updates
func (d *Dispatcher) ListenWebhook(url string, internalPort int) error {
	var response APIResponseUpdate

	response, err := d.api.SetWebhook(url)
	if err != nil {
		return err
	} else if response.Ok {
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			var update Update
			var reader io.ReadCloser
			var err error

			switch request.Header.Get("Content-Encoding") {
			case "gzip":
				reader, err = gzip.NewReader(request.Body)
				if err != nil {
					log.Println(err)
					return
				}
				defer reader.Close()

			default:
				reader = request.Body
			}

			if err = json.NewDecoder(reader).Decode(&update); err != nil {
				log.Println(err)
				return
			}

			d.updates <- &update
		})
		return http.ListenAndServe(fmt.Sprintf(":%d", internalPort), nil)
	} else {
		return errors.New(fmt.Sprintf("could not set webhook: %d %s", response.ErrorCode, response.Description))
	}

	return nil
}
