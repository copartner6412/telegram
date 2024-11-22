package telegram

import (
	"net/http"
)

// UnhideGeneralForumTopicRequest represents a request to unhide the 'General' topic in a forum supergroup chat through the hideGeneralForumTopic through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "unhideGeneralForumTopic" https://core.telegram.org/bots/api#unhidegeneralforumtopic
type UnhideGeneralForumTopicRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`
}

// UnhideGeneralForumTopic unhides the 'General' topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//
// See "unhideGeneralForumTopic" https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (b *Bot) UnhideGeneralForumTopic(chatID any) error {
	request := UnhideGeneralForumTopicRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnhideGeneralForumTopic, request); err != nil {
		return err
	}
	return nil
}
