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

var (
	stickerFile *File
	stickerSet  *StickerSet
)

func TestUploadStickerFile(t *testing.T) {
	resp, err := api.UploadStickerFile(
		chatID,
		StickerFile{
			NewInputFilePath("tests/echotron_test.png"),
			PNGSticker,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	stickerFile = resp.Result
}

func TestCreateNewStickerSet(t *testing.T) {
	resp, err := api.CreateNewStickerSet(
		chatID,
		"echocoverpack_by_echotron_coverage_bot",
		"Echotron Coverage Pack",
		"🤖",
		StickerFile{
			NewInputFileID(stickerFile.FileID),
			PNGSticker,
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestAddStickerToSet(t *testing.T) {
	resp, err := api.AddStickerToSet(
		chatID,
		"echocoverpack_by_echotron_coverage_bot",
		"🤖",
		StickerFile{
			NewInputFilePath("tests/echotron_sticker.png"),
			PNGSticker,
		},
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
	resp, err := api.GetStickerSet("echocoverpack_by_echotron_coverage_bot")

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	stickerSet = resp.Result
}

func TestSetStickerPositionInSet(t *testing.T) {
	resp, err := api.SetStickerPositionInSet(
		stickerSet.Stickers[1].FileID,
		0,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestDeleteStickerFromSet(t *testing.T) {
	resp, err := api.DeleteStickerFromSet(
		stickerSet.Stickers[0].FileID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendSticker(t *testing.T) {
	resp, err := api.SendSticker(
		stickerSet.Stickers[0].FileID,
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

func TestSetStickerSetThumb(t *testing.T) {
	resp, err := api.SetStickerSetThumb(
		"echocoverpack_by_echotron_coverage_bot",
		chatID,
		NewInputFilePath("tests/echotron_thumb.png"),
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}
