package telegram

// ShippingQuery contains information about an incoming shipping query.
//
// See "ShippingQuery" https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	// (Required) Unique query identifier.
	ID string `json:"id"`

	// (Required) User who sent the query.
	FromUser User `json:"from_user"`

	// (Required) Bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// (Required) User specified shipping address.
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
