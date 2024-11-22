package telegram

// ChatMemberRestricted represents a chat member that is under certain restrictions in the chat. Supergroups only.
//
// See "ChatMemberRestricted" https://core.telegram.org/bots/api#chatmemberrestricted
type ChatMemberRestricted struct {
	// (Required) The member's status in the chat, always “restricted”.
	Status string `json:"status"`

	// (Required) Information about the user.
	User User `json:"user"`

	// (Required) True, if the user is a member of the chat at the moment of the request.
	IsMember bool `json:"is_member"`

	// (Required) True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues.
	CanSendMessages bool `json:"can_send_messages"`

	// (Required) True, if the user is allowed to send audios.
	CanSendAudios bool `json:"can_send_audios"`

	// (Required) True, if the user is allowed to send documents.
	CanSendDocuments bool `json:"can_send_documents"`

	// (Required) True, if the user is allowed to send photos.
	CanSendPhotos bool `json:"can_send_photos"`

	// (Required) True, if the user is allowed to send videos.
	CanSendVideos bool `json:"can_send_videos"`

	// (Required) True, if the user is allowed to send video notes.
	CanSendVideoNotes bool `json:"can_send_video_notes"`

	// (Required) True, if the user is allowed to send voice notes.
	CanSendVoiceNotes bool `json:"can_send_voice_notes"`

	// (Required) True, if the user is allowed to send polls.
	CanSendPolls bool `json:"can_send_polls"`

	// (Required) True, if the user is allowed to send animations, games, stickers and use inline bots.
	CanSendOtherMessages bool `json:"can_send_other_messages"`

	// (Required) True, if the user is allowed to add web page previews to their messages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`

	// (Required) True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`

	// (Required) True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users"`

	// (Required) True, if the user is allowed to pin messages.
	CanPinMessages bool `json:"can_pin_messages"`

	// (Required) True, if the user is allowed to create forum topics.
	CanManageTopics bool `json:"can_manage_topics"`

	// (Required) Date when restrictions will be lifted for this user; Unix time. If 0, then the user is restricted forever.
	UntilDate int `json:"until_date"`
}

func (m ChatMemberRestricted) getMemberStatus() string {
	return m.Status
}
