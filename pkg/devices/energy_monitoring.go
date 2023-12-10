package devices

import "github.com/fabiankachlock/tapo-api/pkg/api"

type TapoEnergyMonitoringPlug struct {
	client *api.ApiClient
}

func NewP110(ip, email, password string) (*TapoEnergyMonitoringPlug, error) {
	client, err := api.NewClient(ip, email, password)
	client.Login()
	return &TapoEnergyMonitoringPlug{
		client: client,
	}, err
}

func NewP115(ip, email, password string) (*TapoEnergyMonitoringPlug, error) {
	client, err := api.NewClient(ip, email, password)
	client.Login()
	return &TapoEnergyMonitoringPlug{
		client: client,
	}, err
}

func (t *TapoEnergyMonitoringPlug) RefreshSession() error {
	return t.client.RefreshSession()
}

func (t *TapoEnergyMonitoringPlug) GetDeviceInfo() (api.DeviceInfoPlug, error) {
	return api.DeviceInfoPlug{}, nil
}

func (t *TapoEnergyMonitoringPlug) On() error {
	return nil
}

func (t *TapoEnergyMonitoringPlug) Off() error {
	return nil
}

func (t *TapoEnergyMonitoringPlug) Toggle() error {
	return nil
}

func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info api.PlugDeviceInfoParams) error {
	return nil
}

func (t *TapoEnergyMonitoringPlug) GetDeviceUsage() (api.DeviceUsageEnergyMonitor, error) {
	return api.DeviceUsageEnergyMonitor{}, nil
}

func (t *TapoEnergyMonitoringPlug) GetEnergyUsage(params api.GetEnergyDataParams) (api.EnergyUsage, error) {
	return api.EnergyUsage{}, nil
}

func (t *TapoEnergyMonitoringPlug) GetCurrentPower() (api.CurrentPower, error) {
	return api.CurrentPower{}, nil
}
