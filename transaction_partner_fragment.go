package telegram

import "encoding/json"

// TransactionPartnerFragment describes a withdrawal transaction with Fragment.
//
// See "TransactionPartnerFragment" https://core.telegram.org/bots/api#transactionpartnerfragment
type TransactionPartnerFragment struct {
	// (Required) Type of the transaction partner, always “fragment”.
	Type string `json:"type"`

	// (Optional) State of the transaction if the transaction is outgoing.
	//
	// Its type can be one of:
	//   - RevenueWithdrawalStatePending
	//   - RevenueWithdrawalStateSucceeded
	//   - RevenueWithdrawalStateFailed
	WithdrawalState json.RawMessage `json:"withdrawal_state,omitempty"`
}
