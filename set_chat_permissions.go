package telegram

import (
	"net/http"
)

// SetChatPermissionsRequest represent parameters to set default chat permissions for all members using setChatPermissions method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Permissions
//
// See https://core.telegram.org/bots/api#setchatpermissions
type SetChatPermissionsRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID any `json:"chat_id"`

	// (Required) An object for new default chat permissions
	Permissions ChatPermissions `json:"permissions"`

	// (Optional) Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages
	// and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents,
	// can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	// the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`
}

// SetChatPermissions sets default chat permissions for all members in a group or supergroup using the setChatPermissions through the Telegram bot API.
//
// The bot must be an administrator in the group or a supergroup for this to work
// and must have the can_restrict_members administrator rights.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
//   - permissions: An object for new default chat permissions
//   - useIndependentChatPermissions: Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
//
// See https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(chatId any, permissions ChatPermissions, useIndependentChatPermissions bool) error {
	request := SetChatPermissionsRequest{
		ChatID:                        chatId,
		Permissions:                   permissions,
		UseIndependentChatPermissions: useIndependentChatPermissions,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetChatPermissions, request); err != nil {
		return err
	}

	return nil
}
