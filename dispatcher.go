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


type Bot interface {
	Update(*Update)
}


var sessionMap map[int64]Bot


func DeleteSession(chatId int64) {
	delete(sessionMap, chatId)
}


func RunDispatcher(token string, newBot func(string, int64) Bot) {
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
				chatId = update.Message.Chat.ID

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
