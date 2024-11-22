package telegram

// EncryptedCredentials contains data required to decrypt and authenticate EncryptedPassportElement.
//
// See the Telegram Passport Documentation https://core.telegram.org/passport#receiving-information for a complete description of the data decryption and authentication processes.
//
// See "EncryptedCredentials" https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	// (Required) Base64-encoded encrypted data containing the unique user's payload, data hashes,
	// and secrets required for EncryptedPassportElement decryption and authentication.
	Data string `json:"data"`

	// (Required) Base64-encoded data hash used for data authentication.
	Hash string `json:"hash"`

	// (Required) Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption.
	Secret string `json:"secret"`
}
