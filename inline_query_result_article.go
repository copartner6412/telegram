package telegram

// InlineQueryResultArticle represents a link to an article or web page.
//
// The type of the result must be "article".
//
// Required fields:
//   - Type
//   - ID
//   - Title
//   - InputMessageContent
//
// See "InlineQueryResultArticle" https://core.telegram.org/bots/api#inlinequeryresultarticle
type InlineQueryResultArticle struct {
	// (Required) Type of the result, must be "article".
	Type string `json:"type"`

	// (Required) Unique identifier for this result, 1-64 bytes.
	ID string `json:"id"`

	// (Required) Title of the result.
	Title string `json:"title"`

	// (Required) Content of the message to be sent.
	InputMessageContent InputMessageContent `json:"input_message_content"`

	// (Optional) Inline keyboard attached to the message.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// (Optional) URL of the result.
	URL string `json:"url,omitempty"`

	// (Optional) Pass True if you don't want the URL to be shown in the message.
	HideURL bool `json:"hide_url,omitempty"`

	// (Optional) Short description of the result.
	Description string `json:"description,omitempty"`

	// (Optional) Url of the thumbnail for the result.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`

	// (Optional) Thumbnail width.
	ThumbnailWidth int `json:"thumbnail_width,omitempty"`

	// (Optional) Thumbnail height.
	ThumbnailHeight int `json:"thumbnail_height,omitempty"`
}

func (r InlineQueryResultArticle) getQueryType() string {
	return r.Type
}
