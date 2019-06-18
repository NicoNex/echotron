/*
 * Echotron-GO
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


// Bot is the interface that must be implemented by your definition of
// the struct thus representing each open session with a user on Telegram.
type Bot interface {
	// Update will be called upon receiving any update from Telegram
	// belonging to the associated session.
	Update(*Update)
}


var sessionMap map[int64]Bot


// DelSession deletes the Bot instance, seen as a session, from the
// map with all of them.
func DelSession(chatId int64) {
	delete(sessionMap, chatId)
}


// AddSession allows to create a new Bot instance from within an active one.
func AddSession(token string, chatId int64, newBot func(token string, chatId int64) Bot) {
	if _, isIn := sessionMap[chatId]; !isIn {
		sessionMap[chatId] = newBot(token, chatId)
	}
}


// RunDispatcher is echotron's entry point.
// It uses the bot token to initialise the engine used to communicate
// with Telegram servers, and the newBot function that must return an
// object that implements the Bot interface.
func RunDispatcher(token string, newBot func(token string, chatId int64) Bot) {
	var lastUpdateId int = -1;
	var firstRun bool = true
	var chatId int64
	var response APIResponse

	sessionMap = make(map[int64]Bot)
	engine := NewEngine(token)

	for {
		response = engine.GetResponse(lastUpdateId + 1, 120)

		if response.Ok && len(response.Result) > 0 {
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
					sessionMap[chatId] = newBot(token, chatId)
				}

				if !firstRun {
					go sessionMap[chatId].Update(update)
				}

			}
		}
		firstRun = false
	}
}
