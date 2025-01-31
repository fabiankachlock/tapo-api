# tapo-api

An unofficial TPLink Tapo API Client written in Go.

![GitHub Release](https://img.shields.io/github/v/release/fabiankachlock/tapo-api?style=for-the-badge)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/fabiankachlock/tapo-api?style=for-the-badge)
![GitHub License](https://img.shields.io/github/license/fabiankachlock/tapo-api?style=for-the-badge)

Reverse engineered from https://k4czp3r.xyz/blog/post/reverse-engineering-tp-link-tapo \
Reference implementation: https://github.com/mihai-dinculescu/tapo

Docs: https://pkg.go.dev/github.com/fabiankachlock/tapo-api

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
- H100 (tested)
- H200
- L900
- L920
- L930 (tested)
- L510
- L520
- L530
- L535 (tested)
- L610
- L630

Supported Protocols:
- KLAP (tested)

&#x2705; - Implemented \
&check; - Supported by reference implementation in rust


| Feature<br/><br/><br/>               | GenericDevice<br/><br/><br/> | L510<br/>L520<br/>L610 | L530<br/>L535<br/>L630<br/> | L900<br/><br/><br/> | L920<br/>L930<br/><br/> | P100<br/>P105<br/><br/> | P110<br/>P115<br/><br/> | P300<br/>P304<br/><br/> | H100<br/>H200<br/><br/> |
| ------------------------------------ | :--------------------------: | :--------------------: | :-------------------------: | :-----------------: | :---------------------: | :---------------------: | :---------------------: | :---------------------: | :---------------------: |
| device_reset                         |          &#x2705;\*          |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |                         |                         |
| get_child_device_component_list_json |                              |                        |                             |                     |                         |                         |                         |         &check;         |       &#x2705;\*        |
| get_child_device_list                |                              |                        |                             |                     |                         |                         |                         |         &check;         |        &#x2705;         |
| get_child_device_list_json           |                              |                        |                             |                     |                         |                         |                         |         &check;         |         &check;         |
| get_current_power                    |                              |                        |                             |                     |                         |                         |        &#x2705;         |                         |                         |
| get_device_info                      |           &#x2705;           |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |         &check;         |        &#x2705;         |
| get_device_info_json                 |           &#x2705;           |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |         &check;         |        &#x2705;         |
| get_device_usage                     |                              |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |                         |                         |
| get_energy_data                      |                              |                        |                             |                     |                         |                         |         &check;         |                         |                         |
| get_energy_usage                     |                              |                        |                             |                     |                         |                         |        &#x2705;         |                         |                         |
| get_supported_ringtone_list          |                              |                        |                             |                     |                         |                         |                         |                         |       &#x2705;\*        |
| off                                  |           &#x2705;           |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |                         |                         |
| on                                   |           &#x2705;           |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |                         |                         |
| toggle\*                             |           &#x2705;           |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |                         |        &#x2705;         |                         |                         |
| play_alarm                           |                              |                        |                             |                     |                         |                         |                         |                         |       &#x2705;\*        |
| refresh_session                      |           &#x2705;           |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |        &#x2705;         |        &#x2705;         |         &check;         |        &#x2705;         |
| set_brightness                       |                              |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |                         |                         |                         |                         |
| set_color                            |                              |                        |           &check;           |       &check;       |         &check;         |                         |                         |                         |                         |
| set_color_temperature                |                              |                        |          &#x2705;           |      &#x2705;       |        &#x2705;         |                         |                         |                         |                         |
| set_device_info                      |          &#x2705;\*          |        &#x2705;        |          &#x2705;           |      &#x2705;       |        &#x2705;         |                         |                         |                         |                         |
| set_hue_saturation                   |                              |                        |          &#x2705;           |      &#x2705;       |        &#x2705;         |                         |                         |                         |                         |
| set_lighting_effect                  |                              |                        |                             |                     |        &#x2705;         |                         |                         |                         |                         |
| stop_alarm                           |                              |                        |                             |                     |                         |                         |                         |                         |       &#x2705;\*        |


\* Such APIs are go only and not available in the reference implementation in rust.

## Hub (H100) Child Devices Support

&#x2705; - Implemented \
&check; - Supported by reference implementation in rust

| Feature<br/><br/>                | KE100<br/><br/> | S200B<br/><br/> | T100<br/><br/> | T110<br/><br/> | T300<br/><br/> | T310<br/>T315 |
| -------------------------------- | :-------------: | :-------------: | :------------: | :------------: | :------------: | :-----------: |
| get_device_info \*               |     &check;     |     &check;     |    &check;     |    &check;     |    &#x2705;    |   &#x2705;    |
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
