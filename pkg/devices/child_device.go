package devices

import (
	"encoding/json"

	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
	tapoutil "github.com/fabiankachlock/tapo-api/pkg/util"
)

type TapoChildDevice struct {
	hub *TapoHub
	raw json.RawMessage
}

func (c *TapoChildDevice) GetModel() (string, error) {
	var temp struct {
		Model string `json:"model"`
	}

	err := json.Unmarshal(c.raw, &temp)
	if err != nil {
		return "", err
	}
	return temp.Model, nil
}

func (c *TapoChildDevice) GetDeviceId() (string, error) {
	var temp struct {
		DeviceId string `json:"device_id"`
	}

	err := json.Unmarshal(c.raw, &temp)
	if err != nil {
		return "", err
	}
	return temp.DeviceId, nil
}

func (c *TapoChildDevice) GetNickname() (string, error) {
	var temp struct {
		Nickname string `json:"nickname"`
	}

	err := json.Unmarshal(c.raw, &temp)
	if err != nil {
		return "", err
	}
	return temp.Nickname, nil
}

func (c *TapoChildDevice) hasDeviceId(deviceId string) (bool, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return false, err
	}
	return id == deviceId, nil
}

func (c *TapoChildDevice) hasNickname(name string) (bool, error) {
	id, err := c.GetNickname()
	if err != nil {
		return false, err
	}
	return tapoutil.GetNickname(id) == name, nil
}

func (c TapoChildDevice) GetDeviceInfoGenericChildDevice() (childdevices.DeviceInfoGenericChildDevice, error) {
	var data childdevices.DeviceInfoGenericChildDevice
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) GetDeviceInfoS200B() (childdevices.DeviceInfoS200B, error) {
	var data childdevices.DeviceInfoS200B
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) GetDeviceInfoT100() (childdevices.DeviceInfoT100, error) {
	var data childdevices.DeviceInfoT100
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) GetDeviceInfoT110() (childdevices.DeviceInfoT110, error) {
	var data childdevices.DeviceInfoT110
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) GetDeviceInfoT300() (childdevices.DeviceInfoT300, error) {
	var data childdevices.DeviceInfoT300
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) GetDeviceInfoT310() (childdevices.DeviceInfoT31X, error) {
	var data childdevices.DeviceInfoT31X
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) GetDeviceInfoT315() (childdevices.DeviceInfoT31X, error) {
	var data childdevices.DeviceInfoT31X
	err := json.Unmarshal(c.raw, &data)
	return data, err
}

func (c TapoChildDevice) AsS200B() (*TapoSmartButton, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return nil, err
	}
	return NewS200B(id, c.hub.client)
}

func (c TapoChildDevice) AsT110() (*TapoContactSensor, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return nil, err
	}
	return NewT110(id, c.hub.client)
}

func (c TapoChildDevice) AsT100() (*TapoMotionSensor, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return nil, err
	}
	return NewT100(id, c.hub.client)
}

func (c TapoChildDevice) AsT300() (*TapoWaterLeakSensor, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return nil, err
	}
	return NewT300(id, c.hub.client)
}

func (c TapoChildDevice) AsT310() (*TapoTemperaturHumiditySensor, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return nil, err
	}
	return NewT310(id, c.hub.client)
}

func (c TapoChildDevice) AsT315() (*TapoTemperaturHumiditySensor, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return nil, err
	}
	return NewT315(id, c.hub.client)
}
