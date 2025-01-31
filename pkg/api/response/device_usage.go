package response

type DeviceUsage struct {
	// TimeUsage of the device in minutes.
	TimeUsage UsageByPeriod `json:"time_usage"`
}

// UsageByPeriod holds information about the usage of the device by a certain period.
type UsageByPeriod struct {
	Today  uint64 `json:"today"`
	Past7  uint64 `json:"past7"`
	Past30 uint64 `json:"past30"`
}
