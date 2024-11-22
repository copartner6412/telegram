package telegram

import (
	"net/http"
)

// GetUserProfilePhotosRequest represents the request payload for the getUserProfilePhotos method through the Telegram bot API.
//
// Required fields:
//   - UserID
//
// See https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotosRequest struct {
	// (Required) Unique identifier of the target user.
	UserID int `json:"user_id"`

	*GetUserProfilePhotoOptions
}

// GetUserProfilePhotosRequest represents the optional parameters for request payload of getUserProfilePhotos method through the Telegram bot API.
//
// Required fields:
//   - UserID
//
// See https://core.telegram.org/bots/api#getuserprofilephotos
type GetUserProfilePhotoOptions struct {
	// (Optional) is the sequential number of the first photo to
	// be returned. By default, all photos are returned.
	Offset int `json:"offset,omitempty"`

	// (Optional) is the maximum number of photos to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

// GetUserProfilePhotos retrieves a list of profile pictures for a specified user through the Telegram bot API.
//
// Required parameters:
//   - userID: Unique identifier of the target user.
//
// Results:
//   - photos: Requested profile pictures (in up to 4 sizes each
//   - totalCount: Total number of profile pictures the target user has
//
// See https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(userID int, options *GetUserProfilePhotoOptions) (photos [][]PhotoSize, totalCount int, err error) {
	request := GetUserProfilePhotosRequest{
		UserID:                     userID,
		GetUserProfilePhotoOptions: options,
	}
	result := new(UserProfilePhotos)
	if err := b.sendJsonForResult(result, http.MethodPost, MethodGetUserProfilePhotos, request); err != nil {
		return nil, 0, err
	}
	return result.Photos, result.TotalCount, nil
}
