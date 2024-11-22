package telegram

import (
	"net/http"
)

// SetStickerKeywordsRequest represents to change search keywords assigned to a regular or custom emoji sticker using the setStickerKeywords through the Telegram bot API.
//
// The sticker must belong to a sticker set created by the bot.
//
// Required fields:
//   - Sticker
//
// See "setStickerKeywords" https://core.telegram.org/bots/api#setstickerkeywords
type SetStickerKeywordsRequest struct {
	// (Required) File identifier of the sticker
	Sticker string `json:"sticker"`

	// (Optional) A list of 0-20 search keywords for the sticker with total length of up to 64 characters
	Keywords []string `json:"keywords,omitempty"`
}

// SetStickerKeywords changes the search keywords assigned to a regular or custom emoji sticker through the Telegram bot API.
//
// The sticker must belong to a sticker set created by the bot.
// Parameters:
//   - sticker: File identifier of the sticker
//   - keywords: A list of 0-20 search keywords for the sticker with total length of up to 64 characters
//
// Returns an error if the request was unsuccessful.
//
// See "setStickerKeywords" https://core.telegram.org/bots/api#setstickerkeywords
func (b *Bot) SetStickerKeywords(sticker string, keywords []string) error {
	request := SetStickerKeywordsRequest{
		Sticker:  sticker,
		Keywords: keywords,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetStickerKeywords, request); err != nil {
		return err
	}
	return nil
}
