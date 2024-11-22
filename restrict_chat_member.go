package telegram

import "net/http"

// RestrictChatMemberRequest represents a request to restrict a user in a supergroup using restrictChatMember method through the Telegram bot API.
//
// The bot must be an administrator in the supergroup and have appropriate administrator rights.
// Pass True for all permissions to lift restrictions from a user.
//
// Required fields:
//   - ChatID
//   - UserID
//   - Permissions
//
// See https://core.telegram.org/bots/api#restrictchatmember
type RestrictChatMemberRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`

	// (Required) An object for new user permissions.
	Permissions ChatPermissions `json:"permissions"`

	// (Optional) Pass True if chat permissions are set independently. Otherwise, the
	// can_send_other_messages and can_add_web_page_previews permissions will imply the
	// can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos,
	// can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission
	// will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`

	// (Optional) Date when restrictions will be lifted for the user; Unix time.
	// If the user is restricted for more than 366 days or less than 30 seconds from the
	// current time, they are considered to be restricted forever.
	UntilDate int `json:"until_date,omitempty"`
}

// RestrictChatMember restricts a user in a supergroup through the Telegram bot API.
//
// The bot must be an administrator in the supergroup and have appropriate administrator rights for this to work.
// Passing True for all permissions will lift restrictions from a user.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
//   - userID: Unique identifier of the target user.
//   - permissions: An object for new user permissions.
//   - useIndependentChatPermissions: Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
//   - untilDate: Date when restrictions will be lifted for the user; Unix time. If the user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever. Simply zero if you want to if you want to restrict forever.
//
// See https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(chatID any, userID int, permissions ChatPermissions, useIndependentChatPermissions bool, untilDate int) error {
	request := RestrictChatMemberRequest{
		ChatID:                        chatID,
		UserID:                        userID,
		Permissions:                   permissions,
		UseIndependentChatPermissions: useIndependentChatPermissions,
		UntilDate:                     untilDate,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodRestrictChatMember, request); err != nil {
		return err
	}
	return nil
}
