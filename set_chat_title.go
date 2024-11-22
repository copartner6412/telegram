package telegram

import (
	"net/http"
)

// SetChatTitleRequest represents a request to change the title of a chat using the setChatTitle method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Title
//
// See https://core.telegram.org/bots/api#setchattitle
type SetChatTitleRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) New chat title, 1-128 characters
	Title string `json:"title"`
}

// SetChatTitle changes the title of a chat.
//
// Titles cannot be changed for private chats. The bot must be an administrator in the chat
// for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - title: New chat title, 1-128 characters
//
// Returns an error if the request is unsuccessful.
//
// See https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(chatID any, title string) error {
	request := SetChatTitleRequest{
		ChatID: chatID,
		Title:  title,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetChatTitle, request); err != nil {
		return err
	}
	return nil
}
