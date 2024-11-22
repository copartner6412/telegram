package telegram

import (
	"net/http"
)

// GetBusinessConnectionRequest represents parameters to get information about the connection of the bot with a business account using the getBusinessConnection method through the Telegram bot API.
//
// Required fields:
//   - BusinessConnectionID
//
// See "getBusinessConnection" https://core.telegram.org/bots/api#getbusinessconnection
type GetBusinessConnectionRequest struct {
	// (Required) Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id"`
}

// GetBusinessConnection retrieves information about the connection of the bot with a business account.
//
// Parameters:
//   - businessConnectionID: Unique identifier of the business connection
//
// This method returns a BusinessConnection object on success.
//
// See "getBusinessConnection" https://core.telegram.org/bots/api#getbusinessconnection
func (b *Bot) GetBusinessConnection(businessConnectionID string) (*BusinessConnection, error) {
	request := GetBusinessConnectionRequest{
		BusinessConnectionID: businessConnectionID,
	}
	result := new(BusinessConnection)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetBusinessConnection, request); err != nil {
		return nil, err
	}
	return result, nil
}
