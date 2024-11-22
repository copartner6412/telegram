package telegram

// PassportElementErrorTranslationFile represents an issue with one of the files that constitute the translation of a document.
//
// The error is considered resolved when the file changes.
//
// See "PassportElementErrorTranslationFile" https://core.telegram.org/bots/api#passportelementerrortranslationfile
type PassportElementErrorTranslationFile struct {
	// (Required) Error source, must be "translation_file".
	Source string `json:"source"`

	// (Required) Type of element of the user's Telegram Passport which has the issue.
	// Possible values are:
	//  - "passport"
	//  - "driver_license"
	//  - "identity_card"
	//  - "internal_passport"
	//  - "utility_bill"
	//  - "bank_statement"
	//  - "rental_agreement"
	//  - "passport_registration"
	//  - "temporary_registration"
	Type string `json:"type"`

	// (Required) Base64-encoded file hash.
	FileHash string `json:"file_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorTranslationFile) getErrorType() string {
	return e.Type
}
