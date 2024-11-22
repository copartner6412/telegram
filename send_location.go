package telegram

import "net/http"

// SendLocationRequest represents the parameters for sending a location using the sendLocation method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Latitude
//   - Longitude
//
// See https://core.telegram.org/bots/api#sendlocation
type SendLocationRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	// If the chat is a channel, all Telegram Star proceeds from this location will be credited to the chat's balance.
	// Otherwise, they will be credited to the bot's balance.
	ChatID any `json:"chat_id"`

	// (Required) Latitude of the location.
	Latitude float64 `json:"latitude"`

	// (Required) Longitude of the location.
	Longitude float64 `json:"longitude"`

	*SendLocationOptions
}

// SendLocationOption represents the optional parameters for sending a location using the sendLocation method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#sendlocation
type SendLocationOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`

	// (Optional) Period in seconds during which the location will be updated (see Live Locations).
	// Should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod int `json:"live_period,omitempty"`

	// (Optional) For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// (Optional) For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters.
	// Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	*SendOptions
}

// SendLocation sends a location to a specified chat.
//
// Required parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername). If the chat is a channel, all Telegram Star proceeds from this location will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
//   - Latitude: Latitude of the location.
//   - Longitude: Longitude of the location.
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(toChatID any, Latitude, Longitude float64, options *SendLocationOptions) (*Message, error) {
	request := SendLocationRequest{
		ChatID:              toChatID,
		Latitude:            Latitude,
		Longitude:           Longitude,
		SendLocationOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendLocation, request); err != nil {
		return nil, err
	}
	return result, nil
}
