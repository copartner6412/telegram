package telegram

// InputPaidMediaPhoto represents the paid media to send which is a photo.
//
// See "InputPaidMediaPhoto" https://core.telegram.org/bots/api#inputpaidmediaphoto
type InputPaidMediaPhoto struct {
	// (Required) Type of the media, must be "photo".
	Type string `json:"type"`

	// (Required) File to send.
	//  1. Pass a file_id as a string to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL for Telegram to get a file from the Internet
	//  3. Pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Media string `json:"media"`
}

func (i InputPaidMediaPhoto) getPaidMediaType() string {
	return i.Type
}
