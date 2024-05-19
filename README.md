# tapo-api

An unofficial TPLink Tapo API Client written in Go.

Reverse engineered from https://k4czp3r.xyz/blog/post/reverse-engineering-tp-link-tapo
Reference implementation: https://github.com/mihai-dinculescu/tapo

## Usage

```
go get github.com/fabiankachlock/tapo-api
```

```go
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
```

## Device support

Supported Devices:
- P110
- P115 (tested)

Supported Protocols:
- KLAP

&#x2705; - Implemented \
&check; - Supported by reference implementation in rust

| Feature<br/><br/><br/>               | GenericDevice<br/><br/><br/> | L510<br/>L520<br/>L610 | L530<br/>L630<br/><br/> | L900<br/><br/><br/> | L920<br/>L930<br/><br/> | P100<br/>P105<br/><br/> | P110<br/>P115<br/><br/> | P300<br/><br/><br/> | H100<br/><br/><br/> |
| ------------------------------------ | :--------------------------: | :--------------------: | :---------------------: | :-----------------: | :---------------------: | :---------------------: | :---------------------: | :-----------------: | :-----------------: |
| device_reset                         |                              |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |                     |                     |
| get_child_device_component_list_json |                              |                        |                         |                     |                         |                         |                         |       &check;       |       &check;       |
| get_child_device_list                |                              |                        |                         |                     |                         |                         |                         |       &check;       |       &check;       |
| get_child_device_list_json           |                              |                        |                         |                     |                         |                         |                         |       &check;       |       &check;       |
| get_current_power                    |                              |                        |                         |                     |                         |                         |        &#x2705;         |                     |                     |
| get_device_info                      |           &check;            |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |       &check;       |       &check;       |
| get_device_info_json                 |           &check;            |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |       &check;       |       &check;       |
| get_device_usage                     |                              |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |                     |                     |
| get_energy_data                      |                              |                        |                         |                     |                         |                         |        &#x2705;         |                     |                     |
| get_energy_usage                     |                              |                        |                         |                     |                         |                         |        &#x2705;         |                     |                     |
| off                                  |           &check;            |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |                     |                     |
| on                                   |           &check;            |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |                     |                     |
| refresh_session                      |           &check;            |        &check;         |         &check;         |       &check;       |         &check;         |         &check;         |        &#x2705;         |       &check;       |       &check;       |
| set_brightness                       |                              |        &check;         |         &check;         |       &check;       |         &check;         |                         |                         |                     |                     |
| set_color                            |                              |                        |         &check;         |       &check;       |         &check;         |                         |                         |                     |                     |
| set_color_temperature                |                              |                        |         &check;         |       &check;       |         &check;         |                         |                         |                     |                     |
| set_hue_saturation                   |                              |                        |         &check;         |       &check;       |         &check;         |                         |                         |                     |                     |
| set_lighting_effect                  |                              |                        |                         |                     |         &check;         |                         |                         |                     |                     |
| set() API \*                         |                              |                        |         &check;         |       &check;       |         &check;         |                         |                         |                     |                     |

\* The `set()` API allows multiple properties to be set in a single request.

## Hub (H100) Child Devices Support

| Feature<br/><br/>                | KE100<br/><br/> | S200B<br/><br/> | T100<br/><br/> | T110<br/><br/> | T300<br/><br/> | T310<br/>T315 |
| -------------------------------- | :-------------: | :-------------: | :------------: | :------------: | :------------: | :-----------: |
| get_device_info \*               |     &check;     |     &check;     |    &check;     |    &check;     |    &check;     |    &check;    |
| get_device_info_json             |     &check;     |     &check;     |    &check;     |    &check;     |    &check;     |    &check;    |
| get_temperature_humidity_records |                 |                 |                |                |                |    &check;    |
| get_trigger_logs                 |                 |     &check;     |    &check;     |    &check;     |    &check;     |               |
| set_child_protection             |     &check;     |                 |                |                |                |               |
| set_frost_protection             |     &check;     |                 |                |                |                |               |
| set_max_control_temperature      |     &check;     |                 |                |                |                |               |
| set_min_control_temperature      |     &check;     |                 |                |                |                |               |
| set_target_temperature           |     &check;     |                 |                |                |                |               |
| set_temperature_offset           |     &check;     |                 |                |                |                |               |

\* Obtained by calling `get_child_device_list` on the hub device or `get_device_info` on a child device handler.
