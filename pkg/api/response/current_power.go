package response

// CurrentPower holds information about the current power usage of the device.
type CurrentPower struct {
	// CurrentPower is the current power consumption of the device Watts (W).
	CurrentPower uint64 `json:"current_power"`
}
