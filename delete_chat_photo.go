package telegram

import (
	"net/http"
)

// DeleteChatPhotoRequest represents a request to delete a chat photo using the deleteChatPhoto method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See https://core.telegram.org/bots/api#deletechatphoto
type DeleteChatPhotoRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`
}

// DeleteChatPhoto deletes a chat photo through the Telegram bot API.
//
// Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
//
// Required fields:
//   - ChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//
// See https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(chatID any) error {
	request := DeleteChatPhotoRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteChatPhoto, request); err != nil {
		return err
	}
	return nil
}
