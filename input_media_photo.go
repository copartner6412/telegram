package telegram

// InputMediaPhoto represents a photo to be sent.
//
// See "InputMediaPhoto" https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	// (Required) Type of the result, must be "photo".
	Type string `json:"type"`

	// (Required) File to send.
	//  1. Pass a file_id as a string to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL for Telegram to get a file from the Internet
	//  3. Pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Pass True if the photo needs to be covered with a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

func (i InputMediaPhoto) getMediaType() string {
	return i.Type
}
