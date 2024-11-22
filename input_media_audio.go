package telegram

// InputMediaAudio represents an audio file to be treated as music to be sent.
//
// See "InputMediaAudio" https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	// (Required) Type of the result, must be "audio".
	Type string `json:"type"`

	// (Required) File to send.
	//  1. Pass a file_id as a string to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL for Telegram to get a file from the Internet
	//  3. Pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`

	// (Optional) Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail string `json:"thumbnail,omitempty"`

	*Caption

	// (Optional) Duration of the audio in seconds.
	Duration int `json:"duration,omitempty"`

	// (Optional) Performer of the audio.
	Performer string `json:"performer,omitempty"`

	// (Optional) Title of the audio.
	Title string `json:"title,omitempty"`
}

func (i InputMediaAudio) getMediaType() string {
	return i.Type
}
