package telegram

import (
	"net/http"
)

// GetStarTransactionsRequest represents the request parameters for retrieving the bot's Telegram Star transactions  in chronological order using the getStarTransaction method through the Telegram bot API.
//
// On success, returns a StarTransactions object.
//
// See "getStarTransactions" https://core.telegram.org/bots/api#getstartransactions
type GetStarTransactionsRequest struct {
	// (Optional) Offset for the number of transactions to skip in the response.
	Offset int `json:"offset,omitempty"`

	// (Optional) Maximum number of transactions to retrieve, between 1-100.
	// Defaults to 100 if not specified.
	Limit int `json:"limit,omitempty"`
}

// GetStarTransactions retrieves the bot's Telegram Star transactions in chronological order through the Telegram bot API.
//
// Parameters:
//   - offset: Number of transactions to skip in the response
//   - limit: The maximum number of transactions to be retrieved. Values between 1-100 are accepted. If zero, defaults to 100.
//
// If successful, returns a StarTransactions object.
//
// See "getStarTransactions" https://core.telegram.org/bots/api#getstartransactions
func (b *Bot) GetStarTransactions(offset, limit int) ([]StarTransaction, error) {
	request := GetStarTransactionsRequest{
		Offset: offset,
		Limit:  limit,
	}
	result := new(StarTransactions)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodGetStarTransactions, request); err != nil {
		return nil, err
	}
	return result.Transactions, nil
}
