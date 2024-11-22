package telegram

import "net/http"

// SendPaidMediaRequest represents the parameters for sending paid media using the sendPaidMedia method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - StarCount
//   - Media
//
// See https://core.telegram.org/bots/api#sendpaidmedia
type SendPaidMediaRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	// If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance.
	// Otherwise, they will be credited to the bot's balance.
	ChatID any `json:"chat_id"`

	// (Required) The number of Telegram Stars that must be paid to buy access to the media; 1-2500.
	StarCount int `json:"star_count"`

	// (Required) An array describing the media to be sent; up to 10 items.
	Media []InputPaidMedia `json:"media"`

	*SendPaidMediaOptions
}

// SendPaidMediaOptions represents the optional parameters for sending paid media using the sendPaidMedia method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#sendpaidmedia
type SendPaidMediaOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload,omitempty"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// (Optional) Protects the contents of the sent message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`

	// (Optional) Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`

	// (Optional) Description of the message to reply to.
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// (Optional) Additional interface options.
	//
	// Types:
	//  - InlineKeyboardMarkup
	//  - ReplyKeyboardMarkup
	//  - ReplyKeyboardRemove
	//  - ForceReply
	ReplyMarkup any `json:"reply_markup,omitempty"`
}

// SendPaidMedia sends paid media on behalf of the bot though the Telegram bot API.
//
// Required parameters:
//   - toChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername). If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
//   - media: An array describing the media to be sent; up to 10 items.
//   - starCount: The number of Telegram Stars that must be paid to buy access to the media; 1-2500.
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendpaidmedia
func (b *Bot) SendPaidMedia(toChatID any, media []InputPaidMedia, startCount int, options *SendPaidMediaOptions) (*Message, error) {
	request := SendPaidMediaRequest{
		ChatID:               toChatID,
		StarCount:            startCount,
		Media:                media,
		SendPaidMediaOptions: options,
	}
	result := new(Message)
	// FIXME: create a nested multipart
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendPaidMedia, request, []string{""}, []string{""}); err != nil {
		return nil, err
	}
	return result, nil
}
