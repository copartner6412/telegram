package telegram

// InlineQueryResultVoice represents a link to a voice recording in an .OGG container encoded with OPUS.
//
// By default, this voice recording will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the voice message.
//
// Required fields:
//   - Type
//   - ID
//   - VoiceURL
//   - Title
//
// See "InlineQueryResultVoice" https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResultVoice struct {
	// (Required) Type of the result, must be "voice".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid URL for the voice recording.
	VoiceURL string `json:"voice_url"`

	// (Required) Recording title.
	Title string `json:"title"`

	// (Optional) Caption for the voice recording, 0-1024 characters after entities parsing.
	Caption string `json:"caption,omitempty"`

	// (Optional) Mode for parsing entities in the voice message caption. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// (Optional) List of special entities that appear in the caption, which can be specified instead of parse_mode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// (Optional) Recording duration in seconds.
	VoiceDuration int `json:"voice_duration,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the voice recording.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultVoice) getQueryType() string {
	return r.Type
}
