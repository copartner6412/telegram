package telegram

import (
	"net/http"
)

// GetWebhookInfo retrieves the current status of the webhook through the Telegram bot API.
//
// On success, returns a WebhookInfo object.
//
// If the bot is using getUpdates, will return an object with the url field empty.
//
// See "getWebhookInfo" https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo() (*WebhookInfo, error) {
	result := new(WebhookInfo)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetWebhookInfo, nil); err != nil {
		return nil, err
	}
	return result, nil
}
