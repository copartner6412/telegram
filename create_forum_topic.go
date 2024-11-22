package telegram

import (
	"net/http"
)

// CreateForumTopicRequest represents the parameters required to create a topic in a forum supergroup chat using the createForumTopic through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Name
//
// See "createForumTopic" https://core.telegram.org/bots/api#createforumtopic
type CreateForumTopicRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Topic name, 1-128 characters.
	Name string `json:"name"`

	*CreateForumTopicOptions
}

// CreateForumTopicOptions represents the optional parameters required to create a topic in a forum supergroup chat using the createForumTopic through the Telegram bot API.
//
// See "createForumTopic" https://core.telegram.org/bots/api#createforumtopic
type CreateForumTopicOptions struct {
	// (Optional) Color of the topic icon in RGB format. Currently, must be one of
	// 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB),
	// 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F).
	IconColor int `json:"icon_color,omitempty"`

	// (Optional) Unique identifier of the custom emoji shown as the topic icon.
	// Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// CreateForumTopic creates a topic in a forum supergroup chat through the Telegram bot API.
//
// The bot must be an administrator in the chat and have the can_manage_topics administrator rights.
//
// Required parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - name: Topic name, 1-128 characters.
//
// Returns information about the created topic as a ForumTopic object.
//
// See "createForumTopic" https://core.telegram.org/bots/api#createforumtopic
func (b *Bot) CreateForumTopic(chatID any, name string, options *CreateForumTopicOptions) (*ForumTopic, error) {
	request := CreateForumTopicRequest{
		ChatID:                  chatID,
		Name:                    name,
		CreateForumTopicOptions: options,
	}
	result := new(ForumTopic)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodCreateForumTopic, request); err != nil {
		return nil, err
	}
	return result, nil
}
