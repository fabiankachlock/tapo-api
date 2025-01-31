package childdevices

type DeviceInfoT300 struct {
	DeviceInfoGenericChildDevice

	InAlarm         bool            `json:"in_alarm"`
	WaterLeakStatus WaterLeakStatus `json:"water_leak_status"`
}

type WaterLeakStatus string

const (
	WaterLeakStatusNormal    WaterLeakStatus = "normal"
	WaterLeakStatusWaterDry  WaterLeakStatus = "water_dry"
	WaterLeakStatusWaterLeak WaterLeakStatus = "water_leak"
)

type LogEntryT300 struct {
	Event     T300LogEvent `json:"event"`
	Id        uint64       `json:"id"`
	Timestamp uint64       `json:"timestamp"`
}

type T300LogEvent string

const (
	T300LogEventWaterDry  T300LogEvent = "waterDry"
	T300LogEventWaterLeak T300LogEvent = "waterLeak"
)
