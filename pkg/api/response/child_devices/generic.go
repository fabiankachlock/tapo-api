package childdevices

type DeviceInfoGenericChildDevice struct {
	// Common properties to all Hub child devices.
	AtLowBattery            bool       `json:"at_low_battery"`
	Avatar                  string     `json:"avatar"`
	BindCount               uint32     `json:"bind_count"`
	Category                string     `json:"category"`
	DeviceId                string     `json:"device_id"`
	FWVersion               string     `json:"fw_ver"`
	HWId                    string     `json:"hw_id"`
	HWVersion               string     `json:"hw_ver"`
	JammingRSSI             int16      `json:"jamming_rssi"`
	JammingSignalLevel      uint8      `json:"jamming_signal_level"`
	Mac                     string     `json:"mac"`
	Nickname                string     `json:"nickname"`
	OemId                   string     `json:"oem_id"`
	ParentDeviceId          string     `json:"parent_device_id"`
	Region                  string     `json:"region"`
	RSSI                    int16      `json:"rssi"`
	SignalLevel             uint8      `json:"signal_level"`
	Specs                   string     `json:"specs"`
	Status                  StatusType `json:"status"`
	Type                    string     `json:"type"`
	Model                   string     `json:"model"`
	LastOnboardingTimestamp uint64     `json:"lastOnboardingTimestamp"`
	// The time in seconds between each report.
	ReportInterval   uint32 `json:"report_interval"`
	StatusFollowEdge bool   `json:"status_follow_edge"`
}

type StatusType string

const (
	StatusTypeOffline StatusType = "offline"
	StatusTypeOnline  StatusType = "online"
)
