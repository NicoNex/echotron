package echotron

import "testing"

func TestInlineQueryResultArticle(_ *testing.T) {
	i := InlineQueryResultArticle{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultPhoto(_ *testing.T) {
	i := InlineQueryResultPhoto{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultGif(_ *testing.T) {
	i := InlineQueryResultGif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultMpeg4Gif(_ *testing.T) {
	i := InlineQueryResultMpeg4Gif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultVideo(_ *testing.T) {
	i := InlineQueryResultVideo{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultAudio(_ *testing.T) {
	i := InlineQueryResultAudio{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultVoice(_ *testing.T) {
	i := InlineQueryResultVoice{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultDocument(_ *testing.T) {
	i := InlineQueryResultDocument{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultLocation(_ *testing.T) {
	i := InlineQueryResultLocation{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultVenue(_ *testing.T) {
	i := InlineQueryResultVenue{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultContact(_ *testing.T) {
	i := InlineQueryResultContact{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultGame(_ *testing.T) {
	i := InlineQueryResultGame{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedPhoto(_ *testing.T) {
	i := InlineQueryResultCachedPhoto{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedGif(_ *testing.T) {
	i := InlineQueryResultCachedGif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedMpeg4Gif(_ *testing.T) {
	i := InlineQueryResultCachedMpeg4Gif{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedSticker(_ *testing.T) {
	i := InlineQueryResultCachedSticker{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedDocument(_ *testing.T) {
	i := InlineQueryResultCachedDocument{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedVideo(_ *testing.T) {
	i := InlineQueryResultCachedVideo{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedVoice(_ *testing.T) {
	i := InlineQueryResultCachedVoice{}
	i.ImplementsInlineQueryResult()
}

func TestInlineQueryResultCachedAudio(_ *testing.T) {
	i := InlineQueryResultCachedAudio{}
	i.ImplementsInlineQueryResult()
}

func TestInputTextMessageContent(_ *testing.T) {
	i := InputTextMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputLocationMessageContent(_ *testing.T) {
	i := InputLocationMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputVenueMessageContent(_ *testing.T) {
	i := InputVenueMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputContactMessageContent(_ *testing.T) {
	i := InputContactMessageContent{}
	i.ImplementsInputMessageContent()
}

func TestInputInvoiceMessageContent(_ *testing.T) {
	i := InputInvoiceMessageContent{}
	i.ImplementsInputMessageContent()
}
