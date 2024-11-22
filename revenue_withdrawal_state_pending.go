package telegram

// RevenueWithdrawalStatePending represents the state of a revenue withdrawal that is in progress.
//
// See "RevenueWithdrawalStatePending" https://core.telegram.org/bots/api#revenuewithdrawalstatepending
type RevenueWithdrawalStatePending struct {
	// (Required) Type of the state, always “pending”.
	Type string `json:"type"`
}
