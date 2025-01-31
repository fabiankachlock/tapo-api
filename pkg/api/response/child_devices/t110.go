package childdevices

type DeviceInfoT110 struct {
	DeviceInfoGenericChildDevice

	Open bool `json:"open"`
}

type LogEntryT110 struct {
	Event     T110LogEvent `json:"event"`
	Id        uint64       `json:"id"`
	Timestamp uint64       `json:"timestamp"`
}

type T110LogEvent string

const (
	T110LogEventClose    T300LogEvent = "close"
	T110LogEventOpen     T300LogEvent = "open"
	T110LogEventKeepOpen T300LogEvent = "keepOpen"
)
