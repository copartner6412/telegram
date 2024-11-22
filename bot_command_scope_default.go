package telegram

// BotCommandScopeDefault represents the default scope of bot commands.
// Default commands are used if no commands with a narrower scope are specified for the user.
//
// See "BotCommandScopeDefault" https://core.telegram.org/bots/api#botcommandscopedefault
type BotCommandScopeDefault struct {
	// (Required) Scope type, must be default.
	Type string `json:"type"`
}

func (b BotCommandScopeDefault) getScopeType() string {
	return b.Type
}
