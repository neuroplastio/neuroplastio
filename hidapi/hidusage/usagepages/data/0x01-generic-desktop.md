---
name: Generic Desktop Page
alias: gd
code: 0x01
---
# Usage Table

| Usage ID | Usage Name                                 | Usage Types |
|----------|--------------------------------------------|-------------|
| 00       | Undefined                                  |             |
| 01       | Pointer                                    | CP          |
| 02       | Mouse                                      | CA          |
| 03-03    | Reserved                                   |             |
| 04       | Joystick                                   | CA          |
| 05       | Gamepad                                    | CA          |
| 06       | Keyboard                                   | CA          |
| 07       | Keypad                                     | CA          |
| 08       | Multi-axis  Controller                     | CA          |
| 09       | Tablet  PC  System  Controls               | CA          |
| 0A       | Water  Cooling  Device                     | CA          |
| 0B       | Computer  Chassis  Device                  | CA          |
| 0C       | Wireless  Radio  Controls                  | CA          |
| 0D       | Portable  Device  Control                  | CA          |
| 0E       | System  Multi-Axis  Controller             | CA          |
| 0F       | Spatial  Controller                        | CA          |
| 10       | Assistive  Control                         | CA          |
| 11       | Device  Dock                               | CA          |
| 12       | Dockable  Device                           | CA          |
| 13       | Call  State  Management  Control           | CA          |
| 14-2F    | Reserved                                   |             |
| 30       | X                                          | DV          |
| 31       | Y                                          | DV          |
| 32       | Z                                          | DV          |
| 33       | Rx                                         | DV          |
| 34       | Ry                                         | DV          |
| 35       | Rz                                         | DV          |
| 36       | Slider                                     | DV          |
| 37       | Dial                                       | DV          |
| 38       | Wheel                                      | DV          |
| 39       | Hat Switch                                 | DV          |
| 3A       | Counted  Buffer                            | CL          |
| 3B       | Byte Count                                 | DV          |
| 3C       | Motion Wakeup                              | OSC/DF      |
| 3D       | Start                                      | OOC         |
| 3E       | Select                                     | OOC         |
| 3F-3F    | Reserved                                   |             |
| 40       | Vx                                         | DV          |
| 41       | Vy                                         | DV          |
| 42       | Vz                                         | DV          |
| 43       | Vbrx                                       | DV          |
| 44       | Vbry                                       | DV          |
| 45       | Vbrz                                       | DV          |
| 46       | Vno                                        | DV          |
| 47       | Feature Notification                       | DV/DF       |
| 48       | Resolution Multiplier                      | DV          |
| 49       | Qx                                         | DV          |
| 4A       | Qy                                         | DV          |
| 4B       | Qz                                         | DV          |
| 4C       | Qw                                         | DV          |
| 4D-7F    | Reserved                                   |             |
| 80       | System  Control                            | CA          |
| 81       | System Power Down                          | OSC         |
| 82       | System Sleep                               | OSC         |
| 83       | System Wake Up                             | OSC         |
| 84       | System Context Menu                        | OSC         |
| 85       | System Main Menu                           | OSC         |
| 86       | System App Menu                            | OSC         |
| 87       | System Menu Help                           | OSC         |
| 88       | System Menu Exit                           | OSC         |
| 89       | System Menu Select                         | OSC         |
| 8A       | System Menu Right                          | RTC         |
| 8B       | System Menu Left                           | RTC         |
| 8C       | System Menu Up                             | RTC         |
| 8D       | System Menu Down                           | RTC         |
| 8E       | System Cold Restart                        | OSC         |
| 8F       | System Warm Restart                        | OSC         |
| 90       | D-pad Up                                   | OOC         |
| 91       | D-pad Down                                 | OOC         |
| 92       | D-pad Right                                | OOC         |
| 93       | D-pad Left                                 | OOC         |
| 94       | Index Trigger                              | MC/DV       |
| 95       | Palm Trigger                               | MC/DV       |
| 96       | Thumbstick                                 | CP          |
| 97       | System Function Shift                      | MC          |
| 98       | System Function Shift Lock                 | OOC         |
| 99       | System Function Shift Lock Indicator       | DV          |
| 9A       | System Dismiss Notification                | OSC         |
| 9B       | System Do Not Disturb                      | OOC         |
| 9C-9F    | Reserved                                   |             |
| A0       | System Dock                                | OSC         |
| A1       | System Undock                              | OSC         |
| A2       | System Setup                               | OSC         |
| A3       | System Break                               | OSC         |
| A4       | System Debugger Break                      | OSC         |
| A5       | Application Break                          | OSC         |
| A6       | Application Debugger Break                 | OSC         |
| A7       | System Speaker Mute                        | OSC         |
| A8       | System Hibernate                           | OSC         |
| A9       | System Microphone Mute                     | OOC         |
| AA-AF    | Reserved                                   |             |
| B0       | System Display Invert                      | OSC         |
| B1       | System Display Internal                    | OSC         |
| B2       | System Display External                    | OSC         |
| B3       | System Display Both                        | OSC         |
| B4       | System Display Dual                        | OSC         |
| B5       | System Display Toggle Int/Ext Mode         | OSC         |
| B6       | System Display Swap Primary/Secondary      | OSC         |
| B7       | System Display Toggle LCD Autoscale        | OSC         |
| B8-BF    | Reserved                                   |             |
| C0       | Sensor  Zone                               | CL          |
| C1       | RPM                                        | DV          |
| C2       | Coolant Level                              | DV          |
| C3       | Coolant Critical Level                     | SV          |
| C4       | Coolant Pump                               | US          |
| C5       | Chassis  Enclosure                         | CL          |
| C6       | Wireless Radio Button                      | OOC         |
| C7       | Wireless Radio LED                         | OOC         |
| C8       | Wireless Radio Slider Switch               | OOC         |
| C9       | System Display Rotation Lock Button        | OOC         |
| CA       | System Display Rotation Lock Slider Switch | OOC         |
| CB       | Control Enable                             | DF          |
| CC-CF    | Reserved                                   |             |
| D0       | Dockable Device Unique ID                  | DV          |
| D1       | Dockable Device Vendor ID                  | DV          |
| D2       | Dockable Device Primary Usage Page         | DV          |
| D3       | Dockable Device Primary Usage ID           | DV          |
| D4       | Dockable Device Docking State              | DF          |
| D5       | Dockable Device Display Occlusion          | CL          |
| D6       | Dockable Device Object Type                | DV          |
| D7-DF    | Reserved                                   |             |
| E0       | Call Active LED                            | OOC         |
| E1       | Call Mute Toggle                           | OSC         |
| E2       | Call Mute LED                              | OOC         |
| E3-FFFF  | Reserved                                   |             |

