package telegram

// PassportFile represents a file uploaded to Telegram Passport.
//
// Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
//
// See "PassportFile" https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	// (Required) Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// (Required) Unique identifier for this file, which is supposed to be the same over time
	// and for different bots. Cannot be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// (Required) File size in bytes.
	FileSize int `json:"file_size"`

	// (Required) Unix time when the file was uploaded.
	FileDate int `json:"file_date"`
}
