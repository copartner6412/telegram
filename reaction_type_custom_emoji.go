package telegram

// ReactionTypeCustomEmoji represents a reaction based on a custom emoji.
//
// See "ReactionTypeCustomEmoji" https://core.telegram.org/bots/api#reactiontypecustomemoji
type ReactionTypeCustomEmoji struct {
	// (Required) Type of the reaction, always “custom_emoji”.
	Type string `json:"type"`

	// (Required) Custom emoji identifier.
	CustomEmojiID string `json:"custom_emoji_id"`
}

func (r ReactionTypeCustomEmoji) getReactionType() string {
	return r.Type
}
