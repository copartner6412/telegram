package telegram

// PaidMediaPurchased contains information about a paid media purchase.
//
// See "PaidMediaPurchased" https://core.telegram.org/bots/api#paidmediapurchased
type PaidMediaPurchased struct {
	// (Required) User who purchased the media.
	FromUser User `json:"from_user"`

	// (Required) Bot-specified paid media payload.
	PaidMediaPayload string `json:"paid_media_payload"`
}
