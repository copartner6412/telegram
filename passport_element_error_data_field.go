package telegram

// PassportElementErrorDataField represents an issue in one of the data fields provided by the user in Telegram Passport.
//
// The error is considered resolved when the value of the field in question changes.
//
// See "PassportElementErrorDataField" https://core.telegram.org/bots/api#passportelementerrordatafield
type PassportElementErrorDataField struct {
	// (Required) Error source, must be "data".
	Source string `json:"source"`

	// (Required) The section of the user's Telegram Passport which contains the error.
	//
	// Possible values are:
	//   - "personal_details"
	//   - "passport"
	//   - "driver_license"
	//   - "identity_card"
	//   - "internal_passport"
	//   - "address"
	Type string `json:"type"`

	// (Required) Name of the data field that contains the error.
	FieldName string `json:"field_name"`

	// (Required) Base64-encoded data hash.
	DataHash string `json:"data_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorDataField) getErrorType() string {
	return e.Type
}
