package telegram

import (
	"net/http"
)

// CreateInvoiceLinkRequest represents a request to create a link for an invoice using the createInvoiceLink method through the Telegram bot API.
//
// Required fields:
//   - Title
//   - Description
//   - Payload
//   - Currency
//   - Prices
//
// See "CreateInvoiceLinkRequest" https://core.telegram.org/bots/api#createinvoicelink
type CreateInvoiceLinkRequest struct {
	// (Required) Product name, 1-32 characters.
	Title string `json:"title"`

	// (Required) Product description, 1-255 characters.
	Description string `json:"description"`

	// (Required) Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user; use it for your internal processes.
	Payload string `json:"payload"`

	// (Required) Three-letter ISO 4217 currency code. Pass "XTR" for payments in Telegram Stars.
	Currency string `json:"currency"`

	// (Required) Price breakdown, a JSON-serialized list of components (e.g., product price, tax, discount, delivery cost, delivery tax, bonus, etc.).
	// Must contain exactly one item for payments in Telegram Stars.
	Prices []LabeledPrice `json:"prices"`

	*CreateInvoiceLinkOptions
}

// CreateInvoiceLinkRequest represents optional parameters to create a link for an invoice through the Telegram bot API.
//
// See "CreateInvoiceLinkRequest" https://core.telegram.org/bots/api#createinvoicelink
type CreateInvoiceLinkOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the link will be created.
	// For payments in Telegram Stars only.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken string `json:"provider_token,omitempty"`

	// (Optional) The number of seconds the subscription will be active for before the next payment.
	// The currency must be set to "XTR" (Telegram Stars) if the parameter is used.
	// Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active
	// for a given bot at the same time, including multiple concurrent subscriptions from the same user.
	SubscriptionPeriod int `json:"subscription_period,omitempty"`

	// (Optional) The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double).
	// For example, for a maximum tip of US$ 1.45, pass max_tip_amount = 145. Defaults to 0.
	// Not supported for payments in Telegram Stars.
	MaxTipAmount int `json:"max_tip_amount,omitempty"`

	// (Optional) A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double).
	// At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive,
	// passed in a strictly increased order, and must not exceed max_tip_amount.
	SuggestedTipAmounts *[]int `json:"suggested_tip_amounts,omitempty"`

	// (Optional) JSON-serialized data about the invoice, which will be shared with the payment provider.
	// A detailed description of required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`

	// (Optional) URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoURL string `json:"photo_url,omitempty"`

	// (Optional) Photo size in bytes.
	PhotoSize int `json:"photo_size,omitempty"`

	// (Optional) Photo width.
	PhotoWidth int `json:"photo_width,omitempty"`

	// (Optional) Photo height.
	PhotoHeight int `json:"photo_height,omitempty"`

	// (Optional) Pass True if you require the user's full name to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedName bool `json:"need_name,omitempty"`

	// (Optional) Pass True if you require the user's phone number to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`

	// (Optional) Pass True if you require the user's email address to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedEmail bool `json:"need_email,omitempty"`

	// (Optional) Pass True if you require the user's shipping address to complete the order.
	// Ignored for payments in Telegram Stars.
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`

	// (Optional) Pass True if the user's phone number should be sent to the provider.
	// Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`

	// (Optional) Pass True if the user's email address should be sent to the provider.
	// Ignored for payments in Telegram Stars.
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`

	// (Optional) Pass True if the final price depends on the shipping method.
	// Ignored for payments in Telegram Stars.
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// CreateInvoiceLink generates a link for an invoice through the Telegram bot AOI.
//
// Parameters:
//   - title: Product name, 1-32 characters
//   - description: Product description, 1-255 characters
//   - payload: Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
//   - currency: Three-letter ISO 4217 currency code. Pass “XTR” for payments in Telegram Stars.
//   - prices: Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
//
// Returns the created invoice link as a string upon success.
//
// See "CreateInvoiceLinkRequest" https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(title, description, payload, currency string, prices []LabeledPrice, options *CreateInvoiceLinkOptions) (string, error) {
	request := CreateInvoiceLinkRequest{
		Title:                    title,
		Description:              description,
		Payload:                  payload,
		Currency:                 currency,
		Prices:                   prices,
		CreateInvoiceLinkOptions: options,
	}
	result := new(string)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodCreateInvoiceLink, request); err != nil {
		return "", err
	}
	return *result, nil
}
