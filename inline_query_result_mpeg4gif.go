package telegram

// InlineQueryResultMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound).
//
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// Required fields:
//   - Type
//   - ID
//   - Mpeg4URL
//
// See "InlineQueryResultMpeg4Gif" https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
type InlineQueryResultMpeg4Gif struct {
	// (Required) Type of the result, must be "mpeg4_gif".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid URL for the MPEG4 file. File size must not exceed 1MB.
	Mpeg4URL string `json:"mpeg4_url"`

	// (Optional) Video width.
	Mpeg4Width int `json:"mpeg4_width,omitempty"`

	// (Optional) Video height.
	Mpeg4Height int `json:"mpeg4_height,omitempty"`

	// (Optional) Video duration in seconds.
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`

	// (Optional) URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”. Defaults to “image/jpeg”.
	ThumbnailMimeType string `json:"thumbnail_mime_type,omitempty"`

	// (Optional) Title for the result.
	Title string `json:"title,omitempty"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the video animation.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultMpeg4Gif) getQueryType() string {
	return r.Type
}
