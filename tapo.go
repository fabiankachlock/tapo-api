package tapo

import (
	"github.com/fabiankachlock/tapo-api/pkg/devices"
	tapoutil "github.com/fabiankachlock/tapo-api/pkg/util"
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

// L900 creates a new Tapo L900 device
func (t TapoClient) L900(ip string) (*devices.TapoRgbLightStrip, error) {
	return devices.NewL900(ip, t.username, t.password)
}

// L920 creates a new Tapo L920 device
func (t TapoClient) L920(ip string) (*devices.TapoRgbicLightStrip, error) {
	return devices.NewL920(ip, t.username, t.password)
}

// L930 creates a new Tapo L930 device
func (t TapoClient) L930(ip string) (*devices.TapoRgbicLightStrip, error) {
	return devices.NewL930(ip, t.username, t.password)
}

// L510 creates a new Tapo L510 device
func (t TapoClient) L510(ip string) (*devices.TapoLight, error) {
	return devices.NewL510(ip, t.username, t.password)
}

// L520 creates a new Tapo L520 device
func (t TapoClient) L520(ip string) (*devices.TapoLight, error) {
	return devices.NewL520(ip, t.username, t.password)
}

// L610 creates a new Tapo L610 device
func (t TapoClient) L610(ip string) (*devices.TapoLight, error) {
	return devices.NewL610(ip, t.username, t.password)
}

// L530 creates a new Tapo L530 device
func (t TapoClient) L530(ip string) (*devices.TapoColorLight, error) {
	return devices.NewL530(ip, t.username, t.password)
}

// L535 creates a new Tapo L535 device
func (t TapoClient) L535(ip string) (*devices.TapoColorLight, error) {
	return devices.NewL535(ip, t.username, t.password)
}

// L630 creates a new Tapo L630 device
func (t TapoClient) L630(ip string) (*devices.TapoColorLight, error) {
	return devices.NewL630(ip, t.username, t.password)
}

// GetNickname decodes a nickname of a device
func GetNickname(nickname string) string {
	return tapoutil.GetNickname(nickname)
}

// GetSSID decodes an SSID of an device
func GetSSID(ssid string) string {
	// nicknames and SSIDs are encoded the same way
	return tapoutil.GetNickname(ssid)
}
