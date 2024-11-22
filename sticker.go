package telegram

// Sticker represents a sticker object in Telegram.
//
// Required fields:
//   - FileID
//   - FileUniqueID
//   - Type
//   - Width
//   - Height
//   - IsAnimated
//   - IsVideo
//
// See "sticker" https://core.telegram.org/bots/api#sticker
type Sticker struct {
	// (Required) Identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// (Required) Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// (Required) Type of the sticker, currently one of:
	//   - “regular”
	//   - “mask”
	//   - “custom_emoji”
	// The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	Type string `json:"type"`

	// (Required) Sticker width.
	Width int `json:"width"`

	// (Required) Sticker height.
	Height int `json:"height"`

	// (Required) True, if the sticker is animated.
	IsAnimated bool `json:"is_animated"`

	// (Required) True, if the sticker is a video sticker.
	IsVideo bool `json:"is_video"`

	// (Optional) Sticker thumbnail in the .WEBP or .JPG format.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// (Optional) Emoji associated with the sticker.
	Emoji string `json:"emoji,omitempty"`

	// (Optional) Name of the sticker set to which the sticker belongs.
	SetName string `json:"set_name,omitempty"`

	// (Optional) For premium regular stickers, premium animation for the sticker.
	PremiumAnimation *File `json:"premium_animation,omitempty"`

	// (Optional) For mask stickers, the position where the mask should be placed.
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// (Optional) For custom emoji stickers, unique identifier of the custom emoji.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`

	// (Optional) True, if the sticker must be repainted to a text color in messages,
	// the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places.
	NeedsRepainting bool `json:"needs_repainting,omitempty"`

	// (Optional) File size in bytes.
	FileSize int `json:"file_size,omitempty"`
}
