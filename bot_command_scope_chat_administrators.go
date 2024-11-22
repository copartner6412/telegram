package telegram

// BotCommandScopeChatAdministrators represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
//
// See "BotCommandScopeChatAdministrators" https://core.telegram.org/bots/api#botcommandscopechatadministrators
type BotCommandScopeChatAdministrators struct {
	// (Required) Scope type, must be chat_administrators.
	Type string `json:"type"`

	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`
}

func (b BotCommandScopeChatAdministrators) getScopeType() string {
	return b.Type
}
