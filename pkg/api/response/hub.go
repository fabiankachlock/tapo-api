package response

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
