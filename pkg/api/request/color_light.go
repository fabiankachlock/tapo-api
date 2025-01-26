package request

type ColorLightDiveInfoParams struct {
	DeviceOn         bool   `json:"device_on"`
	Brightness       uint8  `json:"brightness"`
	Hue              uint16 `json:"hue"`
	Saturation       uint16 `json:"saturation"`
	ColorTemperature uint16 `json:"color_temp"`
}
