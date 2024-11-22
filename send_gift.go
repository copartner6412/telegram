package telegram

import "net/http"

// SendGiftRequest represents parameters to send a gift to the given user using the sendGift method through the Telegram bot API.
//
// The gift can't be converted to Telegram Stars by the user.
//
// Required fields:
//   - UserID
//   - GiftID
//
// See "sendGift" https://core.telegram.org/bots/api#sendgift
type SendGiftRequest struct {
	// (Required) Unique identifier of the target user that will receive the gift
	UserID int `json:"user_id"`

	// (Required) Identifier of the gift
	GiftID string `json:"gift_id"`

	*SendGiftOptions
}

// SendGiftOptions represents optional parameters to send a gift to the given user using the sendGift method through the Telegram bot API.
//
// See "sendGift" https://core.telegram.org/bots/api#sendgift
type SendGiftOptions struct {
	// (Optional) Text that will be shown along with the gift; 0-255 characters
	Text string `json:"text,omitempty"`

	// (Optional) Mode for parsing entities in the text. See formatting options for more details. https://core.telegram.org/bots/api#formatting-options
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextParseMode string `json:"text_parse_mode,omitempty"`

	// (Optional) A list of special entities that appear in the gift text. It can be specified instead of text_parse_mode. Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

// SendGift sends a gift to the given user.
//
// Required parameters:
//   - userID: Unique identifier of the target user that will receive the gift
//   - giftID: Identifier of the gift
//
// See "sendGift" https://core.telegram.org/bots/api#sendgift
func (b *Bot) SendGift(userID int, giftID string, options *SendGiftOptions) error {
	request := SendGiftRequest{
		UserID:          userID,
		GiftID:          giftID,
		SendGiftOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSendGift, request); err != nil {
		return err
	}
	return nil
}
