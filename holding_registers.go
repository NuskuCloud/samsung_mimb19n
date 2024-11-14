package samsung_mimb19n

import "github.com/simonvetter/modbus"

func (client *Client) ReadOutdoorTemperature() (float32, error) {
	outdoorTemperature, err := client.ModbusClient.ReadFloat32(5, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return outdoorTemperature / 10, nil
}

func (client *Client) ReadReturnTemperature() (float32, error) {
	returnTemperature, err := client.ModbusClient.ReadFloat32(65, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return returnTemperature / 10, nil
}

func (client *Client) ReadFlowTemperature() (float32, error) {
	flowTemperature, err := client.ModbusClient.ReadFloat32(66, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return flowTemperature / 10, nil
}

func (client *Client) ReadIndoorTemperature() (float32, error) {
	indoorTemperature, err := client.ModbusClient.ReadFloat32(59, modbus.HOLDING_REGISTER)
	if err != nil {
		return -1000, err
	}

	return indoorTemperature / 10, nil
}
