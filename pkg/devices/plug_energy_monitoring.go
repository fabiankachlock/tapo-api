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

func NewEnergyMonitoringPlug(ip string, client api.ApiClient) (*TapoEnergyMonitoringPlug, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoEnergyMonitoringPlug{
		client: &client,
	}, nil
}

// NewP110 creates a new Tapo P110 device.
func NewP110(ip string, client api.ApiClient) (*TapoEnergyMonitoringPlug, error) {
	return NewEnergyMonitoringPlug(ip, client)
}

// NewP115 creates a new Tapo P115 device.
func NewP115(ip string, client api.ApiClient) (*TapoEnergyMonitoringPlug, error) {
	return NewEnergyMonitoringPlug(ip, client)
}

func (t *TapoEnergyMonitoringPlug) RefreshSession() error {
	return t.client.RefreshSession()
}

// ResetDevice resets the device to factory defaults.
func (t *TapoEnergyMonitoringPlug) ResetDevice() error {
	return api.ResetDevice(t.client)
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoEnergyMonitoringPlug) GetDeviceInfo() (response.DeviceInfoEnergyMonitoringPlug, error) {
	return api.GetDeviceInfo[response.DeviceInfoEnergyMonitoringPlug](t.client)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoEnergyMonitoringPlug) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.GetDeviceInfo[map[string]interface{}](t.client)
}

// GetDeviceUsage returns the device usage.
func (t *TapoEnergyMonitoringPlug) GetDeviceUsage() (response.DeviceUsageEnergyMonitoring, error) {
	return api.GetDeviceUsage[response.DeviceUsageEnergyMonitoring](t.client)
}

func (t *TapoEnergyMonitoringPlug) GetEnergyUsage() (response.EnergyUsage, error) {
	return api.GetEnergyUsage(t.client)
}

func (t *TapoEnergyMonitoringPlug) GetEnergyData(params request.EnergyDataParams) (response.EnergyData, error) {
	return api.GetEnergyData(t.client, params)
}

func (t *TapoEnergyMonitoringPlug) GetCurrentPower() (response.CurrentPower, error) {
	return api.GetCurrentPower(t.client)
}

// SetDeviceInfo sets the device information.
func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info request.GenericDeviceInfoParams) error {
	return api.SetDeviceInfo(t.client, info.GetJsonValue())
}

// On turns the device on.
func (t *TapoEnergyMonitoringPlug) On() error {
	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(true))
}

// Off turns the device off.
func (t *TapoEnergyMonitoringPlug) Off() error {
	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(false))
}

// Toggle toggles the device state.
func (t *TapoEnergyMonitoringPlug) Toggle() error {
	state, err := t.GetDeviceInfo()
	if err != nil {
		return err
	}

	return t.SetDeviceInfo(request.NewGenericDeviceInfoParams().SetDeviceOn(!state.DeviceOn))
}
