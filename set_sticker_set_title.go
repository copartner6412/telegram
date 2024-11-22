package telegram

import "net/http"

// SetStickerSetTitleRequest sets the title of a created sticker set through the Telegram bot API.
//
// Required fields:
//   - Name
//   - Title
//
// See "setStickerSetTitle" https://core.telegram.org/bots/api#setstickersettitle
type SetStickerSetTitleRequest struct {
	// (Required) Sticker set name
	Name string `json:"name"`

	// (Required) Sticker set title, 1-64 characters
	Title string `json:"title"`
}

// SetStickerSetTitle sets the title of a created sticker set through the Telegram bot API.
//
// Parameters:
//   - name: Sticker set name
//   - title: Sticker set title, 1-64 characters
//
// See "setStickerSetTitle" https://core.telegram.org/bots/api#setstickersettitle
func (b *Bot) SetStickerSetTitle(name string, title string) error {
	request := SetStickerSetTitleRequest{
		Name:  name,
		Title: title,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetStickerSetTitle, request); err != nil {
		return err
	}
	return nil
}
