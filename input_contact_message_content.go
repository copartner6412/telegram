package telegram

// InputContactMessageContent represents the content of a contact message to be sent as the result of an inline query.
//
// Required fields:
//   - PhoneNumber
//   - FirstName
//
// See "InputContactMessageContent" https://core.telegram.org/bots/api#inputcontactmessagecontent
type InputContactMessageContent struct {
	// (Required) Contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// (Required) Contact's first name.
	FirstName string `json:"first_name"`

	// (Optional) Contact's last name.
	LastName string `json:"last_name,omitempty"`

	// (Optional) Additional data about the contact in the form of a vCard, 0-2048 bytes.
	VCard string `json:"vcard,omitempty"`
}

func (i InputContactMessageContent) isInputMessageContent() bool {
	return true
}
