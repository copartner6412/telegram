package telegram

import (
	"net/http"
)

// GetMe fetches basic information about the bot in form of a User object.
//
// It is used as a simple method for testing the bot's authentication token.
//
// See "getMe" https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (*User, error) {
	result := new(User)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetMe, nil); err != nil {
		return nil, err
	}
	return result, nil
}
