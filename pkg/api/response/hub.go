package response

import "encoding/json"

// DeviceUsageEnergyMonitor holds information about plug devices.
type DeviceInfoHub struct {
	DeviceInfoGeneric

	InAlarmSource string `json:"in_alarm_source"`
	InAlarm       bool   `json:"in_alarm"`
	Overheated    bool   `json:"overheated"`
}

type SupportedAlarmTypeList struct {
	AlarmTypes []string `json:"alarm_type_list"`
}

type ChildDeviceList struct {
	Devices    []json.RawMessage `json:"child_device_list"`
	StartIndex uint16            `json:"start_index"`
	Sum        uint16            `json:"sum"`
}
