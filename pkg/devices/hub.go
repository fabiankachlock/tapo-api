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

func NewHub(ip string, client api.ApiClient) (*TapoHub, error) {
	err := client.Login(ip)
	if err != nil {
		return nil, err
	}

	return &TapoHub{
		client: &client,
	}, nil
}

// NewH100 creates a new Tapo H100 device.
func NewH100(ip string, client api.ApiClient) (*TapoHub, error) {
	return NewHub(ip, client)
}

// NewH200 creates a new Tapo H200 device.
func NewH200(ip string, client api.ApiClient) (*TapoHub, error) {
	return NewHub(ip, client)
}

// RefreshSession refreshes the authentication session of the client.
func (t *TapoHub) RefreshSession() error {
	return t.client.RefreshSession()
}

// GetDeviceInfo returns the device information.
// It is not guaranteed to contain all the properties returned from the Tapo API.
func (t *TapoHub) GetDeviceInfo() (response.DeviceInfoHub, error) {
	return api.RequestData[response.DeviceInfoHub](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoHub) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.RequestData[map[string]interface{}](t.client, request.RequestGetDeviceInfo, request.EmptyParams)
}

func (t *TapoHub) GetChildDeviceList() (ChildDeviceList, error) {
	resp, err := t.client.Request(request.RequestGetChildDeviceList, request.EmptyParams, true)
	if err != nil {
		return ChildDeviceList{}, err
	}

	data := response.TapoResponse[rawDeviceList]{}
	err = json.Unmarshal(resp.Raw(), &data)
	if err != nil {
		return ChildDeviceList{}, err
	}

	devices := []ChildDeviceWrapper{}
	for _, rawDeviceJson := range data.Result.Devices {
		devices = append(devices, ChildDeviceWrapper{rawDeviceJson})
	}
	return ChildDeviceList{
		Devices: devices,
		Start:   data.Result.Start,
		Sum:     data.Result.Sum,
	}, nil
}

func (t *TapoHub) GetChildDeviceComponentList() (response.ChildDeviceComponentList, error) {
	return api.RequestData[response.ChildDeviceComponentList](t.client, request.RequestGetChildDeviceComponentList, request.EmptyParams)
}

func getChild(by func(ChildDeviceWrapper) (bool, error), devices []ChildDeviceWrapper) (bool, ChildDeviceWrapper, error) {
	for _, device := range devices {
		found, err := by(device)
		if err != nil {
			return false, ChildDeviceWrapper{}, err
		}
		if found {
			return true, device, nil
		}
	}
	return false, ChildDeviceWrapper{}, nil
}

func (t *TapoHub) GetChildById(deviceId string) (bool, ChildDeviceWrapper, error) {
	children, err := t.GetChildDeviceList()
	if err != nil {
		return false, ChildDeviceWrapper{}, err
	}

	return getChild(func(c ChildDeviceWrapper) (bool, error) {
		return c.hasDeviceId(deviceId)
	}, children.Devices)
}

func (t *TapoHub) GetChildByNickname(nickname string) (bool, ChildDeviceWrapper, error) {
	children, err := t.GetChildDeviceList()
	if err != nil {
		return false, ChildDeviceWrapper{}, err
	}

	return getChild(func(c ChildDeviceWrapper) (bool, error) {
		return c.hasNickname(nickname)
	}, children.Devices)
}

func (t *TapoHub) GetChild(nicknameOrId string) (bool, ChildDeviceWrapper, error) {
	children, err := t.GetChildDeviceList()
	if err != nil {
		return false, ChildDeviceWrapper{}, err
	}

	return getChild(func(c ChildDeviceWrapper) (bool, error) {
		found, err := c.hasNickname(nicknameOrId)
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
		return c.hasDeviceId(nicknameOrId)
	}, children.Devices)
}

func (t *TapoHub) GetT315(nicknameOrId string) (bool, childdevices.DeviceInfoT31X, error) {
	found, device, err := t.GetChild(nicknameOrId)
	if err != nil {
		return false, childdevices.DeviceInfoT31X{}, err
	}
	if !found {
		return false, childdevices.DeviceInfoT31X{}, nil
	}

	info, err := device.AsT315()
	return true, info, err
}

func (t *TapoHub) GetT310(nicknameOrId string) (bool, childdevices.DeviceInfoT31X, error) {
	found, device, err := t.GetChild(nicknameOrId)
	if err != nil {
		return false, childdevices.DeviceInfoT31X{}, err
	}
	if !found {
		return false, childdevices.DeviceInfoT31X{}, nil
	}

	info, err := device.AsT310()
	return true, info, err
}

func (t *TapoHub) GetT300(nicknameOrId string) (bool, childdevices.DeviceInfoT300, error) {
	found, device, err := t.GetChild(nicknameOrId)
	if err != nil {
		return false, childdevices.DeviceInfoT300{}, err
	}
	if !found {
		return false, childdevices.DeviceInfoT300{}, nil
	}

	info, err := device.AsT300()
	return true, info, err
}

func (t *TapoHub) GetSupportedAlarms() (response.AlarmsList, error) {
	return api.RequestData[response.AlarmsList](t.client, request.RequestSupportedAlarmTypes, request.EmptyParams)
}

func (t *TapoHub) PlayAlarm(alarmType string, volume request.AlarmVolume, duration int) error {
	return api.RequestVoid(t.client, request.RequestPlayAlarm, request.PlayAlarmParams{
		Type:     alarmType,
		Volume:   volume,
		Duration: duration,
	})
}

func (t *TapoHub) StopAlarm() error {
	return api.RequestVoid(t.client, request.RequestStopAlarm, request.EmptyParams)
}
