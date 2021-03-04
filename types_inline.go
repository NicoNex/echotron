/*
 * Echotron
 * Copyright (C) 2021  Michele Dimaggio
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

type InlineQueryResultArticle struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	InputMessageContent *InputMessageContent  `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	Url                 string                `json:"url,omitempty"`
	HideUrl             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	PhotoUrl            string                `json:"photo_url"`
	ThumbUrl            string                `json:"thumb_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultGif struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	GifUrl              string                `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	GifDuration         int                   `json:"gif_duration,omitempty"`
	ThumbUrl            string                `json:"thumb_url"`
	ThumbMimeType       string                `json:"thumb_mime_type,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultMpeg4Gif struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Mpeg4Url            string                `json:"mpeg4_url"`
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                   `json:"mpeg4_duration,omitempty"`
	ThumbUrl            string                `json:"thumb_url"`
	ThumbMimeType       string                `json:"thumb_mime_type,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVideo struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VideoUrl            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultAudio struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	AudioUrl            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultVoice struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VoiceUrl            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultDocument struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	DocumentUrl         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

type InlineQueryResultLocation struct {
	Type                 string                `json:"type"`
	Id                   string                `json:"id"`
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                   `json:"live_period,omitempty"`
	Heading              int                   `json:"heading,omitempty"`
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl             string                `json:"thumb_url,omitempty"`
	ThumbWidth           int                   `json:"thumb_width,omitempty"`
	ThumbHeight          int                   `json:"thumb_height,omitempty"`
}

type InlineQueryResultVenue struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareId        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	GooglePlaceId       string                `json:"google_place_id,omitempty"`
	GooglePlaceType     string                `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

type InlineQueryResultContact struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	Vcard               string                `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

type InlineQueryResultGame struct {
	Type          string                `json:"type"`
	Id            string                `json:"id"`
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboard       `json:"reply_markup,omitempty"`
}

type InlineQueryResultCachedPhoto struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	PhotoFileId         string                `json:"photo_file_id"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedGif struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	GifFileId           string                `json:"gif_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Mpeg4FileId         string                `json:"mpeg4_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedSticker struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	StickerFileId       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedDocument struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	Title               string                `json:"title"`
	DocumentFileId      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVideo struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VideoFileId         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedVoice struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	VoiceFileId         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InlineQueryResultCachedAudio struct {
	Type                string                `json:"type"`
	Id                  string                `json:"id"`
	AudioFileId         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []*MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboard       `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

type InputTextMessageContent struct {
	MessageText           string           `json:"message_text"`
	ParseMode             string           `json:"parse_mode,omitempty"`
	Entities              []*MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool             `json:"disable_web_page_preview,omitempty"`
}

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

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

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}
