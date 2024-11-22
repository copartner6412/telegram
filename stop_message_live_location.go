package telegram

import (
	"net/http"
)

// StopMessageLiveLocationRequest represents parameters used to stop updating a live location message before the live_period expires using the stopMessageLiveLocation method through the Telegram bot API.
//
// Required fields:
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See "stopMessageLiveLocation" https://core.telegram.org/bots/api#stopmessagelivelocation
type StopMessageLiveLocationRequest struct {
	// (Optional) Required if inline_message_id is not specified. Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id,omitempty"`

	// (Optional) Required if inline_message_id is not specified. Identifier of the message with live location to stop.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	*EditMessageOptions
}

// StopMessageLiveLocation stops updating a live location message before the live_period expires through Telegram bot API.
//
// Required parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of the message with live location to stop.
//
// On success, the edited Message is returned.
//
// See "stopMessageLiveLocation" https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(chatID any, messageID int, options *EditMessageOptions) (*Message, error) {
	request := StopMessageLiveLocationRequest{
		ChatID:             chatID,
		MessageID:          messageID,
		EditMessageOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodStopMessageLiveLocation, request); err != nil {
		return nil, err
	}
	return result, nil
}

// StopInlineMessageLiveLocation stops updating a live location inline message before the live_period expires through Telegram bot API.
//
// Required parameters:
//   - inlineMessageID: Identifier of the inline message.
//
// See "stopMessageLiveLocation" https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopInlineMessageLiveLocation(inlineMessageID string, options *EditMessageOptions) error {
	request := StopMessageLiveLocationRequest{
		InlineMessageID:    inlineMessageID,
		EditMessageOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodStopMessageLiveLocation, request); err != nil {
		return err
	}
	return nil
}
