package childdevices

type DeviceInfoPowerStripPlug struct {
	DeviceId           string `json:"device_id"`
	Type               string `json:"type"`
	Model              string `json:"model"`
	HWId               string `json:"hw_id"`
	HWVersion          string `json:"hw_ver"`
	FWId               string `json:"fw_id"`
	FWVersion          string `json:"fw_ver"`
	OEMId              string `json:"oem_id"`
	Mac                string `json:"mac"`
	DeviceOn           bool   `json:"device_on"`
	Nickname           string `json:"nickname"`
	Avatar             string `json:"avatar"`
	HasSetLocationInfo bool   `json:"has_set_location_info"`
	Region             string `json:"region"`
	Latitude           int64  `json:"latitude"`
	Longitude          int64  `json:"longitude"`
	TimeDiff           int64  `json:"time_diff"`
	// The time in seconds this device has been ON since the last state change (On/Off).
	OnTime            uint64         `json:"on_time"`
	StatusFollowEdge  bool           `json:"status_follow_edge"`
	OriginalDeviceId  string         `json:"original_device_id"`
	Category          string         `json:"category"`
	BindCount         uint8          `json:"bind_count"`
	AutoOffRemainTime uint64         `json:"auto_off_remain_time"`
	AutoOffStatus     AutoOffStatus  `json:"auto_off_status"`
	OverheatStatus    OverheatStatus `json:"overheat_status"`
	Position          uint8          `json:"position"`
	SlotNumber        uint8          `json:"slot_number"`
}

type AutoOffStatus string

const (
	AutoOffStatusOn  AutoOffStatus = "on"
	AutoOffStatusOff AutoOffStatus = "off"
)

type OverheatStatus string

const (
	OverheatStatusCoolDown   OverheatStatus = "cool_down"
	OverheatStatusNormal     OverheatStatus = "normal"
	OverheatStatusOverheated OverheatStatus = "overheated"
)
