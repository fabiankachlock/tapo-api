package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
)

// TapoHub is the main struct to interact with the [H100] & [H200] devices.
//
// [H100]: https://www.tapo.com/en/search/?q=H100
// [H200]: https://www.tapo.com/en/search/?q=H200
type TapoHub struct {
	client *api.ApiClient
}

// NewH100 creates a new Tapo H100 device.
func NewH100(ip, email, password string) (*TapoHub, error) {
	client, err := api.NewClient(ip, email, password)
	client.Login()
	return &TapoHub{
		client: client,
	}, err
}

// NewH200 creates a new Tapo H200 device.
func NewH200(ip, email, password string) (*TapoHub, error) {
	client, err := api.NewClient(ip, email, password)
	client.Login()
	return &TapoHub{
		client: client,
	}, err
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoHub) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoHub) GetDeviceInfo() (response.DeviceInfoHub, error) {
	resp, err := t.client.Request(request.RequestGetDeviceInfo, request.EmptyParams)
	if err != nil {
		return response.DeviceInfoHub{}, err
	}

	data, err := response.UnmarshalResponse[response.DeviceInfoHub](resp)
	if err != nil {
		return response.DeviceInfoHub{}, err
	}
	return data.Result, nil
}

// SetDeviceInfo sets the device information.
// func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info request.PlugDeviceInfoParams) error {
// 	_, err := t.client.Request(request.RequestSetDeviceInfo, info)
// 	return err
// }
