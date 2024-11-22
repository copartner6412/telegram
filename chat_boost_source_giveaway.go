package telegram

// ChatBoostSourceGiveaway represents the boost obtained by the creation of a Telegram Premium or a Telegram Star giveaway.
//
// The boost was obtained by the creation of a Telegram Premium or a Telegram Star giveaway. This boosts the chat 4 times for the duration of the corresponding Telegram Premium subscription for Telegram Premium giveaways and prize_star_count / 500 times for one year for Telegram Star giveaways.
//
// See "ChatBoostSourceGiveaway" https://core.telegram.org/bots/api#chatboostsourcegiveaway
type ChatBoostSourceGiveaway struct {
	// (Required) Source of the boost, always “giveaway”.
	Source string `json:"source"`

	// (Required) Identifier of a message in the chat with the giveaway; the message could have been deleted already. May be 0 if the message isn't sent yet.
	GiveawayMessageID int `json:"giveaway_message_id"`

	// (Optional) User that won the prize in the giveaway if any; for Telegram Premium giveaways only.
	User *User `json:"user,omitempty"`

	// (Optional) The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only.
	PrizeStarCount int `json:"prize_star_count,omitempty"`

	// (Optional) True, if the giveaway was completed, but there was no user to win the prize.
	IsUnclaimed bool `json:"is_unclaimed,omitempty"`
}
