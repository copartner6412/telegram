package telegram

type Caption struct {
	// (Optional) Caption (may also be used when resending file by file_id), 0-1024 characters after entities parsing
	Text string `json:"caption,omitempty"`

	// (Optional) Mode for parsing entities in the caption.
	//
	// See formatting options for more details. https://core.telegram.org/bots/api#formatting-options
	ParseMode string `json:"parse_mode,omitempty"`

	// (Optional) A list of special entities that appear in the caption, which can be specified instead of parse_mode
	Entities []MessageEntity `json:"caption_entities,omitempty"`
}

type Text struct {
	// (Optional) Text
	Text string `json:"text,omitempty"`

	// (Optional) Mode for parsing entities in the text.
	//
	// See formatting options for more details. https://core.telegram.org/bots/api#formatting-options
	ParseMode string `json:"parse_mode,omitempty"`

	// (Optional) A list of special entities that appear in the text, which can be specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty"`
}
