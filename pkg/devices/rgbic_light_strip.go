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
func NewRgbicLightStrip(ip, email, password string) (*TapoRgbicLightStrip, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoRgbicLightStrip{
		client: client,
	}, err
}

// NewL920 creates a new Tapo L920 device.
func NewL920(ip, email, password string) (*TapoRgbicLightStrip, error) {
	return NewRgbicLightStrip(ip, email, password)
}

// NewL930 creates a new Tapo L930 device.
func NewL930(ip, email, password string) (*TapoRgbicLightStrip, error) {
	return NewRgbicLightStrip(ip, email, password)
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoRgbicLightStrip) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoRgbicLightStrip) GetDeviceInfo() (response.DeviceInfoRgbicLightStrip, error) {
	return api.RequestData[response.DeviceInfoRgbicLightStrip](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoRgbicLightStrip) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.RequestData[map[string]interface{}](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceUsage returns the device usage.
func (t *TapoRgbicLightStrip) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	return api.RequestData[response.DeviceUsageEnergyMonitor](t.client, request.RequestGetDeviceUsage, request.EmptyParams)
}

// SetDeviceInfo sets the device information.
func (t *TapoRgbicLightStrip) SetDeviceInfo(info request.ColorLightDeviceInfoParams) error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, info.GetJsonValue())
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

// SetBrightness sets the brightness of the rgbic light strip.
func (t *TapoRgbicLightStrip) SetBrightness(brightness uint8) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetBrightness(brightness))
}

// SetHue sets the hue of the rgbic light strip.
func (t *TapoRgbicLightStrip) SetHue(hue uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetHue(hue))
}

// SetSaturation sets the saturation of the rgbic light strip.
func (t *TapoRgbicLightStrip) SetSaturation(saturation uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetSaturation(saturation))
}

// SetColorTemperature sets the color temperature of the rgbic light strip.
func (t *TapoRgbicLightStrip) SetColorTemperature(colorTemperature uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetColorTemperature(colorTemperature))
}
