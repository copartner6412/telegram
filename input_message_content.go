package telegram

// InputMessageContent represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 5 types:
//   - InputTextMessageContent
//   - InputLocationMessageContent
//   - InputVenueMessageContent
//   - InputContactMessageContent
//   - InputInvoiceMessageContent
//
// See "InputMessageContent" https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface {
	isInputMessageContent() bool
}
