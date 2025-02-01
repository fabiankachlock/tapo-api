package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type TapoMotionSensor struct {
	deviceId string
	client   *api.ApiClient
}

func NewMotionSensor(deviceId string, parentClient *api.ApiClient) (*TapoMotionSensor, error) {
	return &TapoMotionSensor{
		deviceId: deviceId,
		client:   parentClient,
	}, nil
}

func NewT100(deviceId string, parentClient *api.ApiClient) (*TapoMotionSensor, error) {
	return NewMotionSensor(deviceId, parentClient)
}

func (t *TapoMotionSensor) RefreshSession() error {
	return t.client.RefreshSession()
}

func (t *TapoMotionSensor) GetDeviceInfo() (childdevices.DeviceInfoT100, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[childdevices.DeviceInfoT100](t.client, t.deviceId, request)
}

func (t *TapoMotionSensor) GetDeviceInfoJSON() (map[string]interface{}, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[map[string]interface{}](t.client, t.deviceId, request)
}

func (t *TapoMotionSensor) GetTriggerLogs(pageSize uint64, startId uint64) (response.TriggerLogs[childdevices.LogEntryT100], error) {
	request := request.NewTapoRequest(request.RequestGetTriggerLogs, request.NewTriggerLogsParams(pageSize, startId))
	return api.ControlChild[response.TriggerLogs[childdevices.LogEntryT100]](t.client, t.deviceId, request)
}
