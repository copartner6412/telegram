package telegram

// TransactionPartnerTelegramAds describes a withdrawal transaction to the Telegram Ads platform.
//
// See "TransactionPartnerTelegramAds" https://core.telegram.org/bots/api#transactionpartnertelegramads
type TransactionPartnerTelegramAds struct {
	// (Required) Type of the transaction partner, always “telegram_ads”
	Type string `json:"type"`
}
