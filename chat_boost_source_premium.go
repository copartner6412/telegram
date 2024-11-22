package telegram

// ChatBoostSourcePremium represents the boost obtained by subscribing to Telegram Premium or by gifting a Telegram Premium subscription to another user.
//
// See "ChatBoostSourcePremium" https://core.telegram.org/bots/api#chatboostsourcepremium
type ChatBoostSourcePremium struct {
	// (Required) Source of the boost, always “premium”.
	Source string `json:"source"`

	// (Required) User that boosted the chat.
	User User `json:"user"`
}
