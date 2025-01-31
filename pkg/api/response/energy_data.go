package response

type EnergyData struct {
	// LocalTime is the local time of the device.
	LocalTime string `json:"local_time"`
	// Data is the energy data for the given interval in watt-hours (Wh).
	Data []uint64 `json:"data"`
	// Start timestamp of the interval in milliseconds. This value is provided
	// in the `get_energy_data` request and is passed through. Note that
	// it may not align with the returned data if the method is used
	// beyond its specified capabilities.
	StartTimestamp uint64 `json:"start_timestamp"`
	// End timestamp of the interval in milliseconds. This value is provided
	// in the `get_energy_data` request and is passed through. Note that
	// it may not align with the returned data for intervals other than hourly.
	EndTimestamp uint64 `json:"end_timestamp"`
	// Interval is the interval of the data in minutes.
	Interval uint64
}
