package main

import (
	"log"
	"os"

	"github.com/fabiankachlock/tapo-api"
)

func main() {
	tapoIp := "192.168.0.2"
	tapoEmail := os.Getenv("TAPO_EMAIL")
	tapoPass := os.Getenv("TAPO_PASS")

	client := tapo.NewClient(tapoEmail, tapoPass)
	device, err := client.P115(tapoIp)
	if err != nil {
		log.Fatalln(err)
	}

	device.Toggle()
}
