package telegram

// RefundedPayment contains basic information about a refunded payment.
//
// See "RefundedPayment" https://core.telegram.org/bots/api#refundedpayment
type RefundedPayment struct {
	// (Required) Three-letter ISO 4217 currency code https://core.telegram.org/bots/payments#supported-currencies, or “XTR” for payments in Telegram Stars.
	// Currently, always “XTR”.
	Currency string `json:"currency"`

	// (Required) Total refunded price in the smallest units of the currency (integer,
	// not float/double). For example, for a price of US$ 1.45, total_amount = 145.
	// See the exp parameter in currencies.json https://core.telegram.org/bots/payments/currencies.json, it shows the number of digits past
	// the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// (Required) Bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// (Required) Telegram payment identifier.
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// (Optional) Provider payment identifier.
	ProviderPaymentChargeID string `json:"provider_payment_charge_id,omitempty"`
}
