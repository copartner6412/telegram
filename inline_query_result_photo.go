package telegram

// InlineQueryResultPhoto represents a link to a photo. By default, this photo will be sent by the user with an optional caption.
//
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//
// Required fields:
//   - Type
//   - ID
//   - PhotoURL
//
// See "InlineQueryResultPhoto" https://core.telegram.org/bots/api#inlinequeryresultphoto
type InlineQueryResultPhoto struct {
	// (Required) Type of the result, must be "photo".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid URL of the photo. Photo must be in JPEG format. Photo size must not exceed 5MB.
	PhotoURL string `json:"photo_url"`

	// (Optional) URL of the thumbnail for the photo.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) Width of the photo.
	PhotoWidth int `json:"photo_width,omitempty"`

	// (Optional) Height of the photo.
	PhotoHeight int `json:"photo_height,omitempty"`

	// (Optional) Title for the result.
	Title string `json:"title,omitempty"`

	// (Optional) Short description of the result.
	Description string `json:"description,omitempty"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the photo.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultPhoto) getQueryType() string {
	return r.Type
}
