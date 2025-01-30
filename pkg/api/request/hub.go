package request

type AlarmVolume string

const (
	AlarmVolumeMute   AlarmVolume = "mute"
	AlarmVolumeLow    AlarmVolume = "low"
	AlarmVolumeNormal AlarmVolume = "normal"
	AlarmVolumeHigh   AlarmVolume = "high"
)

type PlayAlarmParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}

	Duration int         `json:"alarm_duration"`
	Volume   AlarmVolume `json:"alarm_volume"`
	Type     string      `json:"alarm_type"`
}

func NewPlayAlarmParams(alarmType string) PlayAlarmParams {
	return PlayAlarmParams{
		jsonValue: map[string]interface{}{
			"alarm_type":     alarmType,
			"alarm_duration": 0, // by default, play once
			"alarm_volume":   AlarmVolumeNormal,
		},
	}
}

func (c PlayAlarmParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}

func (p PlayAlarmParams) SetType(alarmType string) PlayAlarmParams {
	p.jsonValue["alarm_type"] = alarmType
	return p
}

func (p PlayAlarmParams) SetVolume(volume AlarmVolume) PlayAlarmParams {
	p.jsonValue["alarm_volume"] = volume
	return p
}

func (p PlayAlarmParams) PlanOnce() PlayAlarmParams {
	p.jsonValue["alarm_duration"] = 0
	return p
}

func (p PlayAlarmParams) PlanFor(duration int) PlayAlarmParams {
	p.jsonValue["alarm_duration"] = duration
	return p
}

func (p PlayAlarmParams) PlayContinuous() PlayAlarmParams {
	delete(p.jsonValue, "alarm_duration")
	return p
}
