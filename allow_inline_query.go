package telegram

// AnswerInlineQuery represents the method to send answers to an inline query.
// On success, True is returned. No more than 50 results per query are allowed.
//
// See https://core.telegram.org/bots/api#answerinlinequery
type AnswerInlineQuery struct {
	// (Required) Unique identifier for the answered query.
	InlineQueryID string `json:"inline_query_id"`

	// (Required) An array of results for the inline query.
	Results []InlineQueryResult `json:"results"`

	// (Optional) The maximum amount of time in seconds that the result of the inline query
	// may be cached on the server. Defaults to 300.
	CacheTime int `json:"cache_time,omitempty"`

	// (Optional) Pass True if results may be cached on the server side only for the
	// user that sent the query. By default, results may be returned to any user
	// who sends the same query.
	IsPersonal bool `json:"is_personal,omitempty"`

	// (Optional) Pass the offset that a client should send in the next query with
	// the same text to receive more results. Pass an empty string if there are no more
	// results or if you don't support pagination. Offset length can't exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`

	// (Optional) An object describing a button to be shown above inline query results.
	Button *InlineQueryResultsButton `json:"button,omitempty"`
}

// Note: InlineQueryResult and InlineQueryResultsButton structs should be defined as per their respective Telegram API definitions.
