package telegram

// MenuButtonDefault describes that no specific value for the menu button was set.
//
// See "MenuButtonDefault" https://core.telegram.org/bots/api#menubuttondefault
type MenuButtonDefault struct {
	// (Required) Type of the button, must be default.
	Type string `json:"type"`
}

func (b MenuButtonDefault) getButtonType() string {
	return b.Type
}
