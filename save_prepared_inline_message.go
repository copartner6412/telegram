package telegram

import "net/http"

// SavePreparedInlineMessageRequest represents a request to store a message that can be sent by a user of a Mini App using the savePreparedInlineMessage method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - Result
//
// See "SavePreparedInlineMessage" https://core.telegram.org/bots/api#savepreparedinlinemessage
type SavePreparedInlineMessageRequest struct {
	// (Required) Unique identifier of the target user that can use the prepared message.
	UserID int `json:"user_id"`

	// (Required) A JSON-serialized object describing the message to be sent.
	Result InlineQueryResult `json:"result"`

	*SavePreparedInlineMessageOptions
}

// SavePreparedInlineMessageRequest represents a optional parameters to store a message that can be sent by a user of a Mini App using the savePreparedInlineMessage method through the Telegram bot API.
//
// See "SavePreparedInlineMessage" https://core.telegram.org/bots/api#savepreparedinlinemessage
type SavePreparedInlineMessageOptions struct {
	// (Optional) Pass True if the message can be sent to private chats with users.
	AllowUserChats bool `json:"allow_user_chats,omitempty"`

	// (Optional) Pass True if the message can be sent to private chats with bots.
	AllowBotChats bool `json:"allow_bot_chats,omitempty"`

	// (Optional) Pass True if the message can be sent to group and supergroup chats.
	AllowGroupChats bool `json:"allow_group_chats,omitempty"`

	// (Optional) Pass True if the message can be sent to channel chats.
	AllowChannelChats bool `json:"allow_channel_chats,omitempty"`
}

// SavePreparedInlineMessage stores a message that can be sent by a user of a Mini App.
//
// Required parameters:
//   - userID: Unique identifier of the target user that can use the prepared message
//   - result: An object describing the message to be sent
//
// Returns a PreparedInlineMessage object.
//
// See "SavePreparedInlineMessage" https://core.telegram.org/bots/api#savepreparedinlinemessage
func (b *Bot) SavePreparedInlineMessage(userID int, result InlineQueryResult, options *SavePreparedInlineMessageOptions) (*PreparedInlineMessage, error) {
	request := SavePreparedInlineMessageRequest{
		UserID:                           userID,
		Result:                           result,
		SavePreparedInlineMessageOptions: options,
	}
	resultStruct := new(PreparedInlineMessage)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSavePreparedInlineMessage, request); err != nil {
		return nil, err
	}
	return resultStruct, nil
}
