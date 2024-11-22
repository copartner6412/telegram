package telegram

// ChatPhoto represents a chat photo.
//
// See "ChatPhoto" https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	// (Required) File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`

	// (Required) Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// (Required) File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`

	// (Required) Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}
