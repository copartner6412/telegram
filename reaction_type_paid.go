package telegram

// ReactionTypePaid represents a reaction that is paid.
//
// See "ReactionTypePaid" https://core.telegram.org/bots/api#reactiontypepaid
type ReactionTypePaid struct {
	// (Required) Type of the reaction, always “paid”.
	Type string `json:"type"`
}

func (r ReactionTypePaid) getReactionType() string {
	return r.Type
}
