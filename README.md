# Samsung MIM-B19N Modbus Client

This project provides a Go-based Modbus client for interacting with the Samsung MIM-B19N system.

## Overview

The repository contains a client implementation to communicate with Modbus-enabled MIM-B19N devices. The client handles the configuration and communication protocols required to interface with these devices.

The underlying serial (modbus) connection is handled via [simonvetter/modbus](https://github.com/simonvetter/modbus).

## Installation

`go get github.com/nuskucloud/samsung_mimb19n@v1.0.0`

## Usage

```go
package main

import "github.com/nuskucloud/samsung_mimb19n"
import "fmt"

func main() {
	heatpumpModbus := samsung_mimb19n.NewClient("/dev/ttyUSB0", 500) // Linux
	//heatpumpModbus := samsung_mimb19n.NewClient("COM6", 500) // Windows

	// The modbus slave address of the heat pump.
	err := heatpumpModbus.ModbusClient.SetUnitId(1)
	if err != nil {
		fmt.Println("Error setting unit id")
		return
	}

	err = heatpumpModbus.ModbusClient.Open()
	if err != nil {
		// error out if we failed to connect/open the device
		// note: multiple Open() attempts can be made on the same client until
		// the connection succeeds (i.e. err == nil), calling the constructor again
		// is unnecessary.
		// likewise, a client can be opened and closed as many times as needed.
		fmt.Println("Error opening modbus client")
		panic(err)
		return
	}
	
	heatpumpModbus.CentralHeatingEnable(true)
	heatpumpModbus.SetInsideTargetTemperature(20)
	heatpumpModbus.SetFlowTemperature(40)
}
```

## Fixing hidden register interaction

Some or all of the hidden registers need setting up prior to use, this only needs to be done one time per slave.

```go

err = heatpumpModbus.ModbusClient.WriteRegisters(6000, []uint16{ // outdoor unit
    samsung_mimb19n.OUTDOOR_HIDDEN_REGISTER_OUTSIDE_TEMPERATURE_SENSOR,
    samsung_mimb19n.OUTDOOR_HIDDEN_REGISTER_COMPRESSOR_FREQUENCY,
})
if err != nil {
    Logger.Error("Failed to configure hidden registers")
    panic(err)
}

err = heatpumpModbus.ModbusClient.WriteRegisters(7000, []uint16{ // indoor unit
    // These registers start at 82 and go upwards, so whichever order they are listed will set their register id	
	
    samsung_mimb19n.INDOOR_HIDDEN_REGISTER_FLOW_SENSOR,           // Register 82
    samsung_mimb19n.INDOOR_HIDDEN_REGISTER_TEMPERATURE_REFERENCE, // Register 83
    samsung_mimb19n.INDOOR_HIDDEN_REGISTER_WATER_PWM,             // Register 84
    samsung_mimb19n.INDOOR_HIDDEN_REGISTER_3_WAY_VALVE,           // Register 85
    
    HexToUint16("8010"), // Register 86,  Compressor 1 status
    HexToUint16("8236"), // Register 87,  Compressor 1 Order frequency
    HexToUint16("8237"), // Register 88,  Compressor 1 Target Frequency
    HexToUint16("8276"), // Register 89,  Compressor 2 Operating Frequency
    HexToUint16("8001"), // Register 90,  Heat pump operation mode
})
if err != nil {
    Logger.Error("Failed to configure hidden registers")
    panic(err)
}


func HexToUint16(hexStr string) uint16 {
    // Convert the hex string to a decimal integer (base 16)
    value, err := strconv.ParseUint(hexStr, 16, 16)
    if err != nil {
        panic(err)
    }
    return uint16(value)
}
```

---

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.

## Extra 

Ported from the [glynhudson/samsung-modbus-mim-b19n](https://github.com/glynhudson/samsung-modbus-mim-b19n) project. There is extra documentation in here around the various modbus registers which are not implemented yet.

