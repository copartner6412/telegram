package telegram

// BusinessOpeningHours describes the opening hours of a business.
//
// See "BusinessOpeningHours" https://core.telegram.org/bots/api#businessopeninghours
type BusinessOpeningHours struct {
	// (Required) Unique name of the time zone for which the opening hours are defined.
	TimeZoneName string `json:"time_zone_name"`

	// (Required) List of time intervals describing business opening hours.
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}
