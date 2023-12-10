package tapo

import "github.com/fabiankachlock/tapo-api/pkg/devices"

type TapoClient struct {
	username string
	password string
}

func NewClient(username, password string) TapoClient {
	return TapoClient{username, password}
}

func (t TapoClient) P110(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	return devices.NewP110(ip, t.username, t.password)
}

func (t TapoClient) P115(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	return devices.NewP115(ip, t.username, t.password)
}
