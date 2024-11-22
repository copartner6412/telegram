package telegram

import (
	"net/http"
)

// GetChatMemberCountRequest represents the request payload to get the number of members in a chat using the getChatMemberCount through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "getChatMemberCount" https://core.telegram.org/bots/api#getchatmembercount
type GetChatMemberCountRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`
}

// GetChatMemberCount retrieves the number of members in a specified chat through the Telegram bot API.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
// It returns an integer indicating the member count on success.
//
// See "getChatMemberCount" https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(chatID any) (int, error) {
	request := GetChatAdministratorsRequest{
		ChatID: chatID,
	}
	result := new(int)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetChatMember, request); err != nil {
		return 0, err
	}
	return *result, nil
}
