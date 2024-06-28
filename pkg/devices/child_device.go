package devices

import (
	"encoding/json"

	childdevices "github.com/fabiankachlock/tapo-api/pkg/api/response/child_devices"
	tapoutil "github.com/fabiankachlock/tapo-api/pkg/util"
)

type rawDeviceList struct {
	Devices []json.RawMessage `json:"child_device_list"`
	Start   int32             `json:"start_index"`
	Sum     uint32            `json:"sum"`
}

type ChildDeviceList struct {
	Devices []ChildDeviceWrapper
	Start   int32
	Sum     uint32
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

func (c ChildDeviceWrapper) GetDeviceId() (string, error) {
	var temp struct {
		DeviceId string `json:"device_id"`
	}

	err := json.Unmarshal(c.raw, &temp)
	if err != nil {
		return "", err
	}
	return temp.DeviceId, nil
}

func (c ChildDeviceWrapper) GetNickname() (string, error) {
	var temp struct {
		Nickname string `json:"nickname"`
	}

	err := json.Unmarshal(c.raw, &temp)
	if err != nil {
		return "", err
	}
	return temp.Nickname, nil
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

func (c ChildDeviceWrapper) hasDeviceId(deviceId string) (bool, error) {
	id, err := c.GetDeviceId()
	if err != nil {
		return false, err
	}
	return id == deviceId, nil
}

func (c ChildDeviceWrapper) hasNickname(name string) (bool, error) {
	id, err := c.GetNickname()
	if err != nil {
		return false, err
	}
	return tapoutil.GetNickname(id) == name, nil
}
