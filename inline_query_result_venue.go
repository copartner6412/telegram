package telegram

// InlineQueryResultVenue represents a venue.
//
// By default, the venue will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the venue.
//
// Required fields:
//   - Type
//   - ID
//   - Latitude
//   - Longitude
//   - Title
//   - Address
//
// See "InlineQueryResultVenue" https://core.telegram.org/bots/api#inlinequeryresultvenue
type InlineQueryResultVenue struct {
	// (Required) Type of the result, must be "venue".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Latitude of the venue location in degrees.
	Latitude float64 `json:"latitude"`

	// (Required) Longitude of the venue location in degrees.
	Longitude float64 `json:"longitude"`

	// (Required) Title of the venue.
	Title string `json:"title"`

	// (Required) Address of the venue.
	Address string `json:"address"`

	// (Optional) Foursquare identifier of the venue if known.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// (Optional) Foursquare type of the venue, if known.
	// For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.
	FoursquareType string `json:"foursquare_type,omitempty"`

	// (Optional) Google Places identifier of the venue.
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// (Optional) Google Places type of the venue. See supported types.
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the venue.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// (Optional) URL of the thumbnail for the result.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) Thumbnail width.
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// (Optional) Thumbnail height.
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

func (r InlineQueryResultVenue) getQueryType() string {
	return r.Type
}
