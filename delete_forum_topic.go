package telegram

import "net/http"

// DeleteForumTopicRequest represents parameters to delete a forum topic along with all its messages in a forum supergroup chat using the deleteForumTopic method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageThreadID
//
// See "deleteForumTopic" https://core.telegram.org/bots/api#deleteforumtopic
type DeleteForumTopicRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`
}

// DeleteForumTopic deletes a forum topic along with all its messages in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//   - messageThreadID: Unique identifier for the target message thread of the forum topic
//
// See https://core.telegram.org/bots/api#deleteforumtopic
func (b *Bot) DeleteForumTopic(chatID any, messageThreadID int) error {
	request := DeleteForumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteForumTopic, request); err != nil {
		return err
	}
	return nil
}
