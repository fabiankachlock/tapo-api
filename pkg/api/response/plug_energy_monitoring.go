package response

// DeviceUsageEnergyMonitor holds information about plug devices.
type DeviceInfoPlugEnergyMonitoring struct {
	DeviceInfoGeneric

	DefaultStates         DefaultPlugState      `json:"default_states"`
	OverheatStatus        OverheatStatus        `json:"overheat_status"`
	OvercurrentStatus     OvercurrentStatus     `json:"overcurrent_status"`
	PowerProtectionStatus PowerProtectionStatus `json:"power_protection_status"`
}

// DefaultPlugState holds information about the default state of the plug device.
type DefaultPlugState struct {
	Type  DefaultStateType `json:"type"`
	State PlugState        `json:"state"`
}

// PlugState holds information about the current state of the plug device.
type PlugState struct {
	On bool `json:"on"`
}
