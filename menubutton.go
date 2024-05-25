/*
 * Echotron
 * Copyright (C) 2022 The Echotron Devs
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

// MenuButtonType is a custom type for the various MenuButton*'s Type field.
type MenuButtonType string

// These are all the possible types for the various MenuButton*'s Type field.
const (
	MenuButtonTypeCommands MenuButtonType = "commands"
	MenuButtonTypeWebApp                  = "web_app"
	MenuButtonTypeDefault                 = "default"
)

// MenuButton is a unique type for MenuButtonCommands, MenuButtonWebApp and MenuButtonDefault
type MenuButton struct {
	WebApp *WebAppInfo    `json:"web_app,omitempty"`
	Type   MenuButtonType `json:"type"`
	Text   string         `json:"text,omitempty"`
}

// SetChatMenuButtonOptions contains the optional parameters used by the SetChatMenuButton method.
type SetChatMenuButtonOptions struct {
	MenuButton MenuButton `query:"menu_button"`
	ChatID     int64      `query:"chat_id"`
}

// GetChatMenuButtonOptions contains the optional parameters used by the GetChatMenuButton method.
type GetChatMenuButtonOptions struct {
	ChatID int64 `query:"chat_id"`
}

// SetChatMenuButton is used to change the bot's menu button in a private chat, or the default menu button.
func (a API) SetChatMenuButton(opts *SetChatMenuButtonOptions) (res APIResponseBool, err error) {
	return res, a.client.get(a.base, "setChatMenuButton", urlValues(opts), &res)
}

// GetChatMenuButton is used to get the current value of the bot's menu button in a private chat, or the default menu button.
func (a API) GetChatMenuButton(opts *GetChatMenuButtonOptions) (res APIResponseMenuButton, err error) {
	return res, a.client.get(a.base, "getChatMenuButton", urlValues(opts), &res)
}
