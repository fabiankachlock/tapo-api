package request

type PlayAlarmParams struct {
	Duration int         `json:"alarm_duration"`
	Volume   AlarmVolume `json:"alarm_volume"`
	Type     string      `json:"alarm_type"`
}

type AlarmVolume string

const (
	AlarmVolumeLow    AlarmVolume = "low"
	AlarmVolumeMedium AlarmVolume = "normal"
	AlarmVolumeHigh   AlarmVolume = "high"
)
