/*
 * Echotron
 * Copyright (C) 2018-2022 The Echotron Devs
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

import "fmt"

// APIError represents an error returned by the Telegram API.
type APIError struct {
	desc string
	code int
}

// ErrorCode returns the error code received from the Telegram API.
func (a *APIError) ErrorCode() int {
	return a.code
}

// Description returns the error description received from the Telegram API.
func (a *APIError) Description() string {
	return a.desc
}

// Error returns the error string.
func (a *APIError) Error() string {
	return fmt.Sprintf("API error: %d %s", a.code, a.desc)
}
