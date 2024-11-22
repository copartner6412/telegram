package telegram

import (
	"net/http"
)

// SetCustomEmojiStickerSetThumbnailRequest represents parameters to set the thumbnail of a custom emoji sticker set using the setCustomEmojiStickerSetThumbnail method through the Telegram bot API.
//
// Required fields:
//   - Name
//
// See "setCustomEmojiStickerSetThumbnail" https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
type SetCustomEmojiStickerSetThumbnailRequest struct {
	// (Required) Sticker set name
	Name string `json:"name"`

	// (Optional) Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
	CustomEmojiID string `json:"custom_emoji_id"`
}

// SetCustomEmojiStickerSetThumbnail sets the thumbnail of a custom emoji sticker set.
//
// Parameters:
//   - name: Sticker set name
//   - custom_emoji_id: Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
//
// See "setCustomEmojiStickerSetThumbnail" https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (b *Bot) SetCustomEmojiStickerSetThumbnail(name, customEmojiID string) error {
	request := SetCustomEmojiStickerSetThumbnailRequest{
		Name:          name,
		CustomEmojiID: customEmojiID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetCustomEmojiStickerSetThumbnail, request); err != nil {
		return err
	}
	return nil
}
