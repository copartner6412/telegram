package telegram

// WebhookInfo describes the current status of a webhook.
//
// See "WebhookInfo" https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	// (Required) Webhook URL, may be empty if webhook is not set up.
	URL string `json:"url"`

	// (Required) True, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// (Required) Number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`

	// (Optional) Currently used webhook IP address.
	IPAddress string `json:"ip_address"`

	// (Optional) Unix time for the most recent error that happened when trying to deliver an update via webhook.
	LastErrorDate int `json:"last_error_date"`

	// (Optional) Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook.
	LastErrorMessage string `json:"last_error_message"`

	// (Optional) Unix time of the most recent error that happened when trying to synchronize available updates with Telegram datacenters.
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date"`

	// (Optional) The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery.
	MaxConnections int `json:"max_connections"`

	// (Optional) A list of update types the bot is subscribed to. Defaults to all update types except chat_member.
	AllowedUpdates []string `json:"allowed_updates"`
}
