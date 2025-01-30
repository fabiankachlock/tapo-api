package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoPlug is the main struct to interact with the [P100], [P105] devices.
//
// [P100]: https://www.tapo.com/en/search/?q=P100
// [P105]: https://www.tapo.com/en/search/?q=P105
type TapoPlug struct {
	client *api.ApiClient
}

func NewPlug(ip, email, password string) (*TapoPlug, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoPlug{
		client: client,
	}, err
}

// NewP100 creates a new Tapo P100 device.
func NewP100(ip, email, password string) (*TapoPlug, error) {
	return NewPlug(ip, email, password)
}

// NewP105 creates a new Tapo P105 device.
func NewP105(ip, email, password string) (*TapoPlug, error) {
	return NewPlug(ip, email, password)
}

func (t *TapoPlug) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoPlug) GetDeviceInfo() (response.DeviceInfoPlug, error) {
	return api.RequestData[response.DeviceInfoPlug](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoPlug) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.RequestData[map[string]interface{}](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceUsage returns the device usage.
func (t *TapoPlug) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	return api.RequestData[response.DeviceUsageEnergyMonitor](t.client, request.RequestGetDeviceUsage, request.EmptyParams)
}

// SetDeviceInfo sets the device information.
func (t *TapoPlug) SetDeviceInfo(params request.GenericDeviceInfoParams) error {
	_, err := api.RequestData[response.DeviceInfoGeneric](t.client, request.RequestSetDeviceInfo, params.GetJsonValue())
	return err
}

// On turns the device on.
func (t *TapoPlug) On() error {
	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoPlug) Off() error {
	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state.
func (t *TapoPlug) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}

	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}
