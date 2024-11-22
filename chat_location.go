package telegram

// ChatLocation represents a location to which a chat is connected.
//
// See "ChatLocation" https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	// (Required) The location to which the supergroup is connected. Can't be a live location.
	Location Location `json:"location"`

	// (Required) Location address; 1-64 characters, as defined by the chat owner.
	Address string `json:"address"`
}
