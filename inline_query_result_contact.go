package telegram

// InlineQueryResultContact represents a contact with a phone number.
//
// By default, this contact will be sent by the user.
//
// Alternatively, you can use InputMessageContent to send a message with the specified content instead of the contact.
//
// Required fields:
//   - Type
//   - ID
//   - PhoneNumber
//   - FirstName
//
// See "InlineQueryResultContact" https://core.telegram.org/bots/api#inlinequeryresultcontact
type InlineQueryResultContact struct {
	// (Required) Type of the result, must be "contact".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// (Required) Contact's first name.
	FirstName string `json:"first_name"`

	// (Optional) Contact's last name.
	LastName string `json:"last_name,omitempty"`

	// (Optional) Additional data about the contact in the form of a vCard, 0-2048 bytes.
	VCard string `json:"vcard,omitempty"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) Content of the message to be sent instead of the contact.
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// (Optional) URL of the thumbnail for the result.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) Thumbnail width.
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// (Optional) Thumbnail height.
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

func (r InlineQueryResultContact) getQueryType() string {
	return r.Type
}
