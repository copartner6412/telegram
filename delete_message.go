package telegram

import (
	"net/http"
)

//

// DeleteMessageRequest represent parameters to delete a message, including service messages, sent by the bot using the deleteMessage method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - MessageID
//
// See "deleteMessage" https://core.telegram.org/bots/api#deletemessage
type DeleteMessageRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Identifier of the message to delete
	MessageID int `json:"message_id"`
}

// DeleteMessage deletes a message, with the following limitations:
//   - A message can only be deleted if it was sent less than 48 hours ago.
//   - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
//   - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//   - Bots can delete outgoing messages in private chats, groups, and supergroups.
//   - Bots can delete incoming messages in private chats.
//   - Bots granted can_post_messages permissions can delete outgoing messages in channels.
//   - If the bot is an administrator of a group, it can delete any message there.
//   - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
//
// Parameters:
//   - chatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - messageID: Identifier of the message to delete
//
// See "deleteMessage" https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(chatID any, messageID int) error {
	request := DeleteMessageRequest{
		ChatID:    chatID,
		MessageID: messageID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodDeleteMessage, request); err != nil {
		return err
	}
	return nil
}
