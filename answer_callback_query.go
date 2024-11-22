package telegram

import (
	"net/http"
)

// AnswerCallbackQueryRequest represents the parameters for answering callback queries sent from inline keyboards using the answerCallbackQuery method through the Telegram bot API.
//
// Required fields:
//   - CallbackQueryID
//
// See "answerCallbackQuery" https://core.telegram.org/bots/api#answercallbackquery
type AnswerCallbackQueryRequest struct {
	// (Required) Unique identifier for the query to be answered.
	CallbackQueryID string `json:"callback_query_id"`

	*AnswerCallbackQueryOptions
}

type AnswerCallbackQueryOptions struct {
	// (Optional) Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters.
	Text string `json:"text,omitempty"`

	// (Optional) If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// (Optional) URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url,omitempty"`

	// (Optional) The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}

// AnswerCallbackQuery sends answers to callback queries sent from inline keyboards through the Telegram bot API.
//
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create a game for your bot via @BotFather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
//
// Parameters:
//   - callbackQueryID: Unique identifier for the query to be answered.
//
// See "answerCallbackQuery" https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(callbackQueryID string, options *AnswerCallbackQueryOptions) error {
	request := AnswerCallbackQueryRequest{
		CallbackQueryID:            callbackQueryID,
		AnswerCallbackQueryOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodAnswerCallbackQuery, request); err != nil {
		return err
	}
	return nil
}
