package telegram

import "net/http"

// SetStickerEmojiListRequest represents parameters to change the list of emoji assigned to a regular or custom emoji sticker using the setStickerEmojiList method through the Telegram bot API.
//
// Required fields:
//   - Sticker
//   - EmojiList
//
// See "setStickerEmojiList" https://core.telegram.org/bots/api#setstickeremojilist
type SetStickerEmojiListRequest struct {
	// (Required) File identifier of the sticker
	Sticker string `json:"sticker"`

	// (Required) A list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list"`
}

// SetStickerEmojiList changes the list of emoji assigned to a regular or custom emoji sticker through the Telegram bot API.
//
// The sticker must belong to a sticker set created by the bot.
//
// Parameters:
//   - sticker: File identifier of the sticker
//   - emoji_list: A list of 1-20 emoji associated with the sticker
//
// Returns an error if the request was unsuccessful.
//
// See "setStickerEmojiList" https://core.telegram.org/bots/api#setstickeremojilist
func (b *Bot) SetStickerEmojiList(sticker string, emojiList []string) error {
	request := SetStickerEmojiListRequest{
		Sticker:   sticker,
		EmojiList: emojiList,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetStickerEmojiList, request); err != nil {
		return err
	}
	return nil
}
