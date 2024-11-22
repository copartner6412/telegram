package telegram

// BotCommand represents a bot command.
//
// See "BotCommand" https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	// (Required) Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`

	// (Required) Description of the command; 1-256 characters.
	Description string `json:"description"`
}
