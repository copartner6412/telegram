package telegram

import "encoding/json"

// ChatBoostRemoved represents a boost removed from a chat.
//
// See "ChatBoostRemoved" https://core.telegram.org/bots/api#chatboostremoved
type ChatBoostRemoved struct {
	// (Required) Chat which was boosted.
	Chat Chat `json:"chat"`

	// (Required) Unique identifier of the boost.
	BoostID string `json:"boost_id"`

	// (Required) Point in time (Unix timestamp) when the boost was removed.
	RemoveDate int `json:"remove_date"`

	// (Required) Source of the removed boost.
	//
	// Its type can be one of:
	//   - ChatBoostSourcePremium
	//   - ChatBoostSourceGiftCode
	//   - ChatBoostSourceGiveaway
	Source json.RawMessage `json:"source"`
}
