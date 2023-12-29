/*
 * Echotron
 * Copyright (C) 2018-2022 The Echotron Devs
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
	"encoding/json"
	"net/url"
)

// InlineQueryType is a custom type for the various InlineQueryResult*'s Type field.
type InlineQueryType string

// These are all the possible types for the various InlineQueryResult*'s Type field.
const (
	InlineArticle  InlineQueryType = "article"
	InlinePhoto                    = "photo"
	InlineGIF                      = "gif"
	InlineMPEG4GIF                 = "mpeg4_gif"
	InlineVideo                    = "video"
	InlineAudio                    = "audio"
	InlineVoice                    = "voice"
	InlineDocument                 = "document"
	InlineLocation                 = "location"
	InlineVenue                    = "venue"
	InlineContact                  = "contact"
	InlineGame                     = "game"
	InlineSticker                  = "sticker"
)

// InlineQuery represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	ID       string    `json:"id"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// InlineQueryResult represents an interface that implements all the various InlineQueryResult* types.
type InlineQueryResult interface {
	ImplementsInlineQueryResult()
}

// InlineQueryResultArticle represents a link to an article or web page.
type InlineQueryResultArticle struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	Description         string              `json:"description,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	URL                 string              `json:"url,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
	HideURL             bool                `json:"hide_url,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultArticle) ImplementsInlineQueryResult() {}

// InlineQueryResultPhoto represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Title               string              `json:"title,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url"`
	PhotoURL            string              `json:"photo_url"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	ID                  string              `json:"id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	Type                InlineQueryType     `json:"type"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	PhotoHeight         int                 `json:"photo_height,omitempty"`
	PhotoWidth          int                 `json:"photo_width,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultPhoto) ImplementsInlineQueryResult() {}

// InlineQueryResultGif represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	Title               string              `json:"title,omitempty"`
	GifURL              string              `json:"gif_url"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url"`
	ID                  string              `json:"id"`
	ThumbnailMimeType   string              `json:"thumbnail_mime_type,omitempty"`
	Type                InlineQueryType     `json:"type"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	GifDuration         int                 `json:"gif_duration,omitempty"`
	GifHeight           int                 `json:"gif_height,omitempty"`
	GifWidth            int                 `json:"gif_width,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultGif) ImplementsInlineQueryResult() {}

// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	Title               string              `json:"title,omitempty"`
	Mpeg4URL            string              `json:"mpeg4_url"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url"`
	ID                  string              `json:"id"`
	ThumbnailMimeType   string              `json:"thumbnail_mime_type,omitempty"`
	Type                InlineQueryType     `json:"type"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	Mpeg4Duration       int                 `json:"mpeg4_duration,omitempty"`
	Mpeg4Height         int                 `json:"mpeg4_height,omitempty"`
	Mpeg4Width          int                 `json:"mpeg4_width,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultMpeg4Gif) ImplementsInlineQueryResult() {}

// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	Description         string              `json:"description,omitempty"`
	MimeType            string              `json:"mime_type"`
	ThumbnailURL        string              `json:"thumbnail_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ID                  string              `json:"id"`
	VideoURL            string              `json:"video_url"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	Type                InlineQueryType     `json:"type"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	VideoHeight         int                 `json:"video_height,omitempty"`
	VideoDuration       int                 `json:"video_duration,omitempty"`
	VideoWidth          int                 `json:"video_width,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultVideo) ImplementsInlineQueryResult() {}

// InlineQueryResultAudio represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	AudioURL            string              `json:"audio_url"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	Performer           string              `json:"performer,omitempty"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	AudioDuration       int                 `json:"audio_duration,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultAudio) ImplementsInlineQueryResult() {}

// InlineQueryResultVoice represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	VoiceURL            string              `json:"voice_url"`
	Title               string              `json:"title"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	VoiceDuration       int                 `json:"voice_duration,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultVoice) ImplementsInlineQueryResult() {}

// InlineQueryResultDocument represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	MimeType            string              `json:"mime_type"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	DocumentURL         string              `json:"document_url"`
	Title               string              `json:"title"`
	Description         string              `json:"description,omitempty"`
	ID                  string              `json:"id"`
	Type                InlineQueryType     `json:"type"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultDocument) ImplementsInlineQueryResult() {}

// InlineQueryResultLocation represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	InputMessageContent  InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
	ID                   string              `json:"id"`
	ThumbnailURL         string              `json:"thumbnail_url,omitempty"`
	Title                string              `json:"title"`
	Type                 InlineQueryType     `json:"type"`
	LivePeriod           int                 `json:"live_period,omitempty"`
	HorizontalAccuracy   float64             `json:"horizontal_accuracy,omitempty"`
	ProximityAlertRadius int                 `json:"proximity_alert_radius,omitempty"`
	Longitude            float64             `json:"longitude"`
	Latitude             float64             `json:"latitude"`
	ThumbnailWidth       int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int                 `json:"thumbnail_height,omitempty"`
	Heading              int                 `json:"heading,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultLocation) ImplementsInlineQueryResult() {}

// InlineQueryResultVenue represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	GooglePlaceType     string              `json:"google_place_type,omitempty"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	Title               string              `json:"title"`
	Address             string              `json:"address"`
	FoursquareID        string              `json:"foursquare_id,omitempty"`
	ID                  string              `json:"id"`
	GooglePlaceID       string              `json:"google_place_id,omitempty"`
	FoursquareType      string              `json:"foursquare_type,omitempty"`
	Type                InlineQueryType     `json:"type"`
	Longitude           float64             `json:"longitude"`
	Latitude            float64             `json:"latitude"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultVenue) ImplementsInlineQueryResult() {}

// InlineQueryResultContact represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	ID                  string              `json:"id"`
	PhoneNumber         string              `json:"phone_number"`
	FirstName           string              `json:"first_name"`
	VCard               string              `json:"vcard,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ThumbnailURL        string              `json:"thumbnail_url,omitempty"`
	LastName            string              `json:"last_name,omitempty"`
	ThumbnailWidth      int                 `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                 `json:"thumbnail_height,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultContact) ImplementsInlineQueryResult() {}

// InlineQueryResultGame represents a Game.
type InlineQueryResultGame struct {
	ReplyMarkup   ReplyMarkup     `json:"reply_markup,omitempty"`
	Type          InlineQueryType `json:"type"`
	ID            string          `json:"id"`
	GameShortName string          `json:"game_short_name"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultGame) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	PhotoFileID         string              `json:"photo_file_id"`
	Title               string              `json:"title,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedPhoto) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedGif represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	ID                  string              `json:"id"`
	GifFileID           string              `json:"gif_file_id"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedGif) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	ID                  string              `json:"id"`
	Mpeg4FileID         string              `json:"mpeg4_file_id"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedMpeg4Gif) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	StickerFileID       string              `json:"sticker_file_id"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedSticker) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedDocument represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	Title               string              `json:"title"`
	DocumentFileID      string              `json:"document_file_id"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedDocument) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	VideoFileID         string              `json:"video_file_id"`
	Title               string              `json:"title"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedVideo) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	Type                InlineQueryType     `json:"type"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	ID                  string              `json:"id"`
	VoiceFileID         string              `json:"voice_file_id"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedVoice) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedAudio represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
	AudioFileID         string              `json:"audio_file_id"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedAudio) ImplementsInlineQueryResult() {}

// InputMessageContent represents an interface that implements all the various Input*MessageContent types.
type InputMessageContent interface {
	ImplementsInputMessageContent()
}

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          string              `json:"parse_mode,omitempty"`
	Entities           []*MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

// ImplementsInputMessageContent is used to implement the InputMessageContent interface.
func (i InputTextMessageContent) ImplementsInputMessageContent() {}

// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

// ImplementsInputMessageContent is used to implement the InputMessageContent interface.
func (i InputLocationMessageContent) ImplementsInputMessageContent() {}

// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
}

// ImplementsInputMessageContent is used to implement the InputMessageContent interface.
func (i InputVenueMessageContent) ImplementsInputMessageContent() {}

// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// ImplementsInputMessageContent is used to implement the InputMessageContent interface.
func (i InputContactMessageContent) ImplementsInputMessageContent() {}

// InlineQueryResultsButton represents a button to be shown above inline query results.
// You MUST use exactly one of the fields.
type InlineQueryResultsButton struct {
	WebApp         WebAppInfo `json:"web_app,omitempty"`
	StartParameter string     `json:"start_parameter,omitempty"`
	Text           string     `json:"text"`
}

// InlineQueryOptions is a custom type which contains the various options required by the AnswerInlineQuery method.
type InlineQueryOptions struct {
	Button     InlineQueryResultsButton `query:"button"`
	NextOffset string                   `query:"next_offset"`
	CacheTime  int                      `query:"cache_time"`
	IsPersonal bool                     `query:"is_personal"`
}

// AnswerInlineQuery is used to send answers to an inline query.
func (a API) AnswerInlineQuery(inlineQueryID string, results []InlineQueryResult, opts *InlineQueryOptions) (res APIResponseBase, err error) {
	var vals = make(url.Values)

	jsn, _ := json.Marshal(results)
	vals.Set("inline_query_id", inlineQueryID)
	vals.Set("results", string(jsn))
	return get[APIResponseBase](a.base, "answerInlineQuery", addValues(vals, opts))
}
