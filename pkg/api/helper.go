package api

import "github.com/fabiankachlock/tapo-api/pkg/api/response"

func RequestVoid(d *ApiClient, method string, params interface{}) error {
	_, err := d.Request(method, params)
	return err
}

func RequestData[T any](d *ApiClient, method string, params interface{}) (T, error) {
	var zero T
	resp, err := d.Request(method, params)
	if err != nil {
		return zero, err
	}

	data, err := response.UnmarshalResponse[T](resp)
	if err != nil {
		return zero, err
	}
	return data.Result, data.GetError()
}
