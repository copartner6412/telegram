package telegram

import (
	"net/http"
)

// EditChatSubscriptionInviteLinkRequest represents a request to edit a subscription invite link created by the bot using the EditChatSubscriptionInviteLink method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - InviteLink
//
// See https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
type EditChatSubscriptionInviteLinkRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) The invite link to edit
	InviteLink string `json:"invite_link"`

	// (Optional) Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`
}

// EditChatSubscriptionInviteLink edits a subscription invite link created by the bot.
//
// The bot must have the can_invite_users administrator rights.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - InviteLink: The invite link to edit
//   - name: Invite link name; 0-32 characters
//
// Returns the edited invite link as a ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (b *Bot) EditChatSubscriptionInviteLink(chatID any, inviteLink, name string) (*ChatInviteLink, error) {
	request := EditChatSubscriptionInviteLinkRequest{
		ChatID:     chatID,
		InviteLink: inviteLink,
		Name:       name,
	}
	result := new(ChatInviteLink)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodEditChatSubscriptionInviteLink, request); err != nil {
		return nil, err
	}
	return result, nil
}
