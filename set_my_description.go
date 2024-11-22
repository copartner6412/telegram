package telegram

import "net/http"

// SetMyDescriptionRequest represents parameters to change the bot's description, which is shown in the chat with the bot if the chat is empty. using the setMyDescription method through the Telegram bot API.
//
// See "setMyDescription" https://core.telegram.org/bots/api#setmydescription
type SetMyDescriptionRequest struct {
	// (Optional) New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	Description string `json:"description,omitempty"`

	// A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyDescription changes the bot's description, which is shown in the chat with the bot if the chat is empty through the Telegram bot API.
//
// Parameters:
//   - description: New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
//   - languageCode: A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
//
// Returns an error if the request was unsuccessful or if there was an issue processing the response.
//
// See "setMyDescription" https://core.telegram.org/bots/api#setmydescription
func (b *Bot) SetMyDescription(description string, languageCode string) error {
	request := SetMyDescriptionRequest{
		Description:  description,
		LanguageCode: languageCode,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetMyDescription, request); err != nil {
		return err
	}
	return nil
}
