package telegram

import (
	"net/http"
)

// SendVenueRequest represents the parameters for sending a phone contact using sendContact method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - PhoneNumber
//   - FirstName
//
// See "sendContact" https://core.telegram.org/bots/api#sendcontact
type SendContactRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// (Required) Contact's first name
	FirstName string `json:"first_name"`

	*SendContactOptions
}

// SendVenueOptions represents the optional parameters for sending a phone contact using sendContact method through the Telegram bot API.
//
// See "sendContact" https://core.telegram.org/bots/api#sendcontact
type SendContactOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Contact's last name
	LastName string `json:"last_name,omitempty"`

	// (Optional) Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"vcard,omitempty"`

	*SendOptions
}

// SendContact sends phone contacts to a specified chat.
//
// Required parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - phoneNumber: Contact's phone number
//   - firstName: Additional data about the contact in the form of a vCard, 0-2048 bytes
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(toChatID any, phoneNumber, firstName string, options *SendContactOptions) (*Message, error) {
	request := SendContactRequest{
		ChatID:             toChatID,
		PhoneNumber:        phoneNumber,
		FirstName:          firstName,
		SendContactOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendContact, request); err != nil {
		return nil, err
	}
	return result, nil
}
