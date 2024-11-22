package telegram

// UserChatBoosts represents a list of boosts added to a chat by a user.
//
// See "UserChatBoosts" https://core.telegram.org/bots/api#userchatboosts
type UserChatBoosts struct {
	// (Required) The list of boosts added to the chat by the user.
	Boosts []ChatBoost `json:"boosts"`
}
