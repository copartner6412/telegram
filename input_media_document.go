package telegram

// InputMediaDocument represents a general file to be sent.
//
// See "InputMediaDocument" https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	// (Required) Type of the result, must be "document".
	Type string `json:"type"`

	// (Required) File to send.
	//  1. Pass a file_id as a string to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL for Telegram to get a file from the Internet
	//  3. Pass “attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name.
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

	// (Optional) Disables automatic server-side content type detection for files uploaded using multipart/form-data.
	// Always True, if the document is sent as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

func (i InputMediaDocument) getMediaType() string {
	return i.Type
}
