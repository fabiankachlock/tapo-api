package childdevices

type DeviceInfoT100 struct {
	DeviceInfoGenericChildDevice

	Detected bool `json:"open"`
}

type LogEntryT100 struct {
	Event     T100LogEvent `json:"event"`
	Id        uint64       `json:"id"`
	Timestamp uint64       `json:"timestamp"`
}

type T100LogEvent string

const (
	T100LogEventMotion T300LogEvent = "motion"
)
