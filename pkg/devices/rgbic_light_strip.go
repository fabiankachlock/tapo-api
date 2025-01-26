package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoRgbicLightStrip is the main struct to interact with the [L920] & [L930] devices.
//
// [L920: https://www.tapo.com/en/search/?q=L920
// [L930: https://www.tapo.com/en/search/?q=L930
type TapoRgbicLightStrip struct {
	client *api.ApiClient
}

// NewL920 creates a new Tapo L920 device.
func NewL920(ip, email, password string) (*TapoRgbicLightStrip, error) {
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

// NewL930 creates a new Tapo L930 device.
func NewL930(ip, email, password string) (*TapoRgbicLightStrip, error) {
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

// RefreshSession refreshes the authentication session of the client.
func (t *TapoRgbicLightStrip) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoRgbicLightStrip) GetDeviceInfo() (response.DeviceInfoRgbicLightStrip, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceInfoRgbicLightStrip{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceInfoRgbicLightStrip](resp)
	if err != nil {
		return response.DeviceInfoRgbicLightStrip{}, err
	}
	return data.Result, data.GetError()
}

// GetDeviceUsage returns the device usage.
func (t *TapoRgbicLightStrip) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceUsageEnergyMonitor{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceUsageEnergyMonitor](resp)
	if err != nil {
		return response.DeviceUsageEnergyMonitor{}, err
	}
	return data.Result, data.GetError()
}

// SetDeviceInfo sets the device information.
func (t *TapoRgbicLightStrip) SetDeviceInfo(info request.ColorLightDiveInfoParams) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, info)
	return err
}

func (t *TapoRgbicLightStrip) On() error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.ColorLightDiveInfoParams{
		DeviceOn: true,
	})
	return err
}

func (t *TapoRgbicLightStrip) Off() error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.ColorLightDiveInfoParams{
		DeviceOn: false,
	})
	return err
}

func (t *TapoRgbicLightStrip) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	if state.DeviceOn {
		return t.Off()
	}
	return t.On()
}

func (t *TapoRgbicLightStrip) SetBrightness(brightness uint8) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.ColorLightDiveInfoParams{
		Brightness: brightness,
	})
	return err
}

func (t *TapoRgbicLightStrip) SetHue(hue uint16) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.ColorLightDiveInfoParams{
		Hue: hue,
	})
	return err
}

func (t *TapoRgbicLightStrip) SetSaturation(saturation uint16) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.ColorLightDiveInfoParams{
		Saturation: saturation,
	})
	return err
}

func (t *TapoRgbicLightStrip) SetColorTemperature(colorTemperature uint16) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.ColorLightDiveInfoParams{
		ColorTemperature: colorTemperature,
	})
	return err
}
