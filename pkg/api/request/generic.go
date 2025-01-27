package request

type GenericDeviceInfoParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}
}

func NewGenericDeviceInfoParams() GenericDeviceInfoParams {
	return GenericDeviceInfoParams{
		jsonValue: map[string]interface{}{},
	}
}

func (c GenericDeviceInfoParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}

func (c GenericDeviceInfoParams) SetDeviceOn(deviceOn bool) GenericDeviceInfoParams {
	c.jsonValue["device_on"] = deviceOn
	return c
}
