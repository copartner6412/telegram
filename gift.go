package telegram

// Gift represents a gift that can be sent by the bot.
//
// See "Gift" https://core.telegram.org/bots/api#gift
type Gift struct {
	// (Required) Unique identifier of the gift
	ID string `json:"id"`

	// (Required) The sticker that represents the gift
	Sticker Sticker `json:"sticker"`

	// (Required) The number of Telegram Stars that must be paid to send the sticker
	StarCount int `json:"star_count"`

	// (Optional) The total number of the gifts of this type that can be sent; for limited gifts only
	TotalCount int `json:"total_count,omitempty"`

	// (Optional) The number of remaining gifts of this type that can be sent; for limited gifts only
	RemainingCount int `json:"remaining_count,omitempty"`
}
