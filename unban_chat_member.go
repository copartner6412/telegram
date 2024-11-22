package telegram

import (
	"net/http"
)

// UnbanChatMemberRequest represents a request to unban a previously banned user in a supergroup or channel using unbanChatMember method through the Telegram bot API.
//
// The user will not return to the group or channel automatically but will be able to join via a link, etc.
// The bot must be an administrator for this to work. By default, this method guarantees that after the
// call, the user is not a member of the chat, but will be able to join it. So, if the user is a member of the chat,
// they will also be removed from the chat. If you don't want this, use the parameter only_if_banned.
//
// Required fields:
//   - ChatID
//   - UserID
//
// See https://core.telegram.org/bots/api#unbanchatmember
type UnbanChatMemberRequest struct {
	// (Required) Integer or string. Unique identifier for the target group or username of the target supergroup or channel
	// (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user
	UserID int `json:"user_id"`

	// (Optional) Do nothing if the user is not banned
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// UnbanChatMember unbans a previously banned user in a supergroup or channel.
//
// The user will not return to the group or channel automatically but will be able to join via a link, etc.
// The bot must be an administrator for this to work. By default, this method guarantees that after the call,
// the user is not a member of the chat but will be able to join it. If the user is a member of the chat,
// they will also be removed from the chat. If this behavior is not desired, use the parameter only_if_banned.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target group or username of the target supergroup or channel through the Telegram bot API.
//   - userID: Unique identifier of the target user
//   - onlyIfBanned: Do nothing if the user is not banned.
//
// See https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(chatID any, userID int, onlyIfBanned bool) error {
	request := UnbanChatMemberRequest{
		ChatID:       chatID,
		UserID:       userID,
		OnlyIfBanned: onlyIfBanned,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodUnbanChatMember, request); err != nil {
		return err
	}
	return nil
}
