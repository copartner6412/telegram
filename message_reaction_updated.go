package telegram

import "encoding/json"

// MessageReactionUpdated represents a change of a reaction on a message performed by a user.
//
// See "MessageReactionUpdated" https://core.telegram.org/bots/api#messagereactionupdated
type MessageReactionUpdated struct {
	// (Required) The chat containing the message the user reacted to.
	Chat Chat `json:"chat"`

	// (Required) Unique identifier of the message inside the chat.
	MessageID int `json:"message_id"`

	// (Optional) The user that changed the reaction, if the user isn't anonymous.
	User *User `json:"user,omitempty"`

	// (Optional) The chat on behalf of which the reaction was changed, if the user is anonymous.
	ActorChat *Chat `json:"actor_chat,omitempty"`

	// (Required) Date of the change in Unix time.
	Date int `json:"date"`

	// (Required) Previous list of reaction types that were set by the user.
	//
	// Its type can be one of:
	//   - ReactionTypeEmoji
	//   - ReactionTypeCustomEmoji
	//   - ReactionTypePaid
	OldReaction []json.RawMessage `json:"old_reaction"`

	// (Required) New list of reaction types that have been set by the user.
	//
	// Its type can be one of:
	//   - ReactionTypeEmoji
	//   - ReactionTypeCustomEmoji
	//   - ReactionTypePaid
	NewReaction []json.RawMessage `json:"new_reaction"`
}
