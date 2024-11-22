package telegram

import "net/http"

// SendDocumentRequest represents the parameters for sending a document message using sendDocument method through Telegram bot API.
//
// Required fields:
//   - ChatID
//   - Document
//
// See "sendDocument" https://core.telegram.org/bots/api#senddocument
type SendDocumentRequest struct {
	// (Required) Integer or String. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) File to send.
	//  1. Pass a file_id as a string with "file_id:" prefix to send a file that exists on the Telegram servers (recommended)
	//  2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a file from the Internet
	//  3. Pass an absolute path to upload a new one using multipart/form-data.
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Document string `json:"document"`

	*SendDocumentOptions
}

// SendDocumentOptions represents the optional parameters for sending a document message using sendDocument method through Telegram bot API.
//
// See "sendDocument" https://core.telegram.org/bots/api#senddocument
type SendDocumentOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the message will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread (topic) of the forum; for forum supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// (Optional) Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can  only be uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//
	// More information on Sending Files » https://core.telegram.org/bots/api#sending-files
	Thumbnail string `json:"thumbnail,omitempty"`

	*Caption

	// (Optional) Disables automatic server-side content type detection for files uploaded using multipart/form-data.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`

	*SendOptions
}

// SendDocument sends a general file to a chat.
//
// Required options:
//   - toChatID: Unique identifier for the target chat or username of the target channel (in the format @channelusername).
//   - document: File to send. This method supports files up to 50 MB in size.
//     1. Pass a file_id as a string with "file_id:" prefix to send a file that exists on the Telegram servers (recommended)
//     2. Pass an HTTP URL as a string with "http://" or "https://" prefix for Telegram to get a file from the Internet
//     3. Pass an absolute path to upload a new one using multipart/form-data.
//
// On success, the sent Message is returned.
//
// See https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(toChatID any, document string, options *SendDocumentOptions) (*Message, error) {
	request := SendDocumentRequest{
		ChatID:              toChatID,
		Document:            document,
		SendDocumentOptions: options,
	}
	result := new(Message)
	if err := b.sendMultipartForResult(result, http.MethodPost, MethodSendDocument, request, []string{"document", "thumbnail"}, []string{"thumbnail"}); err != nil {
		return nil, err
	}

	return result, nil
}
