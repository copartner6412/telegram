package telegram

import (
	"net/http"
)

// EditMessageLiveLocationRequest represents the parameters for editing live location messages using the editMessageLiveLocation method through the Telegram bot API.
//
// Required parameters:
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//   - Latitude
//   - Longitude
//
// See "editMessageLiveLocation" https://core.telegram.org/bots/api#editmessagelivelocation
type EditMessageLiveLocationRequest struct {
	// (Optional) Required if inline_message_id is not specified. Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id,omitempty"`

	// (Optional) Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// (Required) Latitude of new location.
	Latitude float64 `json:"latitude"`

	// (Required) Longitude of new location.
	Longitude float64 `json:"longitude"`

	*EditMessageLiveLocationOptions
}

// EditMessageLiveLocationOptions represents the optional parameters for editing live location messages using the editMessageLiveLocation method through the Telegram bot API.
//
// See "editMessageLiveLocation" https://core.telegram.org/bots/api#editmessagelivelocation
type EditMessageLiveLocationOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message to be edited was sent.
	BusinessConnectionID string `json:"business_connection_id"`

	// (Optional) New period in seconds during which the location can be updated, starting
	// from the message send date. If 0x7FFFFFFF is specified, then the location can be
	// updated forever. Otherwise, the new value must not exceed the current live_period
	// by more than a day, and the live location expiration date must remain within the
	// next 90 days. If not specified, then live_period remains unchanged.
	LivePeriod int `json:"live_period,omitempty"`

	// (Optional) The radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`

	// (Optional) Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`

	// (Optional) The maximum distance for proximity alerts about approaching another chat member,
	// in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// (Optional) An object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageLiveLocation edits live location messages through the Telegram bot API.
//
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of the message to edit.
//   - Latitude: Latitude of new location.
//   - Longitude: Longitude of new location.
//
// On success, the edited Message is returned.
//
// See "editMessageLiveLocation" https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(chatID any, messageID int, latitude float64, longitude float64, options *EditMessageLiveLocationOptions) (*Message, error) {
	request := EditMessageLiveLocationRequest{
		ChatID:                         chatID,
		MessageID:                      messageID,
		Latitude:                       latitude,
		Longitude:                      longitude,
		EditMessageLiveLocationOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodEditMessageLiveLocation, request); err != nil {
		return nil, err
	}
	return result, nil
}

// EditInlineMessageLiveLocation edits live location inline messages through the Telegram bot API.
//
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation.
//
// Required parameters:
//   - inlineMessageID: Identifier of the inline message.
//   - Latitude: Latitude of new location.
//   - Longitude: Longitude of new location.
//
// See "editMessageLiveLocation" https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditInlineMessageLiveLocation(inlineMessageID string, latitude, longitude float64, options *EditMessageLiveLocationOptions) error {
	request := EditMessageLiveLocationRequest{
		InlineMessageID:                inlineMessageID,
		Latitude:                       latitude,
		Longitude:                      longitude,
		EditMessageLiveLocationOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditMessageLiveLocation, request); err != nil {
		return err
	}
	return nil
}
