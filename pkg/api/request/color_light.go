package request

type ColorLightDeviceInfoParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}
}

func NewColorLightDiveInfoParams() ColorLightDeviceInfoParams {
	return ColorLightDeviceInfoParams{
		jsonValue: map[string]interface{}{},
	}
}

func (c ColorLightDeviceInfoParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}

func (c ColorLightDeviceInfoParams) SetDeviceOn(deviceOn bool) ColorLightDeviceInfoParams {
	c.jsonValue["device_on"] = deviceOn
	return c
}

func (c ColorLightDeviceInfoParams) SetBrightness(brightness uint8) ColorLightDeviceInfoParams {
	c.jsonValue["brightness"] = brightness
	return c
}

func (c ColorLightDeviceInfoParams) SetHue(hue uint16) ColorLightDeviceInfoParams {
	c.jsonValue["hue"] = hue
	return c
}

func (c ColorLightDeviceInfoParams) SetSaturation(saturation uint16) ColorLightDeviceInfoParams {
	c.jsonValue["saturation"] = saturation
	return c
}

func (c ColorLightDeviceInfoParams) SetColorTemperature(colorTemperature uint16) ColorLightDeviceInfoParams {
	c.jsonValue["color_temp"] = colorTemperature
	return c
}
