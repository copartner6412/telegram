package telegram

import "net/http"

// SetMyDescriptionRequest represents parameters to change the bot's short description using the setMyShortDescription method through the Telegram bot API.
//
// See "setMyShortDescription" https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescriptionRequest struct {
	// (Optional) New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	ShortDescription string `json:"short_description,omitempty"`

	// (Optional) A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyShortDescription changes the bot's short description through the Telegram bot API.
//
// Short description is shown on the bot's profile page and is sent together with the link when users share the bot.
//
// Parameters:
//   - shortDescription: New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
//   - languageCode: A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
//
// See "setMyShortDescription" https://core.telegram.org/bots/api#setmyshortdescription
func (b *Bot) SetMyShortDescription(shortDescription, languageCode string) error {
	request := SetMyShortDescriptionRequest{
		ShortDescription: shortDescription,
		LanguageCode:     languageCode,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetMyShortDescription, request); err != nil {
		return err
	}
	return nil
}
