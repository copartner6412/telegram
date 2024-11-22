package telegram

import (
	"net/http"
)

// StopPoll is used to stop a poll which was sent by the bot using the stopPoll method through the Telegram bot API.
//
// Required fields:
//   - chatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - messageID: Identifier of the original message with the poll
//
// See "stopPoll" https://core.telegram.org/bots/api#stoppoll
type StopPollRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username
	// of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Identifier of the original message with the poll.
	MessageID int `json:"message_id"`

	*EditMessageOptions
}

// StopPoll stops a poll which was sent by the bot through the Telegram bot API.
//
// Required parameters:
//   - chatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - messageID: Identifier of the original message with the poll
//
// On success, the stopped Poll is returned.
//
// See "stopPoll" https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(chatID any, messageID int, options *EditMessageOptions) (*Poll, error) {
	request := StopPollRequest{
		ChatID:             chatID,
		MessageID:          messageID,
		EditMessageOptions: options,
	}
	result := new(Poll)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodStopPoll, request); err != nil {
		return nil, err
	}
	return result, nil
}
