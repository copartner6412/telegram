package telegram

import (
	"net/http"
)

// GetCustomEmojiStickersRequest represents parameters to get information about custom emoji stickers by their identifiers.
//
// Required fields:
//   - CustomEmojiIDs
//
// See "getCustomEmojiStickers" https://core.telegram.org/bots/api#getcustomemojistickers
type GetCustomEmojiStickersRequest struct {
	// (Required) A list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
	CustomEmojiIDs []string `json:"custom_emoji_ids"`
}

// GetCustomEmojiStickers retrieves information about custom emoji stickers by their identifiers through the Telegram bot API.
//
// Parameters:
//   - customEmojiIDs: A list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
//
// Returns an array of Sticker objects on success.
//
// See "getCustomEmojiStickers" https://core.telegram.org/bots/api#getcustomemojistickers
func (b *Bot) GetCustomEmojiStickers(customEmojiIDs []string) ([]Sticker, error) {
	request := GetCustomEmojiStickersRequest{
		CustomEmojiIDs: customEmojiIDs,
	}
	result := new([]Sticker)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetCustomEmojiStickers, request); err != nil {
		return nil, err
	}
	return *result, nil

}
