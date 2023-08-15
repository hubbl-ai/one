package dm

type IOTDevice struct {
	DeviceManufacturer *Company  `json:"iata:iotDevice:deviceManufacturer"` // Manufacturer of the device
	Sensors            []*Sensor `json:"iata:iotDevice:sensors"`            // Reference to the sensors linked to the device
	// Reference of the Logistic Object to which the Connected
	// Device is linked (URI)
	AssociatedObjectId string `json:"iata:iotDevice:associatedObject"`
	// Natural language description of the device. It can describe
	// how and where the device is attached.
	DeviceDescription  string `json:"iata:iotDevice:deviceDescription"`
	DeviceModel        string `json:"iata:iotDevice:deviceModel"`        // Commercial denomination of the device
	DeviceName         string `json:"iata:iotDevice:deviceName"`         // Name of the device defined by the device's owner
	DeviceSerialNumber string `json:"iata:iotDevice:deviceSerialNumber"` // Serial number that allows to uniquely identify the device
}

func (o IOTDevice) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
