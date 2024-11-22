package telegram

// ReactionType represents the type of a reaction.
//
// Currently, it can be one of:
//   - ReactionTypeEmoji
//   - ReactionTypeCustomEmoji
//   - ReactionTypePaid
//
// See "ReactionType" https://core.telegram.org/bots/api#reactiontype
type ReactionType interface {
	getReactionType() string
}
