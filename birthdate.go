package telegram

// Birthdate describes the birthdate of a user.
//
// See "Birthdate" https://core.telegram.org/bots/api#birthdate
type Birthdate struct {
	// (Required) Day of the user's birth; 1-31.
	Day int `json:"day"`

	// (Required) Month of the user's birth; 1-12.
	Month int `json:"month"`

	// (Optional) Year of the user's birth.
	Year int `json:"year,omitempty"`
}
