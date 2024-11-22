package telegram

// InlineQueryResultCachedSticker represents a link to a sticker stored on the Telegram servers.
//
// By default, this sticker will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the sticker.
//
// Required fields:
//   - Type
//   - ID
//   - StickerFileID
//
// See "InlineQueryResultCachedSticker" https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
type InlineQueryResultCachedSticker struct {
	// (Required) Type of the result, must be "sticker".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) A valid file identifier of the sticker.
	StickerFileID string `json:"sticker_file_id"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the sticker.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultCachedSticker) getQueryType() string {
	return r.Type
}
