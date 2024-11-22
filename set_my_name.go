package telegram

import (
	"net/http"
)

// SetMyNameRequest represents the request payload for the setMyName method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#setmyname
type SetMyNameRequest struct {
	// (Optional) New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	Name string `json:"name,omitempty"`

	// (Optional) A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyName changes the bot's name.
//
// Parameters:
//   - name: New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
//   - languageCode: A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
//
// See https://core.telegram.org/bots/api#setmyname
func (b *Bot) SetMyName(name string, languageCode string) error {
	request := SetMyNameRequest{
		Name:         name,
		LanguageCode: languageCode,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetMyName, request); err != nil {
		return err
	}
	return nil
}
