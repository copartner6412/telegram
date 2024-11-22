package telegram

import (
	"net/http"
)

// BanChatMemberRequest represents the parameters to ban a user in a group, supergroup, or channel.
//
// In supergroups and channels, the user will not be able to return to the chat on their own
// using invite links, etc., unless unbanned first. The bot must be an administrator in the chat
// with the appropriate administrator rights for this to work.
//
// Required fields:
//   - ChatID
//   - UserID
//   - UntilDate
//   - RevokeMessages
//
// See https://core.telegram.org/bots/api#banchatmember
type BanChatMemberRequest struct {
	// (Required) Integer or String. Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`

	// (Optional) Date when the user will be unbanned; Unix time.
	// If the user is banned for more than 366 days or less than 30 seconds from the current time, they are considered to be banned forever.
	// Applied for supergroups and channels only.
	UntilDate int `json:"until_date,omitempty"`

	// (Optional) Pass True to delete all messages from the chat for the user being removed.
	// If False, the user will be able to see messages in the group sent before they were removed.
	// Always True for supergroups and channels.
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

// BanChatMember bans a user in a group, supergroup, or channel.
//
// In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first.
// The bot must be an administrator in the chat with the appropriate administrator rights for this to work.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername).
//   - userID: Unique identifier of the target user.
//   - untilDate: Date when the user will be unbanned; Unix time. If the user is banned for more than 366 days or less than 30 seconds from the current time, they are considered to be banned forever. Applied for supergroups and channels only. If pass zero, this optional parameter will be ignored.
//   - revokeMessages: Pass True to delete all messages from the chat for the user being removed. If False, the user will be able to see messages in the group sent before they were removed. Always True for supergroups and channels.
//
// See https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(chatID any, userID int, untilDate int, revokeMessages bool) error {
	request := BanChatMemberRequest{
		ChatID:         chatID,
		UserID:         userID,
		UntilDate:      untilDate,
		RevokeMessages: revokeMessages,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodBanChatMember, request); err != nil {
		return err
	}
	return nil
}
