package telegram

import "net/http"

// SendPollRequest represents the parameters for sending a native poll using sendPoll method through the Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Question
//   - Options
//
// See https://core.telegram.org/bots/api#sendpoll
type SendPollRequest struct {

	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID any `json:"chat_id"`

	// (Required) Poll question, 1-300 characters
	Question string `json:"question"`

	// (Optional) Mode for parsing entities in the question. See formatting options for more details. Currently, only custom emoji entities are allowed
	QuestionParseMode string `json:"question_parse_mode,omitempty"`

	// (Optional) A list of special entities that appear in the poll question. It can be specified instead of question_parse_mode
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`

	// (Required) A list of 2-10 answer options
	Options []InputPollOption `json:"options"`

	*SendPollOptions
}

// SendPollOptions represents the optional parameters for sending a native poll using sendPoll method through the Telegram bot API.
//
// See https://core.telegram.org/bots/api#sendpoll
type SendPollOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) True, if the poll needs to be anonymous, defaults to True
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// (Optional) Poll type, “quiz” or “regular”, defaults to “regular”
	Type string `json:"type,omitempty"`

	// (Optional) True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// (Optional) 0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// (Optional) Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation string `json:"explanation,omitempty"`

	// (Optional) Mode for parsing entities in the explanation.
	//
	// See formatting options for more details. https://core.telegram.org/bots/api#formatting-options
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`

	// (Optional) A list of special entities that appear in the poll explanation. It can be specified instead of explanation_parse_mode
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// (Optional) Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	OpenPeriod int `json:"open_period,omitempty"`

	// (Optional) Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future. Can't be used together with open_period.
	CloseDate int `json:"close_date,omitempty"`

	// (Optional) Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed bool `json:"is_closed,omitempty"`

	*SendOptions
}

// SendPoll sends a native poll to a specified chat through the Telegram bot API.
//
// Required parameters:
//   - toChatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
//   - pollQuestion: Poll question, 1-300 characters
//   - pollOptions: A list of 2-10 answer options
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(toChatID any, pollQuestion Caption, pollOptions []InputPollOption, options *SendPollOptions) (*Message, error) {
	request := SendPollRequest{
		ChatID:            toChatID,
		Question:          pollQuestion.Text,
		QuestionParseMode: pollQuestion.ParseMode,
		QuestionEntities:  pollQuestion.Entities,
		Options:           pollOptions,
		SendPollOptions:   options,
	}
	result := new(Message)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodSendPoll, request); err != nil {
		return nil, err
	}
	return result, nil
}
