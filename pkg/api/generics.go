package api

// TODO: add withToken to method signature
func RequestVoid(d *ApiClient, method string, params interface{}) error {
	_, err := d.Request(method, params, true)
	return err
}

// TODO: add withToken to method signature
func RequestData[T any](d *ApiClient, method string, params interface{}) (T, error) {
	var value T
	resp, err := d.Request(method, params, true)
	if err != nil {
		return value, err
	}

	tapoResponse, err := resp.UnmarshalResponse(&value)
	if err != nil {
		return value, err
	}

	if tapoResponse.HasError() {
		return value, tapoResponse.GetError()
	}
	return value, nil
}
