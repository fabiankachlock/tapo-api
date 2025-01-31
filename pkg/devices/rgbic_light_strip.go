package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoRgbicLightStrip is the main struct to interact with the [L920] & [L930] devices.
//
// [L920]: https://www.tapo.com/en/search/?q=L920
// [L930]: https://www.tapo.com/en/search/?q=L930
type TapoRgbicLightStrip struct {
	client *api.ApiClient
}

// NewRgbicLightStrip creates a new Tapo RGBIC light strip device.
func NewRgbicLightStrip(ip string, client api.ApiClient) (*TapoRgbicLightStrip, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoRgbicLightStrip{
		client: &client,
	}, nil
}

// NewL920 creates a new Tapo L920 device.
func NewL920(ip string, client api.ApiClient) (*TapoRgbicLightStrip, error) {
	return NewRgbicLightStrip(ip, client)
}

// NewL930 creates a new Tapo L930 device.
func NewL930(ip string, client api.ApiClient) (*TapoRgbicLightStrip, error) {
	return NewRgbicLightStrip(ip, client)
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoRgbicLightStrip) RefreshSession() error {
	return t.client.RefreshSession()
}

// ResetDevice resets the device to factory defaults.
func (t *TapoRgbicLightStrip) ResetDevice() error {
	return api.ResetDevice(t.client)
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoRgbicLightStrip) GetDeviceInfo() (response.DeviceInfoRgbicLightStrip, error) {
	return api.GetDeviceInfo[response.DeviceInfoRgbicLightStrip](t.client)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoRgbicLightStrip) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.GetDeviceInfo[map[string]interface{}](t.client)
}

// GetDeviceUsage returns the device usage.
func (t *TapoRgbicLightStrip) GetDeviceUsage() (response.DeviceUsageEnergyMonitoring, error) {
	return api.GetDeviceUsage[response.DeviceUsageEnergyMonitoring](t.client)
}

// SetDeviceInfo sets the device information.
func (t *TapoRgbicLightStrip) SetDeviceInfo(info request.ColorLightDeviceInfoParams) error {
	return api.SetDeviceInfo(t.client, info.GetJsonValue())
}

// On turns the device on.
func (t *TapoRgbicLightStrip) On() error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoRgbicLightStrip) Off() error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state between on and off.
func (t *TapoRgbicLightStrip) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}

// SetBrightness sets the brightness and turns the device on.
// Any pre existing lighting effect will be removed.
func (t *TapoRgbicLightStrip) SetBrightness(brightness uint8) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetBrightness(brightness))
}

// SetHue sets the hue of the and turns the device on.
// Any pre existing lighting effect will be removed.
func (t *TapoRgbicLightStrip) SetHue(hue uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetHue(hue))
}

// SetSaturation sets the saturation and turns the device on.
// Any pre existing lighting effect will be removed.
func (t *TapoRgbicLightStrip) SetSaturation(saturation uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetSaturation(saturation))
}

// SetColorTemperature sets the color temperature and turns the device on.
// Any pre existing lighting effect will be removed.
func (t *TapoRgbicLightStrip) SetColorTemperature(colorTemperature uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetColorTemperature(colorTemperature))
}

// SetLightingEffect sets a lighting effect and turns the device on.
func (t *TapoRgbicLightStrip) SetLightingEffect(effect request.LightingEffect) error {
	return api.SetLightingEffect(t.client, effect)
}
