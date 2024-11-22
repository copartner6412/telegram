package telegram

// ForumTopic represents a forum topic.
//
// See "ForumTopic" https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	// (Required) Unique identifier of the forum topic.
	MessageThreadID int `json:"message_thread_id"`

	// (Required) Name of the topic.
	Name string `json:"name"`

	// (Required) Color of the topic icon in RGB format.
	IconColor int `json:"icon_color"`

	// (Optional) Unique identifier of the custom emoji shown as the topic icon.
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}
