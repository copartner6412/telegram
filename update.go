package telegram

// Update represents an incoming update from Telegram.
//
// At most one of the optional parameters can be present in any given update.
//
// Required fields:
//   - UpdateID
//
// See "Update" https://core.telegram.org/bots/api#update
type Update struct {
	// (Required) The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially.
	// This identifier becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore
	// the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of
	// the next update will be chosen randomly instead of sequentially.
	UpdateID int `json:"update_id"`

	// (Optional) New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message"`

	// (Optional) New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to
	// message fields that are either unavailable or not actively used by your bot.
	EditedMessage *Message `json:"edited_message"`

	// (Optional) New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post"`

	// (Optional) New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes
	// to message fields that are either unavailable or not actively used by your bot.
	EditedChannelPost *Message `json:"edited_channel_post"`

	// (Optional) The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot.
	BusinessConnection *BusinessConnection `json:"business_connection"`

	// (Optional) New message from a connected business account.
	BusinessMessage *Message `json:"business_message"`

	// (Optional) New version of a message from a connected business account.
	EditedBusinessMessage *Message `json:"edited_business_message"`

	// (Optional) Messages were deleted from a connected business account.
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages"`

	// (Optional) A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly
	// specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	MessageReaction *MessageReactionUpdated `json:"message_reaction"`

	// (Optional) Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must
	// explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and
	// can be sent with delay up to a few minutes.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count"`

	// (Optional) New incoming inline query.
	InlineQuery *InlineQuery `json:"inline_query"`

	// (Optional) The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation
	// on the feedback collecting for details on how to enable these updates for your bot. https://core.telegram.org/bots/inline#collecting-feedback
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`

	// (Optional) New incoming callback query.
	CallbackQuery *CallbackQuery `json:"callback_query"`

	// (Optional) New incoming shipping query. Only for invoices with flexible price.
	ShippingQuery *ShippingQuery `json:"shipping_query"`

	// (Optional) New incoming pre-checkout query. Contains full information about checkout.
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query"`

	// (Optional) A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat.
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media"`

	// (Optional) New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot.
	Poll *Poll `json:"poll"`

	// (Optional) A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer"`

	// (Optional) The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is
	// blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member"`

	// (Optional) A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify
	// "chat_member" in the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member"`

	// (Optional) A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request"`

	// (Optional) A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	ChatBoost *ChatBoostUpdated `json:"chat_boost"`

	// (Optional) A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost"`
}
