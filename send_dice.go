package telegram

import (
	"net/http"
)

// SendDiceRequest represents the parameters for sending an animated emoji that will display a random value using sendDice method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See https://core.telegram.org/bots/api#senddice
type SendDiceRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	*SendDiceOptions
}

// SendDiceOptions represents the optional parameters for sending an animated emoji that will display a random value using the sendDice method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#senddice
type SendDiceOptions struct {

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “⚽️”, “🏀”, “🏈”, or “🎳”.
	// Dice can have values 1-6 for “🎲”, “🎯” and “⚽️”, values 1-5 for “🏀” and “🏈”, and values 1-64 for “🎳”. Defaults to “🎲”.
	Emoji string `json:"emoji,omitempty"`

	*SendOptions
}

// SendDice sends an animated emoji that will display a random value through the Telegram bot API.
//
// Required parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(toChatID any, options *SendDiceOptions) (*Message, error) {
	request := SendDiceRequest{
		ChatID:          toChatID,
		SendDiceOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendDice, request); err != nil {
		return nil, err
	}

	return result, nil
}
