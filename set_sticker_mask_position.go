package telegram

import (
	"net/http"
)

// SetStickerMaskPositionRequest represents the parameters  to change the mask position of a mask sticker using the setStickerMaskPosition method through the Telegram bot API.
//
// The sticker must belong to a sticker set that was created by the bot.
//
// Required fields:
//   - Sticker
//
// See "setStickerMaskPosition" https://core.telegram.org/bots/api#setstickermaskposition
type SetStickerMaskPositionRequest struct {
	// (Required) File identifier of the sticker.
	Sticker string `json:"sticker"`

	// (Optional) An object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// SetStickerMaskPosition changes the mask position of a mask sticker through the Telegram bot API.
//
// The sticker must belong to a sticker set that was created by the bot.
//
// Parameters:
//   - sticker: File identifier of the sticker
//   - maskPosition: An object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
//
// See "setStickerMaskPosition" https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerMaskPosition(sticker string, maskPosition *MaskPosition) error {
	request := SetStickerMaskPositionRequest{
		Sticker:      sticker,
		MaskPosition: maskPosition,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetStickerMaskPosition, request); err != nil {
		return err
	}
	return nil
}
