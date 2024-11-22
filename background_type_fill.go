package telegram

import "encoding/json"

// BackgroundTypeFill represents the background automatically filled based on the selected colors.
//
// See "BackgroundTypeFill" https://core.telegram.org/bots/api#backgroundtypefill
type BackgroundTypeFill struct {
	// (Required) Type of the background, always “fill”.
	Type string `json:"type"`

	// (Required) The background fill.
	//
	// Its type can be one of:
	//   - BackgroundFillSolid
	//   - BackgroundFillGradient
	//   - BackgroundFillFreeformGradient
	Fill json.RawMessage `json:"fill"`

	// (Optional) Dimming of the background in dark themes, as a percentage; 0-100.
	DarkThemeDimming int `json:"dark_theme_dimming,omitempty"`
}
