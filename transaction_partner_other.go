package telegram

// TransactionPartnerOther describes a transaction with an unknown source or recipient.
//
// See "TransactionPartnerOther" https://core.telegram.org/bots/api#transactionpartnerother
type TransactionPartnerOther struct {
	// (Required) Type of the transaction partner, always “other”.
	Type string `json:"type"`
}
