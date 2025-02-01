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
	return api.GetDeviceInfo[response.DeviceInfoHub](t.client)
}

// GetDeviceInfoJSON returns the device information in raw JSON format.
func (t *TapoHub) GetDeviceInfoJSON() (map[string]interface{}, error) {
	return api.GetDeviceInfo[map[string]interface{}](t.client)
}

func (t *TapoHub) GetSupportedAlarms() (response.SupportedAlarmTypeList, error) {
	return api.GetSupportedAlarmTypes(t.client)
}

func (t *TapoHub) PlayAlarm(params request.PlayAlarmParams) error {
	return api.PlayAlarm(t.client, params)
}

func (t *TapoHub) StopAlarm() error {
	return api.StopAlarm(t.client)
}

func (t *TapoHub) GetChildDeviceListJSON(startIndex uint16) (map[string]interface{}, error) {
	return api.GetChildDeviceList[map[string]interface{}](t.client, request.NewChildDeviceListParams(startIndex))
}

func (t *TapoHub) GetChildDeviceList(startIndex uint16) (*TapoChildDeviceList, error) {
	list, err := api.GetChildDeviceList[response.ChildDeviceList](t.client, request.NewChildDeviceListParams(startIndex))
	if err != nil {
		return nil, err
	}

	wrappedChildDevices := []*TapoChildDevice{}
	for _, rawDeviceJson := range list.Devices {
		wrappedChildDevices = append(wrappedChildDevices, &TapoChildDevice{t, rawDeviceJson})
	}
	return &TapoChildDeviceList{
		Devices:    wrappedChildDevices,
		StartIndex: list.StartIndex,
		Sum:        list.Sum,
	}, nil
}

func (t *TapoHub) GetAllChildDevices(startIndex uint16) (*TapoChildDeviceList, error) {
	firstPage, err := t.GetChildDeviceList(startIndex)
	if err != nil {
		return nil, err
	}
	err = firstPage.FetchAll()
	if err != nil {
		return nil, err
	}
	return firstPage, nil
}

func (t *TapoHub) GetChildDeviceComponentList() (map[string]interface{}, error) {
	return api.GetChildDeviceComponentList[map[string]interface{}](t.client)
}
