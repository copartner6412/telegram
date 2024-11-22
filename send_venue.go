package telegram

import (
	"net/http"
)

// SendVenueRequest represents the parameters for sending a information about a venue using the sendVenue method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Latitude
//   - Longitude
//   - Title
//   - Address
//
// See https://core.telegram.org/bots/api#sendvenue
type SendVenueRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Latitude of the venue.
	Latitude float64 `json:"latitude"`

	// (Required) Longitude of the venue.
	Longitude float64 `json:"longitude"`

	// (Required) Name of the venue.
	Title string `json:"title"`

	// (Required) Address of the venue.
	Address string `json:"address"`

	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Foursquare identifier of the venue.
	FoursquareID string `json:"foursquare_id,omitempty"`

	// (Optional) Foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// (Optional) Google Places identifier of the venue.
	GooglePlaceID string `json:"google_place_id,omitempty"`

	// (Optional) Google Places type of the venue. (See supported types.)
	GooglePlaceType string `json:"google_place_type,omitempty"`

	*SendOptions
}

// SendVenue sends information about a venue to a specified chat through the Telegram bot API.
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(request SendVenueRequest) (*Message, error) {
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendVenue, request); err != nil {
		return nil, err
	}
	return result, nil
}
