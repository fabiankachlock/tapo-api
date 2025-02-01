package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoLight is the main struct to interact with the [L510], [L520], [L610] devices.
//
// [L510]: https://www.tapo.com/en/search/?q=L510
// [L520]: https://www.tapo.com/en/search/?q=L520
// [L610]: https://www.tapo.com/en/search/?q=L610
type TapoLight struct {
	client *api.ApiClient
}

func NewLight(ip string, client *api.ApiClient) (*TapoLight, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoLight{
		client: client,
	}, nil
}

// NewL510 creates a new Tapo L510 device.
func NewL510(ip string, client *api.ApiClient) (*TapoLight, error) {
	return NewLight(ip, client)
}

// NewL520 creates a new Tapo L520 device.
func NewL520(ip string, client *api.ApiClient) (*TapoLight, error) {
	return NewLight(ip, client)
}

// NewL610 creates a new Tapo L610 device.
func NewL610(ip string, client *api.ApiClient) (*TapoLight, error) {
	return NewLight(ip, client)
}

func (t *TapoLight) RefreshSession() error {
	return t.client.RefreshSession()
}

// ResetDevice resets the device to factory defaults.
func (t *TapoLight) ResetDevice() error {
	return api.ResetDevice(t.client)
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoLight) GetDeviceInfo() (response.DeviceInfoLight, error) {
	return api.GetDeviceInfo[response.DeviceInfoLight](t.client)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoLight) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.GetDeviceInfo[map[string]interface{}](t.client)
}

// GetDeviceUsage returns the device usage.
func (t *TapoLight) GetDeviceUsage() (response.DeviceUsageEnergyMonitoring, error) {
	return api.GetDeviceUsage[response.DeviceUsageEnergyMonitoring](t.client)
}

// SetDeviceInfo sets the device information.
func (t *TapoLight) SetDeviceInfo(info request.LightDeviceInfoParams) error {
	return api.SetDeviceInfo(t.client, info.GetJsonValue())
}

// On turns the device on.
func (t *TapoLight) On() error {
	return t.SetDeviceInfo(request.NewLightDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoLight) Off() error {
	return t.SetDeviceInfo(request.NewLightDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state between on and off.
func (t *TapoLight) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	return t.SetDeviceInfo(request.NewLightDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}

// SetBrightness sets the brightness and turns the device on.
func (t *TapoLight) SetBrightness(brightness uint8) error {
	return t.SetDeviceInfo(request.NewLightDeviceInfoParams().SetBrightness(brightness))
}
