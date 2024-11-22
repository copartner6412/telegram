package telegram

// BusinessOpeningHoursInterval describes an interval of time during which a business is open.
//
// See "BusinessOpeningHoursInterval" https://core.telegram.org/bots/api#businessopeninghoursinterval
type BusinessOpeningHoursInterval struct {
	// (Required) The minute's sequence number in a week, starting on Monday, marking the start of the time interval during which the business is open; 0 - 7 * 24 * 60.
	OpeningMinute int `json:"opening_minute"`

	// (Required) The minute's sequence number in a week, starting on Monday, marking the end of the time interval during which the business is open; 0 - 8 * 24 * 60.
	ClosingMinute int `json:"closing_minute"`
}
