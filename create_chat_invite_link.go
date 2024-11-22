package telegram

import (
	"net/http"
)

// CreateChatInviteLinkRequest represents the request payload for creating an additional invite link for a chat using the createChatInviteLink through the Telegram bot API.
//
// Required fields:
//   - ChatID
//
// See https://core.telegram.org/bots/api#createchatinvitelink
type CreateChatInviteLinkRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	*ChatInviteLinkOptions
}

// See https://core.telegram.org/bots/api#createchatinvitelink
//
// See https://core.telegram.org/bots/api#editchatinvitelink
type ChatInviteLinkOptions struct {
	// (Optional) Invite link name; 0-32 characters.
	Name string `json:"name,omitempty"`

	// (Optional) Point in time (Unix timestamp) when the link will expire.
	ExpireDate int `json:"expire_date,omitempty"`

	// (Optional) The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999.
	MemberLimit int `json:"member_limit,omitempty"`

	// (Optional) True, if users joining the chat via the link need to be approved by chat administrators.
	// If True, member_limit can't be specified.
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// CreateChatInviteLink creates an additional invite link for a chat through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// The link can be revoked using the method revokeChatInviteLink.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//
// Returns the new invite link as a ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(chatID any, options *ChatInviteLinkOptions) (*ChatInviteLink, error) {
	request := CreateChatInviteLinkRequest{
		ChatID:                chatID,
		ChatInviteLinkOptions: options,
	}
	result := new(ChatInviteLink)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodCreateChatInviteLink, request); err != nil {
		return nil, err
	}
	return result, nil
}
