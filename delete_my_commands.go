package telegram

import (
	"net/http"
)

// DeleteMyCommandsRequest represents the request structure for the deleteMyCommands method through the Telegram bot API.
//
// See "deleteMyCommands" https://core.telegram.org/bots/api#deletemycommands
type DeleteMyCommandsRequest struct {
	// (Optional) An object, describing the scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope BotCommandScope `json:"scope,omitempty"`

	// (Optional) A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands.
	LanguageCode string `json:"language_code,omitempty"`
}

// DeleteMyCommands deletes the list of the bot's commands for the given scope and user language through the Telegram bot API.
//
// After deletion, higher level commands will be shown to affected users.
//
// Parameters:
//   - scope: An object, describing the scope of users for which the commands are relevant. If nil, defaults to BotCommandScopeDefault.
//   - languageCode: A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands.
//
// See "deleteMyCommands" https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(scope BotCommandScope, languageCode string) error {
	request := DeleteMyCommandsRequest{
		Scope:        scope,
		LanguageCode: languageCode,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteMyCommands, request); err != nil {
		return err
	}
	return nil
}
