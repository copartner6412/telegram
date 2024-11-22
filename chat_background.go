package telegram

import "encoding/json"

// ChatBackground represents a chat background.
//
// See "ChatBackground" https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	// (Required) Type of the background.
	//
	// Its type can be on of:
	//   - BackgroundTypeFill
	//   - BackgroundTypeWallpaper
	//   - BackgroundTypePattern
	//   - BackgroundTypeChatTheme
	Type json.RawMessage `json:"type"`
}
