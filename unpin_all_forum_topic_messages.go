package telegram

import (
	"net/http"
)

// UnpinAllForumTopicMessagesRequest represents a request to clear the list of pinned messages in a forum topic using the unpinAllForumTopicMessages method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageThreadID
//
// See https://core.telegram.org/bots/api#unpinallforumtopicmessages
type UnpinAllForumTopicMessagesRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or
	// username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier for the target message thread of the forum topic.
	MessageThreadID int `json:"message_thread_id"`
}

// UnpinAllForumTopicMessages clears the list of pinned messages in a forum topic through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have
// the can_pin_messages administrator right in the supergroup. Returns true on success.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - messageThreadID: Unique identifier for the target message thread of the forum topic.
//
// See https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (b *Bot) UnpinAllForumTopicMessages(chatID any, messageThreadID int) error {
	request := UnpinAllForumTopicMessagesRequest{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnpinAllForumTopicMessages, request); err != nil {
		return err
	}
	return nil
}
