package telegram

// InputLocationMessageContent represents the content of a location message to be sent as the result of an inline query.
//
// Required fields:
//   - Latitude
//   - Longitude
//
// See "InputLocationMessageContent" https://core.telegram.org/bots/api#inputlocationmessagecontent
type InputLocationMessageContent struct {
	// (Required) Latitude of the location in degrees.
	Latitude float64 `json:"latitude"`

	// (Required) Longitude of the location in degrees.
	Longitude float64 `json:"longitude"`

	// (Optional) The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`

	// (Optional) Period in seconds during which the location can be updated, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod int `json:"live_period,omitempty"`

	// (Optional) For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// (Optional) For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

func (i InputLocationMessageContent) isInputMessageContent() bool {
	return true
}
