package telegram

// BusinessLocation contains information about the location of a Telegram Business account.
//
// See "BusinessLocation" https://core.telegram.org/bots/api#businesslocation
type BusinessLocation struct {
	// (Required) Address of the business.
	Address string `json:"address"`

	// (Optional) Location of the business.
	Location *Location `json:"location,omitempty"`
}
