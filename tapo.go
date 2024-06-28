package tapo

import (
	"encoding/base64"

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

// P100 creates a new Tapo P100 device
func (t TapoClient) P110(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	return devices.NewP110(ip, t.username, t.password)
}

// P105 creates a new Tapo P105 device
func (t TapoClient) P115(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	return devices.NewP115(ip, t.username, t.password)
}

// H100 creates a new Tapo H100 device
func (t TapoClient) H100(ip string) (*devices.TapoHub, error) {
	return devices.NewH100(ip, t.username, t.password)
}

// H200 creates a new Tapo H200 device
func (t TapoClient) H200(ip string) (*devices.TapoHub, error) {
	return devices.NewH200(ip, t.username, t.password)
}

func GetNickname(nickname string) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(nickname)
	if err != nil {
		return nickname
	}
	return string(decodedBytes)
}
