package telegram

// ChatBoostSourceGiftCode represents the boost obtained by the creation of Telegram Premium gift codes to boost a chat.
//
// See "ChatBoostSourceGiftCode" https://core.telegram.org/bots/api#chatboostsourcegiftcode
type ChatBoostSourceGiftCode struct {
	// (Required) Source of the boost, always “gift_code”.
	Source string `json:"source"`

	// (Required) User for which the gift code was created.
	User User `json:"user"`
}
