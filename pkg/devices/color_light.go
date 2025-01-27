package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoColorLight is the main struct to interact with the [L530], [L535], [L630] devices.
//
// [L530]: https://www.tapo.com/en/search/?q=L530
// [L535]: https://www.tapo.com/en/search/?q=L535
// [L630]: https://www.tapo.com/en/search/?q=L630
type TapoColorLight struct {
	client *api.ApiClient
}

func NewColorLight(ip, email, password string) (*TapoColorLight, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoColorLight{
		client: client,
	}, err
}

// NewL510 creates a new Tapo L510 device.
func NewL530(ip, email, password string) (*TapoColorLight, error) {
	return NewColorLight(ip, email, password)
}

// NewL535 creates a new Tapo L535 device.
func NewL535(ip, email, password string) (*TapoColorLight, error) {
	return NewColorLight(ip, email, password)
}

// NewL630 creates a new Tapo L630 device.
func NewL630(ip, email, password string) (*TapoColorLight, error) {
	return NewColorLight(ip, email, password)
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoColorLight) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
func (t *TapoColorLight) GetDeviceInfo() (response.DeviceInfoColorLight, error) {
	return api.RequestData[response.DeviceInfoColorLight](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceUsage returns the device usage.
func (t *TapoColorLight) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	return api.RequestData[response.DeviceUsageEnergyMonitor](t.client, request.RequestGetDeviceUsage, request.EmptyParams)
}

// SetDeviceInfo sets the device information.
func (t *TapoColorLight) SetDeviceInfo(info request.ColorLightDeviceInfoParams) error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, info.GetJsonValue())
}

// On turns the device on.
func (t *TapoColorLight) On() error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoColorLight) Off() error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state between on and off.
func (t *TapoColorLight) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}

// SetBrightness sets the brightness of the color light.
func (t *TapoColorLight) SetBrightness(brightness uint8) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetBrightness(brightness))
}

// SetHue sets the hue of the color light.
func (t *TapoColorLight) SetHue(hue uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetHue(hue))
}

// SetSaturation sets the saturation of the color light.
func (t *TapoColorLight) SetSaturation(saturation uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetSaturation(saturation))
}

// SetColorTemperature sets the color temperature of the color light.
func (t *TapoColorLight) SetColorTemperature(colorTemperature uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetColorTemperature(colorTemperature))
}
