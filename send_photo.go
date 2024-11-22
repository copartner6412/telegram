package telegram

import "net/http"

// SendPhotoRequest represents the parameters for sending a photo message using sendPhoto method.
//
// Required fields:
//   - ChatID
//   - Photo
//
// See "sendPhoto" https://core.telegram.org/bots/api#sendphoto
type SendPhotoRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Photo to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a photo that exists on the Telegram servers (recommended).
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a photo from the Internet.
	//  3. Pass an absolute path to upload a new photo using multipart/form-data.
	// The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total.
	// Width and height ratio must be at most 20. More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Photo string `json:"photo"`

	*SendPhotoOptions
}

// SendPhotoOptions represents the optional parameters for sending a photo message using sendPhoto method.
//
// See "sendPhoto" https://core.telegram.org/bots/api#sendphoto
type SendPhotoOptions struct {

	// (Optional) Unique identifier of the business connection on behalf of which
	// the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum;
	// for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	*SendOptions
}

// SendPhoto sends a photo to a chat. On success, the sent Message is returned.
//
// The photo can be specified by providing a file path to upload, a file_id of an existing photo on
// Telegram servers, or a URL for Telegram to get the photo from the Internet.
//
// Required parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - photo: Photo to send. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total.
//     1. Pass a file_id as a string with "file_id:" prefix to send a photo that exists on the Telegram servers (recommended).
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a photo from the Internet.
//     3. Pass an absolute path to upload a new photo using multipart/form-data.
//
// See "sendPhoto" https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(toChatID any, photo string, options *SendPhotoOptions) (*Message, error) {
	request := SendPhotoRequest{
		ChatID:           toChatID,
		Photo:            photo,
		SendPhotoOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendPhoto, request, []string{"photo"}, nil); err != nil {
		return nil, err
	}

	return result, nil
}
