package telegram

// ShippingAddress represents a shipping address.
//
// See "ShippingAddress" https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	// (Required) Two-letter ISO 3166-1 alpha-2 country code https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2.
	CountryCode string `json:"country_code"`

	// (Optional) State, if applicable.
	State string `json:"state,omitempty"`

	// (Optional) City.
	City string `json:"city,omitempty"`

	// (Required) First line for the address.
	StreetLine1 string `json:"street_line1"`

	// (Optional) Second line for the address.
	StreetLine2 string `json:"street_line2,omitempty"`

	// (Required) Address post code.
	PostCode string `json:"post_code"`
}
