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


var sessionMap map[int64]Bot
var engine Engine


// DelSession deletes the Bot instance, seen as a session, from the
// map with all of them.
func DelSession(chatId int64) {
	delete(sessionMap, chatId)
}


// AddSession allows to create a new Bot instance from within an active one.
func AddSession(chatId int64, newBot func(engine Engine, chatId int64) Bot) {
	if _, isIn := sessionMap[chatId]; !isIn {
		sessionMap[chatId] = newBot(engine, chatId)
	}
}


// RunDispatcher is echotron's entry point.
// It uses the bot token to initialise the engine used to communicate
// with Telegram servers, and the newBot function to get an instance
// of a user-defined struct that implements the Bot interface.
func RunDispatcher(token string, newBot func(engine Engine, chatId int64) Bot) {
	var timeout int
	var chatId int64
	var response APIResponse

	var firstRun = true
	var lastUpdateId = -1

	sessionMap = make(map[int64]Bot)
	engine = NewEngine(token)

	for {
		response = engine.GetResponse(lastUpdateId+1, timeout)
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

				if _, isIn := sessionMap[chatId]; !isIn {
					sessionMap[chatId] = newBot(engine, chatId)
				}

				if !firstRun {
					go sessionMap[chatId].Update(update)
				}

			}
		}

		if firstRun {
			firstRun = false
			timeout = 120
		}
	}
}
