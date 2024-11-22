package telegram

import "net/http"

// SetStickerSetThumbnail represents a request to set the thumbnail of a regular or mask sticker set using the setStickerSetThumbnail method through the Telegram bot API.
//
// The format of the thumbnail file must match the format of the stickers in the set.
//
// Required fields:
//   - Name
//   - UserID
//   - Format
//   - Thumbnail
//
// See "setStickerSetThumbnail" https://core.telegram.org/bots/api#setstickersetthumbnail
type SetStickerSetThumbnailRequest struct {
	// (Required) Sticker set name.
	Name string `json:"name"`

	// (Required) User identifier of the sticker set owner.
	UserID int `json:"user_id"`

	// (Required) Format of the thumbnail, must be one of:
	//   - “static” for a .WEBP or .PNG image
	//   - “animated” for a .TGS animation
	//   - “video” for a WEBM video
	Format string `json:"format"`

	// (Optional) A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animation-requirements), or a WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a file that already exists on the Telegram servers
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a file from the Internet
	//  3. Pass an absolute path to upload a new one using multipart/form-data.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	// Animated and video sticker set thumbnails can't be uploaded via HTTP URL.
	// If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Thumbnail string `json:"thumbnail,omitempty"`
}

// SetStickerSetThumbnail sets the thumbnail of a regular or mask sticker set through Telegram bot API.
//
// The format of the thumbnail file must match the format of the stickers in the set.
//
// Required parameters:
//   - name: Sticker set name
//   - userID: User identifier of the sticker set owner
//   - format: Format of the thumbnail, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, or “video” for a WEBM video
//   - thumbnail: A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size, or a WEBM video with the thumbnail up to 32 kilobytes in size. Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
//     1. Pass a file_id as a string with "file_id:" prefix to send a file that already exists on the Telegram servers
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a file from the Internet
//     3. Pass an absolute path to upload a new one using multipart/form-data.
//
// See "setStickerSetThumbnail" https://core.telegram.org/bots/api#setstickersetthumbnail
func (b *Bot) SetStickerSetThumbnail(name string, userID int, format, thumbnail string) error {
	request := SetStickerSetThumbnailRequest{
		Name:      name,
		UserID:    userID,
		Format:    format,
		Thumbnail: thumbnail,
	}
	if err := b.sendMultipartForBool(http.MethodPost, MethodSetStickerSetThumbnail, request, []string{"thumbnail"}, nil); err != nil {
		return err
	}
	return nil
}
