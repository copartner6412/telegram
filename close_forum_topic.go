package telegram

import (
	"net/http"
)

// CloseForumTopicRequest represents parameters to close an open topic in a forum supergroup chat using the closeForumTopic method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageThreadID
//
// See "closeForumTopic" https://core.telegram.org/bots/api#closeforumtopic
type CloseForumTopicRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier for the target message thread of the forum topic
	MessageThreadID int `json:"message_thread_id"`
}

// CloseForumTopic closes an open topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have
// the can_manage_topics administrator rights, unless it is the creator of the topic.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//   - messageThreadID: Unique identifier for the target message thread of the forum topic
//
// See https://core.telegram.org/bots/api#closeforumtopic
func (b *Bot) CloseForumTopic(chatID any, messageThreadID int) error {
	request := CloseForumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: messageThreadID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodCloseForumTopic, request); err != nil {
		return err
	}
	return nil
}
