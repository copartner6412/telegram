package telegram

// KeyboardButtonRequestUsers defines the criteria used to request suitable users.
// Information about the selected users will be shared with the bot when the corresponding button is pressed. More about requesting users » https://core.telegram.org/bots/features#chat-and-user-selection
//
// See "KeyboardButtonRequestUsers" https://core.telegram.org/bots/api#keyboardbuttonrequestusers
type KeyboardButtonRequestUsers struct {
	// (Required) Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message.
	RequestID int `json:"request_id"`

	// (Optional) Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsBot bool `json:"user_is_bot,omitempty"`

	// (Optional) Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	UserIsPremium bool `json:"user_is_premium,omitempty"`

	// (Optional) The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity int `json:"max_quantity,omitempty"`

	// (Optional) Pass True to request the users' first and last names.
	RequestName bool `json:"request_name,omitempty"`

	// (Optional) Pass True to request the users' usernames.
	RequestUsername bool `json:"request_username,omitempty"`

	// (Optional) Pass True to request the users' photos.
	RequestPhoto bool `json:"request_photo,omitempty"`
}
