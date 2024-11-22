package telegram

import "encoding/json"

// BackgroundTypePattern represents the background as a PNG or TGV pattern to be combined with the background fill chosen by the user.
//
// See "BackgroundTypePattern" https://core.telegram.org/bots/api#backgroundtypepattern
type BackgroundTypePattern struct {
	// (Required) Type of the background, always “pattern”.
	Type string `json:"type"`

	// (Required) Document with the pattern.
	Document Document `json:"document"`

	// (Required) The background fill that is combined with the pattern.
	//
	// Its type can be one of:
	//   - BackgroundFillSolid
	//   - BackgroundFillGradient
	//   - BackgroundFillFreeformGradient
	Fill json.RawMessage `json:"fill"`

	// (Required) Intensity of the pattern when it is shown above the filled background; 0-100.
	Intensity int `json:"intensity"`

	// (Optional) True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only.
	IsInverted bool `json:"is_inverted,omitempty"`

	// (Optional) True, if the background moves slightly when the device is tilted.
	IsMoving bool `json:"is_moving,omitempty"`
}
