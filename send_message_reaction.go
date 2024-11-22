package telegram

import (
	"net/http"
)

// SetMessageReactionRequest represents the parameters for the setMessageReaction method.
// This method is used to change the chosen reactions on a message.
// Service messages can't be reacted to. Automatically forwarded messages from a channel to its discussion group
// have the same available reactions as messages in the channel. Bots can't use paid reactions.
//
// See https://core.telegram.org/bots/api#setmessagereaction
type SetMessageReactionRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Identifier of the target message. If the message belongs to a media group,
	// the reaction is set to the first non-deleted message in the group instead.
	MessageID int `json:"message_id"`

	// (Optional) A list of reaction types to set on the message.
	// Currently, as non-premium users, bots can set up to one reaction per message.
	// A custom emoji reaction can be used if it is either already present on the message
	// or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	Reaction []ReactionType `json:"reaction,omitempty"`

	// (Optional) Pass True to set the reaction with a big animation.
	IsBig bool `json:"is_big,omitempty"`
}

// SetMessageReaction changes the chosen reactions on a message.
//
// Service messages can't be reacted to. Automatically forwarded messages
// from a channel to its discussion group have the same available reactions
// as messages in the channel. Bots can't use paid reactions.
//
// Parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - toMessageID: Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
//   - reactions: A list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
//   - isBig: Pass True to set the reaction with a big animation.
//
// Returns an error if the request fails.
//
// See https://core.telegram.org/bots/api#setmessagereaction
func (b *Bot) SetMessageReaction(toChatID any, toMessageID int, reactions []ReactionType, isBig bool) error {
	request := SetMessageReactionRequest{
		ChatID:    toChatID,
		MessageID: toMessageID,
		Reaction:  reactions,
		IsBig:     isBig,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetMessageReaction, request); err != nil {
		return err
	}
	return nil
}
