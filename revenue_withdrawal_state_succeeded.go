package telegram

// RevenueWithdrawalStateSucceeded represents the state of a revenue withdrawal that has succeeded.
//
// See "RevenueWithdrawalStateSucceeded" https://core.telegram.org/bots/api#revenuewithdrawalstatesucceeded
type RevenueWithdrawalStateSucceeded struct {
	// (Required) Type of the state, always “succeeded”.
	Type string `json:"type"`

	// (Required) Date the withdrawal was completed in Unix time.
	Date int `json:"date"`

	// (Optional) An HTTPS URL that can be used to see transaction details.
	URL string `json:"url,omitempty"`
}
