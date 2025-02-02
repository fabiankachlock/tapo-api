# Devices

`tapo-api` comes with built in support for a lot of TP-Link Tapo devices. They are grouped by common features and are listed below.

| Group                                                     | Models                                                                                                                                       |
| --------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| [light](/devices/light)                                   | [L510](https://www.tapo.com/en/search/?q=L510) [L520](https://www.tapo.com/en/search/?q=L520) [L610](https://www.tapo.com/en/search/?q=L610) |
| [colored light](/devices/clored-light)                    | [L530](https://www.tapo.com/en/search/?q=L530) [L535](https://www.tapo.com/en/search/?q=L535) [L630](https://www.tapo.com/en/search/?q=L630) |
| [RGB light strip](/devices/rgb-light-strip)               | [L900](https://www.tapo.com/en/search/?q=L900)                                                                                               |
| [RGBIC light strip](/devices/rgbic-light-strip)           | [L920](https://www.tapo.com/en/search/?q=L920) [L930](https://www.tapo.com/en/search/?q=L930)                                                |
| [plug](/devices/plug)                                     | [P100](https://www.tapo.com/en/search/?q=P100) [P105](https://www.tapo.com/en/search/?q=P105)                                                |
| [energy monitoring plug](/devices/energy-monitoring-plug) | [P110](https://www.tapo.com/en/search/?q=P110) [P115](https://www.tapo.com/en/search/?q=P115)                                                |
| [power strip](/devices/power-strip)                       | [P300](https://www.tapo.com/en/search/?q=P300) [P304](https://www.tapo.com/en/search/?q=P304)                                                |
| [hub](/devices/hub)                                       | [H100](https://www.tapo.com/en/search/?q=H100) [H200](https://www.tapo.com/en/search/?q=H200)                                                |
| [generic device](/devices/generic)\*                      | -                                                                                                                                            |

\* The generic device contains basic functionality that is shared be all devices. 

**Hub child devices**

Devices that need a hub to operate can't be accessed directly since they are not connected via Wi-Fi.
Instead, interaction with such always goes through the hub. The following child devices are supported:

| Group                                                            | Models                                                                                        |
| ---------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| [button](/devices/children/button)                               | [S200B](https://www.tapo.com/en/search/?q=S200B)                                              |
| [motion sensor](/devices/children/motion-sensor)                 | [T100](https://www.tapo.com/en/search/?q=T100)                                                |
| [contact sensor](/devices/children/contact-sensor)               | [T110](https://www.tapo.com/en/search/?q=T110)                                                |
| [water leak sensor](/devices/children/water-leak-sensor)         | [T300](https://www.tapo.com/en/search/?q=T300)                                                |
| [temperature&humidity sensor](/devices/children/temp-hum-sensor) | [T310](https://www.tapo.com/en/search/?q=T310) [T315](https://www.tapo.com/en/search/?q=T315) |