package telegram

import (
	"net/http"
)

// GetChatRequest represents parameters to to get up-to-date information about the chat using the getChat method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "getChat" https://core.telegram.org/bots/api#getchat
type GetChatRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`
}

// GetChat retrieves up-to-date information about the specified chat through the Telegram bot API.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
// Returns a ChatFullInfo object on success.
//
// See https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(chatID any) (*ChatFullInfo, error) {
	request := GetChatRequest{
		ChatID: chatID,
	}
	result := new(ChatFullInfo)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetChat, request); err != nil {
		return nil, err
	}
	return result, nil
}
