package telegram

import (
	"net/http"
)

// GetMyNameRequest represents the parameters to get the current bot name for the given user language.
//
// See "getMyName" https://core.telegram.org/bots/api#getmyname
type GetMyNameRequest struct {
	// (Optional) A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyName retrieves the current bot name for the specified user language.
//
// Parameters:
// - languageCode: A two-letter ISO 639-1 language code or an empty string.
//
// Returns BotName on success.
//
// See "getMyName" https://core.telegram.org/bots/api#getmyname
func (b *Bot) GetMyName(languageCode string) (string, error) {
	request := GetMyNameRequest{
		LanguageCode: languageCode,
	}
	result := new(BotName)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetMyName, request); err != nil {
		return "", err
	}
	return result.Name, nil
}
