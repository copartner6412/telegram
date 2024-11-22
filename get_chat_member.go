package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetChatMemberRequest represents a request to obtain information about a member of a chat using the getChatMember through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - UserID
//
// See https://core.telegram.org/bots/api#getchatmember
type GetChatMemberRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`
}

// GetChatMember retrieves information about a member of a chat through the Telegram bot API.
//
// The method is only guaranteed to work for other users if the bot is an administrator in the chat.
//
// Parameters:
//   - chatID: Integer or String. Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername).
//   - userID: Unique identifier of the target user.
//
// On success, returns a ChatMember object.
//
// See https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(chatID any, userID int) (ChatMember, error) {
	request := GetChatMemberRequest{
		ChatID: chatID,
		UserID: userID,
	}
	result := new(json.RawMessage)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetChatMember, request); err != nil {
		return nil, err
	}
	status := struct {
		Status string `json:"status"`
	}{}
	if err := json.Unmarshal(*result, &status); err != nil {
		return nil, fmt.Errorf("error specifying member's status: %w", err)
	}
	switch status.Status {
	case "creator":
		member := new(ChatMemberOwner)
		if err := json.Unmarshal(*result, member); err != nil {
			return nil, err
		}
		return *member, nil
	case "administrator":
		member := new(ChatMemberAdministrator)
		if err := json.Unmarshal(*result, &member); err != nil {
			return nil, err
		}
		return *member, nil
	case "member":
		member := new(ChatMemberMember)
		if err := json.Unmarshal(*result, &member); err != nil {
			return nil, err
		}
		return *member, nil
	case "restricted":
		member := new(ChatMemberRestricted)
		if err := json.Unmarshal(*result, &member); err != nil {
			return nil, err
		}
		return *member, nil
	case "left":
		member := new(ChatMemberLeft)
		if err := json.Unmarshal(*result, &member); err != nil {
			return nil, err
		}
		return *member, nil
	case "kicked":
		member := new(ChatMemberBanned)
		if err := json.Unmarshal(*result, &member); err != nil {
			return nil, err
		}
		return *member, nil
	default:
		return nil, fmt.Errorf("unsupported member status: %s", status.Status)
	}
}
