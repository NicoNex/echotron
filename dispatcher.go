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

package echotron

/* struct type representing the active chat session with a user */
type session struct {
	bot *Bot
	timestamp int64
	chatId int64
	isTimed bool
	lapse int64
}

var sessionMap map[int64]*session 	// table of all the *session instances


func RunDispatcher (token string) {
	sessionMap = make(map[int64]*session)
	
	engine := core.NewEngine(token)
	
	var lastUpdateId int
	var firstRun bool = true
	var chatId int64

	var response core.APIResponse

	for {

		if !firstRun {
			response = engine.GetResponse(lastUpdateId + 1, 120)
		} else {
			response = engine.GetResponse(0, 0)
		}

		if response.Ok && len(response.Result) > 0 {
			for _, update := range response.Result {
				lastUpdateId = update.ID
				
				chatId = update.Message.Chat.ID
				if _, isIn := sessionMap[chatId]; !isIn {
					sessionMap[chatId] = new(session)
					sessionMap[chatId].bot = NewBot(token, chatId)
					sessionMap[chatId].chatId = chatId
				}

				if !firstRun {
					go func() {
						sessionMap[chatId].bot.Update(update)
						sessionMap[chatId].timestamp = time.Now().Unix()
					}()
				}
			}
		}
		firstRun = false
	}
}

