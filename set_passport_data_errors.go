package telegram

import (
	"net/http"
)

// SetPassportDataErrorsRequest represents parameters used to inform a user that some of the Telegram Passport elements they provided contains errors using the setPassportDataErrors method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - Errors
//
// See "setPassportDataErrors" https://core.telegram.org/bots/api#setpassportdataerrors
type SetPassportDataErrorsRequest struct {
	// (Required) User identifier
	UserID int `json:"user_id"`

	// (Required) An array describing the errors
	Errors []PassportElementError `json:"errors"`
}

// SetPassportDataErrors informs a user that some of the Telegram Passport elements they provided contain errors through the Telegram bot API.
//
// The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change).
//
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
//
// Parameters:
//   - userID: User identifier
//   - errors: An array describing the errors
//
// See "setPassportDataErrors" https://core.telegram.org/bots/api#setpassportdataerrors
func (b *Bot) SetPassportDataErrors(userID int, errors []PassportElementError) error {
	request := SetPassportDataErrorsRequest{
		UserID: userID,
		Errors: errors,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetPassportDataErrors, request); err != nil {
		return err
	}
	return nil
}
