package telegram

import "net/http"

// UploadStickerFileRequest represents the parameters required to upload a file with a sticker for later use.
//
// The file can be used multiple times in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods.
//
// Required fields:
//   - UserID
//   - Sticker
//   - StickerFormat
//
// See "uploadStickerFile" https://core.telegram.org/bots/api#uploadstickerfile
type UploadStickerFileRequest struct {
	// (Required) User identifier of sticker file owner.
	UserID int `json:"user_id"`

	// (Required) A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format.
	//
	// See https://core.telegram.org/stickers for technical requirements.
	//
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Sticker string `json:"sticker"`

	// (Required) Format of the sticker, must be one of:
	//   - “static”
	//   - “animated”
	//   - “video”.
	StickerFormat string `json:"sticker_format"`
}

// UploadStickerFile uploads a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods.
// The file can be used multiple times and must meet the required format specifications.
//
// Parameters:
//   - userID: User identifier of sticker file owner.
//   - sticker: A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format.
//   - stickerFormat: Format of the sticker, must be one of “static”, “animated”, “video”.
//
// This method returns the uploaded File on success.
//
// See "uploadStickerFile" https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(userID int, sticker, stickerFormat string) (*File, error) {
	request := UploadStickerFileRequest{
		UserID:        userID,
		Sticker:       sticker,
		StickerFormat: sticker,
	}
	result := new(File)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodUploadStickerFile, request); err != nil {
		return nil, err
	}
	return result, nil
}
