package telegram

import (
	"net/http"
)

// EditMessageReplyMarkup represents a request to edit the reply markup of messages using the editMessageReplyMarkup method through the Telegram bot API.
//
// Required fields:
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See https://core.telegram.org/bots/api#editmessagereplymarkup
type EditMessageReplyMarkup struct {
	// (Optional) Required if inline_message_id is not specified. Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the message to edit.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// (Optional) Unique identifier of the business connection on behalf of which the message to be edited was sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) An object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageReplyMarkup edits only the reply markup of messages using the editMessageReplyMarkup through the Telegram bot API.
//
// Parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of the message to edit.
//   - replyMarkup: An object for a new inline keyboard.
//   - businessConnectionID: Unique identifier of the business connection on behalf of which the message to be edited was sent. Pass empty string, if not specified.
//
// On success, the edited Message is returned.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageReplyMarkup" https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(chatID any, messageID int, replyMarkup InlineKeyboardMarkup, businessConnectionID string) (*Message, error) {
	request := EditMessageReplyMarkup{
		ChatID:               chatID,
		MessageID:            messageID,
		BusinessConnectionID: businessConnectionID,
		ReplyMarkup:          &replyMarkup,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodEditMessageReplyMarkup, request); err != nil {
		return nil, err
	}
	return result, nil
}

// EditInlineMessageReplyMarkup edits only the reply markup of inline messages using the editMessageReplyMarkup through the Telegram bot API.
//
// Parameters:
//   - inlineMessageID: Identifier of the inline message.
//   - replyMarkup: An object for a new inline keyboard.
//   - businessConnectionID: Unique identifier of the business connection on behalf of which the message to be edited was sent. Pass empty string, if not specified.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageReplyMarkup" https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditInlineMessageReplyMarkup(inlineMessageID string, replyMarkup InlineKeyboardMarkup, businessConnectionID string) error {
	request := EditMessageReplyMarkup{
		InlineMessageID:      inlineMessageID,
		BusinessConnectionID: businessConnectionID,
		ReplyMarkup:          &replyMarkup,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditMessageReplyMarkup, request); err != nil {
		return err
	}
	return nil
}
