package telegram

// InlineQueryResultGif represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//
// Required fields:
//   - Type
//   - ID
//   - GIFURL
//
// See "InlineQueryResultGif" https://core.telegram.org/bots/api#inlinequeryresultgif
type InlineQueryResultGif struct {
	// (Required) Type of the result, must be "gif".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid URL for the GIF file. File size must not exceed 1MB.
	GIFURL string `json:"gif_url"`

	// (Optional) Width of the GIF.
	GIFWidth int `json:"gif_width,omitempty"`

	// (Optional) Height of the GIF.
	GIFHeight int `json:"gif_height,omitempty"`

	// (Optional) Duration of the GIF in seconds.
	GIFDuration int `json:"gif_duration,omitempty"`

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

	// (Optional) Content of the message to be sent instead of the GIF animation.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultGif) getQueryType() string {
	return r.Type
}
