package response

type DeviceInfoRgbLightStrip struct {
	DeviceInfoGeneric

	Brightness            uint8                     `json:"brightness"`
	ColorTemperatureRange []uint16                  `json:"color_temp_range"`
	ColorTemperature      uint16                    `json:"color_temp"`
	Hue                   uint16                    `json:"hue"`
	Saturation            uint16                    `json:"saturation"`
	DefaultStates         DefaultRgbLightStripState `json:"default_states"`
	Overheated            bool                      `json:"overheated"`
}

type DefaultRgbLightStripState struct {
	Type  DefaultStateType   `json:"type"`
	State RgbLightStripState `json:"state"`
}

type RgbLightStripState struct {
	Brightness       uint8  `json:"brightness"`
	Hue              uint16 `json:"hue"`
	Saturation       uint16 `json:"saturation"`
	ColorTemperature uint16 `json:"color_temp"`
}
