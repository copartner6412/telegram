package telegram

// BotCommandScope represents the scope to which bot commands are applied.
//
// Currently, the following 7 scopes are supported:
//   - BotCommandScopeDefault
//   - BotCommandScopeAllPrivateChats
//   - BotCommandScopeAllGroupChats
//   - BotCommandScopeAllChatAdministrators
//   - BotCommandScopeChat
//   - BotCommandScopeChatAdministrators
//   - BotCommandScopeChatMember
//
// # Determining list of commands
//
// The following algorithm is used to determine the list of commands for a particular user viewing the bot menu. The first list of commands which is set is returned:
//
// Commands in the chat with the bot:
//   - botCommandScopeChat + language_code
//   - botCommandScopeChat
//   - botCommandScopeAllPrivateChats + language_code
//   - botCommandScopeAllPrivateChats
//   - botCommandScopeDefault + language_code
//   - botCommandScopeDefault
//
// Commands in group and supergroup chats:
//   - botCommandScopeChatMember + language_code
//   - botCommandScopeChatMember
//   - botCommandScopeChatAdministrators + language_code (administrators only)
//   - botCommandScopeChatAdministrators (administrators only)
//   - botCommandScopeChat + language_code
//   - botCommandScopeChat
//   - botCommandScopeAllChatAdministrators + language_code (administrators only)
//   - botCommandScopeAllChatAdministrators (administrators only)
//   - botCommandScopeAllGroupChats + language_code
//   - botCommandScopeAllGroupChats
//   - botCommandScopeDefault + language_code
//   - botCommandScopeDefault
//
// See "BotCommandScope" https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope interface {
	getScopeType() string
}
