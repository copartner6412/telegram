package telegram

import "net/http"

// DeleteMessagesRequest represent parameters to delete multiple messages simultaneously using the deleteMessages method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageID
//
// See "deleteMessages" https://core.telegram.org/bots/api#deletemessages
type DeleteMessagesRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) A list of 1-100 identifiers of messages to delete.
	// See deleteMessage for limitations on which messages can be deleted. https://core.telegram.org/bots/api#deletemessage
	MessageIDs []int `json:"message_ids"`
}

// DeleteMessages deletes multiple messages simultaneously from a chat through the Telegram bot API.
//
// If some of the specified messages can't be found, they are skipped.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - messageIDs: A slice of 1-100 identifiers of messages to delete.
//
// See "deleteMessages" https://core.telegram.org/bots/api#deletemessages
func (b *Bot) DeleteMessages(chatID any, messageIDs []int) error {
	request := DeleteMessagesRequest{
		ChatID:     chatID,
		MessageIDs: messageIDs,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteMessages, request); err != nil {
		return err
	}
	return nil
}
