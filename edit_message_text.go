package telegram

import (
	"net/http"
)

// EditMessageTextRequest represents parameters used to edit text and game messages using the editMessageText method through the Telegram bot API.
//
// Required fields:
//   - Text
//   - ParseMode
//   - Entities
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See "editMessageText" https://core.telegram.org/bots/api#editmessagetext
type EditMessageTextRequest struct {
	// (Required) New text of the message, 1-4096 characters after entities parsing.
	Text

	// (Optional) Integer or string. Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id,omitempty"`

	// (Optional) Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	*EditMessageTextOptions
}

// EditMessageTextOptions represents optional parameters used to edit text and game messages using the editMessageText method through the Telegram bot API.
//
// See "editMessageText" https://core.telegram.org/bots/api#editmessagetext
type EditMessageTextOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message to be edited was sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Link preview generation options for the message.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// (Optional) An object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageText edits text and game messages through the Telegram bot API.
//
// Required parameters:
//   - text: New text of the message, 1-4096 characters after entities parsing.
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of the message to edit.
//
// On success, the edited Message is returned.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageText" https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(chatID any, messageID int, text Text, options *EditMessageTextOptions) (*Message, error) {
	request := EditMessageTextRequest{
		Text:                   text,
		ChatID:                 chatID,
		MessageID:              messageID,
		EditMessageTextOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodEditMessageText, request); err != nil {
		return nil, err
	}
	return result, nil
}

// EditInlineMessageText edits text and game inline messages through the Telegram bot API.
//
// Required parameters:
//   - text: New text of the message, 1-4096 characters after entities parsing.
//   - inlineMessageID: Identifier of the inline message.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageText" https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditInlineMessageText(inlineMessageID string, text Text, options *EditMessageTextOptions) error {
	request := EditMessageTextRequest{
		Text:                   text,
		InlineMessageID:        inlineMessageID,
		EditMessageTextOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditMessageText, request); err != nil {
		return err
	}
	return nil
}
