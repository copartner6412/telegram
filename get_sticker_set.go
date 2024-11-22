package telegram

import (
	"net/http"
)

// GetStickerSetRequest represents parameters to retrieve a sticker using the getStickerSet method through the Telegram bot API.
//
// Required fields:
//   - Name
//
// See "getStickerSet" https://core.telegram.org/bots/api#getstickerset
type GetStickerSetRequest struct {
	// (Required) Name of the sticker set
	Name string `json:"name"`
}

// GetStickerSet retrieves a sticker set by its name through the Telegram bot API.
//
// Parameters:
//
//	name: Name of the sticker set
//
// Returns the StickerSet object on success.
//
// See "getStickerSet" https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(name string) (*StickerSet, error) {
	request := GetStickerSetRequest{
		Name: name,
	}
	result := new(StickerSet)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetStickerSet, request); err != nil {
		return nil, err
	}
	return result, nil
}
