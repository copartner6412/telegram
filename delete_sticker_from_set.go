package telegram

import "net/http"

// DeleteStickerFromSetRequest represents parameters to delete a sticker from a set created by the bot using the deleteStickerFromSet through the Telegram bot API.
//
// Required fields:
//   - Sticker
//
// See "deleteStickerFromSet" https://core.telegram.org/bots/api#deletestickerfromset
type DeleteStickerFromSetRequest struct {
	// (Required) File identifier of the sticker
	Sticker string `json:"sticker"`
}

// DeleteStickerFromSet removes a sticker from a set created by the bot.
//
// Parameters:
//   - sticker: File identifier of the sticker
//
// Returns an error if the request is unsuccessful.
//
// See "deleteStickerFromSet" https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(sticker string) error {
	request := DeleteStickerFromSetRequest{
		Sticker: sticker,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteStickerFromSet, request); err != nil {
		return err
	}
	return nil
}
