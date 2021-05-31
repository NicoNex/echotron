/*
 * Echotron
 * Copyright (C) 2018-2021  The Echotron Devs
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
	"fmt"
)

// InlineQueryType is a custom type for the various InlineQueryResult*'s Type field.
type InlineQueryType string

// These are all the possible types for the various InlineQueryResult*'s Type field.
const (
	ARTICLE  InlineQueryType = "article"
	PHOTO                    = "photo"
	GIF                      = "gif"
	MPEG4GIF                 = "mpeg4_gif"
	VIDEO                    = "video"
	AUDIO                    = "audio"
	VOICE                    = "voice"
	DOCUMENT                 = "document"
	LOCATION                 = "location"
	VENUE                    = "venue"
	CONTACT                  = "contact"
	GAME                     = "game"
	STICKER                  = "sticker"
)

// InlineQuery represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
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
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	URL                 string              `json:"url,omitempty"`
	HideURL             bool                `json:"hide_url,omitempty"`
	Description         string              `json:"description,omitempty"`
	ThumbURL            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultArticle) ImplementsInlineQueryResult() {}

// InlineQueryResultPhoto represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	PhotoURL            string              `json:"photo_url"`
	ThumbURL            string              `json:"thumb_url"`
	PhotoWidth          int                 `json:"photo_width,omitempty"`
	PhotoHeight         int                 `json:"photo_height,omitempty"`
	Title               string              `json:"title,omitempty"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultPhoto) ImplementsInlineQueryResult() {}

// InlineQueryResultGif represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	GifURL              string              `json:"gif_url"`
	GifWidth            int                 `json:"gif_width,omitempty"`
	GifHeight           int                 `json:"gif_height,omitempty"`
	GifDuration         int                 `json:"gif_duration,omitempty"`
	ThumbURL            string              `json:"thumb_url"`
	ThumbMimeType       string              `json:"thumb_mime_type,omitempty"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultGif) ImplementsInlineQueryResult() {}

// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Mpeg4URL            string              `json:"mpeg4_url"`
	Mpeg4Width          int                 `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                 `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                 `json:"mpeg4_duration,omitempty"`
	ThumbURL            string              `json:"thumb_url"`
	ThumbMimeType       string              `json:"thumb_mime_type,omitempty"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultMpeg4Gif) ImplementsInlineQueryResult() {}

// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	VideoURL            string              `json:"video_url"`
	MimeType            string              `json:"mime_type"`
	ThumbURL            string              `json:"thumb_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	VideoWidth          int                 `json:"video_width,omitempty"`
	VideoHeight         int                 `json:"video_height,omitempty"`
	VideoDuration       int                 `json:"video_duration,omitempty"`
	Description         string              `json:"description,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultVideo) ImplementsInlineQueryResult() {}

// InlineQueryResultAudio represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	AudioURL            string              `json:"audio_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	Performer           string              `json:"performer,omitempty"`
	AudioDuration       int                 `json:"audio_duration,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultAudio) ImplementsInlineQueryResult() {}

// InlineQueryResultVoice represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	VoiceURL            string              `json:"voice_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	VoiceDuration       int                 `json:"voice_duration,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultVoice) ImplementsInlineQueryResult() {}

// InlineQueryResultDocument represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	DocumentURL         string              `json:"document_url"`
	MimeType            string              `json:"mime_type"`
	Description         string              `json:"description,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	ThumbURL            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultDocument) ImplementsInlineQueryResult() {}

// InlineQueryResultLocation represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                 InlineQueryType     `json:"type"`
	ID                   string              `json:"id"`
	Latitude             float64             `json:"latitude"`
	Longitude            float64             `json:"longitude"`
	Title                string              `json:"title"`
	HorizontalAccuracy   float64             `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                 `json:"live_period,omitempty"`
	Heading              int                 `json:"heading,omitempty"`
	ProximityAlertRadius int                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *ReplyMarkup        `json:"reply_markup,omitempty"`
	ThumbURL             string              `json:"thumb_url,omitempty"`
	ThumbWidth           int                 `json:"thumb_width,omitempty"`
	ThumbHeight          int                 `json:"thumb_height,omitempty"`
	InputMessageContent  InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultLocation) ImplementsInlineQueryResult() {}

// InlineQueryResultVenue represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Latitude            float64             `json:"latitude"`
	Longitude           float64             `json:"longitude"`
	Title               string              `json:"title"`
	Address             string              `json:"address"`
	FoursquareID        string              `json:"foursquare_id,omitempty"`
	FoursquareType      string              `json:"foursquare_type,omitempty"`
	GooglePlaceID       string              `json:"google_place_id,omitempty"`
	GooglePlaceType     string              `json:"google_place_type,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	ThumbURL            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultVenue) ImplementsInlineQueryResult() {}

// InlineQueryResultContact represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	PhoneNumber         string              `json:"phone_number"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name,omitempty"`
	VCard               string              `json:"vcard,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	ThumbURL            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultContact) ImplementsInlineQueryResult() {}

// InlineQueryResultGame represents a Game.
type InlineQueryResultGame struct {
	Type          InlineQueryType `json:"type"`
	ID            string          `json:"id"`
	GameShortName string          `json:"game_short_name"`
	ReplyMarkup   *ReplyMarkup    `json:"reply_markup,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultGame) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	PhotoFileID         string              `json:"photo_file_id"`
	Title               string              `json:"title,omitempty"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedPhoto) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedGif represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	GifFileID           string              `json:"gif_file_id"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedGif) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Mpeg4FileID         string              `json:"mpeg4_file_id"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedMpeg4Gif) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	StickerFileID       string              `json:"sticker_file_id"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedSticker) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedDocument represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	Title               string              `json:"title"`
	DocumentFileID      string              `json:"document_file_id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedDocument) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	VideoFileID         string              `json:"video_file_id"`
	Title               string              `json:"title"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedVideo) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	VoiceFileID         string              `json:"voice_file_id"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedVoice) ImplementsInlineQueryResult() {}

// InlineQueryResultCachedAudio represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	Type                InlineQueryType     `json:"type"`
	ID                  string              `json:"id"`
	AudioFileID         string              `json:"audio_file_id"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *ReplyMarkup        `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// ImplementsInlineQueryResult is used to implement the InlineQueryResult interface.
func (i InlineQueryResultCachedAudio) ImplementsInlineQueryResult() {}

// InputMessageContent represents an interface that implements all the various Input*MessageContent types.
type InputMessageContent interface {
	ImplementsInputMessageContent()
}

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText           string           `json:"message_text"`
	ParseMode             string           `json:"parse_mode,omitempty"`
	Entities              []*MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool             `json:"disable_web_page_preview,omitempty"`
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
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
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

// InlineQueryOptions is a custom type which contains the various options required by the API.AnswerInlineQuery method.
type InlineQueryOptions struct {
	CacheTime         int    `json:"cache_time"`
	IsPersonal        bool   `json:"is_personal"`
	NextOffset        string `json:"next_offset"`
	SwitchPmText      string `json:"switch_pm_text"`
	SwitchPmParameter string `json:"switch_pm_parameter"`
}

// AnswerInlineQuery is used to send answers to an inline query.
func (a API) AnswerInlineQuery(inlineQueryID string, results []InlineQueryResult, opts *InlineQueryOptions) (APIResponseBase, error) {
	var res APIResponseBase
	jsn, _ := json.Marshal(results)

	var url = fmt.Sprintf(
		"%sanswerInlineQuery?inline_query_id=%s&results=%s&%s",
		string(a),
		inlineQueryID,
		jsn,
		querify(opts),
	)

	content, err := sendGetRequest(url)
	if err != nil {
		return res, err
	}
	json.Unmarshal(content, &res)
	return res, nil
}
