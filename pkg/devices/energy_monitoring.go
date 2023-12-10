package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

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

func (t *TapoEnergyMonitoringPlug) GetDeviceInfo() (response.DeviceInfoPlug, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceInfoPlug{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceInfoPlug](resp)
	if err != nil {
		return response.DeviceInfoPlug{}, err
	}
	return data.Result, nil
}

func (t *TapoEnergyMonitoringPlug) On() error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{
		On: true,
	})
	return err
}

func (t *TapoEnergyMonitoringPlug) Off() error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, request.PlugDeviceInfoParams{
		On: false,
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

func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info request.PlugDeviceInfoParams) error {
	_, err := t.client.Request(request.RequestSetDeviceInfo, info)
	return err
}

func (t *TapoEnergyMonitoringPlug) GetDeviceUsage() (response.DeviceUsageEnergyMonitor, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceUsageEnergyMonitor{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceUsageEnergyMonitor](resp)
	if err != nil {
		return response.DeviceUsageEnergyMonitor{}, err
	}
	return data.Result, nil
}

func (t *TapoEnergyMonitoringPlug) GetEnergyUsage(params request.GetEnergyDataParams) (response.EnergyUsage, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, params)
	if err != nil {
		return response.EnergyUsage{}, err
	}

	data, err := response.UnmarshalResponse[response.EnergyUsage](resp)
	if err != nil {
		return response.EnergyUsage{}, err
	}
	return data.Result, nil
}

func (t *TapoEnergyMonitoringPlug) GetCurrentPower() (response.CurrentPower, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.CurrentPower{}, err
	}

	data, err := response.UnmarshalResponse[response.CurrentPower](resp)
	if err != nil {
		return response.CurrentPower{}, err
	}
	return data.Result, nil
}
