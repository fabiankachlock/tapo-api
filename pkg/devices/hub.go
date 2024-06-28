package devices

import (
	"encoding/json"

	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

// TapoHub is the main struct to interact with the [H100] & [H200] devices.
//
// [H100]: https://www.tapo.com/en/search/?q=H100
// [H200]: https://www.tapo.com/en/search/?q=H200
type TapoHub struct {
	client *api.ApiClient
}

type rawDeviceList struct {
	Devices []json.RawMessage `json:"child_device_list"`
}

type ChildDeviceWrapper struct {
	raw []byte
}

func (c ChildDeviceWrapper) GetModel() (string, error) {
	var temp struct {
		Model string `json:"model"`
	}

	err := json.Unmarshal(c.raw, &temp)
	if err != nil {
		return "", err
	}
	return temp.Model, nil
}

func (c ChildDeviceWrapper) AsT315() (childdevices.DeviceInfoT31X, error) {
	var data childdevices.DeviceInfoT31X
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c ChildDeviceWrapper) AsT310() (childdevices.DeviceInfoT31X, error) {
	var data childdevices.DeviceInfoT31X
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c ChildDeviceWrapper) AsT300() (childdevices.DeviceInfoT300, error) {
	var data childdevices.DeviceInfoT300
	err := json.Unmarshal(c.raw, &data)
	return data, err
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

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoHub) GetChildDeviceList() ([]ChildDeviceWrapper, error) {
	resp, err := t.client.Request(request.RequestGetChildDeviceList, request.EmptyParams)
	if err != nil {
		return []ChildDeviceWrapper{}, err
	}

	data := response.TapoResponse[rawDeviceList]{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return []ChildDeviceWrapper{}, err
	}

	devices := []ChildDeviceWrapper{}
	for _, rawDeviceJson := range data.Result.Devices {
		devices = append(devices, ChildDeviceWrapper{rawDeviceJson})
	}
	return devices, nil
}

// SetDeviceInfo sets the device information.
// func (t *TapoEnergyMonitoringPlug) SetDeviceInfo(info request.PlugDeviceInfoParams) error {
// 	_, err := t.client.Request(request.RequestSetDeviceInfo, info)
// 	return err
// }
