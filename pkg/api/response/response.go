package response

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ResponseOk              = 0
	ErrorInvalidRequest     = -1002
	ErrorMalformedRequest   = -1003
	ErrorInvalidPublicKey   = -1010
	ErrorInvalidCredentials = -1501
	ErrorSessionTimeOut     = 9999
)

var (
	ErrNonSuccessfulResponse = errors.New("non successful response")
)

type TapoResponse[T any] struct {
	Result    T   `json:"result"`
	ErrorCode int `json:"error_code"`
}

func (r TapoResponse[T]) IsOk() bool {
	return r.ErrorCode == ResponseOk
}

func (r TapoResponse[T]) HasError() bool {
	return r.ErrorCode != ResponseOk
}

func (r TapoResponse[T]) GetError() error {
	if r.ErrorCode == ResponseOk {
		return nil
	}
	return fmt.Errorf("%w: error code: %d", ErrNonSuccessfulResponse, r.ErrorCode)
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
type EmptyResponse struct{}
