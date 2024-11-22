package telegram

import (
	"net/http"
)

// GetUpdatesRequest represents a request to receive incoming updates using long polling.
//
// See "getUpdates" https://core.telegram.org/bots/api#getupdates
type GetUpdatesRequest struct {
	// (Optional) Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called
	// with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue.
	// All previous updates will be forgotten.
	Offset int `json:"offset,omitempty"`

	// (Optional) Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`

	// (Optional) Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	Timeout int `json:"timeout,omitempty"`

	// (Optional) A list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	//
	// Available options:
	//  - message
	//  - edited_message
	//  - channel_post
	//  - edited_channel_post
	//  - business_connection
	//  - business_message
	//  - edited_business_message
	//  - deleted_business_messages
	//  - message_reaction
	//  - message_reaction_count
	//  - inline_query
	//  - chosen_inline_result
	//  - callback_query
	//  - shipping_query
	//  - pre_checkout_query
	//  - purchased_paid_media
	//  - poll
	//  - poll_answer
	//  - my_chat_member
	//  - chat_member
	//  - chat_join_request
	//  - chat_boost
	//  - removed_chat_boost
	// Specify an empty list to receive all update types
	// except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// GetUpdates sends a request to the Telegram API to retrieve incoming updates using long polling using the getUpdates method through the Telegram bot API.
//
// Parameters:
//   - offset: Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
//   - limit: Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
//   - timeout: Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
//   - allowedUpdates: A list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
//
// Returns an Array of Update objects.
//
// # Note
//  1. This method will not work if an outgoing webhook is set up.
//  2. In order to avoid getting duplicate updates, recalculate offset after each server response.
//
// See "getUpdates" https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(offset, limit, timeout int, allowedUpdates []string) ([]Update, error) {
	request := GetUpdatesRequest{
		Offset:         offset,
		Limit:          limit,
		Timeout:        timeout,
		AllowedUpdates: allowedUpdates,
	}
	result := new([]Update)
	if err := b.sendJsonForResult(result, http.MethodGet, MethodGetUpdates, request); err != nil {
		return nil, err
	}
	return *result, nil
}
