package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fabiankachlock/tapo-api"
)

func main() {
	tapoIp := "192.168.178.23"
	// tapoIp := "192.168.178.28"
	// tapoIp := "192.168.178.103"
	tapoEmail := os.Getenv("TAPO_EMAIL")
	tapoPass := os.Getenv("TAPO_PASS")

	client := tapo.NewClient(tapoEmail, tapoPass)
	device, err := client.H100(tapoIp)
	if err != nil {
		log.Fatalln(err)
	}

	info, err := device.GetDeviceInfo()
	fmt.Println("Device Info:")
	fmt.Println(err)
	json, _ := json.MarshalIndent(info, "", "  ")
	fmt.Println(string(json))
}
