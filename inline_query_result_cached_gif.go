package telegram

// InlineQueryResultCachedGif represents a link to an animated GIF file stored on the Telegram servers.
//
// By default, this animated GIF file will be sent by the user with an optional caption.
//
// Alternatively, you can use InputMessageContent to send a message with specified content instead of the animation.
//
// Required fields:
//   - Type
//   - ID
//   - GifFileID
//
// See "InlineQueryResultCachedGif" https://core.telegram.org/bots/api#inlinequeryresultcachedgif
type InlineQueryResultCachedGif struct {
	// (Required) Type of the result, must be "gif".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier for the GIF file.
	GifFileID string `json:"gif_file_id"`

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

func (r InlineQueryResultCachedGif) getQueryType() string {
	return r.Type
}
