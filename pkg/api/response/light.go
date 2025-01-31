package response

type DeviceInfoLight struct {
	DeviceInfoGeneric

	Brightness    uint8             `json:"brightness"`
	DefaultStates DefaultLightState `json:"default_states"`
	Overheated    bool              `json:"overheated"`
}

type DefaultLightState struct {
	Brightness  uint8            `json:"brightness"`
	RePowerType DefaultPowerType `json:"re_power_type"`
}
