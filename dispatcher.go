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

import (
	"time"
	"sync"
)


type Bot interface {
	Update(*Update)
}


/* struct type representing the active chat session with a user */
type session struct {
	bot Bot
	timestamp int64
}


var sessionMap 	map[int64]*session
var mutex 		sync.Mutex


/* routine that disposes all sessions been inactive for one hour or more */
func garbageCollector() {
	for {
		mutex.Lock()
		for key, _ := range sessionMap {
			if time.Now().Unix() - sessionMap[key].timestamp > 3600 {
				go delete(sessionMap, key)
			}
		}
		mutex.Unlock()
		time.Sleep(time.Minute)
	}
}


func RunDispatcher(token string, gc bool, newBot func(string, int64) Bot) {
	sessionMap = make(map[int64]*session)
	
	var lastUpdateId int = -1;
	var firstRun bool = true
	var chatId int64
	var response APIResponse

	engine := NewEngine(token)

	if gc {
		go garbageCollector()
	}

	for {
		response = engine.GetResponse(lastUpdateId + 1, 120)

		if response.Ok && len(response.Result) > 0 {
			for _, update := range response.Result {
				lastUpdateId = update.ID
				chatId = update.Message.Chat.ID

				mutex.Lock()
				// if the chatId we're looking for is not in SessionMap
				// we make a new instance of session
				if _, isIn := sessionMap[chatId]; !isIn {
					sessionMap[chatId] = &session{
						bot: newBot(token, chatId),
						timestamp: time.Now().Unix(),
					}
				}

				if !firstRun {
					go func() {
						sessionMap[chatId].bot.Update(update)
						sessionMap[chatId].timestamp = time.Now().Unix()
					}()
				}
				mutex.Unlock()

			}
		}
		firstRun = false
	}
}
