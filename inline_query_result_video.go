package telegram

// InlineQueryResultVideo represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption.
//
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content using input_message_content.
//
// Required fields:
//   - Type
//   - ID
//   - VideoURL
//   - MimeType
//   - Thumbnail
//
// See "InlineQueryResultVideo" https://core.telegram.org/bots/api#inlinequeryresultvideo
type InlineQueryResultVideo struct {
	// (Required) Type of the result, must be "video".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid URL for the embedded video player or video file.
	VideoURL string `json:"video_url"`

	// (Required) MIME type of the content of the video URL, “text/html” or “video/mp4”.
	MimeType string `json:"mime_type"`

	// (Required) URL of the thumbnail (JPEG only) for the video.
	ThumbnailURL string `json:"thumbnail_url"`

	// (Optional) Title for the result.
	Title string `json:"title,omitempty"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Video width.
	VideoWidth int `json:"video_width,omitempty"`

	// (Optional) Video height.
	VideoHeight int `json:"video_height,omitempty"`

	// (Optional) Video duration in seconds.
	VideoDuration int `json:"video_duration,omitempty"`

	// (Optional) Short description of the result.
	Description string `json:"description,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultVideo) getQueryType() string {
	return r.Type
}
