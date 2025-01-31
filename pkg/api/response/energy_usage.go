package response

// EnergyUsage holds information about the energy consumption of the device.
type EnergyUsage struct {
	// LocalTime is the local time of the device.
	LocalTime string `json:"local_time"`
	// CurrentPower is the current power consumption of the device in milliwatts (mW).
	CurrentPower uint64 `json:"current_power"`
	// TodayRuntime is the runtime of the device today in minutes.
	TodayRuntime uint64 `json:"today_runtime"`
	// TodayEnergy is the energy consumption of the device today in watt-hours (Wh).
	TodayEnergy uint64 `json:"today_energy"`
	// MonthRuntime is the runtime of the device this month in minutes.
	MonthRuntime uint64 `json:"month_runtime"`
	// MonthEnergy is the energy consumption of the device this month in watt-hours (Wh).
	MonthEnergy uint64 `json:"month_energy"`
}
