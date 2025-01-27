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

// NewL510 creates a new Tapo L510 device.
func NewL510(ip, email, password string) (*TapoLight, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoLight{
		client: client,
	}, err
}

// NewL520 creates a new Tapo L520 device.
func NewL520(ip, email, password string) (*TapoLight, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoLight{
		client: client,
	}, err
}

// NewL610 creates a new Tapo L610 device.
func NewL610(ip, email, password string) (*TapoLight, error) {
	client, err := api.NewClient(ip, email, password)
	if err != nil {
		return nil, err
	}

	err = client.Login()
	if err != nil {
		return nil, err
	}

	return &TapoLight{
		client: client,
	}, err
}

func (t *TapoLight) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoLight) GetDeviceInfo() (response.DeviceInfoLight, error) {
	return api.RequestData[response.DeviceInfoLight](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceUsage returns the device usage.
func (t *TapoLight) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	return api.RequestData[response.DeviceUsageEnergyMonitor](t.client, request.RequestGetDeviceUsage, request.EmptyParams)
}

func (t *TapoLight) On() error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, request.NewLightDeviceInfoParams().SetDeviceOn(true).GetJsonValue())
}

func (t *TapoLight) Off() error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, request.NewLightDeviceInfoParams().SetDeviceOn(false).GetJsonValue())
}

func (t *TapoLight) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	if state.DeviceOn {
		return t.Off()
	}
	return t.On()
}

func (t *TapoLight) SetBrightness(brightness uint8) error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, request.NewLightDeviceInfoParams().SetBrightness(brightness).GetJsonValue())
}
