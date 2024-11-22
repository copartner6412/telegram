package telegram

import (
	"net/http"
)

// SendAnimationRequest represents the parameters for sending an animation message using the sendAnimation method through the Telegram bot API.
//
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
//
// Required fields:
//   - ChatID
//   - Animation
//
// See https://core.telegram.org/bots/api#sendanimation
type SendAnimationRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Animation to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send an animation that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get an animation from the Internet
	//  3. Pass an absolute path to upload a new animation using multipart/form-data.
	Animation string `json:"animation"`

	*SendAnimationOptions
}

// SendAnimationOptions represents the optional parameters for sending an animation message using the sendAnimation method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#sendanimation
type SendAnimationOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`

	// (Optional) Animation width
	Width int `json:"width,omitempty"`

	// (Optional) Animation height
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

	// (Optional) Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	*SendOptions
}

// SendAnimation sends an animation file (GIF or H.264/MPEG-4 AVC video without sound) to a specified chat or channel through Telegram bot API.
//
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
//
// Required parameters:
//   - toChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - animation: Animation to send.
//     1. Pass a file_id as a string with "file_id:" prefix to send an animation that exists on the Telegram servers (recommended)
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get an animation from the Internet
//     3. Pass an absolute path to upload a new animation using multipart/form-data.
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(toChatID any, animation string, options *SendAnimationOptions) (*Message, error) {
	request := SendAnimationRequest{
		ChatID:               toChatID,
		Animation:            animation,
		SendAnimationOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendAnimation, request, []string{"animation", "thumbnail"}, []string{"thumbnail"}); err != nil {
		return nil, err
	}

	return result, nil
}
