package request

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
	RequestSupportedAlarmTypes           = "get_support_alarm_type_list"
	RequestPlayAlarm                     = "play_alarm"
	RequestStopAlarm                     = "stop_alarm"
)

var (
	EmptyParams = map[string]interface{}{}
)

type TapoRequest struct {
	Method       string                 `json:"method"`
	Params       map[string]interface{} `json:"params"`
	RequestTime  int64                  `json:"requestTimeMilis"`
	TerminalUUID string                 `json:"terminalUUID"`
}
