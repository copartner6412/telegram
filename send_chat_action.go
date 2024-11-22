package telegram

import "net/http"

// SendChatActionRequest represents a request to send a chat action to inform the user that something is happening on the bot's side.
//
// Required fields:
//   - ChatID
//   - Action
//
// See https://core.telegram.org/bots/api#sendchataction
type SendChatActionRequest struct {
	// (Required) Integer or string. Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID any `json:"chat_id"`

	// (Required) Type of action to broadcast. Choose one, depending on what the user is about to receive.
	//
	// Available options:
	//  - typing for text messages
	//  - upload_photo for photos
	//  - record_video or upload_video for videos
	//  - record_voice or upload_voice for voice notes
	//  - upload_document for general files
	//  - choose_sticker for stickers
	//  - find_location for location data
	//  - record_video_note or upload_video_note for video notes
	Action string `json:"action"`

	*SendChatActionOptions
}

type SendChatActionOptions struct {
	// (Optional) Unique identifier of the business connection on behalf of which the action will be sent.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// (Optional) Unique identifier for the target message thread; for supergroups only.
	MessageThreadID int `json:"message_thread_id,omitempty"`
}

// SendChatAction informs the user that something is happening on the bot's side.
// The status is set for 5 seconds or less.
//
// # Example
//
// The ImageBot needs some time to process a request and upload the image.
// Instead of sending a text message along the lines of “Retrieving image, please wait…”,
// the bot may use sendChatAction with action = upload_photo.
// The user will see a “sending photo” status for the bot.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
//
// See https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(toChatID any, Action string, options *SendChatActionOptions) error {
	request := SendChatActionRequest{
		ChatID:                toChatID,
		Action:                Action,
		SendChatActionOptions: options,
	}
	if err := b.sendJsonForBool(http.MethodPost, MethodSendChatAction, request); err != nil {
		return err
	}
	return nil
}
