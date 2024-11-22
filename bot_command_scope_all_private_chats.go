package telegram

// BotCommandScopeAllPrivateChats represents the scope of bot commands, covering all private chats.
//
// See "BotCommandScopeAllPrivateChats" https://core.telegram.org/bots/api#botcommandscopeallprivatechats
type BotCommandScopeAllPrivateChats struct {
	// (Required) Scope type, must be all_private_chats.
	Type string `json:"type"`
}

func (b BotCommandScopeAllPrivateChats) getScopeType() string {
	return b.Type
}
