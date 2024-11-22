package telegram

// InlineQueryResultLocation represents a location on a map.
//
// By default, the location will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the location.
//
// Required fields:
//   - Type
//   - ID
//   - Latitude
//   - Longitude
//
// See "InlineQueryResultLocation" https://core.telegram.org/bots/api#inlinequeryresultlocation
type InlineQueryResultLocation struct {
	// (Required) Type of the result, must be "location".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Location latitude in degrees.
	Latitude float64 `json:"latitude"`

	// (Required) Location longitude in degrees.
	Longitude float64 `json:"longitude"`

	// (Required) Location title.
	Title string `json:"title"`

	// (Optional) The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`

	// (Optional) Period in seconds during which the location can be updated.
	// Should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod int `json:"live_period,omitempty"`

	// (Optional) For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// (Optional) For live locations, a maximum distance for proximity alerts about approaching
	// another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the location.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// (Optional) URL of the thumbnail for the result.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) Thumbnail width.
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// (Optional) Thumbnail height.
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

func (r InlineQueryResultLocation) getQueryType() string {
	return r.Type
}
