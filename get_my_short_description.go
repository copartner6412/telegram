package telegram

import "net/http"

// GetMyShortDescriptionRequest represents parameters to get the current bot short description using the getMyShortDescription method through the Telegram bot API.
//
// See "getMyShortDescription" https://core.telegram.org/bots/api#getmyshortdescription
type GetMyShortDescriptionRequest struct {
	// (Optional) A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyShortDescription gets the current bot short description through the Telegram bot API.
//
// Parameters:
//   - languageCode: A two-letter ISO 639-1 language code or an empty string
//
// Returns BotShortDescription on success.
//
// See "getMyShortDescription" https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) GetMyShortDescription(languageCode string) (string, error) {
	request := GetMyShortDescriptionRequest{
		LanguageCode: languageCode,
	}
	result := new(BotShortDescription)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetMyShortDescription, request); err != nil {
		return "", err
	}
	return result.ShortDescription, nil
}
