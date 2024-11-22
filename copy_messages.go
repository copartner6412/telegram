package telegram

import (
	"net/http"
)

// CopyMessagesRequest represents the parameters for copying multiple messages using copyMessages method through Telegram bot API.
//
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
//
// Required fields:
//   - ChatID
//   - FromChatID
//   - MessageIDs
//
// Send "copyMessages" https://core.telegram.org/bots/api#copymessages
type CopyMessagesRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatID any `json:"from_chat_id"`

	// (Required) A list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
	MessageIDs []int `json:"message_ids"`

	*CopyMessagesOptions
}

// CopyMessagesOptions represents the optional parameters for copying multiple messages using copyMessages method through Telegram bot API.
//
// Send "copyMessages" https://core.telegram.org/bots/api#copymessages
type CopyMessagesOptions struct {
	// (Optional) Unique identifier for the target message thread (topic) of the forum;
	// for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// (Optional) Protects the contents of the sent messages from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// (Optional) Pass True to copy the messages without their captions
	RemoveCaption bool `json:"remove_caption,omitempty"`
}

// CopyMessages copies messages of any kind.
//
// If some of the specified messages can't be found or copied, they are skipped.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message.
// Album grouping is kept for copied messages.
//
// Required parameters:
//   - fromChatID: Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
//   - toChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - messageIDs: A list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
//
// On success, an array of MessageId of the sent messages is returned.
//
// Send "copyMessages" https://core.telegram.org/bots/api#copymessages
func (b *Bot) CopyMessages(fromChatID, toChatID any, messageIDs []int, options *CopyMessagesOptions) ([]int, error) {
	request := CopyMessagesRequest{
		ChatID:              toChatID,
		FromChatID:          fromChatID,
		MessageIDs:          messageIDs,
		CopyMessagesOptions: options,
	}
	result := new([]MessageID)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodCopyMessage, request); err != nil {
		return nil, err
	}

	var ids []int
	for _, messageID := range *result {
		ids = append(ids, messageID.MessageID)
	}

	return ids, nil
}
