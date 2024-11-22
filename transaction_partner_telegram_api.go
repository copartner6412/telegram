package telegram

// TransactionPartnerTelegramApi describes a transaction with payment for paid broadcasting.
//
// See "TransactionPartnerTelegramApi" https://core.telegram.org/bots/api#transactionpartnertelegramapi
type TransactionPartnerTelegramApi struct {
	// (Required) Type of the transaction partner, always “telegram_api”.
	Type string `json:"type"`

	// (Required) The number of successful requests that exceeded regular limits and were therefore billed.
	RequestCount int `json:"request_count"`
}
