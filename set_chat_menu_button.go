package telegram

import (
	"net/http"
)

// SetChatMenuButtonRequest represents the request to change the bot's menu button in a private chat or the default menu button using setChatMenuButton method through the Telegram bot API.
//
// See "setChatMenuButton" https://core.telegram.org/bots/api#setchatmenubutton
type SetChatMenuButtonRequest struct {
	// (Optional) Unique identifier for the target private chat. If not specified, the default bot's menu button will be changed.
	ChatID int `json:"chat_id,omitempty"`

	// (Optional) An object for the bot's new menu button. Defaults to MenuButtonDefault.
	MenuButton *MenuButton `json:"menu_button,omitempty"`
}

// SetChatMenuButton changes the bot's menu button in a private chat or the default menu button through Telegram bot API.
//
// Parameters:
//   - chatID: Unique identifier for the target private chat. If not specified, the default bot's menu button will be changed.
//   - menuButton: An object for the bot's new menu button. If nil, defaults to MenuButtonDefault.
//
// Returns an error if the request was unsuccessful or if there was an issue processing the response.
//
// See "setChatMenuButton" https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(chatID int, menuButton *MenuButton) error {
	request := SetChatMenuButtonRequest{
		ChatID:     chatID,
		MenuButton: menuButton,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetChatMenuButton, request); err != nil {
		return err
	}
	return nil
}
