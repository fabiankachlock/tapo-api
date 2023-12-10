package api

import "encoding/json"

const (
	ResponseOk              = 0
	ErrorInvalidRequest     = -1002
	ErrorMalformedRequest   = -1003
	ErrorInvalidPublicKey   = -1010
	ErrorInvalidCredentials = -1501
	ErrorSessionTimeOut     = 9999
)

type TapoResponse[T any] struct {
	Result    T   `json:"result"`
	ErrorCode int `json:"error_code"`
}

func UnmarshalResponse[T any](data []byte) (TapoResponse[T], error) {
	jsonData := TapoResponse[T]{}
	err := json.Unmarshal(data, &jsonData)
	return jsonData, err
}

type DeviceInfoPlug struct {
	Type               string           `json:"type"`
	DeviceId           string           `json:"device_id"`
	Model              string           `json:"model"`
	HWId               string           `json:"hw_id"`
	HWVersion          string           `json:"hw_ver"`
	FWId               string           `json:"fw_id"`
	FWVersion          string           `json:"fw_ver"`
	OEMId              string           `json:"oem_id"`
	Mac                string           `json:"mac"`
	Ip                 string           `json:"ip"`
	SSID               string           `json:"ssid"`
	SignalLevel        uint8            `json:"signal_level"`
	RSSI               int16            `json:"rssi"`
	Specs              string           `json:"specs"`
	Lang               string           `json:"lang"`
	DeviceOn           bool             `json:"device_on"`
	OnTime             uint64           `json:"on_time"`
	Overheated         bool             `json:"overheated"`
	Nickname           string           `json:"nickname"`
	Avatar             string           `json:"avatar"`
	HasSetLocationInfo bool             `json:"has_set_location_info"`
	Region             string           `json:"region"`
	Latitude           int64            `json:"latitude"`
	Longitude          int64            `json:"longitude"`
	TimeDiff           int64            `json:"time_diff"`
	DefaultStates      DefaultPlugState `json:"default_states"`
}

type DefaultPlugState struct {
	Type  string    `json:"type"`
	State PlugState `json:"state"`
}

type PlugState struct {
	On bool `json:"on"`
}

const (
	DefaultPlugStateCustom     = "custom"
	DefaultPlugStateLastStates = "last_states"
)

type DeviceUsageEnergyMonitor struct {
	TimeUsage  UsageByPeriod `json:"time_usage"`
	PowerUsage UsageByPeriod `json:"power_usage"`
	SavedPower UsageByPeriod `json:"saved_power"`
}

type UsageByPeriod struct {
	Today  uint64 `json:"today"`
	Past7  uint64 `json:"past7"`
	Past30 uint64 `json:"past30"`
}

type EnergyUsage struct {
	LocalTime    string `json:"local_time"`
	CurrentPower uint64 `json:"current_power"`
	TodayRuntime uint64 `json:"today_runtime"`
	TodayEnergy  uint64 `json:"today_energy"`
	MonthRuntime uint64 `json:"month_runtime"`
	MonthEnergy  uint64 `json:"month_energy"`
}

type CurrentPower struct {
	CurrentPower uint64 `json:"current_power"`
}
