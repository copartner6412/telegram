package telegram

import "net/http"

// AnswerPreCheckoutQueryRequest represents the payload for responding to a pre-checkout query using the answerPreCheckoutQuery method through the Telegram bot API.
//
// Required fields:
//   - PreCheckoutQueryID
//   - Ok
//
// See "answerPreCheckoutQueryRequest" https://core.telegram.org/bots/api#answerprecheckoutquery
type AnswerPreCheckoutQueryRequest struct {
	// (Required) Unique identifier for the query to be answered.
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// (Required) Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order.
	// Use False if there are any problems.
	Ok bool `json:"ok"`

	// (Optional) Required if ok is False. Error message in human-readable form that explains the reason for failure to proceed with checkout
	// (e.g., "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!").
	// Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQueryOk responds ok to a pre-checkout query, confirming the user's payment and shipping details through the Telegram bot API.
//
// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query.
//
// Required parameters:
//   - preCheckoutQueryID: Unique identifier for the query to be answered
//
// # Note
//
// The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
//
// See "answerPreCheckoutQuery" https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQueryOk(preCheckoutQueryID string) error {
	request := AnswerPreCheckoutQueryRequest{
		PreCheckoutQueryID: preCheckoutQueryID,
		Ok:                 true,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodAnswerPreCheckoutQuery, request); err != nil {
		return err
	}
	return nil
}

// AnswerPreCheckoutQueryNotOk responds not ok to a pre-checkout query, confirming the user's payment and shipping details through the Telegram bot API.
//
// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query.
//
// Required parameters:
//   - preCheckoutQueryID: Unique identifier for the query to be answered
//   - errorMessage: Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
//
// # Note
//
// The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
//
// See "answerPreCheckoutQuery" https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQueryNotOk(preCheckoutQueryID, errorMessage string) error {
	request := AnswerPreCheckoutQueryRequest{
		PreCheckoutQueryID: preCheckoutQueryID,
		Ok:                 false,
		ErrorMessage:       errorMessage,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodAnswerPreCheckoutQuery, request); err != nil {
		return err
	}
	return nil
}
