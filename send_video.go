package telegram

import (
	"net/http"
)

// Starting December 1, 2024 messages with video that are sent, copied or forwarded to groups and channels with a sufficiently large audience can be automatically scheduled by the server until the respective video is reencoded.
// Such messages will have 0 as their message identifier and can't be used before they are actually sent.

// SendVideoRequest represents the payload for sending video files using sendVideo method.
//
// Telegram clients support MPEG4 format videos (other formats may be sent as Document).
// Bots can currently send video files of up to 50 MB in size; this limit may change in the future.
//
// Required fields:
//   - ChatID
//   - Video
//
// See "sendVideo" https://core.telegram.org/bots/api#sendvideo
type SendVideoRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Video to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a video that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a video from the Internet
	//  3. Pass an absolute path to upload a new video using multipart/form-data.
	Video string `json:"video"`

	*SendVideoOptions
}

// SendVideoOptions represents the optional parameters for sending video files using sendVideo method through Telegram bot API.
//
// See "sendVideo" https://core.telegram.org/bots/api#sendvideo
type SendVideoOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Duration of sent video in seconds.
	Duration int `json:"duration,omitempty"`

	// (Optional) Video width.
	Width int `json:"width,omitempty"`

	// (Optional) Video height.
	Height int `json:"height,omitempty"`

	// (Optional) Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail string `json:"thumbnail,omitempty"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Pass True if the video needs to be covered with a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// (Optional) Pass True if the uploaded video is suitable for streaming.
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	*SendOptions
}

// SendVideo sends video files using sendVideo method through the Telegram Bot API.
//
// Telegram clients support MPEG4 videos (other formats may be sent as Document).
// Bots can currently send video files of up to 50 MB in size.
//
// Required parameters:
//   - toChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - video: Video to send.
//     1. Pass a file_id as a string with "file_id:" prefix to send a video that exists on the Telegram servers (recommended)
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a video from the Internet
//     3. Pass an absolute path to upload a new video using multipart/form-data.
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(toChatID any, video string, options *SendVideoOptions) (*Message, error) {
	request := SendVideoRequest{
		ChatID:           toChatID,
		Video:            video,
		SendVideoOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendVideo, request, []string{"video", "thumbnail"}, []string{"thumbnail"}); err != nil {
		return nil, err
	}

	return result, nil
}
