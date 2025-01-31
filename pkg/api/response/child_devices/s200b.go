package childdevices

type DeviceInfoS200B struct {
	DeviceInfoGenericChildDevice
}

type S200BRotationParams struct {
	RotationDegrees int16 `json:"rotation_deg"`
}

type LogEntryS200B struct {
	Event     S200BLogEvent       `json:"event"`
	Id        uint64              `json:"id"`
	Timestamp uint64              `json:"timestamp"`
	Params    S200BRotationParams `json:"params"`
}

type S200BLogEvent string

const (
	S200BLogEventRotation    T300LogEvent = "rotation"
	S200BLogEventSingleClick T300LogEvent = "singleClick"
	S200BLogEventDoubleCLick T300LogEvent = "doubleClick"
	S200BLogEventLowBattery  T300LogEvent = "lowBattery"
)
