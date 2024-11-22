package telegram

import (
	"net/http"
)

// SetChatDescriptionRequest represents a request to change the description of a group, supergroup, or channel using setChatDescription method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See https://core.telegram.org/bots/api#setchatdescription
type SetChatDescriptionRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Optional) New chat description, 0-255 characters.
	Description string `json:"description,omitempty"`
}

// SetChatDescription changes the description of a group, supergroup, or channel through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - description: New chat description, 0-255 characters.
//
// See https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(chatID any, description string) error {
	request := SetChatDescriptionRequest{
		ChatID:      chatID,
		Description: description,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetChatDescription, request); err != nil {
		return err
	}
	return nil
}
