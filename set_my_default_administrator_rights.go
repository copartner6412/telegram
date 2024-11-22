package telegram

import (
	"net/http"
)

// SetMyDefaultAdministratorRightsRequest represents the parameters for the setMyDefaultAdministratorRights method, which changes the default administrator rights requested by the bot when added as an administrator to groups or channels using the setMyDefaultAdministratorRights method through the Telegram bot API.
//
// See "setMyDefaultAdministratorRights" https://core.telegram.org/bots/api#setmydefaultadministratorrights
type SetMyDefaultAdministratorRightsRequest struct {
	// (Optional) An object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	Rights *ChatAdministratorRights `json:"rights,omitempty"`

	// (Optional) Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels bool `json:"for_channels,omitempty"`
}

// SetMyDefaultAdministratorRights changes the default administrator rights requested by the bot when it is added as an administrator to groups or channels through the Telegram bot API.
//
// These rights will be suggested to users, but they can modify the list before adding the bot.
//
// Parameters:
//   - rights: An object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
//   - forChannels: Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
//
// See "setMyDefaultAdministratorRights" https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(rights *ChatAdministratorRights, forChannels bool) error {
	request := SetMyDefaultAdministratorRightsRequest{
		Rights:      rights,
		ForChannels: forChannels,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetMyDefaultAdministratorRights, request); err != nil {
		return err
	}
	return nil
}
