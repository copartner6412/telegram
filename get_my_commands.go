package telegram

import (
	"net/http"
)

// GetMyCommandsRequest represents the request payload for the getMyCommands method through the Telegram bot API.
//
// See "getMyCommands" https://core.telegram.org/bots/api#getmycommands
type GetMyCommandsRequest struct {
	// (Optional) An object, describing scope of users. Defaults to BotCommandScopeDefault.
	Scope BotCommandScope `json:"scope,omitempty"`

	// (Optional) A two-letter ISO 639-1 language code or an empty string.
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyCommands retrieves the current list of the bot's commands for the specified scope and user language through the Telegram bot API.
//
// Parameters:
//   - scope: An object, describing scope of users. If nil defaults to BotCommandScopeDefault.
//   - languageCode: A two-letter ISO 639-1 language code or an empty string.
//
// See "getMyCommands" https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(scope BotCommandScope, languageCode string) ([]BotCommand, error) {
	request := GetMyCommandsRequest{
		Scope:        scope,
		LanguageCode: languageCode,
	}
	result := new([]BotCommand)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetMyCommands, request); err != nil {
		return nil, err
	}
	return *result, nil
}
