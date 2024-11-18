package samsung_mimb19n

import "time"

// Heating

func (client *Client) SetFlowTemperature(DegreesC uint16) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(68, DegreesC*10)
}

func (client *Client) SetInsideTargetTemperature(DegreesC uint16) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(58, DegreesC*10)
}

func (client *Client) CentralHeatingEnable(value bool) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(52, BoolToInt(value))
}

//  Hot water

func (client *Client) DhwEnable(value bool) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(72, BoolToInt(value))
}

/*
0: Eco
1: Standard
2: Power
3: Force
*/
func (client *Client) SetHotWaterMode(value uint16) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(73, value)
}

func (client *Client) setHotWaterSetTemperature(DegreesC uint16) error {
	time.Sleep(client.sleepDuration)
	return client.ModbusClient.WriteRegister(74, DegreesC*10)
}
