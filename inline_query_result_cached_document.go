package telegram

// InlineQueryResultCachedDocument represents a link to a file stored on the Telegram servers.
//
// By default, this file will be sent by the user with an optional caption.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the file.
//
// Required fields:
//   - Type
//   - ID
//   - Title
//   - DocumentFileID
//
// See "InlineQueryResultCachedDocument" https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
type InlineQueryResultCachedDocument struct {
	// (Required) Type of the result, must be "document".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Title for the result.
	Title string `json:"title"`

	// (Required) A valid file identifier for the file.
	DocumentFileID string `json:"document_file_id"`

	// (Optional) Short description of the result.
	Description string `json:"description,omitempty"`

	*Caption

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the file.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultCachedDocument) getQueryType() string {
	return r.Type
}
