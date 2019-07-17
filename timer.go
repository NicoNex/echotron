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
	"fmt"
	"time"
)


type timer struct {
	timestamp int64
	lapseTime int64
	callback func()
}


var timerMap map[int64]map[string]*timer


// AddTimer creates a timer associated with a name to the current
// bot session that calls a callback once any given amounth of time.
func AddTimer(chatId int64, name string, callback func(), lapse int64) {
	if timerMap[chatId] == nil {
		timerMap[chatId] = make(map[string]*timer)
	}

	timerMap[chatId][name] = &timer{
		time.Now().Unix(),
		lapse,
		callback,
	}
}


// SetTimerLapse sets the lapse time between each call to the callback.
func SetTimerLapse(chatId int64, lapse int64, name string) error {
	timer, ok := timerMap[chatId][name]
	if !ok {
		return fmt.Errorf("Error: cannot find timer %s for instance %d", name, chatId)
	}
	timer.lapseTime = lapse
	return nil
}


// ResetTimer sets the current timestamp associated with the last call
// to the current epoch time in Unix format.
func ResetTimer(chatId int64, name string) error {
	timer, ok := timerMap[chatId][name]
	if !ok {
		return fmt.Errorf("Error: cannot find timer %s for instance %d", name, chatId)
	}
	timer.timestamp = time.Now().Unix()
	return nil
}


// DelTimer deletes the specified timer from the routine.
func DelTimer(chatId int64, name string) {
	if timer, ok := timerMap[chatId]; ok {
		delete(timer, name)
	}
}


func timerRoutine() {
	for {
		for _, m := range timerMap {
			go func(m map[string]*timer) {
				for _, t := range m {
					if time.Now().Unix() - t.timestamp >= t.lapseTime {
						t.timestamp = time.Now().Unix()
						go t.callback()
					}
				}
			}(m)
		}
		time.Sleep(time.Second)
	}
}


func init() {
	timerMap = make(map[int64]map[string]*timer)

	go timerRoutine()
}
