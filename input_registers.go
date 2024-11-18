package samsung_mimb19n

import "time"

func (client *Client) DhwEnable(value bool) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(72, BoolToInt(value))
}

func (client *Client) SetFlowTemperature(DegreesC uint16) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(74, DegreesC*10)
}

func (client *Client) SetInsideTargetTemperature(DegreesC uint16) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(58, DegreesC*10)
}

func (client *Client) CentralHeatingEnable(value bool) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(52, BoolToInt(value))
}
