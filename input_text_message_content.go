package telegram

// InputTextMessageContent represents the content of a text message to be sent as the result of an inline query.
//
// Required fields:
//   - MessageText
//
// See "InputTextMessageContent" https://core.telegram.org/bots/api#inputtextmessagecontent
type InputTextMessageContent struct {
	// (Required) Text of the message to be sent, 1-4096 characters.
	MessageText string `json:"message_text"`

	// (Optional) Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode string `json:"parse_mode,omitempty"`

	// (Optional) List of special entities that appear in message text, which can be specified instead of parse_mode.
	Entities []MessageEntity `json:"entities,omitempty"`

	// (Optional) Link preview generation options for the message.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (i InputTextMessageContent) isInputMessageContent() bool {
	return true
}
