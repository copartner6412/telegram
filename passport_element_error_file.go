package telegram

// PassportElementErrorFile represents an issue with a document scan in the user's Telegram Passport.
//
// The error is considered resolved when the file containing the document scan is updated.
//
// See "PassportElementErrorFile" https://core.telegram.org/bots/api#passportelementerrorfile
type PassportElementErrorFile struct {
	// (Required) Error source, must be "file".
	Source string `json:"source"`

	// (Required) The section of the user's Telegram Passport that contains the issue.
	//
	// Possible values are:
	//   - "utility_bill"
	//   - "bank_statement"
	//   - "rental_agreement"
	//   - "passport_registration"
	//   - "temporary_registration"
	Type string `json:"type"`

	// (Required) Base64-encoded hash of the file with the document scan.
	FileHash string `json:"file_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorFile) getErrorType() string {
	return e.Type
}
