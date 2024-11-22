package telegram

import (
	"net/http"
)

// UnpinAllChatMessagesRequest represents parameters to clear the list of pinned messages in a chat using the unpinAllChatMessages method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "unpinAllChatMessages" https://core.telegram.org/bots/api#unpinallchatmessages
type UnpinAllChatMessagesRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`
}

// UnpinAllChatMessages clears the list of pinned messages in a chat through the Telegram bot API.
//
// If the chat is not a private chat, the bot must be an administrator in the chat and have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
// See https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(chatID any) error {
	request := UnpinAllChatMessagesRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnpinAllForumTopicMessages, request); err != nil {
		return err
	}
	return nil
}
