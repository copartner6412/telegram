package telegram

// InputPaidMedia describes the paid media to be sent.
//
// Currently, it can be one of:
//   - InputPaidMediaPhoto
//   - InputPaidMediaVideo
//
// See "InputPaidMedia" https://core.telegram.org/bots/api#inputpaidmedia
type InputPaidMedia interface {
	getPaidMediaType() string
}
