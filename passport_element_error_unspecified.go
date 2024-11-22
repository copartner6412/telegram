package telegram

// PassportElementErrorUnspecified represents an issue in an unspecified place.
//
// The error is considered resolved when new data is added.
//
// See "PassportElementErrorUnspecified" https://core.telegram.org/bots/api#passportelementerrorunspecified
type PassportElementErrorUnspecified struct {
	// (Required) Error source, must be "unspecified".
	Source string `json:"source"`

	// (Required) Type of element of the user's Telegram Passport which has the issue.
	Type string `json:"type"`

	// (Required) Base64-encoded element hash.
	ElementHash string `json:"element_hash"`

	// (Required) Error message describing the issue.
	Message string `json:"message"`
}

func (e *PassportElementErrorUnspecified) getErrorType() string {
	return e.Type
}
