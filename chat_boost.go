package telegram

import "encoding/json"

// ChatBoost contains information about a chat boost.
//
// See "ChatBoost" https://core.telegram.org/bots/api#chatboost
type ChatBoost struct {
	// (Required) Unique identifier of the boost.
	BoostID string `json:"boost_id"`

	// (Required) Point in time (Unix timestamp) when the chat was boosted.
	AddDate int `json:"add_date"`

	// (Required) Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged.
	ExpirationDate int `json:"expiration_date"`

	// (Required) Source of the added boost.
	//
	// Its type can be one of:
	//   - ChatBoostSourcePremium
	//   - ChatBoostSourceGiftCode
	//   - ChatBoostSourceGiveaway
	Source json.RawMessage `json:"source"`
}
