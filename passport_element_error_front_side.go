package telegram

// PassportElementErrorFrontSide represents an issue with the front side of a document.
//
// The error is considered resolved when the file containing the front side of the document is updated.
//
// See "PassportElementErrorFrontSide" https://core.telegram.org/bots/api#passportelementerrorfrontside
type PassportElementErrorFrontSide struct {
	// (Required) Error source, must be "front_side".
	Source string `json:"source"`

	// (Required) The section of the user's Telegram Passport containing the issue.
	//
	// Possible values are:
	//   - "passport"
	//   - "driver_license"
	//   - "identity_card"
	//   - "internal_passport"
	Type string `json:"type"`

	// (Required) Base64-encoded hash of the file with the front side of the document.
	FileHash string `json:"file_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorFrontSide) getErrorType() string {
	return e.Type
}
