package telegram

import "net/http"

// EditMessageCaption represents the parameters to edit captions of messages using the editMessageCaption method through the Telegram bot API.
//
// Required fields:
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See "editMessageCaption" https://core.telegram.org/bots/api#editmessagecaption
type EditMessageCaptionRequest struct {
	*Caption

	// (Optional) Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id,omitempty"`

	// (Optional) Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	*EditMessageCaptionOptions
}

// EditMessageCaption represents the optional parameters to edit captions of messages using the editMessageCaption method through the Telegram bot API.
//
// See "editMessageCaption" https://core.telegram.org/bots/api#editmessagecaption
type EditMessageCaptionOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which
	// the message to be edited was sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Pass True if the caption must be shown above the message media.
	// Supported only for animation, photo and video messages.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// (Optional) An object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaption edits captions of messages through the Telegram bot API.
//
// Required parameters:
//   - chatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of the message to edit.
//
// On success, the edited Message is returned.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageCaption" https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(chatID any, messageID int, caption Caption, options *EditMessageCaptionOptions) (*Message, error) {
	request := EditMessageCaptionRequest{
		Caption:                   &caption,
		ChatID:                    chatID,
		MessageID:                 messageID,
		EditMessageCaptionOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodEditMessageCaption, request); err != nil {
		return nil, err
	}
	return result, nil
}

// EditInlineMessageCaption edits captions of inline messages through the Telegram bot API.
//
// Required parameters:
//   - inlineMessageID: Identifier of the inline message.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageCaption" https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditInlineMessageCaption(inlineMessageID string, caption Caption, options *EditMessageCaptionOptions) error {
	request := EditMessageCaptionRequest{
		Caption:                   &caption,
		InlineMessageID:           inlineMessageID,
		EditMessageCaptionOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditMessageCaption, request); err != nil {
		return err
	}
	return nil
}
