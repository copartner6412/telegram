package telegram

import "net/http"

// EditChatInviteLinkRequest represents the parameters required to edit a non-primary invite link created by the bot.
//
// Required fields:
//   - ChatID
//   - InviteLink
//
// See https://core.telegram.org/bots/api#editchatinvitelink
type EditChatInviteLinkRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) The invite link to edit.
	InviteLink string `json:"invite_link"`

	*ChatInviteLinkOptions
}

// EditChatInviteLink edits a non-primary invite link created by the bot.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - inviteLink: The invite link to edit.
//
// Returns the edited invite link as a ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(chatID any, inviteLink string, options *ChatInviteLinkOptions) (*ChatInviteLink, error) {
	request := EditChatInviteLinkRequest{
		ChatID:                chatID,
		InviteLink:            inviteLink,
		ChatInviteLinkOptions: options,
	}
	result := new(ChatInviteLink)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodEditChatInviteLink, request); err != nil {
		return nil, err
	}
	return result, nil
}
