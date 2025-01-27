package request

type LightDeviceInfoParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}
}

func NewLightDeviceInfoParams() LightDeviceInfoParams {
	return LightDeviceInfoParams{
		jsonValue: map[string]interface{}{},
	}
}

func (c LightDeviceInfoParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}

func (c LightDeviceInfoParams) SetDeviceOn(deviceOn bool) LightDeviceInfoParams {
	c.jsonValue["device_on"] = deviceOn
	return c
}

func (c LightDeviceInfoParams) SetBrightness(brightness uint8) LightDeviceInfoParams {
	c.jsonValue["brightness"] = brightness
	return c
}
