package telegram

// ChatMemberMember represents a chat member that has no additional privileges or restrictions.
//
// See "ChatMemberMember" https://core.telegram.org/bots/api#chatmembermember
type ChatMemberMember struct {
	// (Required) The member's status in the chat, always “member”.
	Status string `json:"status"`

	// (Required) Information about the user.
	User User `json:"user"`

	// (Optional) Date when the user's subscription will expire; Unix time.
	UntilDate int `json:"until_date,omitempty"`
}

func (m ChatMemberMember) getMemberStatus() string {
	return m.Status
}
