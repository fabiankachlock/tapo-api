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
	response, err := t.client.Request(api.RequestGetDeviceInfo, api.EmptyParams)
	if err != nil {
		return api.DeviceInfoPlug{}, err
	}

	data, err := api.UnmarshalResponse[api.DeviceInfoPlug](response)
	if err != nil {
		return api.DeviceInfoPlug{}, err
	}
	return data.Result, nil
}

func (t *TapoEnergyMonitoringPlug) On() error {
	_, err := t.client.Request(api.RequestSetDeviceInfo, map[string]interface{}{
		"device_on": true,
	})
	return err
}

func (t *TapoEnergyMonitoringPlug) Off() error {
	_, err := t.client.Request(api.RequestSetDeviceInfo, map[string]interface{}{
		"device_on": false,
	})
	return err
}

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

func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info api.PlugDeviceInfoParams) error {
	return nil
}

func (t *TapoEnergyMonitoringPlug) GetDeviceUsage() (api.DeviceUsageEnergyMonitor, error) {
	response, err := t.client.Request(api.RequestGetDeviceInfo, api.EmptyParams)
	if err != nil {
		return api.DeviceUsageEnergyMonitor{}, err
	}

	data, err := api.UnmarshalResponse[api.DeviceUsageEnergyMonitor](response)
	if err != nil {
		return api.DeviceUsageEnergyMonitor{}, err
	}
	return data.Result, nil
}

func (t *TapoEnergyMonitoringPlug) GetEnergyUsage(params api.GetEnergyDataParams) (api.EnergyUsage, error) {
	// response, err := t.client.Request(api.RequestGetDeviceInfo, params)
	// if err != nil {
	return api.EnergyUsage{}, nil
	// }

	// data, err := api.UnmarshalResponse[api.EnergyUsage](response)
	// if err != nil {
	// 	return api.EnergyUsage{}, err
	// }
	// return data.Result, nil
}

func (t *TapoEnergyMonitoringPlug) GetCurrentPower() (api.CurrentPower, error) {
	response, err := t.client.Request(api.RequestGetDeviceInfo, api.EmptyParams)
	if err != nil {
		return api.CurrentPower{}, err
	}

	data, err := api.UnmarshalResponse[api.CurrentPower](response)
	if err != nil {
		return api.CurrentPower{}, err
	}
	return data.Result, nil
}
