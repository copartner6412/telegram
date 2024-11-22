package telegram

import "net/http"

// EditUserStarSubscriptionRequest allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars using the editUserStarSubscription method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - TelegramPaymentChargeID
//   - IsCanceled
//
// See "editUserStarSubscription" https://core.telegram.org/bots/api#edituserstarsubscription
type EditUserStarSubscriptionRequest struct {
	// (Required) Identifier of the user whose subscription will be edited
	UserID int `json:"user_id"`

	// (Required) Telegram payment identifier for the subscription
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// (Required) Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
	IsCanceled bool `json:"bool"`
}

// EditUserStarSubscription allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars through the Telegram bot API.
//
// Parameters:
//   - userID: Identifier of the user whose subscription will be edited
//   - telegramPaymentChargeID: Telegram payment identifier for the subscription
//   - isCanceled: Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
//
// See "editUserStarSubscription" https://core.telegram.org/bots/api#edituserstarsubscription
func (b *Bot) EditUserStarSubscription(userID int, telegramPaymentChargeID string, isCanceled bool) error {
	request := EditUserStarSubscriptionRequest{
		UserID:                  userID,
		TelegramPaymentChargeID: telegramPaymentChargeID,
		IsCanceled:              isCanceled,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodEditUserStarSubscription, request); err != nil {
		return err
	}
	return nil
}
