package telegram

import (
	"net/http"
)

// SendVoiceRequest represents the parameters for sending a voice message using the sendVoice method.
//
// Required fields:
//   - ChatID
//   - Voice
//
// See https://core.telegram.org/bots/api#sendvoice
type SendVoiceRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Audio file to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a file from the Internet
	//  3. Pass an absolute path to upload a new one using multipart/form-data.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Voice string `json:"voice"`

	*SendVoiceOptions
}

type SendVoiceOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	*Caption

	// (Optional) Duration of the voice message in seconds.
	Duration int `json:"duration,omitempty"`

	*SendOptions
}

// SendVoice sends an audio file as a playable voice message.
//
// Your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document).
//
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
//
// Required options:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - voice: Audio file to send.
//     1. Pass a file_id as a string with "file_id:" prefix to send a file that exists on the Telegram servers (recommended)
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a file from the Internet
//     3. Pass an absolute path to upload a new one using multipart/form-data.
//
// See https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(toChatID any, voice string, options *SendVoiceOptions) (*Message, error) {
	request := SendVoiceRequest{
		ChatID:           toChatID,
		Voice:            voice,
		SendVoiceOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendVoice, request, []string{"voice"}, nil); err != nil {
		return nil, err
	}
	return result, nil
}
