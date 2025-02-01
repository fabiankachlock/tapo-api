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

func NewPlug(ip string, client *api.ApiClient) (*TapoPlug, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoPlug{
		client: client,
	}, nil
}

// NewP100 creates a new Tapo P100 device.
func NewP100(ip string, client *api.ApiClient) (*TapoPlug, error) {
	return NewPlug(ip, client)
}

// NewP105 creates a new Tapo P105 device.
func NewP105(ip string, client *api.ApiClient) (*TapoPlug, error) {
	return NewPlug(ip, client)
}

func (t *TapoPlug) RefreshSession() error {
	return t.client.RefreshSession()
}

// ResetDevice resets the device to factory defaults.
func (t *TapoPlug) ResetDevice() error {
	return api.ResetDevice(t.client)
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoPlug) GetDeviceInfo() (response.DeviceInfoPlug, error) {
	return api.GetDeviceInfo[response.DeviceInfoPlug](t.client)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoPlug) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.GetDeviceInfo[map[string]interface{}](t.client)
}

// GetDeviceUsage returns the device usage.
func (t *TapoPlug) GetDeviceUsage() (response.DeviceUsage, error) {
	return api.GetDeviceUsage[response.DeviceUsage](t.client)
}

// SetDeviceInfo sets the device information.
func (t *TapoPlug) SetDeviceInfo(info request.GenericDeviceInfoParams) error {
	return api.SetDeviceInfo(t.client, info.GetJsonValue())
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
