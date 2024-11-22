package telegram

import (
	"net/http"
)

// BanChatSenderChat represents a request to ban a channel chat in a supergroup or a channel using the banChatSenderChat method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - SenderChatID
//
// See https://core.telegram.org/bots/api#banchatsenderchat
type BanChatSenderChatRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target sender chat
	SenderChatID int `json:"sender_chat_id"`
}

// BanChatSenderChat bans a channel chat in a supergroup or a channel through the Telegram bot API.
//
// Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels.
// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - senderChatID: Unique identifier of the target sender chat
//
// See https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(chatID any, senderChatID int) error {
	request := BanChatSenderChatRequest{
		ChatID:       chatID,
		SenderChatID: senderChatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodBanChatSenderChat, request); err != nil {
		return err
	}
	return nil
}
