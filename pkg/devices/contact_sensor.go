package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type TapoContactSensor struct {
	deviceId string
	client   *api.ApiClient
}

func NewContactSensor(deviceId string, parentClient *api.ApiClient) (*TapoContactSensor, error) {
	return &TapoContactSensor{
		deviceId: deviceId,
		client:   parentClient,
	}, nil
}

func NewT110(deviceId string, parentClient *api.ApiClient) (*TapoContactSensor, error) {
	return NewContactSensor(deviceId, parentClient)
}

func (t *TapoContactSensor) RefreshSession() error {
	return t.client.RefreshSession()
}

func (t *TapoContactSensor) GetDeviceInfo() (childdevices.DeviceInfoT110, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[childdevices.DeviceInfoT110](t.client, t.deviceId, request)
}

func (t *TapoContactSensor) GetDeviceInfoJSON() (map[string]interface{}, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[map[string]interface{}](t.client, t.deviceId, request)
}

func (t *TapoContactSensor) GetTriggerLogs(pageSize uint64, startId uint64) (response.TriggerLogs[childdevices.LogEntryT110], error) {
	request := request.NewTapoRequest(request.RequestGetTriggerLogs, request.NewTriggerLogsParams(pageSize, startId))
	return api.ControlChild[response.TriggerLogs[childdevices.LogEntryT110]](t.client, t.deviceId, request)
}
