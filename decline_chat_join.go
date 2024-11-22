package telegram

import (
	"net/http"
)

// DeclineChatJoinRequestRequest represents parameters to decline a chat joint request using the declineChatJoinRequest through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - UserID
//
// See "declineChatJoinRequest" https://core.telegram.org/bots/api#declinechatjoinrequest
type DeclineChatJoinRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`
}

// DeclineChatJoinRequest declines a chat join request through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - userID: Unique identifier of the target user.
//
// See https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(chatID any, userID int) error {
	request := DeclineChatJoinRequest{
		ChatID: chatID,
		UserID: userID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeclineChatJoinRequest, request); err != nil {
		return err
	}
	return nil
}
