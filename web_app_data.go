package telegram

// WebAppData describes data sent from a Web App to the bot.
//
// See "WebAppData" https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	// (Required) The data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// (Required) Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}