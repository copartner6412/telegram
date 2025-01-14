package telegram

// InlineQueryResultsButton represents a button to be shown above inline query results.
//
// You must use exactly one of the optional fields.
//
// Required fields:
//   - Text
//
// See "InlineQueryResultsButton" https://core.telegram.org/bots/api#inlinequeryresultsbutton
type InlineQueryResultsButton struct {
	// (Required) Label text on the button.
	Text string `json:"text"`

	// (Optional) Description of the Web App that will be launched when the user presses
	// the button. The Web App will be able to switch back to the inline mode using the
	// method switchInlineQuery inside the Web App.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// (Optional) Deep-linking parameter for the /start message sent to the bot when a user
	// presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
	StartParameter string `json:"start_parameter,omitempty"`
}
