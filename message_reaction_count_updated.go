package telegram

// MessageReactionCountUpdated represents reaction changes on a message with anonymous reactions.
//
// See "MessageReactionCountUpdated" https://core.telegram.org/bots/api#messagereactioncountupdated
type MessageReactionCountUpdated struct {
	// (Required) The chat containing the message.
	Chat Chat `json:"chat"`

	// (Required) Unique message identifier inside the chat.
	MessageID int `json:"message_id"`

	// (Required) Date of the change in Unix time.
	Date int `json:"date"`

	// (Required) List of reactions that are present on the message.
	Reactions []ReactionCount `json:"reactions"`
}
