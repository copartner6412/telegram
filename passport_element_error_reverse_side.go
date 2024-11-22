package telegram

// PassportElementErrorReverseSide represents an issue with the reverse side of a document.
//
// The error is considered resolved when the file with the reverse side of the document is updated.
//
// See "PassportElementErrorReverseSide" https://core.telegram.org/bots/api#passportelementerrorreverseside
type PassportElementErrorReverseSide struct {
	// (Required) Error source, must be "reverse_side".
	Source string `json:"source"`

	// (Required) The section of the user's Telegram Passport containing the issue.
	//
	// Possible values are:
	//   - "driver_license"
	//   - "identity_card"
	Type string `json:"type"`

	// (Required) Base64-encoded hash of the file with the reverse side of the document.
	FileHash string `json:"file_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorReverseSide) getErrorType() string {
	return e.Type
}
