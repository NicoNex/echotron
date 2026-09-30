package echotron

import (
	"encoding/json"
	"testing"
)

func TestAPIResponseBase(_ *testing.T) {
	a := APIResponseBase{}
	a.Base()
}

func TestAPIResponseUpdate(_ *testing.T) {
	a := APIResponseUpdate{}
	a.Base()
}

func TestAPIResponseUser(_ *testing.T) {
	a := APIResponseUser{}
	a.Base()
}

func TestAPIResponseMessage(_ *testing.T) {
	a := APIResponseMessage{}
	a.Base()
}

func TestAPIResponseMessageArray(_ *testing.T) {
	a := APIResponseMessageArray{}
	a.Base()
}

func TestAPIResponseMessageID(_ *testing.T) {
	a := APIResponseMessageID{}
	a.Base()
}

func TestAPIResponseCommands(_ *testing.T) {
	a := APIResponseCommands{}
	a.Base()
}

func TestAPIResponseBool(_ *testing.T) {
	a := APIResponseBool{}
	a.Base()
}

func TestAPIResponseString(_ *testing.T) {
	a := APIResponseString{}
	a.Base()
}

func TestAPIResponseChat(_ *testing.T) {
	a := APIResponseChat{}
	a.Base()
}

func TestAPIResponseInviteLink(_ *testing.T) {
	a := APIResponseInviteLink{}
	a.Base()
}

func TestAPIResponseStickerSet(_ *testing.T) {
	a := APIResponseStickerSet{}
	a.Base()
}

func TestAPIResponseUserProfile(_ *testing.T) {
	a := APIResponseUserProfile{}
	a.Base()
}

func TestAPIResponseFile(_ *testing.T) {
	a := APIResponseFile{}
	a.Base()
}

func TestAPIResponseAdministrators(_ *testing.T) {
	a := APIResponseAdministrators{}
	a.Base()
}

func TestAPIResponseChatMember(_ *testing.T) {
	a := APIResponseChatMember{}
	a.Base()
}

func TestAPIResponseInteger(_ *testing.T) {
	a := APIResponseInteger{}
	a.Base()
}

func TestAPIResponsePoll(_ *testing.T) {
	a := APIResponsePoll{}
	a.Base()
}

func TestAPIResponseGameHighScore(_ *testing.T) {
	a := APIResponseGameHighScore{}
	a.Base()
}

func TestAPIResponseWebhook(_ *testing.T) {
	a := APIResponseWebhook{}
	a.Base()
}

func TestAPIResponseSentWebAppMessage(_ *testing.T) {
	a := APIResponseSentWebAppMessage{}
	a.Base()
}

func TestAPIResponseMenuButton(_ *testing.T) {
	a := APIResponseMenuButton{}
	a.Base()
}

func TestAPIResponseChatAdministratorRights(_ *testing.T) {
	a := APIResponseChatAdministratorRights{}
	a.Base()
}

func TestAPIResponseBotDescription(_ *testing.T) {
	a := APIResponseBotDescription{}
	a.Base()
}

func TestAPIResponseBotShortDescription(_ *testing.T) {
	a := APIResponseBotShortDescription{}
	a.Base()
}

func TestAPIResponseBusinessConnection(_ *testing.T) {
	a := APIResponseBusinessConnection{}
	a.Base()
}

func TestAPIResponseStarTransactions(_ *testing.T) {
	a := APIResponseStarTransactions{}
	a.Base()
}

func TestAPIResponsePreparedInlineMessage(_ *testing.T) {
	a := APIResponsePreparedInlineMessage{}
	a.Base()
}

func TestAPIResponseGifts(_ *testing.T) {
	a := APIResponseGifts{}
	a.Base()
}

func TestInputMediaPhoto(_ *testing.T) {
	i := InputMediaPhoto{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputMediaVideo(_ *testing.T) {
	i := InputMediaVideo{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputMediaAnimation(_ *testing.T) {
	i := InputMediaAnimation{}
	i.media()
	i.thumbnail()
}

func TestInputMediaAudio(_ *testing.T) {
	i := InputMediaAudio{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputMediaDocument(_ *testing.T) {
	i := InputMediaDocument{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputPaidMediaPhoto(_ *testing.T) {
	i := InputPaidMediaPhoto{}
	i.media()
	i.groupable()
	i.thumbnail()
}

func TestInputPaidMediaVideo(_ *testing.T) {
	i := InputPaidMediaVideo{}
	i.media()
	i.groupable()
	i.thumbnail()
}

func TestBackgroundFillSolid(_ *testing.T) {
	b := BackgroundFillSolid{}
	b.ImplementsBackgroundFill()
}

func TestBackgroundFillGradient(_ *testing.T) {
	b := BackgroundFillGradient{}
	b.ImplementsBackgroundFill()
}

func TestBackgroundFillFreeformGradient(_ *testing.T) {
	b := BackgroundFillFreeformGradient{}
	b.ImplementsBackgroundFill()
}

func TestBackgroundTypeFill(_ *testing.T) {
	b := BackgroundTypeFill{}
	b.ImplementsBackgroundType()
}

func TestBackgroundTypeWallpaper(_ *testing.T) {
	b := BackgroundTypeWallpaper{}
	b.ImplementsBackgroundType()
}

func TestBackgroundTypePattern(_ *testing.T) {
	b := BackgroundTypePattern{}
	b.ImplementsBackgroundType()
}

func TestBackgroundTypeChatTheme(_ *testing.T) {
	b := BackgroundTypeChatTheme{}
	b.ImplementsBackgroundType()
}

func TestAPIResponseUserProfileAudios(_ *testing.T) {
	a := APIResponseUserProfileAudios{}
	a.Base()
}

func TestAPIResponseOwnedGifts(_ *testing.T) {
	a := APIResponseOwnedGifts{}
	a.Base()
}

func TestAPIResponseStarAmount(_ *testing.T) {
	a := APIResponseStarAmount{}
	a.Base()
}

func TestAPIResponseStory(_ *testing.T) {
	a := APIResponseStory{}
	a.Base()
}

func TestStoryAreaTypeLocation(_ *testing.T) {
	s := StoryAreaTypeLocation{}
	s.ImplementsStoryAreaType()
}

func TestStoryAreaTypeSuggestedReaction(_ *testing.T) {
	s := StoryAreaTypeSuggestedReaction{}
	s.ImplementsStoryAreaType()
}

func TestStoryAreaTypeLink(_ *testing.T) {
	s := StoryAreaTypeLink{}
	s.ImplementsStoryAreaType()
}

func TestStoryAreaTypeWeather(_ *testing.T) {
	s := StoryAreaTypeWeather{}
	s.ImplementsStoryAreaType()
}

func TestStoryAreaTypeUniqueGift(_ *testing.T) {
	s := StoryAreaTypeUniqueGift{}
	s.ImplementsStoryAreaType()
}

func TestInputStoryContentPhoto(_ *testing.T) {
	i := InputStoryContentPhoto{}
	i.ImplementsInputStoryContent()
}

func TestInputStoryContentVideo(_ *testing.T) {
	i := InputStoryContentVideo{}
	i.ImplementsInputStoryContent()
}

func TestInputProfilePhotoStatic(_ *testing.T) {
	i := InputProfilePhotoStatic{}
	i.ImplementsInputProfilePhoto()
}

func TestInputProfilePhotoAnimated(_ *testing.T) {
	i := InputProfilePhotoAnimated{}
	i.ImplementsInputProfilePhoto()
}

func TestRichTextJSON(t *testing.T) {
	in := `["plain",{"text":"b","type":"bold"},{"text":["x"],"type":"url","url":"https://t.me"}]`

	var r RichText
	if err := json.Unmarshal([]byte(in), &r); err != nil {
		t.Fatal(err)
	}

	if len(r.Texts) != 3 || r.Texts[0].Plain != "plain" || r.Texts[1].Text.Plain != "b" || r.Texts[2].Text.Texts[0].Plain != "x" {
		t.Fatalf("unexpected RichText: %+v", r)
	}

	out, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != in {
		t.Fatalf("expected %s, got %s", in, out)
	}
}

func TestRichBlockCaption(t *testing.T) {
	in := `[{"type":"table","caption":"t"},{"type":"photo","caption":{"text":"p"}}]`

	var b []RichBlock
	if err := json.Unmarshal([]byte(in), &b); err != nil {
		t.Fatal(err)
	}

	if b[0].TableCaption.Plain != "t" || b[0].Caption != nil || b[1].Caption.Text.Plain != "p" || b[1].TableCaption != nil {
		t.Fatalf("unexpected RichBlocks: %+v", b)
	}
}

func TestInputPollOptionMedia(t *testing.T) {
	o := InputPollOption{
		Text:  "a",
		Media: InputMediaPhoto{Type: MediaTypePhoto, Media: NewInputFileID("id")},
	}

	out, err := json.Marshal(o)
	if err != nil {
		t.Fatal(err)
	}

	if exp := `{"text":"a","media":{"media":"id","type":"photo"}}`; string(out) != exp {
		t.Fatalf("expected %s, got %s", exp, out)
	}
}

func TestInputRichBlockJSON(t *testing.T) {
	blocks := []InputRichBlock{
		{Type: "table", TableCaption: &RichText{Plain: "t"}, Caption: &RichBlockCaption{}},
		{
			Type:    "photo",
			Media:   InputMediaPhoto{Type: MediaTypePhoto, Media: NewInputFileURL("https://t.me/p.jpg")},
			Caption: &RichBlockCaption{Text: RichText{Plain: "c"}},
		},
	}

	out, err := json.Marshal(blocks)
	if err != nil {
		t.Fatal(err)
	}

	exp := `[{"type":"table","caption":"t"},{"caption":{"text":"c"},"photo":{"media":"https://t.me/p.jpg","type":"photo"},"type":"photo"}]`
	if string(out) != exp {
		t.Fatalf("expected %s, got %s", exp, out)
	}
}

func TestRichBlockRoundTrip(t *testing.T) {
	in := `{"type":"photo","caption":{"text":"p"}}`

	var b RichBlock
	if err := json.Unmarshal([]byte(in), &b); err != nil {
		t.Fatal(err)
	}

	out, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}

	if string(out) != in {
		t.Fatalf("expected %s, got %s", in, out)
	}
}

func TestNestedLocalMedia(t *testing.T) {
	_, err := json.Marshal(InputPollOption{
		Text:  "a",
		Media: InputMediaPhoto{Type: MediaTypePhoto, Media: NewInputFilePath("a.jpg")},
	})

	if err == nil {
		t.Fatal("expected an error for a nested local file")
	}
}

func TestLivePhotoEnvelope(t *testing.T) {
	m := InputMediaLivePhoto{Type: MediaTypeLivePhoto, Media: NewInputFileID("v"), Photo: NewInputFileID("p")}

	im, _, err := processMedia(m.media(), m.thumbnail())
	if err != nil {
		t.Fatal(err)
	}
	im.InputMedia = m

	out, err := json.Marshal(im)
	if err != nil {
		t.Fatal(err)
	}

	if exp := `{"media":"v","photo":"p","type":"live_photo"}`; string(out) != exp {
		t.Fatalf("expected %s, got %s", exp, out)
	}
}

func TestUpdateChatIDAnonymousVoter(t *testing.T) {
	u := Update{PollAnswer: &PollAnswer{VoterChat: &Chat{ID: 7}}}

	if id := u.ChatID(); id != 7 {
		t.Fatalf("expected 7, got %d", id)
	}
}
