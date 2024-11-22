package telegram

import "net/http"

// EditGeneralForumTopicRequest edits the name of the 'General' topic in a forum supergroup chat using the editGeneralForumTopic method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Name
//
// See "editGeneralForumTopic" https://core.telegram.org/bots/api#editgeneralforumtopic
type EditGeneralForumTopicRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) New topic name, 1-128 characters.
	Name string `json:"name"`
}

// EditGeneralForumTopic edits the name of the 'General' topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - name: New topic name, 1-128 characters.
//
// See "editGeneralForumTopic" https://core.telegram.org/bots/api#editgeneralforumtopic
func (b *Bot) EditGeneralForumTopic(chatID any, name string) error {
	request := EditGeneralForumTopicRequest{
		ChatID: chatID,
		Name:   name,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditGeneralForumTopic, request); err != nil {
		return err
	}
	return nil
}
