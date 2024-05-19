package response

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

// UnmarshalResponse unmarshals the response from the Tapo API.
func UnmarshalResponse[T any](data []byte) (TapoResponse[T], error) {
	jsonData := TapoResponse[T]{}
	err := json.Unmarshal(data, &jsonData)
	return jsonData, err
}

// DeviceInfoPlug holds information about the device usage.
type DeviceUsageEnergyMonitor struct {
	TimeUsage  UsageByPeriod `json:"time_usage"`
	PowerUsage UsageByPeriod `json:"power_usage"`
	SavedPower UsageByPeriod `json:"saved_power"`
}

// UsageByPeriod holds information about the usage of the device by a certain period.
type UsageByPeriod struct {
	Today  uint64 `json:"today"`
	Past7  uint64 `json:"past7"`
	Past30 uint64 `json:"past30"`
}

// DeviceInfoPlug holds information about the energy usage of the device.
type EnergyUsage struct {
	LocalTime    string `json:"local_time"`
	CurrentPower uint64 `json:"current_power"`
	TodayRuntime uint64 `json:"today_runtime"`
	TodayEnergy  uint64 `json:"today_energy"`
	MonthRuntime uint64 `json:"month_runtime"`
	MonthEnergy  uint64 `json:"month_energy"`
}

// CurrentPower holds information about the current power usage of the device.
type CurrentPower struct {
	CurrentPower uint64 `json:"current_power"`
}
