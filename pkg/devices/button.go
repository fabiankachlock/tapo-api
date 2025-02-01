package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	"github.com/fabiankachlock/tapo-api/pkg/api/response"
	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type TapoSmartButton struct {
	deviceId string
	client   *api.ApiClient
}

func NewSmartButton(deviceId string, parentClient *api.ApiClient) (*TapoSmartButton, error) {
	return &TapoSmartButton{
		deviceId: deviceId,
		client:   parentClient,
	}, nil
}

func NewS200B(deviceId string, parentClient *api.ApiClient) (*TapoSmartButton, error) {
	return NewSmartButton(deviceId, parentClient)
}

func (t *TapoSmartButton) RefreshSession() error {
	return t.client.RefreshSession()
}

func (t *TapoSmartButton) GetDeviceInfo() (childdevices.DeviceInfoS200B, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[childdevices.DeviceInfoS200B](t.client, t.deviceId, request)
}

func (t *TapoSmartButton) GetDeviceInfoJSON() (map[string]interface{}, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[map[string]interface{}](t.client, t.deviceId, request)
}

func (t *TapoSmartButton) GetTriggerLogs(pageSize uint64, startId uint64) (response.TriggerLogs[childdevices.LogEntryS200B], error) {
	request := request.NewTapoRequest(request.RequestGetTriggerLogs, request.NewTriggerLogsParams(pageSize, startId))
	return api.ControlChild[response.TriggerLogs[childdevices.LogEntryS200B]](t.client, t.deviceId, request)
}
