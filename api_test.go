package echotron

import (
	"io"
	"os"
	"reflect"
	"testing"
)

var (
	msgTmp      *Message
	photoTmp    *Message
	api         = NewAPI("1713461126:AAEV5sgVo513Vz4PT33mpp0ZykJqrnSluzM")
	chatID      = int64(14870908)
	groupID     = int64(-1001241973131)
	photoID     = "AgACAgQAAxkDAAMrYFtODxV2LL6-kR_6qSbG9n8dIOIAAti1MRug29lSkNq_9o8PC5uMd7EnXQADAQADAgADbQADeooGAAEeBA"
	animationID = "CgACAgQAAxkDAAICQGBcoGs7GFJ-tR5AkbRRLFTbvdxXAAJ1CAAC1zHgUu-ciZqanytIHgQ"
	audioID     = "CQACAgQAAxkDAAIBCmBbamz_DqKk2GmrzmoM0SrzRN6wAAK9CAACoNvZUgPyk-87OM_YHgQ"
	documentID  = "BQACAgQAAxkDAANmYFtSXcF5kTtwgHeqVUngyuuJMx4AAnQIAAKg29lSb4HP4x-qMT8eBA"
	videoID     = "BAACAgQAAxkDAANxYFtaxF1kfc7nVY_Mtfba3u5dMooAAoYIAAKg29lSpwABJrcveXZlHgQ"
	videoNoteID = "DQACAgQAAxkDAAIBumBbfT5jPC_cvyEcr0_8DpmFDz2PAALVCgACOX7hUjGZ_MmnZVVeHgQ"
	voiceID     = "AwACAgQAAxkDAAPXYFtmoFriwJFVGDgPPpfUBljgnYAAAq8IAAKg29lStEWfrNMMAxgeBA"

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

func TestGetUpdates(t *testing.T) {
	resp, err := api.GetUpdates(nil)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatalf("%d %s", resp.ErrorCode, resp.Description)
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
		t.Fatalf("%d %s", resp.ErrorCode, resp.Description)
	}
}

func TestDeleteWebhook(t *testing.T) {
	resp, err := api.DeleteWebhook(false)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatalf("%d %s", resp.ErrorCode, resp.Description)
	}
}

func TestGetWebhookInfo(t *testing.T) {
	resp, err := api.GetWebhookInfo()

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatalf("%d %s", resp.ErrorCode, resp.Description)
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

func TestSendMessageReply(t *testing.T) {
	resp, err := api.SendMessage(
		"TestSendMessageReply",
		chatID,
		&MessageOptions{
			BaseOptions: BaseOptions{
				ReplyToMessageID: msgTmp.ID,
			},
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
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
		NewInputFilePath("tests/echotron_test.png"),
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
	file, err := os.Open("tests/echotron_test.png")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

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
		NewInputFilePath("tests/echotron_test.png"),
		chatID,
		&PhotoOptions{
			Caption: "TestSendPhotoWithKeyboard",
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
		NewInputFilePath("tests/audio.mp3"),
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
		NewInputFilePath("tests/audio.mp3"),
		chatID,
		&AudioOptions{
			Caption: "TestSendAudioWithKeyboard",
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
	file, err := os.Open("tests/audio.mp3")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

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

func TestSendDocument(t *testing.T) {
	resp, err := api.SendDocument(
		NewInputFilePath("tests/document.pdf"),
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
		NewInputFilePath("tests/document.pdf"),
		chatID,
		&DocumentOptions{
			Caption: "TestSendDocumentWithKeyboard",
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
	file, err := os.Open("tests/document.pdf")

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
		NewInputFilePath("tests/video.webm"),
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
		NewInputFilePath("tests/video.webm"),
		chatID,
		&VideoOptions{
			Caption: "TestSendVideoWithKeyboard",
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
	file, err := os.Open("tests/video.webm")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

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
		NewInputFilePath("tests/animation.mp4"),
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
		NewInputFilePath("tests/animation.mp4"),
		chatID,
		&AnimationOptions{
			Caption: "TestSendAnimationWithKeyboard",
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
	file, err := os.Open("tests/animation.mp4")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

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
		NewInputFilePath("tests/audio.mp3"),
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
		NewInputFilePath("tests/audio.mp3"),
		chatID,
		&VoiceOptions{
			Caption: "TestSendVoiceWithKeyboard",
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
	file, err := os.Open("tests/audio.mp3")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

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
		NewInputFilePath("tests/video_note.mp4"),
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
		NewInputFilePath("tests/video_note.mp4"),
		chatID,
		&VideoNoteOptions{
			BaseOptions: BaseOptions{
				ReplyMarkup: keyboard,
			},
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
	file, err := os.Open("tests/video_note.mp4")

	if err != nil {
		t.Fatal(err)
	}

	data, err := io.ReadAll(file)

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

func TestSendChatAction(t *testing.T) {
	resp, err := api.SendChatAction(Typing, chatID)

	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetChat(t *testing.T) {
	resp, err := api.GetChat(chatID)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Result.Type != "private" && resp.Result.Type != "group" &&
		resp.Result.Type != "supergroup" && resp.Result.Type != "channel" {

		t.Fatalf("wrong chat type, got: %s", resp.Result.Type)
	}
}

func TestGetChatAdministrators(t *testing.T) {
	resp, err := api.GetChatAdministrators(groupID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetChatMemberCount(t *testing.T) {
	resp, err := api.GetChatMemberCount(groupID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetChatMember(t *testing.T) {
	resp, err := api.GetChatMember(groupID, chatID)
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
	resp, err := api.SetMyCommands(nil, commands...)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetMyCommands(t *testing.T) {
	resp, err := api.GetMyCommands(nil)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}

	if !reflect.DeepEqual(resp.Result, commands) {
		t.Logf("expected commands: %v", commands)
		t.Logf("commands from API: %v", resp.Result)
		t.Fatal("error: commands mismatch")
	}
}

func TestDeleteMyCommands(t *testing.T) {
	resp, err := api.DeleteMyCommands(nil)
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
		&MessageMediaOptions{
			Media: InputMediaPhoto{
				Type:  "photo",
				Media: photoID,
			},
		},
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

func TestDeleteMessage(t *testing.T) {
	resp, err := api.DeleteMessage(chatID, msgTmp.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}
