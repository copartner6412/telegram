package telegram

import (
	"net/http"
)

// UnpinAllGeneralForumTopicMessagesRequest represents a request to clear the list of pinned messages in a General forum topic using the unpinAllGeneralForumTopicMessages method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "unpinAllGeneralForumTopicMessages" https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
type UnpinAllGeneralForumTopicMessagesRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`
}

// UnpinAllGeneralForumTopicMessages clears the list of pinned messages in a General forum topic using the "unpinAllGeneralForumTopicMessages" through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//
// See "unpinAllGeneralForumTopicMessages" https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (b *Bot) UnpinAllGeneralForumTopicMessages(chatID any) error {
	request := UnpinAllGeneralForumTopicMessagesRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnpinAllGeneralForumTopicMessages, request); err != nil {
		return err
	}
	return nil
}
