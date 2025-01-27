package response

type DeviceInfoGeneric struct {
	DeviceId           string `json:"device_id"`
	Type               string `json:"type"`
	Model              string `json:"model"`
	HWId               string `json:"hw_id"`
	HWVersion          string `json:"hw_ver"`
	FWId               string `json:"fw_id"`
	FWVersion          string `json:"fw_ver"`
	OEMId              string `json:"oem_id"`
	Mac                string `json:"mac"`
	Ip                 string `json:"ip"`
	SSID               string `json:"ssid"`
	SignalLevel        uint8  `json:"signal_level"`
	RSSI               int16  `json:"rssi"`
	Specs              string `json:"specs"`
	Lang               string `json:"lang"`
	DeviceOn           bool   `json:"device_on"`
	Nickname           string `json:"nickname"`
	Avatar             string `json:"avatar"`
	HasSetLocationInfo bool   `json:"has_set_location_info"`
	Region             string `json:"region"`
	Latitude           int64  `json:"latitude"`
	Longitude          int64  `json:"longitude"`
	TimeDiff           int64  `json:"time_diff"`
}
