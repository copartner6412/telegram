package telegram

import "net/http"

// GetMyDescriptionRequest represents parameters to get the current bot description for the giver user language using the getMyDescription method through the Telegram bot API.
//
// See "getMyDescription" https://core.telegram.org/bots/api#getmydescription
type GetMyDescriptionRequest struct {
	// (Optional) A two-letter ISO 639-1 language code or an empty string
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyDescription gets the current bot description for the given user language through the Telegram bot API.
//
// Parameters:
//   - languageCode: A two-letter ISO 639-1 language code or an empty string
//
// Returns BotDescription on success.
//
// See "getMyDescription" https://core.telegram.org/bots/api#getmydescription
func (b *Bot) GetMyDescription(LanguageCode string) (string, error) {
	request := GetMyDescriptionRequest{
		LanguageCode: LanguageCode,
	}
	result := new(BotDescription)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetMyDescription, request); err != nil {
		return "", nil
	}
	return result.Description, nil
}
