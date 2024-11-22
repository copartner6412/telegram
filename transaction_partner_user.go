package telegram

import "encoding/json"

// TransactionPartnerUser describes a transaction with a user.
//
// This object contains information about the transaction partner, which is a user in this case.
// It includes details about the user, any invoice payloads specified by the bot,
// and information regarding paid media associated with the transaction.
//
// See https://core.telegram.org/bots/api#transactionpartneruser
type TransactionPartnerUser struct {
	// (Required) Type of the transaction partner, always “user”.
	Type string `json:"type"`

	// (Required) Information about the user.
	User User `json:"user"`

	// (Optional) Bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload,omitempty"`

	// (Optional) Information about the paid media bought by the user.
	//
	// Its type can be one of:
	//   - []PaidMediaPreview
	//   - []PaidMediaPhoto
	//   - []PaidMediaVideo
	PaidMedia json.RawMessage `json:"paid_media,omitempty"`

	// (Optional) Bot-specified paid media payload.
	PaidMediaPayload string `json:"paid_media_payload,omitempty"`
}
