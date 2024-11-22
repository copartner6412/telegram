package telegram

import (
	"net/http"
)

// GetGameHighScoresRequest represents the request payload to get data for high score tables using the getGameHighScores method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See "getGameHighScores" https://core.telegram.org/bots/api#getgamehighscores
type GetGameHighScoresRequest struct {
	// (Required) Target user id
	UserID int `json:"user_id"`

	// (Optional) Required if inline_message_id is not specified. Unique identifier for the target chat.
	ChatID int `json:"chat_id,omitempty"`

	// (Optional) Required if inline_message_id is not specified. Identifier of the sent message.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// GetGameHighScores retrieves data for high score tables for a specified user in a game through the Telegram bot API.
//
// Parameters:
//   - userID: Target user id
//   - chatID: Unique identifier for the target chat.
//   - messageID: Identifier of the sent message.
//
// This method will return the score of the specified user and several of their neighbors in a game.
// It returns an array of GameHighScore objects.
//
// # Note
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
//
// See "getGameHighScores" https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(userID int, chatID int, messageID int) ([]GameHighScore, error) {
	request := GetGameHighScoresRequest{
		UserID:    userID,
		ChatID:    chatID,
		MessageID: messageID,
	}
	result := new([]GameHighScore)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetGameScores, request); err != nil {
		return nil, err
	}
	return *result, nil
}

// GetInlineGameHighScores retrieves data for high score tables for a specified user in a game through the Telegram bot API.
//
// Parameters:
//   - userID: Target user id
//   - inlineMessageID: Identifier of the inline message.
//
// This method will return the score of the specified user and several of their neighbors in a game.
// It returns an array of GameHighScore objects.
//
// # Note
//
// This method will currently return scores for the target user, plus two of their closest neighbors on each side. Will also return the top three users if the user and their neighbors are not among them. Please note that this behavior is subject to change.
//
// See "getGameHighScores" https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetInlineGameHighScores(userID int, inlineMessageID string) ([]GameHighScore, error) {
	request := GetGameHighScoresRequest{
		UserID:          userID,
		InlineMessageID: inlineMessageID,
	}
	result := new([]GameHighScore)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetGameScores, request); err != nil {
		return nil, err
	}

	return *result, nil
}
