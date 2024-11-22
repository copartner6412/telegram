package telegram

import (
	"net/http"
)

// SetMyCommandsRequest represents the request to change the list of the bot's commands using the setMyCommands method through the Telegram bot API.
//
// Required fields:
//   - Commands
//
// See "setMyCommands" https://core.telegram.org/bots/api#setmycommands
type SetMyCommandsRequest struct {
	// (Required) A list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []BotCommand `json:"commands"`

	// (Optional) An object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope BotCommandScope `json:"scope,omitempty"`

	// (Optional) A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands.
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyCommands updates the list of commands for the bot through the Telegram bot API.
//
// Parameters:
//   - commands: A list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
//   - scope: An object, describing scope of users for which the commands are relevant. If nil, defaults to BotCommandScopeDefault.
//   - languageCode: A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands.
//
// See https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(commands []BotCommand, scope BotCommandScope, languageCode string) error {
	request := SetMyCommandsRequest{
		Commands:     commands,
		Scope:        scope,
		LanguageCode: languageCode,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetMyCommands, request); err != nil {
		return err
	}
	return nil
}
