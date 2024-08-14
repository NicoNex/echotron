package echotron

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

var (
	msgTmp              *Message
	animationTmp        *Message
	pollTmp             *Message
	locationTmp         *Message
	inviteTmp           *ChatInviteLink
	chatSubInviteTmp    *ChatInviteLink
	filePath            string
	currentBotDesc      string
	currentBotShortDesc string
	msgThreadID         int64
	currentBotName      = "bot name unset"
	api                 = NewAPI("1713461126:AAEV5sgVo513Vz4PT33mpp0ZykJqrnSluzM")
	chatID              = int64(14870908)
	banUserID           = int64(41876271)
	channelID           = int64(-1001563144067)
	groupID             = int64(-1001265771214)
	pinMsgID            = int(11)
	photoID             = "AgACAgQAAxkDAAMrYFtODxV2LL6-kR_6qSbG9n8dIOIAAti1MRug29lSkNq_9o8PC5uMd7EnXQADAQADAgADbQADeooGAAEeBA"
	animationID         = "CgACAgQAAxkDAAICQGBcoGs7GFJ-tR5AkbRRLFTbvdxXAAJ1CAAC1zHgUu-ciZqanytIHgQ"
	audioID             = "CQACAgQAAxkDAAIBCmBbamz_DqKk2GmrzmoM0SrzRN6wAAK9CAACoNvZUgPyk-87OM_YHgQ"
	documentID          = "BQACAgQAAxkDAANmYFtSXcF5kTtwgHeqVUngyuuJMx4AAnQIAAKg29lSb4HP4x-qMT8eBA"
	paidVideoID         = "BAACAgQAAx0EXSuvgwADGGa7vNTyDxfQiyICxWhnLUfhJphkAAIvFAACHozgUVOTwR-Bak97NQQ"
	videoID             = "BAACAgQAAxkDAANxYFtaxF1kfc7nVY_Mtfba3u5dMooAAoYIAAKg29lSpwABJrcveXZlHgQ"
	videoNoteID         = "DQACAgQAAxkDAAIBumBbfT5jPC_cvyEcr0_8DpmFDz2PAALVCgACOX7hUjGZ_MmnZVVeHgQ"
	voiceID             = "AwACAgQAAxkDAAPXYFtmoFriwJFVGDgPPpfUBljgnYAAAq8IAAKg29lStEWfrNMMAxgeBA"
	photoURL            = "https://github.com/NicoNex/echotron/raw/master/assets/tests/echotron_test.png"
	animationURL        = "https://github.com/NicoNex/echotron/raw/master/assets/tests/animation.mp4"
	audioURL            = "https://github.com/NicoNex/echotron/raw/master/assets/tests/audio.mp3"
	documentURL         = "https://github.com/NicoNex/echotron/raw/master/assets/tests/document.pdf"
	logoInvURL          = "https://github.com/NicoNex/echotron/raw/master/assets/tests/echotron_thumb_inv.jpg"
	videoURL            = "https://github.com/NicoNex/echotron/raw/master/assets/tests/video.webm"
	videoNoteURL        = "https://github.com/NicoNex/echotron/raw/master/assets/tests/video_note.mp4"
	voiceURL            = "https://github.com/NicoNex/echotron/raw/master/assets/tests/audio.mp3"

	commands = []BotCommand{
		{Command: "test1", Description: "Test command 1"},
		{Command: "test2", Description: "Test command 2"},
		{Command: "test3", Description: "Test command 3"},
	}

	keyboard = ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{
			{
				{Text: "test 1"},
				{Text: "test 2"},
			},
			{
				{Text: "test 3"},
				{Text: "test 4"},
			},
		},
		ResizeKeyboard: true,
	}

	inlineKeyboard = InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: "test1", CallbackData: "test1"},
				{Text: "test2", CallbackData: "test2"},
			},
			{
				{Text: "test3", CallbackData: "test3"},
			},
		},
	}

	inlineKeyboardEdit = InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: "test1", CallbackData: "test1"},
				{Text: "test2", CallbackData: "test2"},
			},
			{
				{Text: "test3", CallbackData: "test3"},
				{Text: "edit", CallbackData: "edit"},
			},
		},
	}
)

func openBytes(path string) (data []byte, err error) {
	file, err := os.Open(path)

	if err != nil {
		return
	}

	data, err = io.ReadAll(file)

	if err != nil {
		return
	}

	return
}

func TestGetUpdates(t *testing.T) {
	_, err := api.GetUpdates(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetWebhook(t *testing.T) {
	_, err := api.SetWebhook(
		"example.com",
		false,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetWebhookWrongURL(t *testing.T) {
	_, err := api.SetWebhook(
		"example.com_",
		false,
		nil,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDeleteWebhook(t *testing.T) {
	_, err := api.DeleteWebhook(
		false,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetWebhookInfo(t *testing.T) {
	_, err := api.GetWebhookInfo()

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMe(t *testing.T) {
	_, err := api.GetMe()

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMessage(t *testing.T) {
	res, err := api.SendMessage(
		"TestSendMessage *bold* _italic_ `monospace`",
		chatID,
		&MessageOptions{
			ParseMode: MarkdownV2,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	msgTmp = res.Result
}

func TestSetMessageReaction(t *testing.T) {
	_, err := api.SetMessageReaction(
		chatID,
		msgTmp.ID,
		&MessageReactionOptions{
			Reaction: []ReactionType{
				ReactionType{
					Type:  "emoji",
					Emoji: "👍",
				},
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestForwardMessage(t *testing.T) {
	_, err := api.ForwardMessage(
		chatID,
		chatID, // fromChatID
		msgTmp.ID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestForwardMessages(t *testing.T) {
	msg, _ := api.SendMessage(
		"TestForwardMessages",
		chatID,
		nil,
	)

	_, err := api.ForwardMessages(
		chatID,
		chatID, // fromChatID
		[]int{msgTmp.ID, msg.Result.ID},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestForwardMessagesWrongMsgID(t *testing.T) {
	_, err := api.ForwardMessages(
		chatID,
		chatID, // fromChatID
		[]int{},
		nil,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCopyMessage(t *testing.T) {
	_, err := api.CopyMessage(
		chatID,
		chatID, // fromChatID
		msgTmp.ID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyMessages(t *testing.T) {
	msg, _ := api.SendMessage(
		"TestCopyMessages",
		chatID,
		nil,
	)

	_, err := api.CopyMessages(
		chatID,
		chatID, // fromChatID
		[]int{msgTmp.ID, msg.Result.ID},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyMessagesWrongMsgID(t *testing.T) {
	_, err := api.CopyMessages(
		chatID,
		chatID, // fromChatID
		[]int{},
		nil,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSendMessageReply(t *testing.T) {
	_, err := api.SendMessage(
		"TestSendMessageReply",
		chatID,
		&MessageOptions{
			ReplyParameters: ReplyParameters{
				MessageID: msgTmp.ID,
			},
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMessageWithKeyboard(t *testing.T) {
	_, err := api.SendMessage(
		"TestSendMessageWithKeyboard",
		chatID,
		&MessageOptions{
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPhoto(t *testing.T) {
	_, err := api.SendPhoto(
		NewInputFilePath("assets/tests/echotron_test.png"),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhoto",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPhotoByID(t *testing.T) {
	_, err := api.SendPhoto(
		NewInputFileID(photoID),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhotoByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPhotoURL(t *testing.T) {
	_, err := api.SendPhoto(
		NewInputFileURL(photoURL),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhotoURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPhotoBytes(t *testing.T) {
	data, err := openBytes("assets/tests/echotron_test.png")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendPhoto(
		NewInputFileBytes("echotron_test.png", data),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhotoBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPhotoWithKeyboard(t *testing.T) {
	_, err := api.SendPhoto(
		NewInputFilePath("assets/tests/echotron_test.png"),
		chatID,
		&PhotoOptions{
			Caption:     "TestSendPhotoWithKeyboard",
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudio(t *testing.T) {
	_, err := api.SendAudio(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudio",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudioByID(t *testing.T) {
	_, err := api.SendAudio(
		NewInputFileID(audioID),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudioByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudioURL(t *testing.T) {
	_, err := api.SendAudio(
		NewInputFileURL(audioURL),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudioURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudioWithKeyboard(t *testing.T) {
	_, err := api.SendAudio(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&AudioOptions{
			Caption:     "TestSendAudioWithKeyboard",
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudioBytes(t *testing.T) {
	data, err := openBytes("assets/tests/audio.mp3")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendAudio(
		NewInputFileBytes("audio.mp3", data),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudioBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAudioThumbnail(t *testing.T) {
	_, err := api.SendAudio(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&AudioOptions{
			Caption:   "TestSendAudio",
			Thumbnail: NewInputFilePath("assets/tests/echotron_thumb.jpg"),
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendDocument(t *testing.T) {
	_, err := api.SendDocument(
		NewInputFilePath("assets/tests/document.pdf"),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocument",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendDocumentByID(t *testing.T) {
	_, err := api.SendDocument(
		NewInputFileID(documentID),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocumentByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendDocumentURL(t *testing.T) {
	_, err := api.SendDocument(
		NewInputFileURL(documentURL),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocumentURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendDocumentWithKeyboard(t *testing.T) {
	_, err := api.SendDocument(
		NewInputFilePath("assets/tests/document.pdf"),
		chatID,
		&DocumentOptions{
			Caption:     "TestSendDocumentWithKeyboard",
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendDocumentBytes(t *testing.T) {
	file, err := os.Open("assets/tests/document.pdf")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendDocument(
		NewInputFileBytes("document.pdf", data),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocumentBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideo(t *testing.T) {
	_, err := api.SendVideo(
		NewInputFilePath("assets/tests/video.webm"),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideo",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoByID(t *testing.T) {
	_, err := api.SendVideo(
		NewInputFileID(videoID),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideoByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoURL(t *testing.T) {
	_, err := api.SendVideo(
		NewInputFileURL(videoURL),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideoURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoWithKeyboard(t *testing.T) {
	_, err := api.SendVideo(
		NewInputFilePath("assets/tests/video.webm"),
		chatID,
		&VideoOptions{
			Caption:     "TestSendVideoWithKeyboard",
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoBytes(t *testing.T) {
	data, err := openBytes("assets/tests/video.webm")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendVideo(
		NewInputFileBytes("video.webm", data),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideoBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAnimation(t *testing.T) {
	res, err := api.SendAnimation(
		NewInputFilePath("assets/tests/animation.mp4"),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimation",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	animationTmp = res.Result
}

func TestSendAnimationByID(t *testing.T) {
	_, err := api.SendAnimation(
		NewInputFileID(animationID),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimationByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAnimationURL(t *testing.T) {
	res, err := api.SendAnimation(
		NewInputFileURL(animationURL),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimationURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	animationTmp = res.Result
}

func TestSendAnimationWithKeyboard(t *testing.T) {
	_, err := api.SendAnimation(
		NewInputFilePath("assets/tests/animation.mp4"),
		chatID,
		&AnimationOptions{
			Caption:     "TestSendAnimationWithKeyboard",
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAnimationBytes(t *testing.T) {
	data, err := openBytes("assets/tests/animation.mp4")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendAnimation(
		NewInputFileBytes("animation.mp4", data),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimationBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVoice(t *testing.T) {
	_, err := api.SendVoice(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoice",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVoiceByID(t *testing.T) {
	_, err := api.SendVoice(
		NewInputFileID(voiceID),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoiceByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVoiceURL(t *testing.T) {
	_, err := api.SendVoice(
		NewInputFileURL(voiceURL),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoiceURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVoiceWithKeyboard(t *testing.T) {
	_, err := api.SendVoice(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&VoiceOptions{
			Caption:     "TestSendVoiceWithKeyboard",
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVoiceBytes(t *testing.T) {
	data, err := openBytes("assets/tests/audio.mp3")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendVoice(
		NewInputFileBytes("audio.mp3", data),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoiceBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoNote(t *testing.T) {
	_, err := api.SendVideoNote(
		NewInputFilePath("assets/tests/video_note.mp4"),
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoNoteURL(t *testing.T) {
	_, err := api.SendVideoNote(
		NewInputFileURL(videoNoteURL),
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoNoteByID(t *testing.T) {
	_, err := api.SendVideoNote(
		NewInputFileID(videoNoteID),
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoNoteWithKeyboard(t *testing.T) {
	_, err := api.SendVideoNote(
		NewInputFilePath("assets/tests/video_note.mp4"),
		chatID,
		&VideoNoteOptions{
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVideoNoteBytes(t *testing.T) {
	data, err := openBytes("assets/tests/video_note.mp4")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendVideoNote(
		NewInputFileBytes("video_note.mp4", data),
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaPhoto(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFilePath("assets/logo.png"),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaPhoto",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaPhotoByID(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFileID(photoID),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaPhotoByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaPhotoURL(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFileURL(photoURL),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaPhotoURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaPhotoBytes(t *testing.T) {
	data, err := openBytes("assets/tests/echotron_test.png")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFileBytes("echotron_test.png", data),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaPhotoBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaPhotoWithKeyboard(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFilePath("assets/logo.png"),
			},
		},
		&PaidMediaOptions{
			Caption:     "TestSendPaidMediaPhotoWithKeyboard",
			ReplyMarkup: inlineKeyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaVideo(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFilePath("assets/tests/video_note.mp4"),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaVideo",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaVideoByID(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFileID(paidVideoID),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaVideoByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaVideoURL(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFileURL(videoNoteURL),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaVideoURL",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaVideoBytes(t *testing.T) {
	data, err := openBytes("assets/tests/video_note.mp4")

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFileBytes("video_note.mp4", data),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaVideoBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaVideoWithKeyboard(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFilePath("assets/tests/video_note.mp4"),
			},
		},
		&PaidMediaOptions{
			Caption:     "TestSendPaidMediaVideoWithKeyboard",
			ReplyMarkup: inlineKeyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPaidMediaGroup(t *testing.T) {
	_, err := api.SendPaidMedia(
		channelID,
		1,
		[]GroupableInputMedia{
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFilePath("assets/logo.png"),
			},
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFileID(photoID),
			},
			InputPaidMediaPhoto{
				Type:  InputPaidMediaTypePhoto,
				Media: NewInputFileURL(logoInvURL),
			},
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFilePath("assets/tests/video_note.mp4"),
			},
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFileID(paidVideoID),
			},
			InputPaidMediaVideo{
				Type:  InputPaidMediaTypeVideo,
				Media: NewInputFileURL(videoNoteURL),
			},
		},
		&PaidMediaOptions{
			Caption: "TestSendPaidMediaGroup",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMediaGroupPhoto(t *testing.T) {
	_, err := api.SendMediaGroup(
		chatID,
		[]GroupableInputMedia{
			InputMediaPhoto{
				Type:    MediaTypePhoto,
				Media:   NewInputFileID(photoID),
				Caption: "TestSendMediaGroup1",
			},
			InputMediaPhoto{
				Type:    MediaTypePhoto,
				Media:   NewInputFilePath("assets/logo.png"),
				Caption: "TestSendMediaGroup2",
			},
			InputMediaPhoto{
				Type:    MediaTypePhoto,
				Media:   NewInputFileURL(logoInvURL),
				Caption: "TestSendMediaGroup3",
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMediaGroupVideo(t *testing.T) {
	_, err := api.SendMediaGroup(
		chatID,
		[]GroupableInputMedia{
			InputMediaVideo{
				Type:    MediaTypeVideo,
				Media:   NewInputFileID(videoID),
				Caption: "TestSendMediaGroup1",
			},
			InputMediaVideo{
				Type:    MediaTypeVideo,
				Media:   NewInputFilePath("assets/tests/video.webm"),
				Caption: "TestSendMediaGroup2",
			},
			InputMediaVideo{
				Type:    MediaTypeVideo,
				Media:   NewInputFileURL(videoURL),
				Caption: "TestSendMediaGroup3",
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMediaGroupDocument(t *testing.T) {
	_, err := api.SendMediaGroup(
		chatID,
		[]GroupableInputMedia{
			InputMediaDocument{
				Type:    MediaTypeDocument,
				Media:   NewInputFileID(documentID),
				Caption: "TestSendMediaGroup1",
			},
			InputMediaDocument{
				Type:    MediaTypeDocument,
				Media:   NewInputFilePath("assets/tests/document.pdf"),
				Caption: "TestSendMediaGroup2",
			},
			InputMediaDocument{
				Type:    MediaTypeDocument,
				Media:   NewInputFileURL(documentURL),
				Caption: "TestSendMediaGroup3",
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendMediaGroupThumbnail(t *testing.T) {
	_, err := api.SendMediaGroup(
		chatID,
		[]GroupableInputMedia{
			InputMediaAudio{
				Type:      MediaTypeAudio,
				Media:     NewInputFilePath("assets/tests/audio_inv.mp3"),
				Thumbnail: NewInputFilePath("assets/tests/echotron_thumb_inv.jpg"),
				Caption:   "TestSendMediaGroupThumbnail1",
			},
			InputMediaAudio{
				Type:      MediaTypeAudio,
				Media:     NewInputFilePath("assets/tests/audio.mp3"),
				Thumbnail: NewInputFilePath("assets/tests/echotron_thumb.jpg"),
				Caption:   "TestSendMediaGroupThumbnail2",
			},
			InputMediaAudio{
				Type:      MediaTypeAudio,
				Media:     NewInputFileURL(audioURL),
				Thumbnail: NewInputFileURL(logoInvURL),
				Caption:   "TestSendMediaGroupThumbnail3",
			},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendLocation(t *testing.T) {
	res, err := api.SendLocation(
		chatID,
		0.0,
		0.0,
		&LocationOptions{
			LivePeriod:         60,
			HorizontalAccuracy: 50,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	locationTmp = res.Result
}

func TestEditMessageLiveLocation(t *testing.T) {
	_, err := api.EditMessageLiveLocation(
		NewMessageID(chatID, locationTmp.ID),
		0.0,
		0.0,
		&EditLocationOptions{
			HorizontalAccuracy: 50,
			ReplyMarkup:        inlineKeyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestStopMessageLiveLocation(t *testing.T) {
	_, err := api.StopMessageLiveLocation(
		NewMessageID(chatID, locationTmp.ID),
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendVenue(t *testing.T) {
	_, err := api.SendVenue(
		chatID,
		0.0,
		0.0,
		"TestSendVenue",
		"TestSendVenueAddress",
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendContact(t *testing.T) {
	_, err := api.SendContact(
		"1234567890",
		"Name",
		chatID,
		&ContactOptions{
			LastName: "Surname",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendPoll(t *testing.T) {
	res, err := api.SendPoll(
		chatID,
		"TestSendPoll",
		[]InputPollOption{
			{Text: "Option 1"},
			{Text: "Option 2"},
			{Text: "Option 3"},
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	pollTmp = res.Result
}

func TestSendPollWrongOptions(t *testing.T) {
	_, err := api.SendPoll(
		chatID,
		"TestSendPoll",
		[]InputPollOption{},
		nil,
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSendDice(t *testing.T) {
	_, err := api.SendDice(
		chatID,
		Die,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSendChatAction(t *testing.T) {
	_, err := api.SendChatAction(
		Typing,
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	_, err := api.GetUserProfilePhotos(
		chatID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetFile(t *testing.T) {
	res, err := api.GetFile(
		photoID,
	)

	if err != nil {
		t.Fatal(err)
	}

	filePath = res.Result.FilePath
}

func TestDownloadFile(t *testing.T) {
	res, err := api.DownloadFile(
		filePath,
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(res) == 0 {
		t.Fatal("empty file received")
	}
}

func TestBanChatMember(t *testing.T) {
	_, err := api.BanChatMember(
		channelID,
		banUserID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnbanChatMember(t *testing.T) {
	_, err := api.UnbanChatMember(
		channelID,
		banUserID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestRestrictChatMember(t *testing.T) {
	_, err := api.RestrictChatMember(
		groupID,
		banUserID,
		ChatPermissions{
			CanSendMessages: true,
		},
		&RestrictOptions{
			UseIndependentChatPermissions: true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPromoteChatMember(t *testing.T) {
	_, err := api.PromoteChatMember(
		groupID,
		banUserID,
		&PromoteOptions{
			CanManageChat:       true,
			CanPostMessages:     true,
			CanEditMessages:     true,
			CanDeleteMessages:   true,
			CanManageVideoChats: true,
			CanRestrictMembers:  true,
			CanPromoteMembers:   true,
			CanChangeInfo:       true,
			CanInviteUsers:      true,
			CanPinMessages:      true,
			CanPostStories:      true,
			CanEditStories:      true,
			CanDeleteStories:    true,
			CanManageTopics:     true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestBanChatSenderChat(t *testing.T) {
	_, err := api.BanChatSenderChat(
		channelID,
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnbanChatSenderChat(t *testing.T) {
	_, err := api.UnbanChatSenderChat(
		channelID,
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetChatPermissions(t *testing.T) {
	_, err := api.SetChatPermissions(
		groupID,
		ChatPermissions{
			CanSendMessages:       true,
			CanSendAudios:         true,
			CanSendDocuments:      true,
			CanSendPhotos:         true,
			CanSendVideos:         true,
			CanSendVideoNotes:     true,
			CanSendVoiceNotes:     true,
			CanSendPolls:          true,
			CanSendOtherMessages:  true,
			CanAddWebPagePreviews: true,
			CanChangeInfo:         true,
			CanInviteUsers:        true,
			CanPinMessages:        true,
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestExportChatInviteLink(t *testing.T) {
	_, err := api.ExportChatInviteLink(
		channelID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateChatInviteLink(t *testing.T) {
	res, err := api.CreateChatInviteLink(
		channelID, nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	inviteTmp = res.Result
}

func TestEditChatInviteLink(t *testing.T) {
	_, err := api.EditChatInviteLink(
		channelID,
		inviteTmp.InviteLink,
		&InviteLinkOptions{
			ExpireDate: time.Now().Unix() + 300,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateChatSubscriptionInviteLink(t *testing.T) {
	res, err := api.CreateChatSubscriptionInviteLink(
		channelID,
		2592000,
		1,
		&ChatSubscriptionInviteOptions{
			Name: "TestCreateChatSubscriptionInviteLink",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	chatSubInviteTmp = res.Result
}

func TestEditChatSubscriptionInviteLink(t *testing.T) {
	_, err := api.EditChatSubscriptionInviteLink(
		channelID,
		chatSubInviteTmp.InviteLink,
		&ChatSubscriptionInviteOptions{
			Name: "TestEditChatSubscriptionInviteLink",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestRevokeChatInviteLink(t *testing.T) {
	_, err := api.RevokeChatInviteLink(
		channelID,
		inviteTmp.InviteLink,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetChatPhoto(t *testing.T) {
	_, err := api.SetChatPhoto(
		NewInputFilePath("assets/tests/echotron_test.png"),
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteChatPhoto(t *testing.T) {
	_, err := api.DeleteChatPhoto(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetChatTitle(t *testing.T) {
	_, err := api.SetChatTitle(
		groupID,
		"Echotron Coverage Supergroup",
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetChatDescription(t *testing.T) {
	_, err := api.SetChatDescription(
		groupID,
		fmt.Sprintf(
			"This supergroup is used to test some of the methods of the Echotron library for Telegram bots.\n\nLast changed: %d",
			time.Now().Unix(),
		),
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPinChatMessage(t *testing.T) {
	_, err := api.PinChatMessage(
		groupID,
		pinMsgID,
		&PinMessageOptions{
			DisableNotification: true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnpinChatMessage(t *testing.T) {
	_, err := api.UnpinChatMessage(
		groupID,
		&UnpinMessageOptions{
			MessageID: pinMsgID,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnpinAllChatMessages(t *testing.T) {
	_, err := api.UnpinAllChatMessages(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetChat(t *testing.T) {
	res, err := api.GetChat(
		chatID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if res.Result.Type != "private" && res.Result.Type != "group" &&
		res.Result.Type != "supergroup" && res.Result.Type != "channel" {

		t.Fatal("wrong chat type, got:", res.Result.Type)
	}
}

func TestGetChatAdministrators(t *testing.T) {
	_, err := api.GetChatAdministrators(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetChatMemberCount(t *testing.T) {
	_, err := api.GetChatMemberCount(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetChatMember(t *testing.T) {
	_, err := api.GetChatMember(
		groupID,
		chatID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateForumTopic(t *testing.T) {
	res, err := api.CreateForumTopic(
		groupID,
		"Test Topic",
		&CreateTopicOptions{
			IconColor: Green,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	msgThreadID = res.Result.MessageThreadID
}

func TestEditForumTopic(t *testing.T) {
	_, err := api.EditForumTopic(
		groupID,
		msgThreadID,
		&EditTopicOptions{
			Name:              "Testing Topic",
			IconCustomEmojiID: "5411138633765757782",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCloseForumTopic(t *testing.T) {
	_, err := api.CloseForumTopic(
		groupID,
		msgThreadID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestReopenForumTopic(t *testing.T) {
	_, err := api.ReopenForumTopic(
		groupID,
		msgThreadID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnpinAllForumTopicMessages(t *testing.T) {
	res, err := api.SendMessage(
		"Test",
		groupID,
		&MessageOptions{
			MessageThreadID: msgThreadID,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.PinChatMessage(
		groupID,
		res.Result.ID,
		&PinMessageOptions{
			DisableNotification: true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	_, err = api.UnpinAllForumTopicMessages(
		groupID,
		msgThreadID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteForumTopic(t *testing.T) {
	_, err := api.DeleteForumTopic(
		groupID,
		msgThreadID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditGeneralForumTopic(t *testing.T) {
	_, err := api.EditGeneralForumTopic(
		groupID,
		fmt.Sprintf(
			"General | %d",
			time.Now().Unix(),
		),
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestCloseGeneralForumTopic(t *testing.T) {
	_, err := api.CloseGeneralForumTopic(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestHideGeneralForumTopic(t *testing.T) {
	_, err := api.HideGeneralForumTopic(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnhideGeneralForumTopic(t *testing.T) {
	_, err := api.UnhideGeneralForumTopic(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestReopenGeneralForumTopic(t *testing.T) {
	_, err := api.ReopenGeneralForumTopic(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestUnpinAllGeneralForumTopicMessages(t *testing.T) {
	_, err := api.UnpinAllGeneralForumTopicMessages(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUserChatBoosts(t *testing.T) {
	_, err := api.GetUserChatBoosts(
		channelID,
		chatID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetMyCommands(t *testing.T) {
	opts := &CommandOptions{
		LanguageCode: "it",
		Scope:        BotCommandScope{Type: BCSTChat, ChatID: chatID},
	}

	_, err := api.SetMyCommands(
		opts,
		commands...,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMyCommands(t *testing.T) {
	res, err := api.GetMyCommands(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	for i, cmd := range res.Result {
		if !reflect.DeepEqual(*cmd, commands[i]) {
			t.Logf("expected command in %d: %+v", i, commands[i])
			t.Logf("command in %d from API: %+v", i, cmd)
			t.Fatal("error: commands mismatch")
		}
	}
}

func TestDeleteMyCommands(t *testing.T) {
	_, err := api.DeleteMyCommands(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetMyName(t *testing.T) {
	currentBotName = fmt.Sprintf(
		"Echotron Coverage Bot - %d",
		time.Now().Unix(),
	)

	_, err := api.SetMyName(currentBotName, "")

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMyName(t *testing.T) {
	res, err := api.GetMyName("")

	if err != nil {
		t.Fatal(err)
	}

	if res.Result.Name != currentBotName {
		t.Logf("expected bot name [\"%s\"]\n", currentBotName)
		t.Logf("got bot name [\"%s\"]\n", res.Result.Name)
		t.Fatal("error: bot name mismatch")
	}
}

func TestSetMyDescription(t *testing.T) {
	currentBotDesc = fmt.Sprintf(
		"Echotron Coverage Bot - %d",
		time.Now().Unix(),
	)

	_, err := api.SetMyDescription(currentBotDesc, "")

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetMyDescription(t *testing.T) {
	res, err := api.GetMyDescription("")

	if err != nil {
		t.Fatal(err)
	}

	if res.Result.Description != currentBotDesc {
		t.Logf("expected bot description [\"%s\"]\n", currentBotDesc)
		t.Logf("got bot description [\"%s\"]\n", res.Result.Description)
		t.Fatal("error: bot description mismatch")
	}
}

func TestSetMyShortDescription(t *testing.T) {
	currentBotShortDesc = fmt.Sprintf(
		"Echotron Coverage Bot - %d",
		time.Now().Unix(),
	)

	_, err := api.SetMyShortDescription(currentBotShortDesc, "")

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageText(t *testing.T) {
	_, err := api.EditMessageText(
		"edited message",
		NewMessageID(chatID, msgTmp.ID),
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageTextWithKeyboard(t *testing.T) {
	_, err := api.EditMessageText(
		"edited message with keyboard",
		NewMessageID(chatID, msgTmp.ID),
		&MessageTextOptions{
			ReplyMarkup: inlineKeyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageCaption(t *testing.T) {
	_, err := api.EditMessageCaption(
		NewMessageID(chatID, animationTmp.ID),
		&MessageCaptionOptions{
			Caption: "TestEditMessageCaption",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageMedia(t *testing.T) {
	_, err := api.EditMessageMedia(
		NewMessageID(chatID, animationTmp.ID),
		InputMediaAnimation{
			Type:    MediaTypeAnimation,
			Media:   NewInputFileID(animationID),
			Caption: "TestEditMessageMedia",
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageMediaBytes(t *testing.T) {
	_, err := api.EditMessageMedia(
		NewMessageID(chatID, animationTmp.ID),
		InputMediaAnimation{
			Type:    MediaTypeAnimation,
			Media:   NewInputFilePath("assets/tests/animation.mp4"),
			Caption: "TestEditMessageMediaBytes",
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageMediaURL(t *testing.T) {
	_, err := api.EditMessageMedia(
		NewMessageID(chatID, animationTmp.ID),
		InputMediaAnimation{
			Type:    MediaTypeAnimation,
			Media:   NewInputFileURL(animationURL),
			Caption: "TestEditMessageMediaURL",
		},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestEditMessageReplyMarkup(t *testing.T) {
	_, err := api.EditMessageReplyMarkup(
		NewMessageID(chatID, msgTmp.ID),
		&MessageReplyMarkupOptions{
			ReplyMarkup: inlineKeyboardEdit,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestStopPoll(t *testing.T) {
	_, err := api.StopPoll(
		chatID,
		pollTmp.ID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMessage(t *testing.T) {
	_, err := api.DeleteMessage(
		chatID,
		msgTmp.ID,
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMessages(t *testing.T) {
	msg, _ := api.SendMessage(
		"TestDeleteMessages",
		chatID,
		nil,
	)

	_, err := api.DeleteMessages(
		chatID,
		[]int{msgTmp.ID, msg.Result.ID},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteMessagesWrongMsgIDs(t *testing.T) {
	_, err := api.DeleteMessages(
		chatID,
		[]int{},
	)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
