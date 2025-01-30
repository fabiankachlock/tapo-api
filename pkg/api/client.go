package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

const TerminalUUID = "00-00-00-00-00-00"

// ApiClient is the main struct to interact with the raw Tapo API.
type ApiClient struct {
	username string
	password string
	protocol Protocol
}

// NewClient creates a new ApiClient.
func NewClient(email, password string, protocol Protocol) ApiClient {
	return ApiClient{
		username: email,
		password: password,
		protocol: protocol,
	}
}

// Login logs in to the Tapo API.
func (d *ApiClient) Login(ip string) error {
	url := fmt.Sprintf("http://%s/app", ip)
	err := d.protocol.Login(url, d.username, d.password)
	if err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}
	return nil
}

// RefreshSession refreshes the authentication session of the client.
func (d *ApiClient) RefreshSession() error {
	err := d.protocol.RefreshSession(d.username, d.password)
	if err != nil {
		return fmt.Errorf("failed to refresh session: %w", err)
	}
	return nil
}

func (d *ApiClient) RequestRaw(method string, params interface{}, withToken bool) ([]byte, error) {
	request := map[string]interface{}{
		"method":           method,
		"params":           params,
		"requestTimeMilis": time.Now().UnixMilli(),
		"terminalUUID":     "00-00-00-00-00-00",
	}
	requestData, err := json.Marshal(request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	response, err := d.protocol.ExecuteRequest(requestData, withToken)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to execute request: %w", err)
	}
	return response, nil
}

func (d *ApiClient) Request(method string, params interface{}, withToken bool) (response.GenericResponse, error) {
	raw, err := d.RequestRaw(method, params, withToken)
	if err != nil {
		return response.GenericResponse{}, err
	}
	return response.NewGenericResponse(raw), nil
}
