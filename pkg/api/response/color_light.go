package response

type DeviceInfoColorLight struct {
	DeviceInfoGeneric

	Brightness       uint8                  `json:"brightness"`
	ColorTemperature uint16                 `json:"color_temp"`
	Hue              uint16                 `json:"hue"`
	Saturation       uint16                 `json:"saturation"`
	DefaultStates    DefaultColorLightState `json:"default_states"`
	Overheated       bool                   `json:"overheated"`
}

type DefaultColorLightState struct {
	Type  DefaultStateType   `json:"type"`
	State RgbLightStripState `json:"state"`
}

type ColorLightState struct {
	Brightness       uint8  `json:"brightness"`
	Hue              uint16 `json:"hue"`
	Saturation       uint16 `json:"saturation"`
	ColorTemperature uint16 `json:"color_temp"`
}
