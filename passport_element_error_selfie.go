package telegram

// PassportElementErrorSelfie represents an issue with the selfie associated with a document.
//
// The error is considered resolved when the file containing the selfie is updated.
//
// See "PassportElementErrorSelfie" https://core.telegram.org/bots/api#passportelementerrorselfie
type PassportElementErrorSelfie struct {
	// (Required) Error source, must be "selfie".
	Source string `json:"source"`

	// (Required) The section of the user's Telegram Passport that contains the issue.
	//
	// Possible values are:
	//   - "passport"
	//   - "driver_license"
	//   - "identity_card"
	//   - "internal_passport"
	Type string `json:"type"`

	// (Required) Base64-encoded hash of the file containing the selfie.
	FileHash string `json:"file_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorSelfie) getErrorType() string {
	return e.Type
}
