package telegram

// MenuButtonCommands represents a menu button, which opens the bot's list of commands.
//
// See "MenuButtonCommands" https://core.telegram.org/bots/api#menubuttoncommands
type MenuButtonCommands struct {
	// (Required) Type of the button, must be commands.
	Type string `json:"type"`
}

func (b MenuButtonCommands) getButtonType() string {
	return b.Type
}
