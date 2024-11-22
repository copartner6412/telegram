package telegram

// ChatBoostUpdated represents a boost added to a chat or changed.
//
// See "ChatBoostUpdated" https://core.telegram.org/bots/api#chatboostupdated
type ChatBoostUpdated struct {
	// (Required) Chat which was boosted.
	Chat Chat `json:"chat"`

	// (Required) Information about the chat boost.
	Boost ChatBoost `json:"boost"`
}
