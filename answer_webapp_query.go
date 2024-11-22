package telegram

import (
	"net/http"
)

// AnswerWebAppQuery represents the result of an interaction with a Web App and sends a corresponding message on behalf of the user to the chat from which the query originated using the answerWebAppQuery method through the Telegram bot API.
//
// Required fields:
//   - WebAppQueryID
//   - Result
//
// See "answerWebAppQuery" https://core.telegram.org/bots/api#answerwebappquery
type AnswerWebAppQueryRequest struct {
	// (Required) Unique identifier for the query to be answered.
	WebAppQueryID string `json:"web_app_query_id"`

	// (Required) An object describing the message to be sent.
	Result InlineQueryResult `json:"result"`
}

// AnswerWebAppQuery sets the result of an interaction with a Web App and sends a corresponding message on behalf of the user to the chat from which the query originated.
//
// Parameters:
//   - webAppQueryID: Unique identifier for the query to be answered
//   - inlineQueryResult: An object describing the message to be sent
//
// On success, a SentWebAppMessage object is returned.
//
// See "answerWebAppQuery" https://core.telegram.org/bots/api#answerwebappquery
func (b *Bot) AnswerWebAppQuery(webAppQueryID string, result InlineQueryResult) (*SentWebAppMessage, error) {
	request := AnswerWebAppQueryRequest{
		WebAppQueryID: webAppQueryID,
		Result:        result,
	}
	resultStruct := new(SentWebAppMessage)
	if err := b.sendJsonForResult(resultStruct, http.MethodPost, MethodAnswerWebAppQuery, request); err != nil {
		return nil, err
	}
	return resultStruct, nil
}
