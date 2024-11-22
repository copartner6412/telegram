package telegram

import (
	"net/http"
)

// SendMessageRequest represents the payload for sending a message using sendMessage method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Text
//
// See "sendMessage" https://core.telegram.org/bots/api#sendmessage
type SendMessageRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Text of the message to be sent, 1-4096 characters after entities parsing.
	Text

	*SendMessageOptions
}

// SendMessageOptions represents optional parameters for sending a message using sendMessage method through the Telegram bot API.
//
// See "sendMessage" https://core.telegram.org/bots/api#sendmessage
type SendMessageOptions struct {

	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Link preview generation options for the message.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	*SendOptions
}

// SendMessage sends a text message using the Telegram Bot API through the Telegram bot API.
//
// Required parameters:
//   - toChatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - text: Text of the message to be sent, 1-4096 characters after entities parsing.
//
// On success, the sent Message is returned.
//
// See "sendMessage" https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(toChatID any, text Text, options *SendMessageOptions) (*Message, error) {
	request := SendMessageRequest{
		ChatID:             toChatID,
		Text:               text,
		SendMessageOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendMessage, request); err != nil {
		return nil, err
	}

	return result, nil
}
