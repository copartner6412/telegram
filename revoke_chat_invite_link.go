package telegram

import (
	"net/http"
)

// RevokeChatInviteLinkRequest represents a request to revoke an invite link created by the bot.
//
// Required fields:
//   - ChatID
//   - InviteLink
//
// See https://core.telegram.org/bots/api#revokechatinvitelink
type RevokeChatInviteLinkRequest struct {
	// (Required) Integer or string. Unique identifier of the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) The invite link to revoke.
	InviteLink string `json:"invite_link"`
}

// RevokeChatInviteLink revokes an invite link created by the bot.
//
// If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier of the target chat or username of the target channel (in the format @channelusername).
//   - inviteLink: The invite link to revoke.
//
// Returns the revoked invite link as a ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(chatID any, inviteLink string) (*ChatInviteLink, error) {
	request := RevokeChatInviteLinkRequest{
		ChatID:     chatID,
		InviteLink: inviteLink,
	}
	result := new(ChatInviteLink)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodRevokeChatInviteLink, request); err != nil {
		return nil, err
	}
	return result, nil
}
