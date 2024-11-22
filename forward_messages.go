package telegram

import (
	"net/http"
)

// ForwardMessagesRequest represents the parameters for forwarding multiple messages using forwardMessages method.
//
// Required fields:
//   - ChatID
//   - FromChatID
//   - MessageIDs
//
// See "forwardMessages" https://core.telegram.org/bots/api#forwardmessages
type ForwardMessagesRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatID any `json:"from_chat_id"`

	// (Required) A list of 1-100 identifiers of messages in the chat from_chat_id to forward.
	// The identifiers must be specified in a strictly increasing order.
	MessageIDs []int `json:"message_ids"`

	*ForwardMessageOptions
}

// ForwardMessages forwards multiple messages of any kind. If some of the specified messages can't be
// found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
//
// Required parameters:
//   - fromChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - toChatID: Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
//   - messageIDs: A list of 1-100 identifiers of messages in the chat from_chat_id to forward. The identifiers must be specified in a strictly increasing order.
//
// On success, an array of MessageId of the sent messages is returned.
//
// See "forwardMessages" https://core.telegram.org/bots/api#forwardmessages
func (b *Bot) ForwardMessages(fromChatID, toChatID any, messageIDs []int, options *ForwardMessageOptions) ([]int, error) {
	request := ForwardMessagesRequest{
		ChatID:                toChatID,
		FromChatID:            fromChatID,
		MessageIDs:            messageIDs,
		ForwardMessageOptions: options,
	}
	result := new([]MessageID)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodForwardMessages, request); err != nil {
		return nil, err
	}

	var ids []int
	for _, messageID := range *result {
		ids = append(ids, messageID.MessageID)
	}

	return ids, nil
}
