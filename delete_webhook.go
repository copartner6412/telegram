package telegram

import (
	"net/http"
)

// DeleteWebhookRequest represents parameters used to remove webhook integration if you decide to switch back to getUpdates.
//
// See "deleteWebhook" https://core.telegram.org/bots/api#deletewebhook
type DeleteWebhookRequest struct {
	// (Optional) Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// DeleteWebhook removes webhook integration if the bot decides to switch back to getUpdates.
//
// Parameters:
//   - dropPendingUpdates: Pass True to drop all pending updates
//
// See "deleteWebhook" https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(dropPendingUpdates bool) error {
	request := DeleteWebhookRequest{
		DropPendingUpdates: dropPendingUpdates,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteWebhook, request); err != nil {
		return err
	}
	return nil
}
