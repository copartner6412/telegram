package telegram

import "net/http"

// GetAvailableGifts returns the list of gifts that can be sent by the bot to users.
//
// Returns a slice of Gift.
//
// See "getAvailableGifts" https://core.telegram.org/bots/api#getavailablegifts
func (b *Bot) GetAvailableGifts() ([]Gift, error) {
	result := new(Gifts)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetAvailableGifts, nil); err != nil {
		return nil, err
	}
	return result.Gifts, nil
}
