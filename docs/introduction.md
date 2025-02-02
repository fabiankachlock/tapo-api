# Introduction

`tapo-api` is an unofficial API client for TP-Link Tapo devices written in Golang. It can be used to communicate with a wide range of Tapo smart home devices locally.

## Getting Started

A minimum go version of `1.20` is required. You can install the package using `go get`:

```sh
go get -u github.com/fabiankachlock/tapo-api
```

## Quick Start

Here is a simple example of how to use the API client:

```go
func main() {
    // 1. Information needed to connect to a device 
    tapoIp := "192.168.172.31" // the ip address of your device
	tapoEmail := "user@example.com" // the email address of your Tapo account
	tapoPass := "securepassword" // the password of your Tapo account

    // 2. Create a new API client
	client := tapo.NewClient(tapoEmail, tapoPass)

    // 3. Get your device
    light, err := client.L535(tapoIp) // create a instance of a device of your model
	if err != nil {
		panic(err)
	}

    // 4. Interact with your device
    err = light.On()
	if err != nil {
		panic(err)
	}
}
```

## Next Steps

See [supported devices](/devices/) for a list of devices that can be used or see [advanced usage](/advanced/) if your device is not listed or you want to add support for a new device.