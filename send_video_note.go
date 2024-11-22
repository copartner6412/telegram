package telegram

import "net/http"

// SendVideoNoteRequest represents the parameters for sending a video note message using the sendVideoNote method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - VideoNote
//
// See https://core.telegram.org/bots/api#sendvideonote
type SendVideoNoteRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Video note to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a video note that exists on the Telegram servers (recommended)
	//  2. Pass an absolute path to upload a new video using multipart/form-data
	// Sending video notes by a URL is currently unsupported.
	VideoNote string `json:"video_note"`

	*SendVideoNoteOptions
}

// SendVideoNoteOptions represents the optional parameters for sending a video note message using the sendVideoNote method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#sendvideonote
type SendVideoNoteOptions struct {

	// (Optional) Duration of sent video in seconds.
	Duration int `json:"duration,omitempty"`

	// (Optional) Video width and height, i.e. diameter of the video message.
	Length int `json:"length,omitempty"`

	// (Optional) Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	Thumbnail string `json:"thumbnail,omitempty"`

	*SendOptions
}

// SendVideoNote sends video messages through the Telegram bot API.
//
// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long.
//
// Required parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - videoNote: Video note to send. Sending video notes by a URL is currently unsupported.
//     1. Pass a file_id as a string with "file_id:" prefix to send a video note that exists on the Telegram servers (recommended)
//     2. Pass an absolute path to upload a new video using multipart/form-data
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(toChatID any, videoNote string, options *SendVideoNoteOptions) (*Message, error) {
	request := SendVideoNoteRequest{
		ChatID:               toChatID,
		VideoNote:            videoNote,
		SendVideoNoteOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendVideoNote, request, []string{"video_note", "thumbnail"}, []string{"thumbnail"}); err != nil {
		return nil, err
	}

	return result, nil
}
