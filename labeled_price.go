package telegram

// LabeledPrice represents a portion of the price for goods or services.
//
// See "LabeledPrice" https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	// (Required) Portion label.
	Label string `json:"label"`

	// (Required) Price of the product in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json https://core.telegram.org/bots/payments/currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}
