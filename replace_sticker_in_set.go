package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

// ReplaceStickerInSetRequest represents parameters used to replace an existing sticker in a sticker set with a new one using the replaceStickerInSet method through the Telegram bot API.
//
// Required fields:
//   - UserID
//   - Name
//   - OldSticker
//   - Sticker
//
// See "replaceStickerInSet" https://core.telegram.org/bots/api#replacestickerinset
type ReplaceStickerInSet struct {
	// (Required) User identifier of the sticker set owner.
	UserID int `json:"user_id"`

	// (Required) Sticker set name.
	Name string `json:"name"`

	// (Required) File identifier of the replaced sticker.
	OldSticker string `json:"old_sticker"`

	// (Required) An object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
	Sticker InputSticker `json:"sticker"`
}

// ReplaceStickerInSet replaces an existing sticker in a sticker set with a new one through the Telegram bot API.
//
// The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet.
//
// Parameters:
//   - userID: User identifier of the sticker set owner.
//   - name: Sticker set name.
//   - oldSticker: File identifier of the replaced sticker.
//   - sticker: An object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
//
// See "replaceStickerInSet" https://core.telegram.org/bots/api#replacestickerinset
func (b *Bot) ReplaceStickerInSet(userID int, name, oldSticker string, sticker InputSticker) error {
	requestPayload := &bytes.Buffer{}
	writer := multipart.NewWriter(requestPayload)

	if err := writer.WriteField("user_id", fmt.Sprintf("%d", userID)); err != nil {
		return fmt.Errorf("error adding user_id field: %w", err)
	}

	if err := writer.WriteField("name", name); err != nil {
		return fmt.Errorf("error adding name field: %w", err)
	}

	if err := writer.WriteField("old_sticker", oldSticker); err != nil {
		return fmt.Errorf("error adding old_sticker field: %w", err)
	}

	nestedWriter, err := writer.CreateFormField("sticker")
	if err != nil {
		return fmt.Errorf("error creating form field for sticker: %w", err)
	}

	if err := writeNestedMultipart(nestedWriter, sticker, []string{"sticker"}, []string{"sticker"}); err != nil {
		return fmt.Errorf("error writing nested multipart form for sticker: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("error closing multipart writer: %w", err)
	}

	httpResponse, err := b.send(http.MethodPost, MethodAddStickerToSet, writer.FormDataContentType(), requestPayload)
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
