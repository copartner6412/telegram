package telegram

// PassportElementError represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of:
//   - PassportElementErrorDataField
//   - PassportElementErrorFrontSide
//   - PassportElementErrorReverseSide
//   - PassportElementErrorSelfie
//   - PassportElementErrorFile
//   - PassportElementErrorFiles
//   - PassportElementErrorTranslationFile
//   - PassportElementErrorTranslationFiles
//   - PassportElementErrorUnspecified
//
// See "PassportElementError" https://core.telegram.org/bots/api#passportelementerror
type PassportElementError interface {
	getErrorType() string
}
