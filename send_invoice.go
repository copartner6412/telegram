package telegram

import (
	"net/http"
)

// SendInvoiceRequest represents the parameters used to send invoices using the sendInvoice method through the Telegram bot API.
//
// Required fields:
//   - Title
//   - Description
//   - Payload
//   - Currency
//   - Prices
//
// On success, the sent Message is returned.
//
// See "sendInvoice" https://core.telegram.org/bots/api#sendinvoice
type SendInvoiceRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Product name, 1-32 characters.
	Title string `json:"title"`

	// (Required) Product description, 1-255 characters.
	Description string `json:"description"`

	// (Required) Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`

	// (Required) Three-letter ISO 4217 currency code, see more on currencies https://core.telegram.org/bots/payments#supported-currencies. Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`

	// (Required) Price breakdown, a list of components. Must contain exactly one item for payments in Telegram Stars.
	Prices []LabeledPrice `json:"prices"`

	*SendInvoiceOptions
}

// SendInvoiceOptions represents the optional parameters used to send invoices using the sendInvoice method through the Telegram bot API.
//
// See "sendInvoice" https://core.telegram.org/bots/api#sendinvoice
type SendInvoiceOptions struct {
	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken string `json:"provider_token,omitempty"`

	// (Optional) The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json https://core.telegram.org/bots/payments/currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// (Optional) An array of suggested amounts of tips in the smallest units of the currency (integer). At most 4 suggested tip amounts can be specified.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// (Optional) Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button. Otherwise, a URL button with a deep link will be shown.
	StartParameter string `json:"start_parameter,omitempty"`

	// (Optional) JSON-serialized data about the invoice to be shared with the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// (Optional) URL of the product photo for the invoice.
	PhotoURL string `json:"photo_url,omitempty"`

	// (Optional) Size of the product photo in bytes.
	PhotoSize int `json:"photo_size,omitempty"`

	// (Optional) Width of the product photo.
	PhotoWidth int `json:"photo_width,omitempty"`

	// (Optional) Height of the product photo.
	PhotoHeight int `json:"photo_height,omitempty"`

	// (Optional) Pass True if you require the user's full name to complete the order.
	NeedName bool `json:"need_name,omitempty"`

	// (Optional) Pass True if you require the user's phone number to complete the order.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// (Optional) Pass True if you require the user's email address to complete the order.
	NeedEmail bool `json:"need_email,omitempty"`

	// (Optional) Pass True if you require the user's shipping address to complete the order.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// (Optional) Pass True if the user's phone number should be sent to the provider.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// (Optional) Pass True if the user's email address should be sent to the provider.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// (Optional) Pass True if the final price depends on the shipping method.
	IsFlexible bool `json:"is_flexible,omitempty"`

	*SendOptions
}

// SendInvoice sends an invoice to a specified chat.
//
// Required parameters:
//   - chatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - title: Product name, 1-32 characters
//   - description: Product description, 1-255 characters
//   - payload: Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
//   - currency: Three-letter ISO 4217 currency code. Pass “XTR” for payments in Telegram Stars.
//   - prices: Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
//
// On success, the sent Message is returned.
//
// See "sendInvoice" https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(chatID any, title, description, payload, currency string, prices []LabeledPrice, options *SendInvoiceOptions) (*Message, error) {
	request := SendInvoiceRequest{
		ChatID:             chatID,
		Title:              title,
		Description:        description,
		Payload:            payload,
		Currency:           currency,
		Prices:             prices,
		SendInvoiceOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendInvoice, request); err != nil {
		return nil, err
	}
	return result, nil
}
