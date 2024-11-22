package telegram

import (
	"net/http"
)

// GetUserChatBoostsRequest represents parameters to get the list of boosts added to a chat by a user using the getUserChatBoosts method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - UserID
//
// See "getUserChatBoosts" https://core.telegram.org/bots/api#getuserchatboosts
type GetUserChatBoostsRequest struct {
	// (Required) Integer or string. Unique identifier for the chat or username of the channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user
	UserID int `json:"user_id"`
}

// GetUserChatBoosts retrieves the list of boosts added to a chat by a specific user through the Telegram bot API.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the chat or username of the channel (in the format @channelusername)
//   - userID: Unique identifier of the target user
//
// This method requires administrator rights in the chat.
//
// See "getUserChatBoosts" https://core.telegram.org/bots/api#getuserchatboosts
func (b *Bot) GetUserChatBoosts(chatID any, userID int) (*UserChatBoosts, error) {
	request := GetUserChatBoostsRequest{
		ChatID: chatID,
		UserID: userID,
	}
	result := new(UserChatBoosts)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetUserChatBoosts, request); err != nil {
		return nil, err
	}
	return result, nil
}
