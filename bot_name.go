package telegram

// BotName represents the bot's name.
//
// See "BotName" https://core.telegram.org/bots/api#botname
type BotName struct {
	// (Required) The bot's name.
	Name string `json:"name"`
}
