package telegram

import (
	"net/http"
)

// LeaveChatRequest represents parameters to for you bot to leave a group, supergroup or channel using the leaveChat method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See "leaveChat" https://core.telegram.org/bots/api#leavechat
type LeaveChatRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`
}

// LeaveChat allows the bot to leave a group, supergroup, or channel through the Telegram bot API.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//
// See "leaveChat" https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(chatID any) error {
	request := LeaveChatRequest{
		ChatID: chatID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodLeaveChat, request); err != nil {
		return err
	}
	return nil
}
