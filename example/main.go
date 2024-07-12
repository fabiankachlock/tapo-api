package main

import (
	"fmt"
	"os"

	"github.com/fabiankachlock/tapo-api/pkg/api"
)

func main() {
	tapoIp := "192.168.4.2"
	tapoEmail := os.Getenv("TAPO_EMAIL")
	tapoPass := os.Getenv("TAPO_PASS")

	client, err := api.NewClient(tapoIp, tapoEmail, tapoPass)
	if err != nil {
		panic(err)
	}
	err = client.Login()
	if err != nil {
		panic(err)
	}

	resp, err := client.Request("get_support_alarm_type_list", map[string]interface{}{})
	// resp, err := client.Request("play_alarm", map[string]interface{}{
	// 	"alarm_duration": 2,
	// 	"alarm_volume":   "low",
	// 	"alarm_type":     "Connection 1",
	// })
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

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
