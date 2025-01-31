package request

type ControlChildParams struct {
	// jsonValue is the map that will be marshaled into the JSON body of the request.
	// A map must be used explicitly, because otherwise there is no way of differentiating
	// between an empty value and a value that was not set.
	jsonValue map[string]interface{}
}

func NewControlChildParams(deviceId string, request TapoRequest) ControlChildParams {
	return ControlChildParams{
		jsonValue: map[string]interface{}{
			"device_id":   deviceId,
			"requestData": request,
		},
	}
}
