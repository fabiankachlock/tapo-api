package response

import "github.com/fabiankachlock/tapo-api/pkg/api/request"

type DeviceInfoRgbicLightStrip struct {
	DeviceInfoGeneric

	Brightness            uint8                       `json:"brightness"`
	ColorTemperatureRange []uint16                    `json:"color_temp_range"`
	ColorTemperature      uint16                      `json:"color_temp"`
	Hue                   uint16                      `json:"hue"`
	Saturation            uint16                      `json:"saturation"`
	DefaultStates         DefaultRgbicLightStripState `json:"default_states"`
	Overheated            bool                        `json:"overheated"`
}

type DefaultRgbicLightStripState struct {
	Type  DefaultStateType     `json:"type"`
	State RgbicLightStripState `json:"state"`
}

type RgbicLightStripState struct {
	Brightness       uint8                  `json:"brightness"`
	Hue              uint16                 `json:"hue"`
	Saturation       uint16                 `json:"saturation"`
	ColorTemperature uint16                 `json:"color_temp"`
	LightingEffect   request.LightingEffect `json:"lighting_effect"`
}
