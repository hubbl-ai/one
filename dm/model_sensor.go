package dm

type Sensor struct {
	IOTDevice          *IOTDevice `json:"iata:iotDevice:sensor"`          // Reference to the IoT Device to which the sensor is linked
	SensorDescription  string     `json:"iata:sensor:sensorDescription"`  // Natural language description of the sensor
	SensorName         string     `json:"iata:sensor:sensorName"`         // Name of the sensor defined by the sensor's manufacturer
	SensorSerialNumber string     `json:"iata:sensor:sensorSerialNumber"` // Serial number that allows to uniquely identify the sensor
	SensorType         SensorType `json:"iata:sensor:sensorType"`         // Type of sensor as described in Interactive Cargo Recommended Practice
}

func (o Sensor) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
