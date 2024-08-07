---
name: Auxiliary Display Page
alias: ad
code: 0x14
---
# Usage Table

| Usage ID | Usage Name                   | Usage Types    |
|----------|------------------------------|----------------|
| 00       | Undefined                    |                |
| 01       | Alphanumeric  Display        | CA             |
| 02       | Auxiliary  Display           | CA             |
| 03-1F    | Reserved                     |                |
| 20       | Display  Attributes  Report  | CL             |
| 21       | ASCII Character Set          | SF             |
| 22       | Data Read Back               | SF             |
| 23       | Font Read Back               | SF             |
| 24       | Display  Control  Report     | CL             |
| 25       | Clear Display                | DF             |
| 26       | Display Enable               | DF             |
| 27       | Screen Saver Delay           | SV/DV          |
| 28       | Screen Saver Enable          | DF             |
| 29       | Vertical Scroll              | SF/DF          |
| 2A       | Horizontal Scroll            | SF/DF          |
| 2B       | Character  Report            | CL             |
| 2C       | Display Data                 | DV             |
| 2D       | Display  Status              | CL             |
| 2E       | Stat Not Ready               | Sel            |
| 2F       | Stat Ready                   | Sel            |
| 30       | Err Not a loadable character | Sel            |
| 31       | Err Font data cannot be read | Sel            |
| 32       | Cursor Position Report       | Sel            |
| 33       | Row                          | DV             |
| 34       | Column                       | DV             |
| 35       | Rows                         | SV             |
| 36       | Columns                      | SV             |
| 37       | Cursor Pixel Positioning     | SF             |
| 38       | Cursor Mode                  | DF             |
| 39       | Cursor Enable                | DF             |
| 3A       | Cursor Blink                 | DF             |
| 3B       | Font  Report                 | CL             |
| 3C       | Font Data                    | Buffered Bytes |
| 3D       | Character Width              | SV             |
| 3E       | Character Height             | SV             |
| 3F       | Character Spacing Horizontal | SV             |
| 40       | Character Spacing Vertical   | SV             |
| 41       | Unicode Character Set        | SF             |
| 42       | Font 7-Segment               | SF             |
| 43       | 7-Segment Direct Map         | SF             |
| 44       | Font 14-Segment              | SF             |
| 45       | 14-Segment Direct Map        | SF             |
| 46       | Display Brightness           | DV             |
| 47       | Display Contrast             | DV             |
| 48       | Character  Attribute         | CL             |
| 49       | Attribute Readback           | SF             |
| 4A       | Attribute Data               | DV             |
| 4B       | Char Attr Enhance            | OOC            |
| 4C       | Char Attr Underline          | OOC            |
| 4D       | Char Attr Blink              | OOC            |
| 4E-7F    | Reserved                     |                |
| 80       | Bitmap Size X                | SV             |
| 81       | Bitmap Size Y                | SV             |
| 82       | Max Blit Size                | SV             |
| 83       | Bit Depth Format             | SV             |
| 84       | Display Orientation          | DV             |
| 85       | Palette  Report              | CL             |
| 86       | Palette Data Size            | SV             |
| 87       | Palette Data Offset          | SV             |
| 88       | Palette Data                 | Buffered Bytes |
| 89-89    | Reserved                     |                |
| 8A       | Blit  Report                 | CL             |
| 8B       | Blit Rectangle X1            | SV             |
| 8C       | Blit Rectangle Y1            | SV             |
| 8D       | Blit Rectangle X2            | SV             |
| 8E       | Blit Rectangle Y2            | SV             |
| 8F       | Blit Data                    | Buffered Bytes |
| 90       | Soft  Button                 | CL             |
| 91       | Soft Button ID               | SV             |
| 92       | Soft Button Side             | SV             |
| 93       | Soft Button Offset 1         | SV             |
| 94       | Soft Button Offset 2         | SV             |
| 95       | Soft Button Report           | SV             |
| 96-C1    | Reserved                     |                |
| C2       | Soft Keys                    | SV             |
| C3-CB    | Reserved                     |                |
| CC       | Display Data Extensions      | SF             |
| CD-CE    | Reserved                     |                |
| CF       | Character Mapping            | SV             |
| D0-DC    | Reserved                     |                |
| DD       | Unicode Equivalent           | SV             |
| DE-DE    | Reserved                     |                |
| DF       | Character Page Mapping       | SV             |
| E0-FE    | Reserved                     |                |
| FF       | Request Report               | DV             |
| 100-FFFF | Reserved                     |                |
