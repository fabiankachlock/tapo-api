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

func UnmarshalResponse[T any](data []byte) (TapoResponse[T], error) {
	jsonData := TapoResponse[T]{}
	err := json.Unmarshal(data, &jsonData)
	return jsonData, err
}

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
