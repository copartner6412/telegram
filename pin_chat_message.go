package telegram

import (
	"net/http"
)

// PinChatMessage represents the parameters required to pin a message in a chat using the pinChatMessage method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageID
//
// See https://core.telegram.org/bots/api#pinchatmessage
type PinChatMessageRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Identifier of a message to pin.
	MessageID int `json:"message_id"`

	*PinChatMessageOptions
}

type PinChatMessageOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be pinned.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Pass True if it is not necessary to send a notification to all chat members about the new pinned message.
	// Notifications are always disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// PinChatMessage adds a message to the list of pinned messages in a chat through the Telegram bot API.
//
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of a message to pin.
//
// See https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(chatID any, messageID int, options *PinChatMessageOptions) error {
	request := PinChatMessageRequest{
		ChatID:                chatID,
		MessageID:             messageID,
		PinChatMessageOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodPinChatMessage, request); err != nil {
		return err
	}
	return nil
}
