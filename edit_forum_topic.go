package telegram

import (
	"net/http"
)

// EditForumTopicRequest represents the parameters required to edit the name and icon of a topic in a forum supergroup chat using the editForumTopic through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageThreadID
//
// See "editForumTopic" https://core.telegram.org/bots/api#editforumtopic
type EditForumTopicRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Integer. Unique identifier for the target message thread of the forum topic.
	MessageThreadID int `json:"message_thread_id"`

	*EditForumTopicOptions
}

// EditForumTopicOptions represents the optional parameters required to edit the name and icon of a topic in a forum supergroup chat using the editForumTopic through the Telegram bot API.
//
// See "editForumTopic" https://core.telegram.org/bots/api#editforumtopic
type EditForumTopicOptions struct {
	// (Optional) New topic name, 0-128 characters. If not specified or empty,
	// the current name of the topic will be kept.
	Name string `json:"name,omitempty"`

	// (Optional) New unique identifier of the custom emoji shown as the topic icon.
	// Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	// Pass an empty string to remove the icon. If not specified, the current icon will be kept.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// EditForumTopic edits the name and icon of a topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic.
//
// Required parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - messageThreadID: Integer. Unique identifier for the target message thread of the forum topic.
//
// See https://core.telegram.org/bots/api#editforumtopic
func (b *Bot) EditForumTopic(chatID any, messageThreadID int, options *EditForumTopicOptions) error {
	request := EditForumTopicRequest{
		ChatID:                chatID,
		MessageThreadID:       messageThreadID,
		EditForumTopicOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditForumTopic, request); err != nil {
		return err
	}
	return nil
}
