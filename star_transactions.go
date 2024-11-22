package telegram

// StarTransactions contains a list of Telegram Star transactions.
//
// See "StarTransactions" https://core.telegram.org/bots/api#startransactions
type StarTransactions struct {
	// (Required) The list of transactions.
	Transactions []StarTransaction `json:"transactions"`
}
