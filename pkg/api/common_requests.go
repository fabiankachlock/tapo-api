package api

import (
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// GetSupportedAlarmTypes returns a list of supported alarm types by the device.
func GetSupportedAlarmTypes(c *ApiClient) (response.SupportedAlarmTypeList, error) {
	value := response.SupportedAlarmTypeList{}
	err := c.RequestData(request.RequestSupportedAlarmTypes, request.EmptyParams, true, &value)
	if err != nil {
		return response.SupportedAlarmTypeList{}, err
	}
	return value, nil
}

// PlayAlarm plays an alarm on the device.
func PlayAlarm(c *ApiClient, params request.PlayAlarmParams) error {
	return c.RequestVoid(request.RequestPlayAlarm, params.GetJsonValue(), true)
}

// StopAlarm stops the playing alarm on the device.
func StopAlarm(c *ApiClient) error {
	return c.RequestVoid(request.RequestStopAlarm, request.EmptyParams, true)
}

// ResetDevice resets the device.
func ResetDevice(c *ApiClient) error {
	return c.RequestVoid(request.RequestDeviceReset, request.EmptyParams, true)
}

// GetDeviceInfo returns core information about the device like the device name, model, etc.
func GetDeviceInfo[T any](c *ApiClient) (T, error) {
	var value T
	err := c.RequestData(request.RequestGetDeviceInfo, request.EmptyParams, true, &value)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetDeviceUsage returns data about the usage of the device like on/off times or power usage.
func GetDeviceUsage[T any](c *ApiClient) (T, error) {
	var value T
	err := c.RequestData(request.RequestGetDeviceUsage, request.EmptyParams, true, &value)
	if err != nil {
		return value, err
	}
	return value, nil
}

// GetEnergyUsage returns data about the energy consumption of the device.
func GetEnergyUsage(c *ApiClient) (response.EnergyUsage, error) {
	value := response.EnergyUsage{}
	err := c.RequestData(request.RequestGetEnergyUsage, request.EmptyParams, true, &value)
	if err != nil {
		return response.EnergyUsage{}, err
	}
	return value, nil
}

// GetEnergyData returns data about the energy consumption of the device for a certain period.
func GetEnergyData(c *ApiClient, params request.EnergyDataParams) (response.EnergyData, error) {
	value := response.EnergyData{}
	err := c.RequestData(request.RequestGetEnergyData, params.GetJsonValue(), true, &value)
	if err != nil {
		return response.EnergyData{}, err
	}
	return value, nil
}

// GetCurrentPower returns the current power consumption of the device.
func GetCurrentPower(c *ApiClient) (response.CurrentPower, error) {
	value := response.CurrentPower{}
	err := c.RequestData(request.RequestGetCurrentPower, request.EmptyParams, true, &value)
	if err != nil {
		return response.CurrentPower{}, err
	}
	return value, nil
}
