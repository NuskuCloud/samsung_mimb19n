package samsung_mimb19n

func (client *Client) DhwEnable(value bool) error {
	return client.ModbusClient.WriteRegister(72, BoolToInt(value))
}

func (client *Client) SetFlowTemperature(DegreesC uint16) error {
	return client.ModbusClient.WriteRegister(74, DegreesC*10)
}

func (client *Client) SetInsideTargetTemperature(DegreesC uint16) error {
	return client.ModbusClient.WriteRegister(58, DegreesC*10)
}

func (client *Client) CentralHeatingEnable(value bool) error {
	return client.ModbusClient.WriteRegister(52, BoolToInt(value))
}
