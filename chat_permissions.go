package telegram

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
//
// See "ChatPermissions" https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	// (Optional) True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues.
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// (Optional) True, if the user is allowed to send audios.
	CanSendAudios bool `json:"can_send_audios,omitempty"`

	// (Optional) True, if the user is allowed to send documents.
	CanSendDocuments bool `json:"can_send_documents,omitempty"`

	// (Optional) True, if the user is allowed to send photos.
	CanSendPhotos bool `json:"can_send_photos,omitempty"`

	// (Optional) True, if the user is allowed to send videos.
	CanSendVideos bool `json:"can_send_videos,omitempty"`

	// (Optional) True, if the user is allowed to send video notes.
	CanSendVideoNotes bool `json:"can_send_video_notes,omitempty"`

	// (Optional) True, if the user is allowed to send voice notes.
	CanSendVoiceNotes bool `json:"can_send_voice_notes,omitempty"`

	// (Optional) True, if the user is allowed to send polls.
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// (Optional) True, if the user is allowed to send animations, games, stickers and use inline bots.
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// (Optional) True, if the user is allowed to add web page previews to their messages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// (Optional) True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups.
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// (Optional) True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// (Optional) True, if the user is allowed to pin messages. Ignored in public supergroups.
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// (Optional) True, if the user is allowed to create forum topics. If omitted defaults to the value of can_pin_messages.
	CanManageTopics bool `json:"can_manage_topics,omitempty"`
}
