package telegram

import (
	"net/http"
)

// UnbanChatSenderChatRequest represents a request to unban a previously banned channel chat in a supergroup or channel using the unbanChatSenderChat method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - SenderChatID
//
// See https://core.telegram.org/bots/api#unbanchatsenderchat
type UnbanChatSenderChatRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target sender chat
	SenderChatID int `json:"sender_chat_id"`
}

// UnbanChatSenderChat unbans a previously banned channel chat in a supergroup or channel through the Telegram bot API.
//
// The bot must be an administrator for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - senderChatID: Unique identifier of the target sender chat
//
// Returns an error if the request fails or if the Telegram API response indicates failure.
//
// See https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(chatID any, senderChatID int) error {
	request := UnbanChatSenderChatRequest{
		ChatID:       chatID,
		SenderChatID: senderChatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnbanChatSenderChat, request); err != nil {
		return err
	}

	return nil
}
