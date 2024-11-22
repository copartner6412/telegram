package telegram

import (
	"net/http"
)

// GetMyDefaultAdministratorRightsRequest represents parameters to get the current default administrator rights of the bot using the getMyDefaultAdministratorRights method through the Telegram bot API.
//
// See "getMyDefaultAdministratorRights" https://core.telegram.org/bots/api#getmydefaultadministratorrights
type GetMyDefaultAdministratorRightsRequest struct {
	// (Optional) Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
	ForChannels bool `json:"for_channels"`
}

// GetMyDefaultAdministratorRights retrieves the current default administrator rights of the bot.
//
// Parameters:
//   - ForChannels: Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
//
// Returns ChatAdministratorRights on success.
//
// See https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(forChannels bool) (*ChatAdministratorRights, error) {
	request := GetMyDefaultAdministratorRightsRequest{
		ForChannels: forChannels,
	}
	result := new(ChatAdministratorRights)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetMyDefaultAdministratorRights, request); err != nil {
		return nil, err
	}
	return result, nil
}
