package telegram

// BotDescription represents the bot's description.
//
// See "BotDescription" https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	// (Required) The bot's description.
	Description string `json:"description"`
}
