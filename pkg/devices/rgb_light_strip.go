package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoRgbicLightStrip is the main struct to interact with the [L900] devices.
//
// [L900]: https://www.tapo.com/en/search/?q=L900
type TapoRgbLightStrip struct {
	client *api.ApiClient
}

func NewRgbLightStrip(ip, email, password string) (*TapoRgbLightStrip, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoRgbLightStrip{
		client: client,
	}, err
}

// NewL920 creates a new Tapo L920 device.
func NewL900(ip, email, password string) (*TapoRgbLightStrip, error) {
	return NewRgbLightStrip(ip, email, password)
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoRgbLightStrip) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoRgbLightStrip) GetDeviceInfo() (response.DeviceInfoRgbLightStrip, error) {
	return api.RequestData[response.DeviceInfoRgbLightStrip](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoRgbLightStrip) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.RequestData[map[string]interface{}](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceUsage returns the device usage.
func (t *TapoRgbLightStrip) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	return api.RequestData[response.DeviceUsageEnergyMonitor](t.client, request.RequestGetDeviceUsage, request.EmptyParams)
}

// SetDeviceInfo sets the device information.
func (t *TapoRgbLightStrip) SetDeviceInfo(info request.ColorLightDeviceInfoParams) error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, info.GetJsonValue())
}

// On turns the device on.
func (t *TapoRgbLightStrip) On() error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoRgbLightStrip) Off() error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state between on and off.
func (t *TapoRgbLightStrip) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}

// SetBrightness sets the brightness of the rgb light strip.
func (t *TapoRgbLightStrip) SetBrightness(brightness uint8) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetBrightness(brightness))
}

// SetHue sets the hue of the rgb light strip.
func (t *TapoRgbLightStrip) SetHue(hue uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetHue(hue))
}

// SetSaturation sets the saturation of the rgb light strip.
func (t *TapoRgbLightStrip) SetSaturation(saturation uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetSaturation(saturation))
}

// SetColorTemperature sets the color temperature of the rgb light strip.
func (t *TapoRgbLightStrip) SetColorTemperature(colorTemperature uint16) error {
	return t.SetDeviceInfo(request.NewColorLightDeviceInfoParams().SetColorTemperature(colorTemperature))
}
