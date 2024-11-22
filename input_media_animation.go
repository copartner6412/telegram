package telegram

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//
// See "InputMediaAnimation" https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	// (Required) Type of the result, must be "animation".
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

	// (Optional) Pass True, if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) Animation width.
	Width int `json:"width,omitempty"`

	// (Optional) Animation height.
	Height int `json:"height,omitempty"`

	// (Optional) Animation duration in seconds.
	Duration int `json:"duration,omitempty"`

	// (Optional) Pass True if the animation needs to be covered with a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

func (i InputMediaAnimation) getMediaType() string {
	return i.Type
}
