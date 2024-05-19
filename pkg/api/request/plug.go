package request

// PlugDeviceInfoParams is the request parameters for the [RequestSetDeviceInfo] requests to plugs.
type PlugDeviceInfoParams struct {
	On bool `json:"device_on"`
}
