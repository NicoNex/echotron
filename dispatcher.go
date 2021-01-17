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
	updates    chan []*Update
}

// NewDispatcher returns a new instance of the Dispatcher object;
// useful for polling telegram and dispatch every update to the
// corresponding Bot instance.
func NewDispatcher(token string, newBot NewBotFn) Dispatcher {
	d := Dispatcher{
		NewApi(token),
		make(map[int64]Bot),
		newBot,
		make(chan []*Update),
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
func (d *Dispatcher) Poll() {
	var timeout int
	var firstRun = true
	var lastUpdateId = -1

	for {
		response := d.api.GetUpdates(lastUpdateId+1, timeout)
		if response.Ok {
			if !firstRun {
				d.updates <- response.Result
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
}

func (d *Dispatcher) listen() {
	for uList := range d.updates {
		for _, update := range uList {
			var chatId int64

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

			if bot, ok := d.sessionMap[chatId]; ok {
				go bot.Update(update)
			}
		}
	}
}
