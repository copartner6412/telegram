package telegram

import (
	"net/http"
)

// DeleteChatStickerSetRequest represents a request to delete a group sticker set from a supergroup using the deleteChatStickerSet method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "deleteChatStickerSetRequest" https://core.telegram.org/bots/api#deletechatstickerset
type DeleteChatStickerSetRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`
}

// DeleteChatStickerSet deletes a group sticker set from a supergroup through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//
// See "deleteChatStickerSetRequest" https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(chatID any) error {
	request := DeleteChatPhotoRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteChatStickerSet, request); err != nil {
		return err
	}
	return nil
}
