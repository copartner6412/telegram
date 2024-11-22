package telegram

// InlineQueryResultCachedVoice represents a link to a voice message stored on the Telegram servers.
//
// By default, this voice message will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the voice message.
//
// Required fields:
//   - Type
//   - ID
//   - VoiceFileID
//   - Title
//
// See "InlineQueryResultCachedVoice" https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
type InlineQueryResultCachedVoice struct {
	// (Required) Type of the result, must be "voice".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier for the voice message.
	VoiceFileID string `json:"voice_file_id"`

	// (Required) Voice message title.
	Title string `json:"title"`

	*Caption

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the voice message.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultCachedVoice) getQueryType() string {
	return r.Type
}
