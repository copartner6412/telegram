package telegram

// ChosenInlineResult represents a result of an inline query that was chosen by the user and sent to their chat partner.
//
// Required fields:
//   - ResultID
//   - From
//   - Query
//
// # Note
//
// It is necessary to enable inline feedback via @BotFather in order to receive these objects in updates.
//
// See "ChosenInlineResult" https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	// (Required) The unique identifier for the result that was chosen.
	ResultID string `json:"result_id"`

	// (Required) The user that chose the result.
	From User `json:"from"`

	// (Required) The query that was used to obtain the result.
	Query string `json:"query"`

	// (Optional) Sender location, only for bots that require user location.
	Location *Location `json:"location,omitempty"`

	// (Optional) Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
}
