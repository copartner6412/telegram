package telegram

import (
	"net/http"
)

// Close closes the bot instance before moving it from one local server to another.
//
// You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart.
// The method will return error 429 in the first 10 minutes after the bot is launched.
//
// See "Close" https://core.telegram.org/bots/api#close
func (b *Bot) Close() error {
	if err := b.sendJsonForBool(http.MethodPost, MethodClose, nil); err != nil {
		return err
	}
	return nil
}
