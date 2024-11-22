package telegram

// WebAppInfo describes a Web App.
//
// See "WebAppInfo" https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	// (Required) An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
	URL string `json:"url"`
}
