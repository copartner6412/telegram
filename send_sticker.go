package telegram

import (
	"net/http"
)

// SendStickerRequest represents the parameters used to send static .WEBP, animated .TGS, or video .WEBM a sticker using the sendSticker method via the Telegram bot API.
//
// Required fields:
//   - ChatID:
//   - Sticker
//
// See "sendSticker" https://core.telegram.org/bots/api#sendSticker
type SendStickerRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Sticker to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a .WEBP sticker from the Internet
	//  3. Pass an absolute path to upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	//
	// Video and animated stickers can't be sent via an HTTP URL.
	Sticker string `json:"sticker"`

	*SendStickerOptions
}

type SendStickerOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Emoji associated with the sticker; only for just uploaded stickers.
	Emoji string `json:"emoji,omitempty"`

	*SendOptions
}

// SendSticker sends a static .WEBP, animated .TGS, or video .WEBM sticker via the Telegram API.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - sticker: Sticker to send. Video and animated stickers can't be sent via an HTTP URL.
//     1. Pass a file_id as a string with "file_id:" prefix to send a file that exists on the Telegram servers (recommended)
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a .WEBP sticker from the Internet
//     3. Pass an absolute path to upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data.
//
// On success, the sent Message is returned.
//
// See "sendSticker" https://core.telegram.org/bots/api#sendSticker
func (b *Bot) SendSticker(chatID any, sticker string, options *SendStickerOptions) (*Message, error) {
	request := SendStickerRequest{
		ChatID:             chatID,
		Sticker:            sticker,
		SendStickerOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendSticker, request, []string{"sticker"}, nil); err != nil {
		return nil, err
	}
	return result, nil
}
