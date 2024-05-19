package tapo

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/klap"
	"github.com/fabiankachlock/tapo-api/pkg/devices"
)

// TapoClient is the main struct to interact with the Tapo API
type TapoClient struct {
	username string
	password string
}

// NewClient creates a new TapoClient
func NewClient(username, password string) TapoClient {
	return TapoClient{username, password}
}

func createClient(ip, email, password string) (api.Protocol, error) {
	protocol, err := klap.NewProtocol(api.NewOptions(ip, email, password))
	if err != nil {
		return nil, err
	}

	err = protocol.Login()
	return protocol, err
}

// P100 creates a new Tapo P100 device
func (t TapoClient) P110(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	client, err := createClient(ip, t.username, t.password)
	if err != nil {
		return nil, err
	}
	return devices.NewP110(client), nil
}

// P105 creates a new Tapo P105 device
func (t TapoClient) P115(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	client, err := createClient(ip, t.username, t.password)
	if err != nil {
		return nil, err
	}
	return devices.NewP115(client), nil
}
