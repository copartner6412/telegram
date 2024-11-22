package telegram

import "encoding/json"

// ReactionCount represents a reaction added to a message along with the number of times it was added.
//
// See "ReactionCount" https://core.telegram.org/bots/api#reactioncount
type ReactionCount struct {
	// (Required) Type of the reaction.
	//
	// Its type can be one of:
	//   - ReactionTypeEmoji
	//   - ReactionTypeCustomEmoji
	//   - ReactionTypePaid
	Type json.RawMessage `json:"type"`

	// (Required) Number of times the reaction was added.
	TotalCount int `json:"total_count"`
}
