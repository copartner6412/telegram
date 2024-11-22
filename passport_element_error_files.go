package telegram

// PassportElementErrorFiles represents an issue with a list of document scans in the user's Telegram Passport.
//
// The error is considered resolved when the list of files containing the scans is updated.
//
// See "PassportElementErrorFiles" https://core.telegram.org/bots/api#passportelementerrorfiles
type PassportElementErrorFiles struct {
	// (Required) Error source, must be "files".
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

	// (Required) List of base64-encoded file hashes.
	FileHashes []string `json:"file_hashes"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorFiles) getErrorType() string {
	return e.Type
}
