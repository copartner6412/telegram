package telegram

// InlineQueryResultCachedVideo represents a link to a video file stored on the Telegram servers.
//
// By default, this video file will be sent by the user with an optional caption.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the video.
//
// Required fields:
//   - Type
//   - ID
//   - VideoFileID
//   - Title
//
// See "InlineQueryResultCachedVideo" https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
type InlineQueryResultCachedVideo struct {
	// (Required) Type of the result, must be "video".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier for the video file.
	VideoFileID string `json:"video_file_id"`

	// (Required) Title for the result.
	Title string `json:"title"`

	// (Optional) Short description of the result.
	Description string `json:"description,omitempty"`

	*Caption

	// (Optional) Pass True if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the video.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultCachedVideo) getQueryType() string {
	return r.Type
}
