package samsung_mimb19n

import (
	"github.com/simonvetter/modbus"
	"time"
)

func (client *Client) ReadOutdoorTemperature() (float32, error) {
	time.Sleep(client.sleepDuration)
	outdoorTemperature, err := client.ModbusClient.ReadRegister(5, modbus.HOLDING_REGISTER)
	if err != nil {
		return 10123, err
	}

	return float32(outdoorTemperature / 10), nil
}

func (client *Client) ReadIndoorTemperature() (float32, error) {
	time.Sleep(client.sleepDuration)
	indoorTemperature, err := client.ModbusClient.ReadRegister(59, modbus.HOLDING_REGISTER)
	if err != nil {
		return 10123, err
	}

	return float32(indoorTemperature) / 10, nil
}

func (client *Client) ReadReturnTemperature() (float32, error) {
	time.Sleep(client.sleepDuration)
	returnTemperature, err := client.ModbusClient.ReadRegister(65, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return float32(returnTemperature / 10), nil
}

func (client *Client) ReadFlowTemperature() (float32, error) {
	time.Sleep(client.sleepDuration)
	flowTemperature, err := client.ModbusClient.ReadRegister(66, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return float32(flowTemperature / 10), nil
}

func (client *Client) ReadOperatingMode() (float32, error) {
	time.Sleep(client.sleepDuration)
	indoorTemperature, err := client.ModbusClient.ReadRegister(53, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return float32(indoorTemperature / 10), nil
}

func (client *Client) ReadWaterOutTemperature() (float32, error) {
	time.Sleep(client.sleepDuration)
	waterOutTemperature, err := client.ModbusClient.ReadRegister(53, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return float32(waterOutTemperature / 10), nil
}

func (client *Client) ReadHotWaterTemperature() (float32, error) {
	time.Sleep(client.sleepDuration)
	waterOutTemperature, err := client.ModbusClient.ReadRegister(75, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return float32(waterOutTemperature / 10), nil
}
