package telegram

// ChatMemberBanned represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
//
// See "ChatMemberBanned" https://core.telegram.org/bots/api#chatmemberbanned
type ChatMemberBanned struct {
	// (Required) The member's status in the chat, always “kicked”.
	Status string `json:"status"`

	// (Required) Information about the user.
	User User `json:"user"`

	// (Required) Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever.
	UntilDate int `json:"until_date"`
}

func (m ChatMemberBanned) getMemberStatus() string {
	return m.Status
}
