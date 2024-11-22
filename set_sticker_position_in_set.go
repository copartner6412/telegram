package telegram

import (
	"net/http"
)

// SetStickerPositionInSetRequest represents parameters to move a sticker in a set created by the bot to a specific position using the setStickerPositionInSet method through the Telegram bot API
//
// Required fields:
//   - Sticker
//   - Position
//
// See "setStickerPositionInSet" https://core.telegram.org/bots/api#setstickerpositioninset
type SetStickerPositionInSetRequest struct {
	// (Required) File identifier of the sticker
	Sticker string `json:"sticker"`

	// (Required) New sticker position in the set, zero-based
	Position int `json:"position"`
}

// SetStickerPositionInSet moves a sticker in a set created by the bot to a specific position.
//
// Parameters:
//   - sticker: File identifier of the sticker
//   - position: New sticker position in the set, zero-based
//
// Returns an error if the request is unsuccessful.
//
// See "setStickerPositionInSet" https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(sticker string, position int) error {
	request := SetStickerPositionInSetRequest{
		Sticker:  sticker,
		Position: position,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetStickerPositionInSet, request); err != nil {
		return err
	}
	return nil
}
