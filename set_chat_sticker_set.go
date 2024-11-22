package telegram

import (
	"net/http"
)

// SetChatStickerSetRequest represents a request to set a new group sticker set for a supergroup using the setChatStickerSet method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - StickerSetName
//
// See "setChatStickerSet" https://core.telegram.org/bots/api#setchatstickerset
type SetChatStickerSetRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Name of the sticker set to be set as the group sticker set.
	StickerSetName string `json:"sticker_set_name"`
}

// SetChatStickerSet sets a new group sticker set for a supergroup through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - stickerSetName: Name of the sticker set to be set as the group sticker set.
//
// See "setChatStickerSet" https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(chatID any, stickerSetName string) error {
	request := SetChatStickerSetRequest{
		ChatID:         chatID,
		StickerSetName: stickerSetName,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetChatStickerSet, request); err != nil {
		return err
	}
	return nil
}
