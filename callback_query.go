package telegram

import "encoding/json"

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
//
// If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
//
// # Note
//
// After the user presses a callback button, Telegram clients will display a progress bar until you call answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if no notification to the user is needed (e.g., without specifying any of the optional parameters).
//
// See "CallbackQuery" https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	// (Required) Unique identifier for this query.
	ID string `json:"id"`

	// (Required) Sender of the callback query.
	From User `json:"from"`

	// (Optional) Message sent by the bot with the callback button that originated the query.
	//
	// Its type can be one of:
	//   - Massage
	//   - InaccessibleMessage
	Message json.RawMessage `json:"message,omitempty"`

	// (Optional) Identifier of the message sent via the bot in inline mode, that originated the query.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// (Required) Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent.
	ChatInstance string `json:"chat_instance"`

	// (Optional) Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	Data string `json:"data,omitempty"`

	// (Optional) Short name of a Game to be returned, serves as the unique identifier for the game.
	GameShortName string `json:"game_short_name,omitempty"`
}
