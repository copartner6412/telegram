package telegram

// InlineQueryResultCachedAudio represents a link to an MP3 audio file stored on the Telegram servers.
//
// By default, this audio file will be sent by the user.
//
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//
// Required fields:
//   - Type
//   - ID
//   - AudioFileID
//
// See "InlineQueryResultCachedAudio" https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
type InlineQueryResultCachedAudio struct {
	// (Required) Type of the result, must be "audio".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier for the audio file.
	AudioFileID string `json:"audio_file_id"`

	*Caption

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the audio.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultCachedAudio) getQueryType() string {
	return r.Type
}
