package echotron

import (
	"reflect"
	"testing"
)

var (
	msgTmp      *Message
	api         = NewAPI("1713461126:AAEV5sgVo513Vz4PT33mpp0ZykJqrnSluzM")
	chatID      = int64(41876271)
	photoID     = "AgACAgQAAxkDAAMrYFtODxV2LL6-kR_6qSbG9n8dIOIAAti1MRug29lSkNq_9o8PC5uMd7EnXQADAQADAgADbQADeooGAAEeBA"
	animationID = "CgACAgQAAxkDAAICQGBcoGs7GFJ-tR5AkbRRLFTbvdxXAAJ1CAAC1zHgUu-ciZqanytIHgQ"
	audioID     = "CQACAgQAAxkDAAIBCmBbamz_DqKk2GmrzmoM0SrzRN6wAAK9CAACoNvZUgPyk-87OM_YHgQ"
	documentID  = "BQACAgQAAxkDAANmYFtSXcF5kTtwgHeqVUngyuuJMx4AAnQIAAKg29lSb4HP4x-qMT8eBA"
	stickerID   = "CAACAgIAAxkBAAICImBclqwLQdFHZo15R1zU0vYC1JMFAAImAwACtXHaBj4ZC4vnHBlAHgQ"
	videoID     = "BAACAgQAAxkDAANxYFtaxF1kfc7nVY_Mtfba3u5dMooAAoYIAAKg29lSpwABJrcveXZlHgQ"
	videoNoteID = "DQACAgQAAxkDAAIBumBbfT5jPC_cvyEcr0_8DpmFDz2PAALVCgACOX7hUjGZ_MmnZVVeHgQ"
	voiceID     = "AwACAgQAAxkDAAPXYFtmoFriwJFVGDgPPpfUBljgnYAAAq8IAAKg29lStEWfrNMMAxgeBA"

	commands = []BotCommand{
		{Command: "test1", Description: "Test command 1"},
		{Command: "test2", Description: "Test command 2"},
		{Command: "test3", Description: "Test command 3"},
	}
)

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

func TestGetStickerSet(t *testing.T) {
	resp, err := api.GetStickerSet("RickAndMorty")
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendMessage(t *testing.T) {
	resp, err := api.SendMessage("TestSendMessage", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
	msgTmp = resp.Result
}

func TestSendMessageReply(t *testing.T) {
	resp, err := api.SendMessageReply("TestSendMessageReply", chatID, msgTmp.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendMessageWithKeyboard(t *testing.T) {
	kbd := api.KeyboardMarkup(false, true, false,
		api.KeyboardRow(
			api.KeyboardButton("test 1", false, false),
			api.KeyboardButton("test 2", false, false),
		),
		api.KeyboardRow(
			api.KeyboardButton("test 3", false, false),
			api.KeyboardButton("test 4", false, false),
		),
	)

	resp, err := api.SendMessageWithKeyboard("TestSendMessageWithKeyboard", chatID, kbd)
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

func TestSendPhoto(t *testing.T) {
	resp, err := api.SendPhoto("tests/echotron_test.png", "TestSendPhoto", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendPhotoByID(t *testing.T) {
	resp, err := api.SendPhotoByID(photoID, "TestSendPhotoByID", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendPhotoWithKeyboard(t *testing.T) {
	kbd := api.KeyboardMarkup(false, true, false,
		api.KeyboardRow(
			api.KeyboardButton("test 1", false, false),
			api.KeyboardButton("test 2", false, false),
		),
		api.KeyboardRow(
			api.KeyboardButton("test 3", false, false),
			api.KeyboardButton("test 4", false, false),
		),
	)

	resp, err := api.SendPhotoWithKeyboard("tests/echotron_test.png", "TestSendPhotoWithKeyboard", chatID, kbd)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudio(t *testing.T) {
	resp, err := api.SendAudio("tests/audio.mp3", "TestSendAudio", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAudioByID(t *testing.T) {
	resp, err := api.SendAudioByID(audioID, "TestSendAudioByID", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendDocument(t *testing.T) {
	resp, err := api.SendDocument("tests/document.pdf", "TestSendDocument", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendDocumentByID(t *testing.T) {
	resp, err := api.SendDocumentByID(documentID, "TestSendDocumentByID", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideo(t *testing.T) {
	resp, err := api.SendVideo("tests/video.webm", "TestSendVideo", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoByID(t *testing.T) {
	resp, err := api.SendVideoByID(videoID, "TestSendVideoByID", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoNote(t *testing.T) {
	resp, err := api.SendVideoNote("tests/video_note.mp4", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVideoNoteByID(t *testing.T) {
	resp, err := api.SendVideoNoteByID(videoNoteID, chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVoice(t *testing.T) {
	resp, err := api.SendVoice("tests/audio.mp3", "TestSendVoice", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendVoiceByID(t *testing.T) {
	resp, err := api.SendVoiceByID(voiceID, "TestSendVoiceByID", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendContact(t *testing.T) {
	resp, err := api.SendContact("1234567890", "Name", "Surname", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendSticker(t *testing.T) {
	resp, err := api.SendSticker(stickerID, chatID)
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

func TestSetMyCommands(t *testing.T) {
	resp, err := api.SetMyCommands(commands...)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestGetMyCommands(t *testing.T) {
	resp, err := api.GetMyCommands()
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

func TestSendAnimation(t *testing.T) {
	resp, err := api.SendAnimation("tests/animation.mp4", "TestSendAnimation", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}

func TestSendAnimationByID(t *testing.T) {
	resp, err := api.SendAnimationByID(animationID, "TestSendAnimationByID", chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !resp.Ok {
		t.Fatal(resp.ErrorCode, resp.Description)
	}
}
