package request

type ChildDeviceListParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}
}

func NewChildDeviceListParams(startIndex uint16) ChildDeviceListParams {
	return ChildDeviceListParams{
		jsonValue: map[string]interface{}{
			"start_index": startIndex,
		},
	}
}

func (c ChildDeviceListParams) GetJsonValue() map[string]interface{} {
	return c.jsonValue
}
