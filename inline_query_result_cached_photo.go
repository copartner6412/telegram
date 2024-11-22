package telegram

// InlineQueryResultCachedPhoto represents a link to a photo stored on the Telegram servers.
//
// By default, this photo will be sent by the user with an optional caption.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the photo.
//
// Required fields:
//   - Type
//   - ID
//   - PhotoFileID
//
// See "InlineQueryResultCachedPhoto" https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
type InlineQueryResultCachedPhoto struct {
	// (Required) Type of the result, must be "photo".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier of the photo.
	PhotoFileID string `json:"photo_file_id"`

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

func (r InlineQueryResultCachedPhoto) getQueryType() string {
	return r.Type
}
