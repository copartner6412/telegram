package telegram

// InlineQueryResultGame represents a game.
//
// By default, the game will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the game.
//
// Required fields:
//   - Type
//   - ID
//   - GameShortName
//
// See "InlineQueryResultGame" https://core.telegram.org/bots/api#inlinequeryresultgame
type InlineQueryResultGame struct {
	// (Required) Type of the result, must be "game".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Short name of the game.
	GameShortName string `json:"game_short_name"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (r InlineQueryResultGame) getQueryType() string {
	return r.Type
}
