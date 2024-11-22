package telegram

// InlineKeyboardButton represents one button of an inline keyboard. Exactly one of the optional fields must be used to specify type of the button.
//
// See "InlineKeyboardButton" https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	// (Required) Label text on the button
	Text string `json:"text"`

	// (Optional) HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	URL string `json:"url,omitempty"`

	// (Optional) Data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// (Optional) Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// (Optional) An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	LoginURL *LoginUrl `json:"login_url,omitempty"`

	// (Optional) If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// (Optional) If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// (Optional) If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field. Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	// (Optional) Description of the button that copies the specified text to the clipboard.
	CopyText *CopyTextButton `json:"copy_text,omitempty"`

	// (Optional) Description of the game that will be launched when the user presses the button.
	//
	// # NOTE
	//
	// This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// (Optional) Specify True, to send a Pay button. Substrings “” and “XTR” in the buttons's text will be replaced with a Telegram Star icon. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
	Pay bool `json:"pay,omitempty"`
}
