package telegram

type SendOptions struct {
	// (Optional) Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// (Optional) Protects the contents of the sent message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`

	// (Optional) Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast bool `json:"allow_paid_broadcast,omitempty"`

	// (Optional) Unique identifier of the message effect to be added to the message; for private chats only.
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// (Optional) Description of the message to reply to.
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty"`

	// (Optional) Additional interface options.
	//
	// Types:
	//  - InlineKeyboardMarkup
	//  - ReplyKeyboardMarkup
	//  - ReplyKeyboardRemove
	//  - ForceReply
	ReplyMarkup any `json:"reply_markup,omitempty"`
}
