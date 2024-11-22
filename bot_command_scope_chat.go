package telegram

// BotCommandScopeChat represents the scope of bot commands, covering a specific chat.
//
// See "BotCommandScopeChat" https://core.telegram.org/bots/api#botcommandscopechat
type BotCommandScopeChat struct {
	// (Required) Scope type, must be chat.
	Type string `json:"type"`

	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`
}

func (b BotCommandScopeChat) getScopeType() string {
	return b.Type
}
