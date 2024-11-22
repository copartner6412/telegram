package telegram

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
//
// See "InlineKeyboardMarkup" https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	// (Required) Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
