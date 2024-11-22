package telegram

// InputInvoiceMessageContent represents the content of an invoice message to be sent as the result of an inline query.
//
// Required fields:
//   - Title
//   - Description
//   - Payload
//   - Currency
//   - Prices
//
// See "InputInvoiceMessageContent" https://core.telegram.org/bots/api#inputinvoicemessagecontent
type InputInvoiceMessageContent struct {
	// (Required) Product name, 1-32 characters.
	Title string `json:"title"`

	// (Required) Product description, 1-255 characters.
	Description string `json:"description"`

	// (Required) Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload"`

	// (Required) Three-letter ISO 4217 currency code, see more on currencies https://core.telegram.org/bots/payments#supported-currencies. Pass “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`

	// (Required) Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	Prices []LabeledPrice `json:"prices"`

	// (Optional) Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken string `json:"provider_token,omitempty"`

	// (Optional) The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json https://core.telegram.org/bots/payments/currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// (Optional) An array of suggested amounts of tip in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`

	// (Optional) A JSON-serialized object for data about the invoice, which will be shared with the payment provider. A detailed description of the required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// (Optional) URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoUrl string `json:"photo_url,omitempty"`

	// (Optional) Photo size in bytes.
	PhotoSize int `json:"photo_size,omitempty"`

	// (Optional) Photo width.
	PhotoWidth int `json:"photo_width,omitempty"`

	// (Optional) Photo height.
	PhotoHeight int `json:"photo_height,omitempty"`

	// (Optional) Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedName bool `json:"need_name,omitempty"`

	// (Optional) Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// (Optional) Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail bool `json:"need_email,omitempty"`

	// (Optional) Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// (Optional) Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// (Optional) Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// (Optional) Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
	IsFlexible bool `json:"is_flexible,omitempty"`
}

func (i InputInvoiceMessageContent) isInputMessageContent() bool {
	return true
}
