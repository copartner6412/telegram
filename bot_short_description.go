package telegram

// BotShortDescription represents the bot's short description.
//
// See "BotShortDescription" https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	// (Required) The bot's short description.
	ShortDescription string `json:"short_description"`
}
