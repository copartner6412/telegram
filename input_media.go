package telegram

// InputMedia represents the content of a media message to be sent.
//
// It should be one of:
//   - InputMediaAnimation
//   - InputMediaDocument
//   - InputMediaAudio
//   - InputMediaPhoto
//   - InputMediaVideo
type InputMedia interface {
	getMediaType() string
}
