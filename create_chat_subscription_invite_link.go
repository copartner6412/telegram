package telegram

import "net/http"

// CreateChatSubscriptionInviteLinkRequest represents a request to create a subscription invite link for a channel chat using the createChatSubscriptionInviteLink method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - SubscriptionPeriod
//   - SubscriptionPrice
//
// See https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
type CreateChatSubscriptionInviteLinkRequest struct {
	// (Required) Integer or string. Unique identifier for the target channel chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPeriod int `json:"subscription_period"`

	// (Required) The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500.
	SubscriptionPrice int `json:"subscription_price"`

	// (Optional) Invite link name; 0-32 characters.
	Name string `json:"name,omitempty"`
}

// CreateChatSubscriptionInviteLink creates a subscription invite link for a channel chat.
//
// The bot must have the can_invite_users administrator rights. The link can be edited using the method
// editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target channel chat or username of the target channel (in the format @channelusername).
//   - subscriptionPeriod: The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
//   - subscriptionPrice: The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500.
//   - name: Invite link name; 0-32 characters.
//
// Returns the new invite link as a ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (b *Bot) CreateChatSubscriptionInviteLink(chatID any, subscriptionPeriod int, subscriptionPrice int, name string) (*ChatInviteLink, error) {
	request := CreateChatSubscriptionInviteLinkRequest{
		ChatID:             chatID,
		SubscriptionPeriod: subscriptionPeriod,
		SubscriptionPrice:  subscriptionPrice,
		Name:               name,
	}
	result := new(ChatInviteLink)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodCreateChatSubscriptionInviteLink, request); err != nil {
		return nil, err
	}
	return result, nil
}
