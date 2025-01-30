package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoGenericDevice is a handler for generic devices.
// It provides basic functionality that is common on all Tapo devices.
type TapoGenericDevice struct {
	client *api.ApiClient
}

// NewGenericDevice create a handler for a generic Tapo device.
func NewGenericDevice(ip string, client api.ApiClient) (*TapoGenericDevice, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoGenericDevice{
		client: &client,
	}, nil
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoGenericDevice) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoGenericDevice) GetDeviceInfo() (response.DeviceInfoGeneric, error) {
	return api.RequestData[response.DeviceInfoGeneric](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoGenericDevice) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.RequestData[map[string]interface{}](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// SetDeviceInfo sets the device information.
func (t *TapoGenericDevice) SetDeviceInfo(params request.GenericDeviceInfoParams) error {
	_, err := api.RequestData[response.DeviceInfoGeneric](t.client, request.RequestSetDeviceInfo, params.GetJsonValue())
	return err
}

// On turns the device on.
func (t *TapoGenericDevice) On() error {
	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoGenericDevice) Off() error {
	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state.
func (t *TapoGenericDevice) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}

	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}
