package telegram

// PassportData describes Telegram Passport data shared with the bot by the user.
//
// See "PassportData" https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	// (Required) Array with information about documents and other Telegram Passport elements
	// that were shared with the bot.
	Data []EncryptedPassportElement `json:"data"`

	// (Required) Encrypted credentials required to decrypt the data.
	Credentials EncryptedCredentials `json:"credentials"`
}
