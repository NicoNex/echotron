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
	"errors"
)


type timer struct {
	timestamp int64
	lapseTime int64
	callback func()
}


var timerMap map[int64]map[string]*timer


func AddTimer(chatId int64, lapse int64, name string, callback func()) {
	if timerMap[chatId] == nil {
		timerMap[chatId] = make(map[string]*timer)
	}

	timerMap[chatId][name] = &timer{
		time.Now().Unix(),
		lapse,
		callback,
	}
}


func SetTimerLapse(chatId int64, lapse int64, name string) error {
	timer, ok := timerMap[chatId][name]
	if !ok {
		return errors.New("Error: cannot find timer")
	}
	timer.lapseTime = lapse
	return nil
}


func DelTimer(chatId int64, name string) {
	if timer, ok := timerMap[chatId]; ok {
		delete(timer, name)
	}
}


func timerRoutine() {
	for {
		for _, m := range timerMap {
			go func() {
				for _, t := range m {
					if time.Now().Unix() - t.timestamp >= t.lapseTime {
						t.timestamp = time.Now().Unix()
						go t.callback()
					}
				}
			}()
		}
		time.Sleep(time.Second)
	}
}


func init() {
	timerMap = make(map[int64]map[string]*timer)

	go timerRoutine()
}
