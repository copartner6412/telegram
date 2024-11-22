package telegram

// PreCheckoutQuery contains information about an incoming pre-checkout query.
//
// See "PreCheckoutQuery" https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	// (Required) Unique query identifier.
	ID string `json:"id"`

	// (Required) User who sent the query.
	FromUser User `json:"from_user"`

	// (Required) Three-letter ISO 4217 currency code https://core.telegram.org/bots/payments#supported-currencies, or “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`

	// (Required) Total price in the smallest units of the currency
	// (integer, not float/double). For example, for a price of US$ 1.45
	// pass amount = 145. See the exp parameter in currencies.json https://core.telegram.org/bots/payments/currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// (Required) Bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// (Optional) Identifier of the shipping option chosen by the user.
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// (Optional) Order information provided by the user.
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}
