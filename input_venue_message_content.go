package telegram

// InputVenueMessageContent represents the content of a venue message to be sent as the result of an inline query.
//
// Required fields:
//   - Latitude
//   - Longitude
//   - Title
//   - Address
//
// See "InputVenueMessageContent" https://core.telegram.org/bots/api#inputvenuemessagecontent
type InputVenueMessageContent struct {
	// (Required) Latitude of the venue in degrees.
	Latitude float64 `json:"latitude"`

	// (Required) Longitude of the venue in degrees.
	Longitude float64 `json:"longitude"`

	// (Required) Name of the venue.
	Title string `json:"title"`

	// (Required) Address of the venue.
	Address string `json:"address"`

	// (Optional) Foursquare identifier of the venue, if known.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// (Optional) Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// (Optional) Google Places identifier of the venue.
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// (Optional) Google Places type of the venue. (See supported types. https://developers.google.com/places/web-service/supported_types)
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

func (i InputVenueMessageContent) isInputMessageContent() bool {
	return true
}
