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

import "testing"

func TestSendSticker(t *testing.T) {
    resp, err := api.SendSticker(
        stickerID,
        chatID,
        nil,
    )

    if err != nil {
        t.Fatal(err)
    }

    if !resp.Ok {
        t.Fatal(resp.ErrorCode, resp.Description)
    }
}

func TestGetStickerSet(t *testing.T) {
    resp, err := api.GetStickerSet("RickAndMorty")
    if err != nil {
        t.Fatal(err)
    }

    if !resp.Ok {
        t.Fatal(resp.ErrorCode, resp.Description)
    }
}
