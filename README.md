# Samsung MIM-B19N Modbus Client

This project provides a Go-based Modbus client for interacting with the Samsung MIM-B19N system.

## Overview

The repository contains a client implementation to communicate with Modbus-enabled MIM-B19N devices. The client handles the configuration and communication protocols required to interface with these devices.

## Installation

`go get github.com/nuskucloud/samsung_mimb19n@v0.5.0`

## Usage

```go
package main

import "github.com/nuskucloud/samsung_mimb19n"

func main() {
	heatpumpModbus := samsung_mimb19n.NewClient("/dev/ttyUSB0", 500)

	heatpumpModbus.DhwEnable(true)
	heatpumpModbus.SetInsideTargetTemperature(20)
	heatpumpModbus.SetFlowTemperature(40)
}
```

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Extra 

Ported from the [glynhudson/samsung-modbus-mim-b19n](https://github.com/glynhudson/samsung-modbus-mim-b19n) project. There is extra documentation in here around the various modbus registers which are not implemented yet.

