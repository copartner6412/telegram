package telegram

// BusinessMessagesDeleted represents an object received when messages are deleted from a connected business account.
//
// See "BusinessMessagesDeleted" https://core.telegram.org/bots/api#businessmessagesdeleted
type BusinessMessagesDeleted struct {
	// (Required) Unique identifier of the business connection.
	BusinessConnectionID string `json:"business_connection_id"`

	// (Required) Information about a chat in the business account.
	// The bot may not have access to the chat or the corresponding user.
	Chat Chat `json:"chat"`

	// (Required) The list of identifiers of deleted messages in the chat of the business account.
	MessageIDs []int `json:"message_ids"`
}
