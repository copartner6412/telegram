package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

// CreateNewStickerSetRequest represents the parameters required to create a new sticker set owned by a user using the createNewStickerSet method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - Name
//   - Title
//   - Stickers
//
// See "createNewStickerSetRequest" https://core.telegram.org/bots/api#createNewStickerSet
type CreateNewStickerSetRequest struct {
	// (Required) User identifier of created sticker set owner.
	UserID int `json:"user_id"`

	// (Required) Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals).
	// Can contain only English letters, digits and underscores.
	// Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>".
	// <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`

	// (Required) Sticker set title, 1-64 characters.
	Title string `json:"title"`

	// (Required) A list of 1-50 initial stickers to be added to the sticker set.
	Stickers []InputSticker `json:"stickers"`

	*CreateNewStickerSetOptions
}

// CreateNewStickerSetOptions represents the optional parameters required to create a new sticker set owned by a user using the createNewStickerSet method through the Telegram bot API.
//
// See "createNewStickerSetRequest" https://core.telegram.org/bots/api#createNewStickerSet
type CreateNewStickerSetOptions struct {
	// (Optional) Type of stickers in the set, pass one of:
	//   - “regular”
	//   - “mask”
	//   - “custom_emoji”
	//
	// By default, a regular sticker set is created.
	StickerType string `json:"sticker_type,omitempty"`

	// (Optional) Pass True if stickers in the sticker set must be repainted to the color of text when used in messages,
	// the accent color if used as emoji status, white on chat photos, or another appropriate color based on context;
	// for custom emoji sticker sets only.
	NeedsRepainting bool `json:"needs_repainting,omitempty"`
}

// CreateNewStickerSet creates a new sticker set owned by a user through the Telegram bot API.
//
// The bot will be able to edit the sticker set thus created.
//
// Required parameters:
//   - userID: User identifier of created sticker set owner.
//   - name: Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
//   - title: Sticker set title, 1-64 characters
//   - stickers: A list of 1-50 initial stickers to be added to the sticker set.
//
// See "createNewStickerSet" https://core.telegram.org/bots/api#createNewStickerSet
func (b *Bot) CreateNewStickerSet(userID int, name, title string, stickers []InputSticker, options *CreateNewStickerSetOptions) error {
	requestPayload := new(bytes.Buffer)
	writer := multipart.NewWriter(requestPayload)

	if err := writer.WriteField("user_id", fmt.Sprintf("%d", userID)); err != nil {
		return fmt.Errorf("error adding user_id field: %w", err)
	}

	if err := writer.WriteField("name", name); err != nil {
		return fmt.Errorf("error adding name field: %w", err)
	}

	if err := writer.WriteField("title", title); err != nil {
		return fmt.Errorf("error adding title field: %w", err)
	}

	if options != nil {
		if options.NeedsRepainting {
			if err := writer.WriteField("needs_repainting", fmt.Sprintf("%v", options.NeedsRepainting)); err != nil {
				return fmt.Errorf("error adding needs_repainting field: %w", err)
			}

		}
		if options.StickerType != "" {
			if err := writer.WriteField("sticker_type", options.StickerType); err != nil {
				return fmt.Errorf("error adding sticker_type field: %w", err)
			}
		}
	}

	for i, sticker := range stickers {
		nestedWriter, err := writer.CreateFormField(fmt.Sprintf("stickers[%d]", i))
		if err != nil {
			return fmt.Errorf("error creating form field for stickers[%d]: %w", i, err)
		}
		if err := writeNestedMultipart(nestedWriter, sticker, []string{"sticker"}, []string{"sticker"}); err != nil {
			return fmt.Errorf("error writing nested multipart form for field stickers[%d]: %w", i, err)
		}
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("error closing multipart writer: %w", err)
	}

	httpResponse, err := b.send(http.MethodPost, MethodCreateNewStickerSet, writer.FormDataContentType(), requestPayload)
	if err != nil {
		return err
	}

	var response struct {
		Ok          bool    `json:"ok"`
		Result      bool    `json:"result"`
		Description *string `json:"description"`
		ErrorCode   *int    `json:"error_code"`
		// Parameters  ResponseParameters `json:"parameters"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	if !response.Ok || !response.Result || response.ErrorCode != nil || response.Description != nil {
		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
		if !response.Ok || !response.Result {
			errorMessage += ", result \"false\" (request unsuccessful)"
		} else {
			errorMessage += ", result \"true\" (request successful)"
		}
		if response.ErrorCode != nil {
			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", *response.ErrorCode)
		}
		if response.Description != nil {
			errorMessage += fmt.Sprintf(", description: %s", *response.Description)
		}
		return errors.New(errorMessage)
	}
	return nil
}
