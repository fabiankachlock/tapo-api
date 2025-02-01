package main

import (
	"os"

	"github.com/fabiankachlock/tapo-api/pkg/api"
	"github.com/fabiankachlock/tapo-api/pkg/klap"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	tapoIp := "192.168.4.21"
	tapoEmail := os.Getenv("TAPO_EMAIL")
	tapoPass := os.Getenv("TAPO_PASS")

	protocol, err := klap.NewProtocol()
	if err != nil {
		panic(err)
	}

	client := api.NewClient(tapoEmail, tapoPass, protocol)
	err = client.Login(tapoIp)
	if err != nil {
		panic(err)
	}

	// d, err := devices.NewL535(tapoIp, tapoEmail, tapoPass)
	// if err != nil {
	// 	panic(err)
	// }
	// err = d.SetDeviceInfo(request.NewColorLightDeviceInfoParams().
	// 	SetDeviceOn(true).
	// 	SetBrightness(20).
	// 	SetColorTemperature(2800))

	// if err != nil {
	// 	panic(err)
	// }
	// resp, err := d.GetDeviceInfoJSON()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", resp)

	// resp, err := client.Request("get_device_info", map[string]interface{}{}, true)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(resp.Raw()))
	// resp1 := response.TapoResponse[response.DeviceInfoLight]{}
	// err = resp.UnmarshalRaw(&resp1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", resp1)

	// resp2 := response.DeviceInfoLight{}
	// _, err = resp.UnmarshalResponse(&resp2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", resp2.Avatar)

	// resp, err := client.Request("play_alarm", map[string]interface{}{
	// 	"alarm_duration": 2,
	// 	"alarm_volume":   "low",
	// 	"alarm_type":     "Connection 1",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(resp))

	// client := tapo.NewClient(tapoEmail, tapoPass)
	// device, err := client.H100(tapoIp)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// _, _ = device.GetChildDeviceComponentList()
	// devices, _ := device.GetChildDeviceList()
	// for _, device := range devices.Devices {
	// 	model, _ := device.GetModel()
	// 	if model == "T315" {
	// 		info, _ := device.AsT315()
	// 		fmt.Printf("device: %s tmp: %.2f hum: %d\n", tapo.GetNickname(info.Nickname), info.CurrentTemperature, info.CurrentHumidity)
	// 	}
	// }

	// fmt.Println("get child:")
	// ok, t315, err := device.GetT315("Sens001")
	// fmt.Printf("ok: %v, err: %v\n", ok, err)
	// fmt.Printf("t315: %+v\n", t315)
	// info, err := device.GetDeviceInfo()
	// fmt.Println("Device Info:")
	// fmt.Println(err)
	// json, _ := json.MarshalIndent(info, "", "  ")
	// fmt.Println(string(json))
}
