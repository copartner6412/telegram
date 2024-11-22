package telegram

import (
	"net/http"
)

// SetChatAdministratorCustomTitleRequest represents a request to set a custom title for an administrator in a supergroup using the setChatAdministratorCustomTitle method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - UserID
//   - CustomTitle
//
// See https://core.telegram.org/bots/api#setchatadministratorcustomtitle
type SetChatAdministratorCustomTitleRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`

	// (Required) New custom title for the administrator; 0-16 characters, emoji are not allowed.
	CustomTitle string `json:"custom_title"`
}

// SetChatAdministratorCustomTitle sets a custom title for an administrator in a supergroup promoted by the bot through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - userID: Unique identifier of the target user.
//   - customTitle: New custom title for the administrator; 0-16 characters, emoji are not allowed.
//
// See https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(chatID any, userID int, customTitle string) error {
	request := SetChatAdministratorCustomTitleRequest{
		ChatID:      chatID,
		UserID:      userID,
		CustomTitle: customTitle,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetChatAdministratorCustomTitle, request); err != nil {
		return err
	}

	return nil
}
