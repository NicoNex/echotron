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
	msgTmp       *Message
	photoTmp     *Message
	pollTmp      *Message
	locationTmp  *Message
	inviteTmp    *ChatInviteLink
	expInviteTmp string
	filePath     string
	api          = NewAPI("1713461126:AAEV5sgVo513Vz4PT33mpp0ZykJqrnSluzM")
	chatID       = int64(14870908)
	banUserID    = int64(41876271)
	botID        = int64(1713461126)
	channelID    = int64(-1001563144067)
	groupID      = int64(-1001265771214)
	pinMsgID     = int(11)
	photoID      = "AgACAgQAAxkDAAMrYFtODxV2LL6-kR_6qSbG9n8dIOIAAti1MRug29lSkNq_9o8PC5uMd7EnXQADAQADAgADbQADeooGAAEeBA"
	animationID  = "CgACAgQAAxkDAAICQGBcoGs7GFJ-tR5AkbRRLFTbvdxXAAJ1CAAC1zHgUu-ciZqanytIHgQ"
	audioID      = "CQACAgQAAxkDAAIBCmBbamz_DqKk2GmrzmoM0SrzRN6wAAK9CAACoNvZUgPyk-87OM_YHgQ"
	documentID   = "BQACAgQAAxkDAANmYFtSXcF5kTtwgHeqVUngyuuJMx4AAnQIAAKg29lSb4HP4x-qMT8eBA"
	videoID      = "BAACAgQAAxkDAANxYFtaxF1kfc7nVY_Mtfba3u5dMooAAoYIAAKg29lSpwABJrcveXZlHgQ"
	videoNoteID  = "DQACAgQAAxkDAAIBumBbfT5jPC_cvyEcr0_8DpmFDz2PAALVCgACOX7hUjGZ_MmnZVVeHgQ"
	voiceID      = "AwACAgQAAxkDAAPXYFtmoFriwJFVGDgPPpfUBljgnYAAAq8IAAKg29lStEWfrNMMAxgeBA"

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
	resp, err := api.GetUpdates(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSetWebhook(t *testing.T) {
	resp, err := api.SetWebhook(
		"example.com",
		false,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestDeleteWebhook(t *testing.T) {
	resp, err := api.DeleteWebhook(
		false,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetWebhookInfo(t *testing.T) {
	resp, err := api.GetWebhookInfo()

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetMe(t *testing.T) {
	resp, err := api.GetMe()

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendMessage(t *testing.T) {
	resp, err := api.SendMessage(
		"TestSendMessage *bold* _italic_ `monospace`",
		chatID,
		&MessageOptions{
			ParseMode: MarkdownV2,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	msgTmp = resp.Result
}

func TestForwardMessage(t *testing.T) {
	resp, err := api.ForwardMessage(
		chatID,
		chatID, // fromChatID
		msgTmp.ID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestCopyMessage(t *testing.T) {
	resp, err := api.CopyMessage(
		chatID,
		chatID, // fromChatID
		msgTmp.ID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendMessageReply(t *testing.T) {
	resp, err := api.SendMessage(
		"TestSendMessageReply",
		chatID,
		&MessageOptions{
			ReplyToMessageID: msgTmp.ID,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendMessageWithKeyboard(t *testing.T) {
	resp, err := api.SendMessage(
		"TestSendMessageWithKeyboard",
		chatID,
		&MessageOptions{
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendPhoto(t *testing.T) {
	resp, err := api.SendPhoto(
		NewInputFilePath("assets/tests/echotron_test.png"),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhoto",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	photoTmp = resp.Result
}

func TestSendPhotoByID(t *testing.T) {
	resp, err := api.SendPhoto(
		NewInputFileID(photoID),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhotoByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendPhotoBytes(t *testing.T) {
	data, err := openBytes("assets/tests/echotron_test.png")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := api.SendPhoto(
		NewInputFileBytes("echotron_test.png", data),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhotoBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendPhotoWithKeyboard(t *testing.T) {
	resp, err := api.SendPhoto(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudio(t *testing.T) {
	resp, err := api.SendAudio(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudio",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudioByID(t *testing.T) {
	resp, err := api.SendAudio(
		NewInputFileID(audioID),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudioByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudioWithKeyboard(t *testing.T) {
	resp, err := api.SendAudio(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudioBytes(t *testing.T) {
	data, err := openBytes("assets/tests/audio.mp3")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := api.SendAudio(
		NewInputFileBytes("audio.mp3", data),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudioBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudioThumb(t *testing.T) {
	resp, err := api.SendAudio(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudio",
			Thumb:   NewInputFilePath("assets/tests/echotron_thumb.jpg"),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendDocument(t *testing.T) {
	resp, err := api.SendDocument(
		NewInputFilePath("assets/tests/document.pdf"),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocument",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendDocumentByID(t *testing.T) {
	resp, err := api.SendDocument(
		NewInputFileID(documentID),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocumentByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendDocumentWithKeyboard(t *testing.T) {
	resp, err := api.SendDocument(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
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

	resp, err := api.SendDocument(
		NewInputFileBytes("document.pdf", data),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocumentBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideo(t *testing.T) {
	resp, err := api.SendVideo(
		NewInputFilePath("assets/tests/video.webm"),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideo",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoByID(t *testing.T) {
	resp, err := api.SendVideo(
		NewInputFileID(videoID),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideoByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoWithKeyboard(t *testing.T) {
	resp, err := api.SendVideo(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoBytes(t *testing.T) {
	data, err := openBytes("assets/tests/video.webm")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := api.SendVideo(
		NewInputFileBytes("video.webm", data),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideoBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAnimation(t *testing.T) {
	resp, err := api.SendAnimation(
		NewInputFilePath("assets/tests/animation.mp4"),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimation",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAnimationByID(t *testing.T) {
	resp, err := api.SendAnimation(
		NewInputFileID(animationID),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimationByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAnimationWithKeyboard(t *testing.T) {
	resp, err := api.SendAnimation(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAnimationBytes(t *testing.T) {
	data, err := openBytes("assets/tests/animation.mp4")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := api.SendAnimation(
		NewInputFileBytes("animation.mp4", data),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimationBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVoice(t *testing.T) {
	resp, err := api.SendVoice(
		NewInputFilePath("assets/tests/audio.mp3"),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoice",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVoiceByID(t *testing.T) {
	resp, err := api.SendVoice(
		NewInputFileID(voiceID),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoiceByID",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVoiceWithKeyboard(t *testing.T) {
	resp, err := api.SendVoice(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVoiceBytes(t *testing.T) {
	data, err := openBytes("assets/tests/audio.mp3")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := api.SendVoice(
		NewInputFileBytes("audio.mp3", data),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoiceBytes",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoNote(t *testing.T) {
	resp, err := api.SendVideoNote(
		NewInputFilePath("assets/tests/video_note.mp4"),
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

func TestSendMediaGroup(t *testing.T) {
	resp, err := api.SendMediaGroup(
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

func TestSendMediaGroupThumb(t *testing.T) {
	resp, err := api.SendMediaGroup(
		chatID,
		[]GroupableInputMedia{
			InputMediaAudio{
				Type:    MediaTypeAudio,
				Media:   NewInputFilePath("assets/tests/audio_inv.mp3"),
				Thumb:   NewInputFilePath("assets/tests/echotron_thumb_inv.jpg"),
				Caption: "TestSendMediaGroupThumb1",
			},
			InputMediaAudio{
				Type:    MediaTypeAudio,
				Media:   NewInputFilePath("assets/tests/audio.mp3"),
				Thumb:   NewInputFilePath("assets/tests/echotron_thumb.jpg"),
				Caption: "TestSendMediaGroupThumb2",
			},
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

func TestSendVideoNoteByID(t *testing.T) {
	resp, err := api.SendVideoNote(
		NewInputFileID(videoNoteID),
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

func TestSendVideoNoteWithKeyboard(t *testing.T) {
	resp, err := api.SendVideoNote(
		NewInputFilePath("assets/tests/video_note.mp4"),
		chatID,
		&VideoNoteOptions{
			ReplyMarkup: keyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoNoteBytes(t *testing.T) {
	data, err := openBytes("assets/tests/video_note.mp4")

	if err != nil {
		t.Fatal(err)
	}

	resp, err := api.SendVideoNote(
		NewInputFileBytes("video_note.mp4", data),
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

func TestSendLocation(t *testing.T) {
	resp, err := api.SendLocation(
		chatID,
		0.0,
		0.0,
		&LocationOptions{
			LivePeriod: 60,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	locationTmp = resp.Result
}

func TestEditMessageLiveLocation(t *testing.T) {
	resp, err := api.EditMessageLiveLocation(
		NewMessageID(chatID, locationTmp.ID),
		0.0,
		0.0,
		&EditLocationOptions{
			ReplyMarkup: inlineKeyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestStopMessageLiveLocation(t *testing.T) {
	resp, err := api.StopMessageLiveLocation(
		NewMessageID(chatID, locationTmp.ID),
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVenue(t *testing.T) {
	resp, err := api.SendVenue(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendContact(t *testing.T) {
	resp, err := api.SendContact(
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

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendPoll(t *testing.T) {
	resp, err := api.SendPoll(
		chatID,
		"TestSendPoll",
		[]string{"Option 1", "Option 2", "Option 3"},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	pollTmp = resp.Result
}

func TestSendDice(t *testing.T) {
	resp, err := api.SendDice(
		chatID,
		Die,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendChatAction(t *testing.T) {
	resp, err := api.SendChatAction(
		Typing,
		chatID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	resp, err := api.GetUserProfilePhotos(
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

func TestGetFile(t *testing.T) {
	resp, err := api.GetFile(
		photoID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	filePath = resp.Result.FilePath
}

func TestDownloadFile(t *testing.T) {
	resp, err := api.DownloadFile(
		filePath,
	)

	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("empty file received")
	}
}

func TestBanChatMember(t *testing.T) {
	resp, err := api.BanChatMember(
		channelID,
		banUserID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestUnbanChatMember(t *testing.T) {
	resp, err := api.UnbanChatMember(
		channelID,
		banUserID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestRestrictChatMember(t *testing.T) {
	resp, err := api.RestrictChatMember(
		groupID,
		banUserID,
		ChatPermissions{},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestPromoteChatMember(t *testing.T) {
	resp, err := api.PromoteChatMember(
		groupID,
		banUserID,
		&PromoteOptions{
			CanDeleteMessages:   true,
			CanManageVoiceChats: true,
			CanRestrictMembers:  true,
			CanPromoteMembers:   true,
			CanChangeInfo:       true,
			CanInviteUsers:      true,
			CanPinMessages:      true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSetChatAdministratorCustomTitle(t *testing.T) {
	resp, err := api.SetChatAdministratorCustomTitle(
		groupID,
		banUserID,
		"TestCustomTitle",
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSetChatPermissions(t *testing.T) {
	resp, err := api.SetChatPermissions(
		groupID,
		ChatPermissions{
			CanSendMessages:       true,
			CanSendMediaMessages:  true,
			CanSendPolls:          true,
			CanSendOtherMessages:  true,
			CanAddWebPagePreviews: true,
			CanChangeInfo:         true,
			CanInviteUsers:        true,
			CanPinMessages:        true,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestExportChatInviteLink(t *testing.T) {
	resp, err := api.ExportChatInviteLink(
		channelID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	expInviteTmp = resp.Result
}

func TestCreateChatInviteLink(t *testing.T) {
	resp, err := api.CreateChatInviteLink(
		channelID, nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	inviteTmp = resp.Result
}

func TestEditChatInviteLink(t *testing.T) {
	resp, err := api.EditChatInviteLink(
		channelID,
		inviteTmp.InviteLink,
		&InviteLinkOptions{
			ExpireDate: time.Now().Unix() + 20,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestRevokeChatInviteLink(t *testing.T) {
	resp, err := api.RevokeChatInviteLink(
		channelID,
		expInviteTmp,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSetChatPhoto(t *testing.T) {
	resp, err := api.SetChatPhoto(
		NewInputFilePath("assets/tests/echotron_test.png"),
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestDeleteChatPhoto(t *testing.T) {
	resp, err := api.DeleteChatPhoto(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSetChatTitle(t *testing.T) {
	resp, err := api.SetChatTitle(
		groupID,
		"Echotron Coverage Supergroup",
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSetChatDescription(t *testing.T) {
	resp, err := api.SetChatDescription(
		groupID,
		fmt.Sprintf(
			"This supergroup is used to test some of the methods of the Echotron library for Telegram bots.\n\nLast changed: %d",
			time.Now().Unix(),
		),
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestPinChatMessage(t *testing.T) {
	resp, err := api.PinChatMessage(
		groupID,
		pinMsgID,
		&DisableNotificationOptions{true},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestUnpinChatMessage(t *testing.T) {
	resp, err := api.UnpinChatMessage(
		groupID,
		pinMsgID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestUnpinAllChatMessages(t *testing.T) {
	resp, err := api.UnpinAllChatMessages(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetChat(t *testing.T) {
	resp, err := api.GetChat(
		chatID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if resp.Result.Type != "private" && resp.Result.Type != "group" &&
		resp.Result.Type != "supergroup" && resp.Result.Type != "channel" {

		t.Fatal("wrong chat type, got:", resp.Result.Type)
	}
}

func TestGetChatAdministrators(t *testing.T) {
	resp, err := api.GetChatAdministrators(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetChatMemberCount(t *testing.T) {
	resp, err := api.GetChatMemberCount(
		groupID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetChatMember(t *testing.T) {
	resp, err := api.GetChatMember(
		groupID,
		chatID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestAnswerCallbackQuery(t *testing.T) {
	_, err := api.AnswerCallbackQuery(
		"test",
		&CallbackQueryOptions{
			Text: "test",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}

func TestSetMyCommands(t *testing.T) {
	resp, err := api.SetMyCommands(
		nil,
		commands...,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetMyCommands(t *testing.T) {
	resp, err := api.GetMyCommands(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	for i, res := range resp.Result {
		if !reflect.DeepEqual(*res, commands[i]) {
			t.Logf("expected command in %d: %v", i, commands[i])
			t.Logf("command in %d from API: %v", i, res)
			t.Fatal("error: commands mismatch")
		}
	}
}

func TestDeleteMyCommands(t *testing.T) {
	resp, err := api.DeleteMyCommands(
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestEditMessageText(t *testing.T) {
	resp, err := api.EditMessageText(
		"edited message",
		NewMessageID(chatID, msgTmp.ID),
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestEditMessageTextWithKeyboard(t *testing.T) {
	resp, err := api.EditMessageText(
		"edited message with keyboard",
		NewMessageID(chatID, msgTmp.ID),
		&MessageTextOptions{
			ReplyMarkup: inlineKeyboard,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestEditMessageCaption(t *testing.T) {
	resp, err := api.EditMessageCaption(
		NewMessageID(chatID, photoTmp.ID),
		&MessageCaptionOptions{
			Caption: "TestEditMessageCaption",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestEditMessageMedia(t *testing.T) {
	resp, err := api.EditMessageMedia(
		NewMessageID(chatID, photoTmp.ID),
		InputMediaPhoto{
			Type:    MediaTypePhoto,
			Media:   NewInputFileID(photoID),
			Caption: "TestEditMessageMedia",
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

func TestEditMessageMediaBytes(t *testing.T) {
	resp, err := api.EditMessageMedia(
		NewMessageID(chatID, photoTmp.ID),
		InputMediaPhoto{
			Type:    MediaTypePhoto,
			Media:   NewInputFilePath("assets/logo.png"),
			Caption: "TestEditMessageMediaBytes",
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

func TestEditMessageReplyMarkup(t *testing.T) {
	resp, err := api.EditMessageReplyMarkup(
		NewMessageID(chatID, msgTmp.ID),
		&MessageReplyMarkup{
			ReplyMarkup: inlineKeyboardEdit,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestStopPoll(t *testing.T) {
	resp, err := api.StopPoll(
		chatID,
		pollTmp.ID,
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestDeleteMessage(t *testing.T) {
	resp, err := api.DeleteMessage(
		chatID,
		msgTmp.ID,
	)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}
