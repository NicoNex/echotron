/*
 * Echotron
 * Copyright (C) 2018 The Echotron Contributors
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
	"testing"
	"time"
)

var (
	stickerFile    *File
	stickerSet     *StickerSet
	stickerSetName = fmt.Sprintf("set%d_by_echotron_coverage_bot", time.Now().Unix())
)

func TestUploadStickerFile(t *testing.T) {
	resp, err := api.UploadStickerFile(
		chatID,
		NewInputFilePath("assets/tests/echotron_test.png"),
		StaticFormat,
	)

	if err != nil {
		t.Fatal(err)
	}

	stickerFile = resp.Result
}

func TestCreateNewStickerSet(t *testing.T) {
	_, err := api.CreateNewStickerSet(
		chatID,
		stickerSetName,
		"Echotron Coverage Pack",
		[]InputSticker{
			{
				Sticker:   NewInputFileID(stickerFile.FileID),
				EmojiList: []string{"🤖"},
				Format:    StaticFormat,
			},
			{
				Sticker:   NewInputFilePath("assets/tests/echotron_test.png"),
				EmojiList: []string{"🤖"},
				Format:    StaticFormat,
			},
			{
				Sticker:   NewInputFileURL(photoURL),
				EmojiList: []string{"🤖"},
				Format:    StaticFormat,
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestAddStickerToSet(t *testing.T) {
	_, err := api.AddStickerToSet(
		chatID,
		stickerSetName,
		InputSticker{
			Sticker:   NewInputFilePath("assets/tests/echotron_sticker.png"),
			EmojiList: []string{"🤖"},
			Format:    StaticFormat,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCustomEmojiStickers(t *testing.T) {
	_, err := api.GetCustomEmojiStickers(
		"5407041870620531251",
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetStickerSet(t *testing.T) {
	resp, err := api.GetStickerSet(
		stickerSetName,
	)

	if err != nil {
		t.Fatal(err)
	}

	stickerSet = resp.Result
}

func TestSetStickerPositionInSet(t *testing.T) {
	_, err := api.SetStickerPositionInSet(
		stickerSet.Stickers[1].FileID,
		0,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetStickerEmojiList(t *testing.T) {
	_, err := api.SetStickerEmojiList(
		stickerSet.Stickers[0].FileID,
		[]string{"🤖", "👾"},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetStickerKeywords(t *testing.T) {
	_, err := api.SetStickerKeywords(
		stickerSet.Stickers[0].FileID,
		[]string{"echotron"},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetStickerSetTitle(t *testing.T) {
	_, err := api.SetStickerSetTitle(
		stickerSetName,
		fmt.Sprintf("new_%s", stickerSetName),
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestReplaceStickerInSet(t *testing.T) {
	_, err := api.ReplaceStickerInSet(
		chatID,
		stickerSetName,
		stickerSet.Stickers[0].FileID,
		InputSticker{
			Sticker:   NewInputFileURL(photoURL),
			EmojiList: []string{"🤖"},
			Format:    StaticFormat,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteStickerFromSet(t *testing.T) {
	_, err := api.DeleteStickerFromSet(
		stickerSet.Stickers[1].FileID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendSticker(t *testing.T) {
	_, err := api.SendSticker(
		stickerSet.Stickers[0].FileID,
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetStickerSetThumbnail(t *testing.T) {
	_, err := api.SetStickerSetThumbnail(
		stickerSetName,
		chatID,
		NewInputFilePath("assets/tests/echotron_thumb.png"),
		StaticFormat,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteStickerSet(t *testing.T) {
	_, err := api.DeleteStickerSet(stickerSetName)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetForumTopicIconStickers(t *testing.T) {
	res, err := api.GetForumTopicIconStickers()

	if err != nil {
		t.Fatal(err)
	}

	if len(res.Result) == 0 {
		t.Fatal("error: Telegram returned no forum topic icon stickers")
	}
}
