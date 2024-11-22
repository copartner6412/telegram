package telegram

import "encoding/json"

// ChatMemberUpdated represents changes in the status of a chat member.
//
// See "ChatMemberUpdated" https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	// (Required) Chat the user belongs to.
	Chat Chat `json:"chat"`

	// (Required) Performer of the action, which resulted in the change.
	From User `json:"from"`

	// (Required) Date the change was done in Unix time.
	Date int `json:"date"`

	// (Required) Previous information about the chat member.
	//
	// It type can be one of:
	//   - ChatMemberOwner
	//   - ChatMemberAdministrator
	//   - ChatMemberMember
	//   - ChatMemberRestricted
	//   - ChatMemberLeft
	//   - ChatMemberBanned
	OldChatMember json.RawMessage `json:"old_chat_member"`

	// (Required) New information about the chat member.
	//
	// It type can be one of:
	//   - ChatMemberOwner
	//   - ChatMemberAdministrator
	//   - ChatMemberMember
	//   - ChatMemberRestricted
	//   - ChatMemberLeft
	//   - ChatMemberBanned
	NewChatMember json.RawMessage `json:"new_chat_member"`

	// (Optional) Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`

	// (Optional) True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator.
	ViaJoinRequest bool `json:"via_join_request,omitempty"`

	// (Optional) True, if the user joined the chat via a chat folder invite link.
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}
