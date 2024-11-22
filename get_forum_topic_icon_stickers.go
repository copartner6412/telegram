package telegram

import (
	"net/http"
)

// GetForumTopicIconStickers retrieves custom emoji stickers that can be used as a forum topic icon by any user through the Telegram bot API.
//
// This method requires no parameters.
//
// Returns an Array of Sticker objects.
//
// See "getForumTopicIconStickers" https://core.telegram.org/bots/api#getforumtopiciconstickers
func (b *Bot) GetForumTopicIconStickers() ([]Sticker, error) {
	result := new([]Sticker)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetForumTopicIconStickers, nil); err != nil {
		return nil, err
	}
	return *result, nil
}
