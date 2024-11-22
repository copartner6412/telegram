package telegram

import (
	"net/http"
)

// LogOut logs the bot out from the cloud Bot API server before launching it locally.
//
// After a successful call, you can immediately log in on a local server, but will not be able to log in back
// to the cloud Bot API server for 10 minutes.
//
// See "logOut" https://core.telegram.org/bots/api#logout
func (b *Bot) Logout() error {
	if err := b.sendJsonForBool(http.MethodPost, MethodLogOut, nil); err != nil {
		return err
	}
	return nil
}
