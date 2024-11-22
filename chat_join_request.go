package telegram

// ChatJoinRequest represents a join request sent to a chat.
//
// See "ChatJoinRequest" https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	// (Required) Chat to which the request was sent.
	Chat Chat `json:"chat"`

	// (Required) User that sent the join request.
	From User `json:"from"`

	// (Required) Identifier of a private chat with the user who sent the join request.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	// The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	UserChatID int `json:"user_chat_id"`

	// (Required) Date the request was sent in Unix time.
	Date int `json:"date"`

	// (Optional) Bio of the user.
	Bio string `json:"bio,omitempty"`

	// (Optional) Chat invite link that was used by the user to send the join request.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
