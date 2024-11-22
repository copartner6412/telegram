package telegram

import (
	"net/http"
)

// CopyMessageRequest represents the parameters for copying a message.
//
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
//
// Required fields:
//   - ChatID
//   - FromChatID
//   - MessageID
//
// See "copyMessage" https://core.telegram.org/bots/api#copymessage
type CopyMessageRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	FromChatID any `json:"from_chat_id"`

	// (Required) Message identifier in the chat specified in from_chat_id
	MessageID int `json:"message_id"`

	*CopyMessageOptions
}

// CopyMessageOptions represents the optional parameters for copying a message.
//
// See "copyMessage" https://core.telegram.org/bots/api#copymessage
type CopyMessageOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	*Caption

	// (Optional) Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	*SendOptions
}

// CopyMessages copies messages of any kind.
//
// The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
//
// Required parameters:
//   - Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
//   - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - Message identifier in the chat specified in from_chat_id
//
// Returns the MessageId of the sent message on success.
//
// See "copyMessage" https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(toChatID, fromChatID any, messageID int, options *CopyMessageOptions) (int, error) {
	request := CopyMessageRequest{
		ChatID:             toChatID,
		FromChatID:         fromChatID,
		MessageID:          messageID,
		CopyMessageOptions: options,
	}
	result := new(MessageID)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodCopyMessage, request); err != nil {
		return 0, err
	}

	return result.MessageID, nil
}
