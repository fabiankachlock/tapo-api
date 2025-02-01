package devices

type TapoChildDeviceList struct {
	hub        *TapoHub
	Devices    []*TapoChildDevice
	StartIndex uint16
	Sum        uint16
}

func (t *TapoChildDeviceList) FetchNextPage() error {
	newPage, err := t.hub.GetChildDeviceList(t.StartIndex + uint16(len(t.Devices)))
	if err != nil {
		return err
	}

	t.Devices = append(t.Devices, newPage.Devices...)
	return nil
}

func (t *TapoChildDeviceList) FetchAll() error {
	for {
		err := t.FetchNextPage()
		if err != nil {
			return err
		}
		if t.IsFullyLoaded() {
			break
		}
	}
	return nil
}

func (t *TapoChildDeviceList) IsFullyLoaded() bool {
	return len(t.Devices) >= int(t.Sum)
}

func containsChild(devices []*TapoChildDevice, by func(*TapoChildDevice) (bool, error)) (bool, *TapoChildDevice, error) {
	for _, device := range devices {
		found, err := by(device)
		if err != nil {
			return false, nil, err
		}
		if found {
			return true, device, nil
		}
	}
	return false, nil, nil
}

func (t *TapoChildDeviceList) SearchChild(by func(*TapoChildDevice) (bool, error)) (bool, *TapoChildDevice, error) {
	// check if the currently loaded devices contain the child
	found, child, err := containsChild(t.Devices, by)
	if err != nil {
		return false, nil, err
	}
	if found {
		return true, child, nil
	}
	if !t.IsFullyLoaded() {
		// if the child is not present and the list not fully loaded, fetch the next page
		err := t.FetchNextPage()
		if err != nil {
			return false, nil, err
		}
		return containsChild(t.Devices, by)
	}
	return false, nil, nil
}

func (t *TapoChildDeviceList) GetChildById(deviceId string) (bool, *TapoChildDevice, error) {
	return t.SearchChild(func(c *TapoChildDevice) (bool, error) {
		return c.hasDeviceId(deviceId)
	})
}

func (t *TapoChildDeviceList) GetChildByNickname(nickname string) (bool, *TapoChildDevice, error) {
	return t.SearchChild(func(c *TapoChildDevice) (bool, error) {
		return c.hasNickname(nickname)
	})
}

func (t *TapoChildDeviceList) GetChild(nicknameOrId string) (bool, *TapoChildDevice, error) {
	return t.SearchChild(func(c *TapoChildDevice) (bool, error) {
		found, err := c.hasNickname(nicknameOrId)
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
		return c.hasDeviceId(nicknameOrId)
	})
}

func (c TapoChildDeviceList) GetS200B(nicknameOrId string) (bool, *TapoSmartButton, error) {
	found, device, err := c.GetChild(nicknameOrId)
	if found {
		handler, err := device.AsS200B()
		return true, handler, err
	}
	return false, nil, err
}

func (c TapoChildDeviceList) GetT100(nicknameOrId string) (bool, *TapoMotionSensor, error) {
	found, device, err := c.GetChild(nicknameOrId)
	if found {
		handler, err := device.AsT100()
		return true, handler, err
	}
	return false, nil, err
}

func (c TapoChildDeviceList) GetT110(nicknameOrId string) (bool, *TapoContactSensor, error) {
	found, device, err := c.GetChild(nicknameOrId)
	if found {
		handler, err := device.AsT110()
		return true, handler, err
	}
	return false, nil, err
}

func (c TapoChildDeviceList) GetT300(nicknameOrId string) (bool, *TapoWaterLeakSensor, error) {
	found, device, err := c.GetChild(nicknameOrId)
	if found {
		handler, err := device.AsT300()
		return true, handler, err
	}
	return false, nil, err
}

func (c TapoChildDeviceList) GetT310(nicknameOrId string) (bool, *TapoTemperaturHumiditySensor, error) {
	found, device, err := c.GetChild(nicknameOrId)
	if found {
		handler, err := device.AsT310()
		return true, handler, err
	}
	return false, nil, err
}

func (c TapoChildDeviceList) GetT315(nicknameOrId string) (bool, *TapoTemperaturHumiditySensor, error) {
	found, device, err := c.GetChild(nicknameOrId)
	if found {
		handler, err := device.AsT315()
		return true, handler, err
	}
	return false, nil, err
}
