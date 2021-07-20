/*
 * Echotron
 * Copyright (C) 2021  The Echotron Devs
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
	"encoding/json"
	"fmt"
)

// Game represents a game.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    Animation       `json:"animation,omitempty"`
}

// CallbackGame is a placeholder, currently holds no information.
type CallbackGame struct{}

// GameHighScore represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

// GameScoreOptions contains the optional parameters used in API.SetGameScore method.
type GameScoreOptions struct {
	Force              bool `query:"force"`
	DisableEditMessage bool `query:"disable_edit_message"`
}

// SendGame is used to send a Game.
func (a API) SendGame(gameShortName string, chatID int64, opts *BaseOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssendGame?game_short_name=%s&chat_id=%d&%s",
		a.base,
		encode(gameShortName),
		chatID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// SetGameScore is used to set the score of the specified user in a game.
func (a API) SetGameScore(userID int64, score int, msgID MessageIDOptions, opts *GameScoreOptions) (APIResponseMessage, error) {
	var res APIResponseMessage
	var url = fmt.Sprintf(
		"%ssetGameScore?user_id=%d&score=%d&%s&%s",
		a.base,
		userID,
		score,
		querify(msgID),
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}

// GetGameHighScores is used to get data for high score tables.
func (a API) GetGameHighScores(userID int64, opts MessageIDOptions) (APIResponseGameHighScore, error) {
	var res APIResponseGameHighScore
	var url = fmt.Sprintf(
		"%sgetGameHighScores?user_id=%d&%s",
		a.base,
		userID,
		querify(opts),
	)

	cnt, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}

	return res, json.Unmarshal(cnt, &res)
}
