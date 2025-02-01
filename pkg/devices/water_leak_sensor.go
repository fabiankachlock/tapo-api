package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type TapoWaterLeakSensor struct {
	deviceId string
	client   *api.ApiClient
}

func NewWaterLeakSensor(deviceId string, parentClient *api.ApiClient) (*TapoWaterLeakSensor, error) {
	return &TapoWaterLeakSensor{
		deviceId: deviceId,
		client:   parentClient,
	}, nil
}

func NewT300(deviceId string, parentClient *api.ApiClient) (*TapoWaterLeakSensor, error) {
	return NewWaterLeakSensor(deviceId, parentClient)
}

func (t *TapoWaterLeakSensor) RefreshSession() error {
	return t.client.RefreshSession()
}

func (t *TapoWaterLeakSensor) GetDeviceInfo() (childdevices.DeviceInfoT300, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[childdevices.DeviceInfoT300](t.client, t.deviceId, request)
}

func (t *TapoWaterLeakSensor) GetDeviceInfoJSON() (map[string]interface{}, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[map[string]interface{}](t.client, t.deviceId, request)
}

func (t *TapoWaterLeakSensor) GetTriggerLogs(pageSize uint64, startId uint64) (response.TriggerLogs[childdevices.LogEntryT300], error) {
	request := request.NewTapoRequest(request.RequestGetTriggerLogs, request.NewTriggerLogsParams(pageSize, startId))
	return api.ControlChild[response.TriggerLogs[childdevices.LogEntryT300]](t.client, t.deviceId, request)
}
