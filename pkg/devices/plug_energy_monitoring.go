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
func (t *TapoEnergyMonitoringPlug) GetDeviceInfo() (response.DeviceInfoPlugEnergyMonitoring, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceInfoPlugEnergyMonitoring{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceInfoPlugEnergyMonitoring](resp)
	if err != nil {
		return response.DeviceInfoPlugEnergyMonitoring{}, err
	}
	return data.Result, data.GetError()
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoEnergyMonitoringPlug) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.RequestData[map[string]interface{}](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// On turns the device on.
func (t *TapoEnergyMonitoringPlug) On() error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{
		On: true,
	})
}

// Off turns the device off.
func (t *TapoEnergyMonitoringPlug) Off() error {
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{
		On: false,
	})
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
	return api.RequestVoid(t.client, request.RequestSetDeviceInfo, info)
}

// GetDeviceUsage returns the device usage.
func (t *TapoEnergyMonitoringPlug) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	return api.RequestData[response.DeviceUsageEnergyMonitor](t.client, request.RequestGetDeviceUsage, request.EmptyParams)
}

// GetEnergyUsage returns the energy usage of the device.
func (t *TapoEnergyMonitoringPlug) GetEnergyUsage(params request.GetEnergyDataParams) (response.EnergyUsage, error) {
	return api.RequestData[response.EnergyUsage](t.client, request.RequestGetEnergyUsage, params)
}

// GetCurrentPower returns the current power usage of the device.
func (t *TapoEnergyMonitoringPlug) GetCurrentPower() (response.CurrentPower, error) {
	return api.RequestData[response.CurrentPower](t.client, request.RequestGetCurrentPower, request.EmptyParams)
}
