package telegram

import (
	"net/http"
)

// SendGame represents a request to send a game using the sendGame method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - GameShortName
//
// See "SendGameRequest" https://core.telegram.org/bots/api#sendgame
type SendGameRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat.
	ChatID int `json:"chat_id"`

	// (Required) Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	GameShortName string `json:"game_short_name"`

	*SendGameOptions
}

// SendGame represents a optional parameters to send a game using the sendGame method through the Telegram bot API.
//
// See "SendGameRequest" https://core.telegram.org/bots/api#sendgame
type SendGameOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// (Optional) Protects the contents of the sent message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`

	// (Optional) Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`

	// (Optional) Unique identifier of the message effect to be added to the message; for private chats only.
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// (Optional) Description of the message to reply to.
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// (Optional) An object for an inline keyboard.
	// If empty, one 'Play game_title' button will be shown.
	// If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// SendGame sends a game to a specified chat trough the Telegram bot API.
//
// Required parameters:
//   - chatID
//   - gameShortName
//
// On success, the sent Message is returned.
//
// See "SendGameRequest" https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(chatID int, gameShortName string, options *SendGameOptions) (*Message, error) {
	request := SendGameRequest{
		ChatID:          chatID,
		GameShortName:   gameShortName,
		SendGameOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendGame, request); err != nil {
		return nil, err
	}
	return result, nil
}
