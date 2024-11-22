package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

// EditMessageMediaRequest represents the parameters to edit animation, audio, document, photo, or video messages, or to add media to text messages using the editMessageMedia method through the Telegram bot API.
//
// Required fields:
//   - Media
//   - ChatID (Required if inline_message_id is not specified.)
//   - MessageID (Required if inline_message_id is not specified.)
//   - InlineMessageID (Required if chat_id and message_id are not specified.)
//
// See "editMessageMedia" https://core.telegram.org/bots/api#editmessagemedia
type EditMessageMediaRequest struct {
	// (Required) An object for a new media content of the message.
	Media InputMedia `json:"media"`

	// (Optional) Required if inline_message_id is not specified. Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id,omitempty"`

	// (Optional) Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageID int `json:"message_id,omitempty"`

	// (Optional) Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	*EditMessageOptions
}

// EditMessageOptions represents the optional parameters to edit messages through the Telegram bot API.
//
// See "editMessageMedia" https://core.telegram.org/bots/api#editmessagemedia
type EditMessageOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message to be edited was sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) An object for a new inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMedia edits animation, audio, document, photo, or video messages, or adds media to text messages through the Telegram bot API.
//
// If a message is part of a message album, it can only be edited to an audio for audio albums, only to a document for document albums, and to a photo or a video otherwise.
//
// When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL.
//
// Required parameters:
//   - chatID: Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - messageID: Identifier of the message to edit.
//   - media: An object for a new media content of the message.
//
// On success, the edited Message is returned.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageMedia" https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(chatID any, messageID int, media InputMedia, options *EditMessageOptions) (*Message, error) {
	requestPayload := new(bytes.Buffer)
	writer := multipart.NewWriter(requestPayload)
	if err := writer.WriteField("chat_id", fmt.Sprintf("%v", chatID)); err != nil {
		return nil, fmt.Errorf("error adding chat_id field: %w", err)
	}
	if err := writer.WriteField("message_id", fmt.Sprintf("%d", messageID)); err != nil {
		return nil, fmt.Errorf("error adding message_id field: %w", err)
	}
	if options != nil {
		if options.BusinessConnectionID != "" {
			if err := writer.WriteField("business_connection_id", options.BusinessConnectionID); err != nil {
				return nil, fmt.Errorf("error adding business_connection_id field: %w", err)
			}
		}
		if options.ReplyMarkup != nil {
			jsonData, err := json.Marshal(options.ReplyMarkup)
			if err != nil {
				return nil, fmt.Errorf("error JSON marshaling reply markup: %w", err)
			}
			if err := writer.WriteField("reply_markup", string(jsonData)); err != nil {
				return nil, fmt.Errorf("error adding reply_markup field: %w", err)
			}
		}
	}

	nestedWriter, err := writer.CreateFormField("media")
	if err != nil {
		return nil, fmt.Errorf("error creating a nested multipart form field for media: %w", err)
	}

	if err := writeNestedMultipart(nestedWriter, media, []string{"media", "thumbnail"}, []string{"media", "thumbnail"}); err != nil {
		return nil, fmt.Errorf("error writing media multipart section: %w", err)
	}

	httpResponse, err := b.send(http.MethodPost, MethodEditMessageMedia, writer.FormDataContentType(), requestPayload)
	if err != nil {
		return nil, fmt.Errorf("error sending POST request to Telegram method endpoint editMessageMedia in multipart format: %w", err)
	}
	defer httpResponse.Body.Close()

	var response struct {
		Ok          bool     `json:"ok"`
		Result      *Message `json:"result"`
		Description *string  `json:"description"`
		ErrorCode   *int     `json:"error_code"`
		// Parameters  ResponseParameters `json:"parameters"`
	}

	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if !response.Ok || response.ErrorCode != nil || response.Description != nil {
		errorMessage := fmt.Sprintf("Telegram API error: HTTP status \"%s\"", httpResponse.Status)
		if !response.Ok {
			errorMessage += ", request unsuccessful"
		} else {
			errorMessage += ", request successful"
		}
		if response.ErrorCode != nil {
			errorMessage += fmt.Sprintf(", Telegram error code \"%d\"", *response.ErrorCode)
		}
		if response.Description != nil {
			errorMessage += fmt.Sprintf(", description: %s", *response.Description)
		}
		return nil, errors.New(errorMessage)
	}

	return response.Result, nil
}

// EditInlineMessageMedia edits animation, audio, document, photo, or video inline messages, or adds media to text inline messages through the Telegram bot API.
//
// When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL.
//
// Required parameters:
//   - inlineMessageID: Identifier of the inline message.
//   - media: An object for a new media content of the message.
//
// # Note
//
// Business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
//
// See "editMessageMedia" https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditInlineMessageMedia(inlineMessageID string, media InputMedia, options *EditMessageOptions) error {
	requestPayload := new(bytes.Buffer)
	writer := multipart.NewWriter(requestPayload)
	if err := writer.WriteField("inline_message_id", inlineMessageID); err != nil {
		return fmt.Errorf("error adding inline_message_id field: %w", err)
	}
	if options != nil {
		if options.BusinessConnectionID != "" {
			if err := writer.WriteField("business_connection_id", options.BusinessConnectionID); err != nil {
				return fmt.Errorf("error adding business_connection_id field: %w", err)

			}
		}
		if options.ReplyMarkup != nil {
			jsonData, err := json.Marshal(options.ReplyMarkup)
			if err != nil {
				return fmt.Errorf("error JSON marshaling reply markup: %w", err)
			}
			if err := writer.WriteField("reply_markup", string(jsonData)); err != nil {
				return fmt.Errorf("error adding reply_markup field: %w", err)
			}
		}
	}

	nestedWriter, err := writer.CreateFormField("media")
	if err != nil {
		return fmt.Errorf("error creating a nested multipart form field for media: %w", err)
	}

	if err := writeNestedMultipart(nestedWriter, media, []string{"media", "thumbnail"}, []string{"media", "thumbnail"}); err != nil {
		return fmt.Errorf("error writing media multipart section: %w", err)
	}

	httpResponse, err := b.send(http.MethodPost, MethodEditMessageMedia, writer.FormDataContentType(), requestPayload)
	if err != nil {
		return fmt.Errorf("error sending POST request to Telegram method endpoint editMessageMedia in multipart format: %w", err)
	}
	defer httpResponse.Body.Close()

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
