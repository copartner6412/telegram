package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

// AddStickerToSetRequest represents the parameters for adding a sticker to a set using the addStickerToSet method through the Telegram bot API.
//
// Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers.
//
// Required fields:
//   - UserID
//   - Name
//   - Sticker
//
// See "addStickerToSet" https://core.telegram.org/bots/api#addstickertoset
type AddStickerToSetRequest struct {
	// (Required) User identifier of sticker set owner.
	UserID int `json:"user_id"`

	// (Required) Sticker set name.
	Name string `json:"name"`

	// (Required) InputSticker. An object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
	Sticker InputSticker `json:"sticker"`
}

// AddStickerToSet adds a new sticker to a set created by the bot through the Telegram bot API.
//
// Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers.
//
// Parameters:
//   - userID: User identifier of sticker set owner
//   - name: Sticker set name
//   - sticker: An object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
//
// See "addStickerToSet" https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(userID int, name string, sticker InputSticker) error {
	requestPayload := new(bytes.Buffer)
	writer := multipart.NewWriter(requestPayload)

	if err := writer.WriteField("user_id", fmt.Sprintf("%d", userID)); err != nil {
		return fmt.Errorf("error adding user_id field: %w", err)
	}

	if err := writer.WriteField("name", name); err != nil {
		return fmt.Errorf("error adding name field: %w", err)
	}

	nestedWriter, err := writer.CreateFormField("sticker")
	if err != nil {
		return fmt.Errorf("error creating for field for sticker: %w", err)
	}
	if err := writeNestedMultipart(nestedWriter, sticker, []string{"sticker"}, []string{"sticker"}); err != nil {
		return fmt.Errorf("error writing nested multipart for field sticker: %w", err)
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
