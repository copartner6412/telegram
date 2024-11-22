package telegram

import (
	"net/http"
)

// UnpinChatMessageRequest represents the request payload to remove a message from the list of pinned messages in a chat using the unpinChatMessage method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See https://core.telegram.org/bots/api#unpinchatmessage
type UnpinChatMessageRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	*UnpinChatMessageOptions
}

type UnpinChatMessageOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be unpinned.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Identifier of the message to unpin. Required if
	// business_connection_id is specified. If not specified, the
	// most recent pinned message (by sending date) will be unpinned.
	MessageID int `json:"message_id,omitempty"`
}

// UnpinChatMessage removes a message from the list of pinned messages in a chat through the Telegram bot API.
//
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//
// See https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(chatID any, options *UnpinChatMessageOptions) error {
	request := UnpinChatMessageRequest{
		ChatID:                  chatID,
		UnpinChatMessageOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnpinChatMessage, request); err != nil {
		return err
	}
	return nil
}
