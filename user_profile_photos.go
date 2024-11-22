package telegram

// UserProfilePhotos represent a user's profile pictures.
//
// See "UserProfilePhotos" https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	// (Required) Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// (Required) Requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}
