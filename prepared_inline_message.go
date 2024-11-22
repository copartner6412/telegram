package telegram

// PreparedInlineMessage describes an inline message to be sent by a user of a Mini App.
//
// See "PreparedInlineMessage" https://core.telegram.org/bots/api#preparedinlinemessage
type PreparedInlineMessage struct {
	// (Required) Unique identifier of the prepared message
	ID string `json:"id"`

	// (Required) Expiration date of the prepared message, in Unix time. Expired prepared messages can no longer be used
	ExpirationDate string `json:"expiration_date"`
}
