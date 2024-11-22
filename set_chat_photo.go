package telegram

import (
	"net/http"
)

// SetChatPhotoRequest represents a request to set a new profile photo for a chat using the setChatPhoto method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Photo
//
// See https://core.telegram.org/bots/api#setchatphoto
type SetChatPhotoRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) New chat photo, uploaded using multipart/form-data.
	Photo string `json:"photo"`
}

// SetChatPhoto sets a new profile photo for the chat through the Telegram bot API.
//
// Photos cannot be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - photo: Absolute local path to new chat photo, uploaded using multipart/form-data.
//
// See https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(chatID any, photo string) error {
	request := SetChatPhotoRequest{
		ChatID: chatID,
		Photo:  photo,
	}
	if err := b.sendMultipartForBool(http.MethodPost, MethodSetChatPhoto, request, []string{"photo"}, nil); err != nil {
		return err
	}
	return nil
}
