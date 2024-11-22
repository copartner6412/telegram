package telegram

import "encoding/json"

// StarTransaction describes a Telegram Star transaction.
//
// See "StarTransaction" https://core.telegram.org/bots/api#startransaction
type StarTransaction struct {
	// (Required) Unique identifier of the transaction.
	// Coincides with the identifier of the original transaction for refund transactions.
	// Coincides with SuccessfulPayment.telegram_payment_charge_id for successful incoming payments from users.
	ID string `json:"id"`

	// (Required) Number of Telegram Stars transferred by the transaction.
	Amount int `json:"amount"`

	// (Optional) Source of an incoming transaction (e.g., a user purchasing goods or services,
	// Fragment refunding a failed withdrawal). Only for incoming transactions.
	//
	// Its type can be one of:
	//   - TransactionPartnerUser
	//   - TransactionPartnerFragment
	//   - TransactionPartnerTelegramAds
	//   - TransactionPartnerTelegramApi
	//   - TransactionPartnerOther
	Source json.RawMessage `json:"source,omitempty"`

	// (Optional) Receiver of an outgoing transaction (e.g., a user for a purchase refund,
	// Fragment for a withdrawal). Only for outgoing transactions.
	//
	// Its type can be one of:
	//   - TransactionPartnerUser
	//   - TransactionPartnerFragment
	//   - TransactionPartnerTelegramAds
	//   - TransactionPartnerTelegramApi
	//   - TransactionPartnerOther
	Receiver json.RawMessage `json:"receiver,omitempty"`
}
