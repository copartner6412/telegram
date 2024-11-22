package telegram

// PassportElementErrorTranslationFiles represents an issue with the translated version of a document.
//
// The error is considered resolved when a file with the document translation changes.
//
// See "PassportElementErrorTranslationFiles" https://core.telegram.org/bots/api#passportelementerrortranslationfiles
type PassportElementErrorTranslationFiles struct {
	// (Required) Error source, must be "translation_files".
	Source string `json:"source"`

	// (Required) Type of element of the user's Telegram Passport which has the issue.
	// Possible values are:
	//   - "passport"
	//   - "driver_license"
	//   - "identity_card"
	//   - "internal_passport"
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

func (e *PassportElementErrorTranslationFiles) getErrorType() string {
	return e.Type
}
