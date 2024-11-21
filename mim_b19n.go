package samsung_mimb19n

import (
	"github.com/simonvetter/modbus"
	"time"
)

type Client struct {
	ModbusClient  *modbus.ModbusClient
	sleepDuration time.Duration
}

const (
	// Outdoor unit
	OUTDOOR_HIDDEN_REGISTER_COMPRESSOR_FREQUENCY       = 33411 // modbus:4, MessageSetID/hex: "8283"
	OUTDOOR_HIDDEN_REGISTER_OUTSIDE_TEMPERATURE_SENSOR = 33284 // modbus:5, MessageSetID/hex: "8204"

	// Indoor unit
	INDOOR_HIDDEN_REGISTER_FLOW_SENSOR           = 17129 // modbus:82, MessageSetID/hex: "42E9"
	INDOOR_HIDDEN_REGISTER_3_WAY_VALVE           = 16487 // modbus:90, MessageSetID/hex: "4067"
	INDOOR_HIDDEN_REGISTER_WATER_PWM             = 16580 // modbus:?, MessageSetID/hex: "40C4"
	INDOOR_HIDDEN_REGISTER_TEMPERATURE_REFERENCE = 16495 // modbus:?, MessageSetID/hex: "406F"
)

// NewClient creates and returns a new Modbus client configured for the specified serial port.
const (
	defaultModbusSpeed    = 9600
	defaultModbusDataBits = 8
	defaultModbusParity   = modbus.PARITY_EVEN
	defaultModbusStopBits = 1
	defaultModbusTimeout  = 1 * time.Second
	serialProtocol        = "rtu://"
)

// NewClient creates and returns a new Modbus client configured for the specified serial port.
func NewClient(serialPort string, sleepAfterModbusCommand time.Duration) *Client {
	modbusClient := createModbusClient(serialPort)
	client := &Client{
		ModbusClient:  modbusClient,
		sleepDuration: sleepAfterModbusCommand,
	}
	return client
}

// createModbusClient creates and returns a new Modbus client.
func createModbusClient(serialPort string) *modbus.ModbusClient {
	modbusClient, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:      serialProtocol + serialPort,
		Speed:    defaultModbusSpeed,
		DataBits: defaultModbusDataBits,
		Parity:   defaultModbusParity,
		StopBits: defaultModbusStopBits,
		Timeout:  defaultModbusTimeout,
	})
	if err != nil {
		panic(err)
	}
	return modbusClient
}

func BoolToInt(bool bool) uint16 {
	if bool {
		return 1
	} else {
		return 0
	}
}
