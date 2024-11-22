package telegram

import (
	"net/http"
)

// SendAudioRequest represents the parameters to send an audio file to be displayed in the music player using sendAudio method through Telegram bot API.
//
// Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the sendVoice method instead.
//
// Required fields:
//   - ChatID
//   - Audio
//
// See "sendAudio" https://core.telegram.org/bots/api#sendaudio
type SendAudioRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Audio file to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send an audio file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get an audio file from the Internet
	//  3. Pass an absolute path to upload a new one using multipart/form-data.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Audio string `json:"audio"`

	*SendAudioOptions
}

// SendAudioOptions represents the optional parameters to send an audio file to be displayed in the music player using sendAudio method through Telegram bot API.
//
// See "sendAudio" https://core.telegram.org/bots/api#sendaudio
type SendAudioOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum;
	// for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	*Caption

	// (Optional) Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`

	// (Optional) Performer
	Performer string `json:"performer,omitempty"`

	// (Optional) Track name
	Title string `json:"title,omitempty"`

	// (Optional) Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail string `json:"thumbnail,omitempty"`

	*SendOptions
}

// SendAudio sends an audio file to be displayed in the music player.
//
// Your audio must be in the .MP3 or .M4A format.
// Bots can currently send audio files of up to 50 MB in size.
//
// Required parameters:
//   - toChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - audio: Audio file to send.
//     1. Pass a file_id as a string with "file_id:" prefix to send an audio file that exists on the Telegram servers (recommended)
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get an audio file from the Internet
//     3. Pass an absolute path to upload a new one using multipart/form-data.
//
// On success, the sent Message is returned.
//
// See "sendAudio" https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(toChatID any, audio string, options *SendAudioOptions) (*Message, error) {
	request := SendAudioRequest{
		ChatID:           toChatID,
		Audio:            audio,
		SendAudioOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendAudio, request, []string{"audio", "thumbnail"}, []string{"thumbnail"}); err != nil {
		return nil, err
	}

	return result, nil
}
