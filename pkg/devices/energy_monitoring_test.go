package devices

import (
	"fmt"
	"time"

	"github.com/fabiankachlock/tapo-api/pkg/api/request"
)

func Example() {
	plug, _ := NewP115("127.0.0.1", "user", "pass")
	_ = plug.On()
	info, _ := plug.GetDeviceInfo()
	fmt.Println(info)
}

func ExampleTapoEnergyMonitoringPlug_SetDeviceInfo() {
	plug, _ := NewP115("127.0.0.1", "user", "pass")
	_ = plug.SetDeviceInfo(request.PlugDeviceInfoParams{
		On: true,
	})
}

func ExampleTapoEnergyMonitoringPlug_GetEnergyUsage() {
	plug, _ := NewP115("127.0.0.1", "user", "pass")
	usage, _ := plug.GetEnergyUsage(request.GetEnergyDataParamsDaily(time.Now()))
	fmt.Println(usage)
}
