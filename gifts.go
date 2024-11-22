package telegram

// Gifts represents a list of gifts.
//
// See "Gifts" https://core.telegram.org/bots/api#gifts
type Gifts struct {
	// (Required) The list of gifts.
	Gifts []Gift
}
