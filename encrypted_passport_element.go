package telegram

// EncryptedPassportElement describes documents or other Telegram Passport elements shared with the bot by the user.
//
// See "EncryptedPassportElement" https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	// (Required) Element type, which can be one of "personal_details", "passport", "driver_license", "identity_card",
	// "internal_passport", "address", "utility_bill", "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration", "phone_number", "email".
	Type string `json:"type"`

	// (Optional) Base64-encoded encrypted Telegram Passport element data provided by the user. Available only for
	// types "personal_details", "passport", "driver_license", "identity_card", "internal_passport", and "address".
	// This data can be decrypted and verified using the accompanying EncryptedCredentials.
	Data string `json:"data,omitempty"`

	// (Optional) User's verified phone number, available only for the "phone_number" type.
	PhoneNumber string `json:"phone_number,omitempty"`

	// (Optional) User's verified email address, available only for the "email" type.
	Email string `json:"email,omitempty"`

	// (Optional) Array of encrypted files with documents provided by the user. Available only for types "utility_bill",
	// "bank_statement", "rental_agreement", "passport_registration", and "temporary_registration". These files can be
	// decrypted and verified using the accompanying EncryptedCredentials.
	Files []PassportFile `json:"files,omitempty"`

	// (Optional) Encrypted file with the front side of the document provided by the user. Available only for types
	// "passport", "driver_license", "identity_card", and "internal_passport". This file can be decrypted and verified
	// using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// (Optional) Encrypted file with the reverse side of the document provided by the user. Available only for types
	// "driver_license" and "identity_card". This file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// (Optional) Encrypted file with the selfie of the user holding a document, provided by the user. Available if
	// requested for types "passport", "driver_license", "identity_card", and "internal_passport". This file can be
	// decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`

	// (Optional) Array of encrypted files with translated versions of documents provided by the user. Available if
	// requested for types "passport", "driver_license", "identity_card", "internal_passport", "utility_bill",
	// "bank_statement", "rental_agreement", "passport_registration", and "temporary_registration". These files can
	// be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation,omitempty"`

	// (Required) Base64-encoded element hash for using in PassportElementErrorUnspecified.
	Hash string `json:"hash"`
}
