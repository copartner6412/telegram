package telegram

// BotCommandScopeChatMember represents the scope of bot commands, covering a specific member of a group or supergroup chat.
//
// See "BotCommandScopeChatMember" https://core.telegram.org/bots/api#botcommandscopechatmember
type BotCommandScopeChatMember struct {
	// (Required) Scope type, must be chat_member.
	Type string `json:"type"`

	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`
}

func (b BotCommandScopeChatMember) getScopeType() string {
	return b.Type
}
