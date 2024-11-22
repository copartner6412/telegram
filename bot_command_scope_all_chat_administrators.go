package telegram

// BotCommandScopeAllChatAdministrators represents the scope of bot commands, covering all group and supergroup chat administrators.
//
// See "BotCommandScopeAllChatAdministrators" https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
type BotCommandScopeAllChatAdministrators struct {
	// (Required) Scope type, must be all_chat_administrators.
	Type string `json:"type"`
}

func (b BotCommandScopeAllChatAdministrators) getScopeType() string {
	return b.Type
}
