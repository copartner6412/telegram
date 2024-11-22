package telegram

import (
	"net/http"
)

// GetChatMenuButtonRequest represents parameters to get the current value of the bot's menu button in a private chat, or the default menu button using the getChatMenuButton method through the Telegram bot API.
//
// See "getChatMenuButton" https://core.telegram.org/bots/api#getchatmenubutton
type GetChatMenuButtonRequest struct {
	// (Optional) Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
	ChatID int `json:"chat_id,omitempty"`
}

// GetChatMenuButton retrieves the current value of the bot's menu button in a private chat, or the default menu button through the Telegram bot API.
//
// Parameters:
//   - chatID: Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
//
// Returns the MenuButton on success.
//
// See "getChatMenuButton" https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(chatID int) (*MenuButton, error) {
	request := GetChatMenuButtonRequest{
		ChatID: chatID,
	}
	result := new(MenuButton)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetChatMenuButton, request); err != nil {
		return nil, err
	}

	return result, nil
}
