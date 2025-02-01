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

func NewColorLight(ip string, client *api.ApiClient) (*TapoColorLight, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoColorLight{
		client: client,
	}, nil
}

// NewL510 creates a new Tapo L510 device.
func NewL530(ip string, client *api.ApiClient) (*TapoColorLight, error) {
	return NewColorLight(ip, client)
}

// NewL535 creates a new Tapo L535 device.
func NewL535(ip string, client *api.ApiClient) (*TapoColorLight, error) {
	return NewColorLight(ip, client)
}

// NewL630 creates a new Tapo L630 device.
func NewL630(ip string, client *api.ApiClient) (*TapoColorLight, error) {
	return NewColorLight(ip, client)
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoColorLight) RefreshSession() error {
	return t.client.RefreshSession()
}

// ResetDevice resets the device to factory defaults.
func (t *TapoColorLight) ResetDevice() error {
	return api.ResetDevice(t.client)
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoColorLight) GetDeviceInfo() (response.DeviceInfoColorLight, error) {
	return api.GetDeviceInfo[response.DeviceInfoColorLight](t.client)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoColorLight) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.GetDeviceInfo[map[string]interface{}](t.client)
}

// GetDeviceUsage returns the device usage.
func (t *TapoColorLight) GetDeviceUsage() (response.DeviceUsageEnergyMonitoring, error) {
	return api.GetDeviceUsage[response.DeviceUsageEnergyMonitoring](t.client)
}

// SetDeviceInfo sets the device information.
func (t *TapoColorLight) SetDeviceInfo(info request.ColorLightDeviceInfoParams) error {
	return api.SetDeviceInfo(t.client, info.GetJsonValue())
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

// SetBrightness sets the brightness and turns the device on.
func (t *TapoColorLight) SetBrightness(brightness uint8) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetBrightness(brightness))
}

// SetHue sets the hue of the and turns the device on.
func (t *TapoColorLight) SetHue(hue uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetHue(hue))
}

// SetSaturation sets the saturation and turns the device on.
func (t *TapoColorLight) SetSaturation(saturation uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetSaturation(saturation))
}

// SetColorTemperature sets the color temperature and turns the device on.
func (t *TapoColorLight) SetColorTemperature(colorTemperature uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetColorTemperature(colorTemperature))
}
