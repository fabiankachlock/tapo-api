package childdevices

type DeviceInfoT300 struct {
	// Common properties to all Hub child devices.
	AtLowBattery       bool   `json:"at_low_battery"`
	Avatar             string `json:"avatar"`
	BindCOunt          uint32 `json:"bind_count"`
	Category           string `json:"category"`
	DeviceId           string `json:"device_id"`
	FWVersion          string `json:"fw_ver"`
	HWId               string `json:"hw_id"`
	HWVersion          string `json:"hw_ver"`
	JammingRSSI        int16  `json:"jamming_rssi"`
	JammingSignalLevel uint8  `json:"jamming_signal_level"`
	Mac                string `json:"mac"`
	Nickname           string `json:"nickname"`
	OemId              string `json:"oem_id"`
	ParentDeviceId     string `json:"parent_device_id"`
	Region             string `json:"region"`
	RSSI               int16  `json:"rssi"`
	SignalLevel        uint8  `json:"signal_level"`
	Specs              string `json:"specs"`
	Status             string `json:"status"`
	Type               string `json:"type"`

	// Specific properties to this device.
	InAlarm bool `json:"in_alarm"`

	LastOnboardingTimestamp uint64 `json:"lastOnboardingTimestamp"`

	// The time in seconds between each report.
	ReportInterval   uint32          `json:"report_interval"`
	StatusFollowEdge bool            `json:"status_follow_edge"`
	WaterLeakStatus  WaterLeakStatus `json:"water_leak_status"`
}

type WaterLeakStatus string

const (
	WaterLeakStatusNormal    WaterLeakStatus = "normal"
	WaterLeakStatusWaterDry  WaterLeakStatus = "water_dry"
	WaterLeakStatusWaterLeak WaterLeakStatus = "water_leak"
)
