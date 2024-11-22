package telegram

// ChatMemberOwner represents a chat member that owns the chat and has all administrator privileges.
//
// See "ChatMemberOwner" https://core.telegram.org/bots/api#chatmemberowner
type ChatMemberOwner struct {
	// (Required) The member's status in the chat, always “creator”.
	Status string `json:"status"`

	// (Required) Information about the user.
	User User `json:"user"`

	// (Required) True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`

	// (Optional) Custom title for this user.
	CustomTitle string `json:"custom_title,omitempty"`
}

func (m ChatMemberOwner) getMemberStatus() string {
	return m.Status
}
