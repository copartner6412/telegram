package telegram

// ShippingOption represents one shipping option.
//
// See "ShippingOption" https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	// (Required) Shipping option identifier.
	ID string `json:"id"`

	// (Required) Option title.
	Title string `json:"title"`

	// (Required) List of price portions.
	Prices []LabeledPrice `json:"prices"`
}
