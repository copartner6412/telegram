package telegram

// CopyTextButton represents an inline keyboard button that copies specified text to the clipboard.
//
// See "CopyTextButton" https://core.telegram.org/bots/api#copytextbutton
type CopyTextButton struct {
	// (Required) The text to be copied to the clipboard; 1-256 characters
	Text string `json:"text"`
}
