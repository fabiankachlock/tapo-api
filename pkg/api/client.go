package api

import (
	"encoding/json"
	"fmt"

	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// ApiClient is the main struct to interact with the raw Tapo API.
type ApiClient struct {
	username string
	password string
	protocol Protocol
}

// NewClient creates a new ApiClient.
func NewClient(email, password string, protocol Protocol) *ApiClient {
	return &ApiClient{
		username: email,
		password: password,
		protocol: protocol,
	}
}

// Login creates a login session with the Tapo device.
func (c *ApiClient) Login(ip string) error {
	url := fmt.Sprintf("http://%s/app", ip)
	err := c.protocol.Login(url, c.username, c.password)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	return nil
}

// RefreshSession refreshes the authentication session of the client.
func (c *ApiClient) RefreshSession() error {
	err := c.protocol.RefreshSession(c.username, c.password)
	if err != nil {
		return fmt.Errorf("failed to refresh session: %w", err)
	}
	return nil
}

// RequestRaw sends a request to the device and returns the raw response.
// The response is not checked for errors.
//
// It can be unmarshaled either by wrapping in [response.GenericResponse] or by using the generic [response.UnmarshalResponse[T]] function.
func (c *ApiClient) RequestRaw(method string, params interface{}, withToken bool) ([]byte, error) {
	request := request.NewTapoRequest(method, params)
	requestData, err := json.Marshal(request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	response, err := c.protocol.ExecuteRequest(requestData, withToken)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to execute request: %w", err)
	}
	return response, nil
}

// Request sends a request to the device and returns the response. The response is not checked for errors.
// It returns a [response.GenericResponse] which can be used for further processing of the response.
func (c *ApiClient) Request(method string, params interface{}, withToken bool) (response.GenericResponse, error) {
	raw, err := c.RequestRaw(method, params, withToken)
	if err != nil {
		return response.GenericResponse{}, err
	}
	return response.NewGenericResponse(raw), nil
}

// RequestVoid sends a request to the device and returns an error if the response contains an error.
func (c *ApiClient) RequestVoid(method string, params interface{}, withToken bool) error {
	resp, err := c.Request(method, params, true)
	if err != nil {
		return err
	}

	// response must be unmarshaled to check for the responses ErrorCode field
	tapoResponse, err := resp.UnmarshalResponse(struct{}{})
	if err != nil {
		return err
	}
	if tapoResponse.HasError() {
		return tapoResponse.GetError()
	}
	return nil
}

// RequestData sends a request to the device and returns the response. The response is checked for errors.
// The provided response must be a pointer of the expected response type. It will be wrapped within a [response.TapoResponse]
// representing [response.TapoResponse.Result] to check for errors.
func (c *ApiClient) RequestData(method string, params interface{}, withToken bool, responseValue interface{}) error {
	raw, err := c.Request(method, params, true)
	if err != nil {
		return err
	}
	tapoResponse, err := raw.UnmarshalResponse(&responseValue)
	if err != nil {
		return err
	}

	if tapoResponse.HasError() {
		return tapoResponse.GetError()
	}

	if decodable, ok := tapoResponse.Result.(response.Decodable); ok {
		err := decodable.Decode()
		if err != nil {
			return err
		}
	}
	return nil
}
