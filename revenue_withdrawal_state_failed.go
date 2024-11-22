package telegram

// RevenueWithdrawalStateFailed represents the state of a revenue withdrawal that has failed.
//
// See "RevenueWithdrawalStateFailed" https://core.telegram.org/bots/api#revenuewithdrawalstatefailed
type RevenueWithdrawalStateFailed struct {
	// (Required) Type of the state, always “failed”.
	Type string `json:"type"`
}
