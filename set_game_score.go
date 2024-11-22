package telegram

import "net/http"

// SetGameScoreRequest represents parameters used to set the score of the specified user in a game message using the setGameScore method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - Score
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See "setGameScore" https://core.telegram.org/bots/api#setgamescore
type SetGameScoreRequest struct {
	// (Optional) Unique identifier for the target chat. Required if inline_message_id is not specified.
	ChatID int `json:"chat_id,omitempty"`

	// (Optional) Identifier of the sent message. Required if inline_message_id is not specified.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Identifier of the inline message. Required if chat_id and message_id are not specified.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// (Required) User identifier.
	UserID int `json:"user_id"`

	// (Required) New score, must be non-negative.
	Score int `json:"score"`

	*SetGameScoreOptions
}

type SetGameScoreOptions struct {
	// (Optional) Pass True if the high score is allowed to decrease.
	// This can be useful when fixing mistakes or banning cheaters.
	Force bool `json:"force,omitempty"`

	// (Optional) Pass True if the game message should not be automatically edited to include the current scoreboard.
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
}

// SetGameScore sets the score of the specified user in a game message through the Telegram bot API.
//
// Required parameters:
//   - chatID: Unique identifier for the target chat
//   - messageID: Identifier of the sent message
//   - userID: User identifier
//   - score: New score, must be non-negative
//
// On success, the Message is returned.
//
// Returns an error if the new score is not greater than the user's current score in the chat and force is False.
//
// See "setGameScore" https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(chatID, messageID, userID, score int, options *SetGameScoreOptions) (*Message, error) {
	request := SetGameScoreRequest{
		ChatID:              chatID,
		MessageID:           messageID,
		UserID:              userID,
		Score:               score,
		SetGameScoreOptions: options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSetGameScore, request); err != nil {
		return nil, err
	}
	return result, nil
}

// SetInlineGameScore sets the score of the specified user in a inline game message through the Telegram bot API.
//
// Required parameters:
//   - inlineMessageID: Identifier of the inline message
//   - userID: User identifier
//   - score: New score, must be non-negative
//
// Returns an error if the new score is not greater than the user's current score in the chat and force is False.
//
// See "setGameScore" https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetInlineGameScore(inlineMessageID string, userID, score int, options *SetGameScoreOptions) error {
	request := SetGameScoreRequest{
		InlineMessageID:     inlineMessageID,
		UserID:              userID,
		Score:               score,
		SetGameScoreOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSetGameScore, request); err != nil {
		return err
	}
	return nil
}
