package telegram

// SentWebAppMessage describes an inline message sent by a Web App on behalf of a user using the SentWebAppMessage method through the Telegram bot API.
//
// Required fields:
//   - InlineMessageID
//
// See "SentWebAppMessage" https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	// (Optional) Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
}
