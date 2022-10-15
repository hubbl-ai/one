package dm

type LogisticsObject struct {
	Events            []*Event     `json:"iata:logisticsObject:events,omitempty"`            // Events object
	IOTDevices        []*IOTDevice `json:"iata:logisticsObject:iotDevices,omitempty"`        // Allows to link Logistic Objects with IoT Devices
	CompanyIdentifier string       `json:"iata:logisticsObject:companyIdentifier,omitempty"` // Company identifier from the Internet of Logistics of the entity that hosts the Logistics Object.
}
