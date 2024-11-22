package telegram

import (
	"net/http"
)

// DeleteStickerSetRequest represents to delete a sticker set that was created by the bot using the deleteStickerSet method through the deleteStickerSet method through the Telegram bot API.
//
// Required fields:
//   - Name
//
// See "deleteStickerSet" https://core.telegram.org/bots/api#deletestickerset
type DeleteStickerSetRequest struct {
	// (Required) Sticker set name
	Name string `json:"name"`
}

// DeleteStickerSet deletes a sticker set that was created by the bot.
//
// Parameters:
//   - name: Sticker set name
//
// See "deleteStickerSet" https://core.telegram.org/bots/api#deletestickerset
func (b *Bot) DeleteStickerSet(name string) error {
	request := DeleteStickerSetRequest{
		Name: name,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteStickerSet, request); err != nil {
		return err
	}
	return nil
}
