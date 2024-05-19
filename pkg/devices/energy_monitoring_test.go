package devices

import (
	"fmt"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/api/klap"
	"github.com/fabiankachlock/tapo-api/pkg/api/request"
)

func Example() {
	protocol, _ := klap.NewProtocol(api.NewOptions("127.0.0.1", "user", "pass"))
	plug := NewP115(protocol)
	_ = plug.On()
	info, _ := plug.GetDeviceInfo()
	fmt.Println(info)
}

func ExampleTapoEnergyMonitoringPlug_SetDeviceInfo() {
	protocol, _ := klap.NewProtocol(api.NewOptions("127.0.0.1", "user", "pass"))
	plug := NewP115(protocol)
	_ = plug.SetDeviceInfo(request.PlugDeviceInfoParams{
		On: true,
	})
}

func ExampleTapoEnergyMonitoringPlug_GetEnergyUsage() {
	protocol, _ := klap.NewProtocol(api.NewOptions("127.0.0.1", "user", "pass"))
	plug := NewP115(protocol)
	usage, _ := plug.GetEnergyUsage(request.GetEnergyDataParamsDaily(time.Now()))
	fmt.Println(usage)
}
