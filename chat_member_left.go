package telegram

// ChatMemberLeft represents a chat member that isn't currently a member of the chat, but may join it themselves.
//
// See "ChatMemberLeft" https://core.telegram.org/bots/api#chatmemberleft
type ChatMemberLeft struct {
	// (Required) The member's status in the chat, always “left”.
	Status string `json:"status"`

	// (Required) Information about the user.
	User User `json:"user"`
}

func (m ChatMemberLeft) getMemberStatus() string {
	return m.Status
}
