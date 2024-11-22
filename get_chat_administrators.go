package telegram

import (
	"net/http"
)

// GetChatAdministratorsRequest represents parameters to get a list of administrators in a chat, which aren't bots.
//
// Required fields:
//   - ChatID
//
// See "getChatAdministrators" https://core.telegram.org/bots/api#getchatadministrators
type GetChatAdministratorsRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`
}

// GetChatAdministrators retrieves a list of administrators in a specified chat, excluding bots.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
// The method returns an array of ChatMember objects on success.
//
// See https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(chatID any) ([]ChatMemberAdministrator, error) {
	request := GetChatAdministratorsRequest{
		ChatID: chatID,
	}
	result := new([]ChatMemberAdministrator)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetChatAdministrators, request); err != nil {
		return nil, err
	}
	return *result, nil
}
