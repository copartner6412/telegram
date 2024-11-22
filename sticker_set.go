package telegram

// StickerSet represents a sticker set object in Telegram.
//
// Required fields:
//   - Name
//   - Title
//   - StickerType
//   - Stickers
//
// See "StickerSet" https://core.telegram.org/bots/api#stickerSet
type StickerSet struct {
	// (Required) Sticker set name.
	Name string `json:"name"`

	// (Required) Sticker set title.
	Title string `json:"title"`

	// (Required) Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”.
	StickerType string `json:"sticker_type"`

	// (Required) List of all set stickers.
	Stickers []Sticker `json:"stickers"`

	// (Optional) Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
