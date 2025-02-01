package devices

import (
	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
)

type TapoTemperaturHumiditySensor struct {
	deviceId string
	client   *api.ApiClient
}

func NewTemperaturHumiditySensor(deviceId string, parentClient *api.ApiClient) (*TapoTemperaturHumiditySensor, error) {
	return &TapoTemperaturHumiditySensor{
		deviceId: deviceId,
		client:   parentClient,
	}, nil
}

func NewT310(deviceId string, parentClient *api.ApiClient) (*TapoTemperaturHumiditySensor, error) {
	return NewTemperaturHumiditySensor(deviceId, parentClient)
}

func NewT315(deviceId string, parentClient *api.ApiClient) (*TapoTemperaturHumiditySensor, error) {
	return NewTemperaturHumiditySensor(deviceId, parentClient)
}

func (t *TapoTemperaturHumiditySensor) RefreshSession() error {
	return t.client.RefreshSession()
}

func (t *TapoTemperaturHumiditySensor) GetDeviceInfo() (childdevices.DeviceInfoT300, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[childdevices.DeviceInfoT300](t.client, t.deviceId, request)
}

func (t *TapoTemperaturHumiditySensor) GetDeviceInfoJSON() (map[string]interface{}, error) {
	request := request.NewTapoRequest(request.RequestGetDeviceInfo, request.EmptyParams)
	return api.ControlChild[map[string]interface{}](t.client, t.deviceId, request)
}

func (t *TapoWaterLeakSensor) GetTemperatureHumidityRecords() (childdevices.TemperaturHumidityRecords, error) {
	request := request.NewTapoRequest(request.RequestGetTemperatureHumidityRecords, request.EmptyParams)
	return api.ControlChild[childdevices.TemperaturHumidityRecords](t.client, t.deviceId, request)
}
