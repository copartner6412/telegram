package telegram

import (
	"net/http"
)

// HideGeneralForumTopicRequest represents a request to hide the 'General' topic in a forum supergroup chat through the hideGeneralForumTopic through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "hideGeneralForumTopic" https://core.telegram.org/bots/api#hidegeneralforumtopic
type HideGeneralForumTopicRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`
}

// HideGeneralForumTopic hides the 'General' topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have
// the can_manage_topics administrator rights. The topic will be automatically
// closed if it was open.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//
// See "hideGeneralForumTopic" https://core.telegram.org/bots/api#hidegeneralforumtopic
func (b *Bot) HideGeneralForumTopic(chatID any) error {
	request := HideGeneralForumTopicRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodHideGeneralForumTopic, request); err != nil {
		return err
	}
	return nil
}
