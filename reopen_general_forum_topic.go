package telegram

import (
	"net/http"
)

// ReopenGeneralForumTopicRequest represents parameters to reopen a closed 'General' topic in a forum supergroup chat using the closeGeneralForumTopic method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "reopenGeneralForumTopic" https://core.telegram.org/bots/api#closegeneralforumtopic
type ReopenGeneralForumTopicRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID any `json:"chat_id"`
}

// ReopenGeneralForumTopic reopens a closed 'General' topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//
// See "reopenGeneralForumTopic" https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) ReopenGeneralForumTopic(chatID any) error {
	request := ReopenGeneralForumTopicRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodReopenGeneralForumTopic, request); err != nil {
		return err
	}
	return nil
}
