package telegram

// MenuButtonWebApp represents a menu button, which launches a Web App.
//
// See "MenuButtonWebApp" https://core.telegram.org/bots/api#menubuttonwebapp
type MenuButtonWebApp struct {
	// (Required) Type of the button, must be web_app.
	Type string `json:"type"`
	// (Required) Text on the button.
	Text string `json:"text"`
	// (Required) Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Alternatively, a t.me link to a Web App of the bot can be specified in the object instead of the Web App's URL, in which case the Web App will be opened as if the user pressed the link.
	WebApp WebAppInfo `json:"web_app"`
}

func (b MenuButtonWebApp) getButtonType() string {
	return b.Type
}
