package echotron

import "testing"

func TestInlineQueryResultArticleImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultArticle{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultPhotoImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultPhoto{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultGifImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultGif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultMpeg4GifImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultMpeg4Gif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultVideoImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultVideo{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultAudioImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultAudio{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultVoiceImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultVoice{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultDocumentImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultDocument{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultLocationImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultLocation{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultVenueImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultVenue{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultContactImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultContact{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultGameImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultGame{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedPhotoImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedPhoto{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedGifImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedGif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedMpeg4GifImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedMpeg4Gif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedStickerImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedSticker{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedDocumentImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedDocument{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedVideoImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedVideo{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedVoiceImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedVoice{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedAudioImplementsInlineQueryResult(_ *testing.T) {
	i := InlineQueryResultCachedAudio{}
	i.ImplementsInlineQueryResult()
}

func TestInputTextMessageContentImplementsInputMessageContent(_ *testing.T) {
	i := InputTextMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputLocationMessageContentImplementsInputMessageContent(_ *testing.T) {
	i := InputLocationMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputVenueMessageContentImplementsInputMessageContent(_ *testing.T) {
	i := InputVenueMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputContactMessageContentImplementsInputMessageContent(_ *testing.T) {
	i := InputContactMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestAnswerInlineQuery(t *testing.T) {
	_, err := api.AnswerInlineQuery(
		"test",
		[]InlineQueryResult{},
		nil,
	)

	if err != nil {
		t.Fatal(err)
	}
}
