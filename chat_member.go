package telegram

// ChatMember contains information about one member of a chat.
//
// Currently, the following 6 types of chat members are supported:
//   - ChatMemberOwner
//   - ChatMemberAdministrator
//   - ChatMemberMember
//   - ChatMemberRestricted
//   - ChatMemberLeft
//   - ChatMemberBanned
//
// See "ChatMember" https://core.telegram.org/bots/api#chatmember
type ChatMember interface {
	getMemberStatus() string
}
