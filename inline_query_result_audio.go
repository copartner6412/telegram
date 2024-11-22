package telegram

// InlineQueryResultAudio represents a link to an MP3 audio file.
//
// By default, this audio file will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the audio.
//
// Required fields:
//   - Type
//   - ID
//   - AudioURL
//   - Title
//
// See "InlineQueryResultAudio" https://core.telegram.org/bots/api#inlinequeryresultaudio
type InlineQueryResultAudio struct {
	// (Required) Type of the result, must be "audio".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid URL for the audio file.
	AudioURL string `json:"audio_url"`

	// (Required) Title of the audio file.
	Title string `json:"title"`

	*Caption

	// (Optional) Performer of the audio file.
	Performer string `json:"performer,omitempty"`

	// (Optional) Audio duration in seconds.
	AudioDuration int `json:"audio_duration,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the audio.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultAudio) getQueryType() string {
	return r.Type
}
