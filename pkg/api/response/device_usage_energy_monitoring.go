package response

type DeviceUsageEnergyMonitoring struct {
	// TimeUsage of the device in minutes.
	TimeUsage UsageByPeriod `json:"time_usage"`
	// PowerUsage of the device in watt-hours (Wh).
	PowerUsage UsageByPeriod `json:"power_usage"`
	// SavedPower of the device in watt-hours (Wh).
	SavedPower UsageByPeriod `json:"saved_power"`
}
