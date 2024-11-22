package telegram

// InlineQueryResultDocument represents a link to a file.
//
// By default, this file will be sent by the user with an optional caption.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
//
// Currently, only .PDF and .ZIP files can be sent using this method.
//
// Required fields:
//   - Type
//   - ID
//   - Title
//   - MimeType
//
// See "InlineQueryResultDocument" https://core.telegram.org/bots/api#inlinequeryresultdocument
type InlineQueryResultDocument struct {
	// (Required) Type of the result, must be "document".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Title for the result.
	Title string `json:"title"`

	// (Required) A valid URL for the file.
	DocumentURL string `json:"document_url"`

	// (Required) MIME type of the content of the file, either "application/pdf" or "application/zip".
	MimeType string `json:"mime_type"`

	*Caption

	// (Optional) Short description of the result.
	Description string `json:"description,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the file.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// (Optional) URL of the thumbnail (JPEG only) for the file.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) Thumbnail width.
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// (Optional) Thumbnail height.
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

func (r InlineQueryResultDocument) getQueryType() string {
	return r.Type
}
