package telegram

// ForumTopicClosed represents a service message about a forum topic closed in the chat.
type ForumTopicClosed struct{}

// ForumTopicReopened represents a service message about a forum topic reopened in the chat.
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about General forum topic hidden in the chat.
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about General forum topic unhidden in the chat.
type GeneralForumTopicUnhidden struct{}

// VideoChatStarted represents a service message about a video chat started in the chat.
type VideoChatStarted struct{}

// CallbackGame represents a placeholder, currently holding no information. Use BotFather to set up your game.
type CallbackGame struct{}
