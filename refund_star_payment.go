package telegram

import "net/http"

// RefundStarPaymentRequest represents parameters to refund a successful payment in Telegram Stars using the refundStarPayment method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - TelegramPaymentChargeID
//
// See "refundStarPayment" https://core.telegram.org/bots/api#refundstarpayment
type RefundStarPaymentRequest struct {
	// (Required) Identifier of the user whose payment will be refunded
	UserID int `json:"user_id"`

	// (Required) Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

// RefundStarPayment refunds a successful payment in Telegram Stars using the refundStarPayment through the Telegram bot API.
//
// Parameters:
//   - userID: Identifier of the user whose payment will be refunded
//   - telegramPaymentChargeID: Telegram payment identifier
//
// See "refundStarPayment" https://core.telegram.org/bots/api#refundstarpayment
func (b *Bot) RefundStarPayment(userID int, telegramPaymentChargeID string) error {
	request := RefundStarPaymentRequest{
		UserID:                  userID,
		TelegramPaymentChargeID: telegramPaymentChargeID,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodRefundStarPayment, request); err != nil {
		return err
	}
	return nil
}
