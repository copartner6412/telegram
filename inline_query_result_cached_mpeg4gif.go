package telegram

// InlineQueryResultCachedMpeg4Gif represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the animation.
//
// Required fields:
//   - Type
//   - ID
//   - Mpeg4FileID
//
// See "InlineQueryResultCachedMpeg4Gif" https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
type InlineQueryResultCachedMpeg4Gif struct {
	// (Required) Type of the result, must be "mpeg4_gif".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier for the MPEG-4 file.
	Mpeg4FileID string `json:"mpeg4_file_id"`

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

func (r InlineQueryResultCachedMpeg4Gif) getQueryType() string {
	return r.Type
}
