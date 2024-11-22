package telegram

// BotCommandScopeAllGroupChats represents the scope of bot commands, covering all group and supergroup chats.
//
// See "BotCommandScopeAllGroupChats" https://core.telegram.org/bots/api#botcommandscopeallgroupchats
type BotCommandScopeAllGroupChats struct {
	// (Required) Scope type, must be all_group_chats.
	Type string `json:"type"`
}

func (b BotCommandScopeAllGroupChats) getScopeType() string {
	return b.Type
}
