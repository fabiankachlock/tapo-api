package tapo

import (
	"errors"
	"fmt"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/devices"
	"github.com/fabiankachlock/tapo-api/pkg/klap"
)

// TapoClient is the main struct to interact with the Tapo API
type TapoClient struct {
	username string
	password string
	options  TapoClientOption
}

// NewClient creates a new TapoClient
func NewClient(username, password string, options ...func(*TapoClientOption)) TapoClient {
	clientOptions := DefaultOptions
	for _, option := range options {
		option(&clientOptions)
	}

	return TapoClient{username, password, clientOptions}
}

func (t TapoClient) createProtocol() (api.Protocol, error) {
	if t.options.Timeout <= 0 {
		return nil, errors.New("failed to create protocol: timeout must be greater than 0")
	}

	if t.options.Protocol != TapoProtocolKLAP {
		// KLAP is the only supported protocol currently
		return nil, errors.New("failed to create protocol: unsupported protocol")
	}

	protocol, err := klap.NewProtocol(time.Duration(t.options.Timeout) * time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to create protocol: %w", err)
	}
	return protocol, nil
}

func (t TapoClient) GetApiClient() (*api.ApiClient, error) {
	protocol, err := t.createProtocol()
	if err != nil {
		return nil, err
	}

	return api.NewClient(t.username, t.password, protocol), nil
}

// Generic creates a new generic Tapo device
func (t TapoClient) Generic(ip string) (*devices.TapoGenericDevice, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewGenericDevice(ip, client)
}

// Plug creates a new Tapo plug device
func (t TapoClient) Plug(ip string) (*devices.TapoPlug, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewPlug(ip, client)
}

// Light creates a new Tapo light device
func (t TapoClient) Light(ip string) (*devices.TapoLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewLight(ip, client)
}

// ColorLight creates a new Tapo color light device
func (t TapoClient) ColorLight(ip string) (*devices.TapoColorLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewColorLight(ip, client)
}

// RgbLightStrip creates a new Tapo RGB light strip device
func (t TapoClient) RgbLightStrip(ip string) (*devices.TapoRgbLightStrip, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewRgbLightStrip(ip, client)
}

// RgbicLightStrip creates a new Tapo RGBIC light strip device
func (t TapoClient) RgbicLightStrip(ip string) (*devices.TapoRgbicLightStrip, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewRgbicLightStrip(ip, client)
}

// P100 creates a new Tapo P100 device
func (t TapoClient) P100(ip string) (*devices.TapoPlug, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewP100(ip, client)
}

// P105 creates a new Tapo P105 device
func (t TapoClient) P105(ip string) (*devices.TapoPlug, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewP105(ip, client)
}

// P100 creates a new Tapo P100 device
func (t TapoClient) P110(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewP110(ip, client)
}

// P105 creates a new Tapo P105 device
func (t TapoClient) P115(ip string) (*devices.TapoEnergyMonitoringPlug, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewP115(ip, client)
}

// H100 creates a new Tapo H100 device
func (t TapoClient) H100(ip string) (*devices.TapoHub, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewH100(ip, client)
}

// H200 creates a new Tapo H200 device
func (t TapoClient) H200(ip string) (*devices.TapoHub, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewH200(ip, client)
}

// L900 creates a new Tapo L900 device
func (t TapoClient) L900(ip string) (*devices.TapoRgbLightStrip, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL900(ip, client)
}

// L920 creates a new Tapo L920 device
func (t TapoClient) L920(ip string) (*devices.TapoRgbicLightStrip, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL920(ip, client)
}

// L930 creates a new Tapo L930 device
func (t TapoClient) L930(ip string) (*devices.TapoRgbicLightStrip, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL930(ip, client)
}

// L510 creates a new Tapo L510 device
func (t TapoClient) L510(ip string) (*devices.TapoLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL510(ip, client)
}

// L520 creates a new Tapo L520 device
func (t TapoClient) L520(ip string) (*devices.TapoLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL520(ip, client)
}

// L610 creates a new Tapo L610 device
func (t TapoClient) L610(ip string) (*devices.TapoLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL610(ip, client)
}

// L530 creates a new Tapo L530 device
func (t TapoClient) L530(ip string) (*devices.TapoColorLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL530(ip, client)
}

// L535 creates a new Tapo L535 device
func (t TapoClient) L535(ip string) (*devices.TapoColorLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL535(ip, client)
}

// L630 creates a new Tapo L630 device
func (t TapoClient) L630(ip string) (*devices.TapoColorLight, error) {
	client, err := t.GetApiClient()
	if err != nil {
		return nil, err
	}
	return devices.NewL630(ip, client)
}
