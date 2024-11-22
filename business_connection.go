package telegram

// BusinessConnection describes the connection of the bot with a business account.
//
// See "BusinessConnection" https://core.telegram.org/bots/api#businessconnection
type BusinessConnection struct {
	// (Required) Unique identifier of the business connection.
	ID string `json:"id"`

	// (Required) Business account user that created the business connection.
	User User `json:"user"`

	// (Required) Identifier of a private chat with the user who created the business connection.
	// This number may have more than 32 significant bits and some programming languages may have
	// difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a
	// 64-bit integer or double-precision float type are safe for storing this identifier.
	UserChatID int `json:"user_chat_id"`

	// (Required) Date the connection was established in Unix time.
	Date int `json:"date"`

	// (Required) True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours.
	CanReply bool `json:"can_reply"`

	// (Required) True, if the connection is active.
	IsEnabled bool `json:"is_enabled"`
}
