package telegram

import "net/http"

// ForwardMessageRequest represents the payload to forward messages using forwardMessage method.
//
// Required fields:
//   - ChatID
//   - FromChatID
//   - MessageID
//
// See "forwardMessage" https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Integer or String. Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername).
	FromChatID any `json:"from_chat_id"`

	// (Required) Message identifier in the chat specified in from_chat_id.
	MessageID int `json:"message_id"`

	*ForwardMessageOptions
}

// ForwardMessageOptions represents optional parameters to forward messages using forwardMessage method.
//
// See "forwardMessage" https://core.telegram.org/bots/api#forwardmessage
type ForwardMessageOptions struct {
	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// (Optional) Protects the contents of the forwarded message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
}

// ForwardMessage sends a message to a specified chat by forwarding it from another chat.
//
// Service messages and messages with protected content can't be forwarded.
//
// Required parameters:
//   - fromChatID: Integer or String. Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername).
//   - toChatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - MessageID: Message identifier in the chat specified in from_chat_id.
//
// See "forwardMessage" https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(fromChatID, toChatID any, messageID int, options *ForwardMessageOptions) (*Message, error) {
	request := ForwardMessageRequest{
		ChatID:                toChatID,
		FromChatID:            fromChatID,
		MessageID:             messageID,
		ForwardMessageOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodForwardMessage, request); err != nil {
		return nil, err
	}

	return result, nil
}
