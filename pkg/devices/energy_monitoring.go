package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoEnergyMonitoringPlug is the main struct to interact with the [P110] & [P115] devices.
//
// [P110]: https://www.tapo.com/en/search/?q=P110
// [P115]: https://www.tapo.com/en/search/?q=P115
type TapoEnergyMonitoringPlug struct {
	client *api.ApiClient
}

// NewP110 creates a new Tapo P110 device.
func NewP110(ip, email, password string) (*TapoEnergyMonitoringPlug, error) {
	client, err := api.NewClient(ip, email, password)
	client.Login()
	return &TapoEnergyMonitoringPlug{
		client: client,
	}, err
}

// NewP115 creates a new Tapo P115 device.
func NewP115(ip, email, password string) (*TapoEnergyMonitoringPlug, error) {
	client, err := api.NewClient(ip, email, password)
	client.Login()
	return &TapoEnergyMonitoringPlug{
		client: client,
	}, err
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoEnergyMonitoringPlug) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoEnergyMonitoringPlug) GetDeviceInfo() (response.DeviceInfoPlug, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceInfoPlug{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceInfoPlug](resp)
	if err != nil {
		return response.DeviceInfoPlug{}, err
	}
	return data.Result, data.GetError()
}

// On turns the device on.
func (t *TapoEnergyMonitoringPlug) On() error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{
		On: true,
	})
	return err
}

// Off turns the device off.
func (t *TapoEnergyMonitoringPlug) Off() error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{
		On: false,
	})
	return err
}

// Toggle toggles the device state.
func (t *TapoEnergyMonitoringPlug) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}
	if state.DeviceOn {
		return t.Off()
	}
	return t.On()
}

// SetDeviceInfo sets the device information.
func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info request.PlugDeviceInfoParams) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, info)
	return err
}

// GetDeviceUsage returns the device usage.
func (t *TapoEnergyMonitoringPlug) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	resp, err := t.client.Request(request.RequestGetDeviceUsage, request.EmptyParams)
	if err != nil {
		return response.DeviceUsageEnergyMonitor{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceUsageEnergyMonitor](resp)
	if err != nil {
		return response.DeviceUsageEnergyMonitor{}, err
	}
	return data.Result, data.GetError()
}

// GetEnergyUsage returns the energy usage of the device.
func (t *TapoEnergyMonitoringPlug) GetEnergyUsage(params request.GetEnergyDataParams) (response.EnergyUsage, error) {
	resp, err := t.client.Request(request.RequestGetEnergyUsage, params)
	if err != nil {
		return response.EnergyUsage{}, err
	}

	data, err := response.UnmarshalResponse[response.EnergyUsage](resp)
	if err != nil {
		return response.EnergyUsage{}, err
	}
	return data.Result, data.GetError()
}

// GetCurrentPower returns the current power usage of the device.
func (t *TapoEnergyMonitoringPlug) GetCurrentPower() (response.CurrentPower, error) {
	resp, err := t.client.Request(request.RequestGetCurrentPower, request.EmptyParams)
	if err != nil {
		return response.CurrentPower{}, err
	}

	data, err := response.UnmarshalResponse[response.CurrentPower](resp)
	if err != nil {
		return response.CurrentPower{}, err
	}
	return data.Result, data.GetError()
}
