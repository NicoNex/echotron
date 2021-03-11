/*
 * Echotron
 * Copyright (C) 2018-2021  Michele Dimaggio
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

type InlineQueryType string

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

// This object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	Id       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

type InlineQueryResult interface {
	ImplementsInlineQueryResult()
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	Title               string              `json:"title"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	Url                 string              `json:"url,omitempty"`
	HideUrl             bool                `json:"hide_url,omitempty"`
	Description         string              `json:"description,omitempty"`
	ThumbUrl            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultArticle) ImplementsInlineQueryResult() {}

// Represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	PhotoUrl            string              `json:"photo_url"`
	ThumbUrl            string              `json:"thumb_url"`
	PhotoWidth          int                 `json:"photo_width,omitempty"`
	PhotoHeight         int                 `json:"photo_height,omitempty"`
	Title               string              `json:"title,omitempty"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultPhoto) ImplementsInlineQueryResult() {}

// Represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	GifUrl              string              `json:"gif_url"`
	GifWidth            int                 `json:"gif_width,omitempty"`
	GifHeight           int                 `json:"gif_height,omitempty"`
	GifDuration         int                 `json:"gif_duration,omitempty"`
	ThumbUrl            string              `json:"thumb_url"`
	ThumbMimeType       string              `json:"thumb_mime_type,omitempty"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultGif) ImplementsInlineQueryResult() {}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	Mpeg4Url            string              `json:"mpeg4_url"`
	Mpeg4Width          int                 `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                 `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                 `json:"mpeg4_duration,omitempty"`
	ThumbUrl            string              `json:"thumb_url"`
	ThumbMimeType       string              `json:"thumb_mime_type,omitempty"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultMpeg4Gif) ImplementsInlineQueryResult() {}

// Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	VideoUrl            string              `json:"video_url"`
	MimeType            string              `json:"mime_type"`
	ThumbUrl            string              `json:"thumb_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	VideoWidth          int                 `json:"video_width,omitempty"`
	VideoHeight         int                 `json:"video_height,omitempty"`
	VideoDuration       int                 `json:"video_duration,omitempty"`
	Description         string              `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVideo) ImplementsInlineQueryResult() {}

// Represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	AudioUrl            string              `json:"audio_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	Performer           string              `json:"performer,omitempty"`
	AudioDuration       int                 `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultAudio) ImplementsInlineQueryResult() {}

// Represents a link to a voice recording in an .OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	VoiceUrl            string              `json:"voice_url"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	VoiceDuration       int                 `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVoice) ImplementsInlineQueryResult() {}

// Represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
// Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	DocumentUrl         string              `json:"document_url"`
	MimeType            string              `json:"mime_type"`
	Description         string              `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	ThumbUrl            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultDocument) ImplementsInlineQueryResult() {}

// Represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                 InlineQueryType     `json:"type"`
	Id                   string              `json:"id"`
	Latitude             float64             `json:"latitude"`
	Longitude            float64             `json:"longitude"`
	Title                string              `json:"title"`
	HorizontalAccuracy   float64             `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                 `json:"live_period,omitempty"`
	Heading              int                 `json:"heading,omitempty"`
	ProximityAlertRadius int                 `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboard     `json:"reply_markup,omitempty"`
	ThumbUrl             string              `json:"thumb_url,omitempty"`
	ThumbWidth           int                 `json:"thumb_width,omitempty"`
	ThumbHeight          int                 `json:"thumb_height,omitempty"`
	InputMessageContent  InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultLocation) ImplementsInlineQueryResult() {}

// Represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	Latitude            float64             `json:"latitude"`
	Longitude           float64             `json:"longitude"`
	Title               string              `json:"title"`
	Address             string              `json:"address"`
	FoursquareId        string              `json:"foursquare_id,omitempty"`
	FoursquareType      string              `json:"foursquare_type,omitempty"`
	GooglePlaceId       string              `json:"google_place_id,omitempty"`
	GooglePlaceType     string              `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	ThumbUrl            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultVenue) ImplementsInlineQueryResult() {}

// Represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	PhoneNumber         string              `json:"phone_number"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name,omitempty"`
	Vcard               string              `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	ThumbUrl            string              `json:"thumb_url,omitempty"`
	ThumbWidth          int                 `json:"thumb_width,omitempty"`
	ThumbHeight         int                 `json:"thumb_height,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultContact) ImplementsInlineQueryResult() {}

// Represents a Game.
type InlineQueryResultGame struct {
	Type          InlineQueryType `json:"type"`
	Id            string          `json:"id"`
	GameShortName string          `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboard `json:"reply_markup,omitempty"`
}

func (i InlineQueryResultGame) ImplementsInlineQueryResult() {}

// Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	PhotoFileId         string              `json:"photo_file_id"`
	Title               string              `json:"title,omitempty"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedPhoto) ImplementsInlineQueryResult() {}

// Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	GifFileId           string              `json:"gif_file_id"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedGif) ImplementsInlineQueryResult() {}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	Mpeg4FileId         string              `json:"mpeg4_file_id"`
	Title               string              `json:"title,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedMpeg4Gif) ImplementsInlineQueryResult() {}

// Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	StickerFileId       string              `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedSticker) ImplementsInlineQueryResult() {}

// Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	Title               string              `json:"title"`
	DocumentFileId      string              `json:"document_file_id"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedDocument) ImplementsInlineQueryResult() {}

// Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	VideoFileId         string              `json:"video_file_id"`
	Title               string              `json:"title"`
	Description         string              `json:"description,omitempty"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedVideo) ImplementsInlineQueryResult() {}

// Represents a link to a voice message stored on the Telegram servers.
// By default, this voice message will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	VoiceFileId         string              `json:"voice_file_id"`
	Title               string              `json:"title"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedVoice) ImplementsInlineQueryResult() {}

// Represents a link to an MP3 audio file stored on the Telegram servers.
// By default, this audio file will be sent by the user.
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	Type                InlineQueryType     `json:"type"`
	Id                  string              `json:"id"`
	AudioFileId         string              `json:"audio_file_id"`
	Caption             string              `json:"caption,omitempty"`
	ParseMode           string              `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity    `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard     `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (i InlineQueryResultCachedAudio) ImplementsInlineQueryResult() {}

type InputMessageContent interface {
	ImplementsInputMessageContent()
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText           string           `json:"message_text"`
	ParseMode             string           `json:"parse_mode,omitempty"`
	Entities              []*MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool             `json:"disable_web_page_preview,omitempty"`
}

func (i InputTextMessageContent) ImplementsInputMessageContent() {}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

func (i InputLocationMessageContent) ImplementsInputMessageContent() {}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareId    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceId   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

func (i InputVenueMessageContent) ImplementsInputMessageContent() {}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

func (i InputContactMessageContent) ImplementsInputMessageContent() {}
