package telegram

// InlineQueryResult represents one result of an inline query. Telegram clients currently support results of the following 20 types:
//   - InlineQueryResultCachedAudio
//   - InlineQueryResultCachedDocument
//   - InlineQueryResultCachedGif
//   - InlineQueryResultCachedMpeg4Gif
//   - InlineQueryResultCachedPhoto
//   - InlineQueryResultCachedSticker
//   - InlineQueryResultCachedVideo
//   - InlineQueryResultCachedVoice
//   - InlineQueryResultArticle
//   - InlineQueryResultAudio
//   - InlineQueryResultContact
//   - InlineQueryResultGame
//   - InlineQueryResultDocument
//   - InlineQueryResultGif
//   - InlineQueryResultLocation
//   - InlineQueryResultMpeg4Gif
//   - InlineQueryResultPhoto
//   - InlineQueryResultVenue
//   - InlineQueryResultVideo
//   - InlineQueryResultVoice
//
// # Note
//
// All URLs passed in inline query results will be available to end users and therefore must be assumed to be public.
//
// See "InlineQueryResult" https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult interface {
	getQueryType() string
}
