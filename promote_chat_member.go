package telegram

import (
	"net/http"
)

// PromoteChatMemberRequest represents a request to promote or demote a user in a supergroup or a channel using the promoteChatMember method through the Telegram bot API.
//
// The bot must be an administrator in the chat with the appropriate administrator rights for this to work.
// Pass False for all boolean parameters to demote a user.
//
// Required fields:
//   - ChatID
//   - UserID
//
// See https://core.telegram.org/bots/api#promotechatmember
type PromoteChatMemberRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`

	*PromoteChatMemberOptions
}

// PromoteChatMemberOptions represents a optional parameters for a request to promote or demote a user in a supergroup or a channel using the promoteChatMember method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#promotechatmember
type PromoteChatMemberOptions struct {
	// (Optional) Pass True if the administrator's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// (Optional) Pass True if the administrator can access the chat event log, get boost list,
	// see hidden supergroup and channel members, report spam messages, and ignore slow mode.
	// Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat,omitempty"`

	// (Optional) Pass True if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// (Optional) Pass True if the administrator can manage video chats.
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`

	// (Optional) Pass True if the administrator can restrict, ban or unban chat members,
	// or access supergroup statistics.
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// (Optional) Pass True if the administrator can add new administrators with a subset
	// of their own privileges or demote administrators that they have promoted, directly
	// or indirectly (promoted by administrators that were appointed by him).
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// (Optional) Pass True if the administrator can change chat title, photo, and other settings.
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// (Optional) Pass True if the administrator can invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// (Optional) Pass True if the administrator can post stories to the chat.
	CanPostStories bool `json:"can_post_stories,omitempty"`

	// (Optional) Pass True if the administrator can edit stories posted by other users,
	// post stories to the chat page, pin chat stories, and access the chat's story archive.
	CanEditStories bool `json:"can_edit_stories,omitempty"`

	// (Optional) Pass True if the administrator can delete stories posted by other users.
	CanDeleteStories bool `json:"can_delete_stories,omitempty"`

	// (Optional) Pass True if the administrator can post messages in the channel or
	// access channel statistics; for channels only.
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// (Optional) Pass True if the administrator can edit messages of other users
	// and can pin messages; for channels only.
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// (Optional) Pass True if the administrator can pin messages; for supergroups only.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// (Optional) Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only.
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}

// PromoteChatMember promotes or demotes a user in a supergroup or a channel through the Telegram bot API.
//
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Pass False for all boolean parameters to demote a user.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - userID: Unique identifier of the target user.
//
// See https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(chatID any, userID int, options *PromoteChatMemberOptions) error {
	request := PromoteChatMemberRequest{
		ChatID:                   chatID,
		UserID:                   userID,
		PromoteChatMemberOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodPromoteChatMember, request); err != nil {
		return err
	}
	return nil
}
