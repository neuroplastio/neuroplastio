---
name: Lighting and Illumination Page
alias: lig
code: 0x59
---
# Usage Table

| Usage ID | Usage Name                      | Usage Types |
|----------|---------------------------------|-------------|
| 00       | Undefined                       |             |
| 01       | LampArray                       | CA          |
| 02       | LampArrayAttributesReport       | CL          |
| 03       | LampCount                       | SV/DV       |
| 04       | BoundingBoxWidthInMicrometers   | SV          |
| 05       | BoundingBoxHeightInMicrometers  | SV          |
| 06       | BoundingBoxDepthInMicrometers   | SV          |
| 07       | LampArrayKind                   | SV          |
| 08       | MinUpdateIntervalInMicroseconds | SV          |
| 09-1F    | Reserved                        |             |
| 20       | LampAttributesRequestReport     | CL          |
| 21       | LampId                          | SV/DV       |
| 22       | LampAttributesResponseReport    | CL          |
| 23       | PositionXInMicrometers          | DV          |
| 24       | PositionYInMicrometers          | DV          |
| 25       | PositionZInMicrometers          | DV          |
| 26       | LampPurposes                    | DV          |
| 27       | UpdateLatencyInMicroseconds     | DV          |
| 28       | RedLevelCount                   | DV          |
| 29       | GreenLevelCount                 | DV          |
| 2A       | BlueLevelCount                  | DV          |
| 2B       | IntensityLevelCount             | DV          |
| 2C       | IsProgrammable                  | DV          |
| 2D       | InputBinding                    | DV          |
| 2E-4F    | Reserved                        |             |
| 50       | LampMultiUpdateReport           | CL          |
| 51       | RedUpdateChannel                | DV          |
| 52       | GreenUpdateChannel              | DV          |
| 53       | BlueUpdateChannel               | DV          |
| 54       | IntensityUpdateChannel          | DV          |
| 55       | LampUpdateFlags                 | DV          |
| 56-5F    | Reserved                        |             |
| 60       | LampRangeUpdateReport           | CL          |
| 61       | LampIdStart                     | DV          |
| 62       | LampIdEnd                       | DV          |
| 63-6F    | Reserved                        |             |
| 70       | LampArrayControlReport          | CL          |
| 71       | AutonomousMode                  | DV          |
| 72-FFFF  | Reserved                        |             |
