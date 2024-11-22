package telegram

// InputSticker describes a sticker to be added to a sticker set.
//
// Required fields:
//   - Sticker
//   - Format
//   - EmojiList
//
// See "InputSticker" https://core.telegram.org/bots/api#inputSticker
type InputSticker struct {
	// (Required) The added sticker.
	//  1. Pass a file_id as a string to send a file that already exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string for Telegram to get a file from the Internet
	//  3. Pass a path to upload a new one using multipart/form-data
	//  4. Pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// Animated and video stickers can't be uploaded via HTTP URL.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Sticker string `json:"sticker"`

	// (Required) Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a WEBM video.
	Format string `json:"format"`

	// (Required) List of 1-20 emoji associated with the sticker.
	EmojiList []string `json:"emoji_list"`

	// (Optional) Position where the mask should be placed on faces. For
	// “mask” stickers only.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// (Optional) List of 0-20 search keywords for the sticker with total
	// length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
	Keywords []string `json:"keywords,omitempty"`
}
