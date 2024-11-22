package telegram

import "net/http"

// ApproveChatJoinRequestRequest represents a request to approve a chat join request using the approveChatJoinRequest through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - UserID
//
// See https://core.telegram.org/bots/api#approvechatjoinrequest
type ApproveChatJoinRequestRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`
}

// ApproveChatJoinRequest approves a chat join request.
//
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - userID: Unique identifier of the target user.
//
// See https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(chatID any, userID int) error {
	request := ApproveChatJoinRequestRequest{
		ChatID: chatID,
		UserID: userID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodApproveChatJoinRequest, request); err != nil {
		return err
	}
	return nil
}
