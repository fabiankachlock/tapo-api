package main

import (
	"log"
	"os"

	"github.com/fabiankachlock/tapo-api/pkg/api"
)

func main() {
	tapoIp := "192.168.0.2"
	tapoEmail := os.Getenv("TAPO_EMAIL")
	tapoPass := os.Getenv("TAPO_PASS")

	client, err := api.NewClient(tapoIp, tapoEmail, tapoPass)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Handshake()
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Request("get_device_info", map[string]interface{}{})
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Request("get_device_usage", map[string]interface{}{})
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Request("get_energy_usage", map[string]interface{}{})
	if err != nil {
		log.Fatalln(err)
	}

	// err = client.Request("get_energy_data", map[string]interface{}{
	// 	"start_timestamp": 0,
	// 	"end_timestamp": 1,
	// 	"interval":        60,
	// })
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	err = client.Request("get_current_power", map[string]interface{}{})
	if err != nil {
		log.Fatalln(err)
	}
}
