package telegram

// UsersShared contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
//
// See "UsersShared" https://core.telegram.org/bots/api#usersshared
type UsersShared struct {
	// (Required) Identifier of the request.
	RequestID int `json:"request_id"`

	// (Required) Information about users shared with the bot.
	Users []SharedUser `json:"users"`
}
