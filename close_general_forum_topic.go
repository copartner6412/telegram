package telegram

import (
	"net/http"
)

// CloseGeneralForumRequest represents parameters to close an open 'General' topic in a forum supergroup chat using the closeGeneralForumTopic method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "closeGeneralForumTopic" https://core.telegram.org/bots/api#closegeneralforumtopic
type CloseGeneralForumRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID any `json:"chat_id"`
}

// CloseGeneralForumTopic closes an open 'General' topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//
// See "closeGeneralForumTopic" https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) CloseGeneralForumTopic(chatID any) error {
	request := CloseGeneralForumRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodCloseGeneralForumTopic, request); err != nil {
		return err
	}
	return nil
}
