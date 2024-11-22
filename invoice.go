package telegram

// Invoice contains basic information about an invoice.
//
// See "Invoice" https://core.telegram.org/bots/api#invoice
type Invoice struct {
	// (Required) Product name.
	Title string `json:"title"`

	// (Required) Product description.
	Description string `json:"description"`

	// (Optional) Unique bot deep-linking parameter that can be used to generate this invoice.
	StartParameter string `json:"start_parameter,omitempty"`

	// (Required) Three-letter ISO 4217 currency code https://core.telegram.org/bots/payments#supported-currencies, or “XTR” for payments in Telegram Stars.
	Currency string `json:"currency"`

	// (Required) Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json https://core.telegram.org/bots/payments/currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}
