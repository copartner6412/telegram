package telegram

// ChatInviteLink represents an invite link for a chat.
//
// See "ChatInviteLink" https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	// (Required) The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`

	// (Required) Creator of the link.
	Creator User `json:"creator"`

	// (Required) True, if users joining the chat via the link need to be approved by chat administrators.
	CreatesJoinRequest bool `json:"creates_join_request"`

	// (Required) True, if the link is primary.
	IsPrimary bool `json:"is_primary"`

	// (Required) True, if the link is revoked.
	IsRevoked bool `json:"is_revoked"`

	// (Optional) Invite link name.
	Name string `json:"name,omitempty"`

	// (Optional) Point in time (Unix timestamp) when the link will expire or has been expired.
	ExpireDate int `json:"expire_date,omitempty"`

	// (Optional) The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999.
	MemberLimit int `json:"member_limit,omitempty"`

	// (Optional) Number of pending join requests created using this link.
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`

	// (Optional) The number of seconds the subscription will be active for before the next payment.
	SubscriptionPeriod int `json:"subscription_period,omitempty"`

	// (Optional) The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link.
	SubscriptionPrice int `json:"subscription_price,omitempty"`
}
