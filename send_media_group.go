package telegram

import "net/http"

// SendMediaGroupRequest represents the parameters for sending a group of photos, videos, documents or audios as an album using the sendMediaGroup method through the Telegram bot API.
//
// Documents and audio files can be only grouped in an album with messages of the same type.
//
// Required fields:
//   - ChatID
//   - Media
//
// See https://core.telegram.org/bots/api#sendmediagroup
type SendMediaGroupRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	// If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance.
	// Otherwise, they will be credited to the bot's balance.
	ChatID any `json:"chat_id"`

	// (Required) An array describing messages to be sent, must include 2-10 items. Can not be InputMediaAnimation.
	Media []InputMedia `json:"media"`

	*SendMediaGroupOptions
}

// SendMediaGroupOptions represents the optional parameters for sending a group of photos, videos, documents or audios as an album using the sendMediaGroup method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#sendmediagroup
type SendMediaGroupOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Sends messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// (Optional) Protects the contents of the sent messages from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`

	// (Optional) Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`

	// (Optional) Unique identifier of the message effect to be added to the message; for private chats only.
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// (Optional) Description of the message to reply to.
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`
}

// SendMediaGroup sends a group of photos, videos, documents, or audios as an album.
//
// Documents and audio files can be only grouped in an album with messages of the same type.
//
// Required parameters:
//
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername). If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
//   - media: An array describing messages to be sent, must include 2-10 items.
//
// On success, an array of Messages that were sent is returned.
//
// See https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(toChatId any, media []InputMedia, options *SendMediaGroupOptions) ([]Message, error) {
	request := SendMediaGroupRequest{
		ChatID:                toChatId,
		Media:                 media,
		SendMediaGroupOptions: options,
	}
	result := new([]Message)
	// FIXME: nested multipart
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendMediaGroup, request, []string{""}, []string{""}); err != nil {
		return nil, err
	}

	return *result, nil
}
