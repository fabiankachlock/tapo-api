package api

import (
	"time"
)

const (
	RequestComponentNegotiation          = "component_nego"
	RequestHandshake                     = "handshake"
	RequestLoginDevice                   = "login_device"
	RequestSecurePassthrough             = "securePassthrough"
	RequestSetDeviceInfo                 = "set_device_info"
	RequestSetLightingEffect             = "set_lighting_effect"
	RequestGetDeviceInfo                 = "get_device_info"
	RequestGetDeviceUsage                = "get_device_usage"
	RequestGetEnergyUsage                = "get_energy_usage"
	RequestGetEnergyData                 = "get_energy_data"
	RequestGetCurrentPower               = "get_current_power"
	RequestGetChildDeviceList            = "get_child_device_list"
	RequestGetChildDeviceComponentList   = "get_child_device_component_list"
	RequestControlChild                  = "control_child"
	RequestMultiple                      = "multipleRequest"
	RequestGetTriggerLog                 = "get_trigger_logs"
	RequestGetTemperatureHumidityRecords = "get_temp_humidity_records"
)

var (
	EmptyParams = map[string]interface{}{}
)

type PlugDeviceInfoParams struct {
	On bool `json:"device_on"`
}

type GetEnergyDataParams struct {
	StartTimestamp uint64 `json:"start_timestamp"`
	EndTimestamp   uint64 `json:"end_timestamp"`
	Interval       uint64 `json:"interval"`
}

func GetEnergyDataParamsHourly(start time.Time, end time.Time) GetEnergyDataParams {
	return GetEnergyDataParams{
		StartTimestamp: uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix()),
		EndTimestamp:   uint64(time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 0, end.Location()).Unix()),
		Interval:       60,
	}
}

func GetEnergyDataParamsDaily(start time.Time) GetEnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return GetEnergyDataParams{
		StartTimestamp: ts,
		EndTimestamp:   ts,
		Interval:       1440,
	}
}

func GetEnergyDataParamsMonthly(start time.Time) GetEnergyDataParams {
	ts := uint64(time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).Unix())
	return GetEnergyDataParams{
		StartTimestamp: ts,
		EndTimestamp:   ts,
		Interval:       43200,
	}
}
