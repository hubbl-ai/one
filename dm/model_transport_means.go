package dm

type TransportMeans struct {
	TransportCompany  *Company           `json:"iata:transportMeans:transportCompany,omitempty"`  // Company operating the transport means
	TransportMovement *TransportMovement `json:"iata:transportMeans:transportMovement,omitempty"` // Transport Movement on which the Transport Means is used
	// TransportSegment *TransportSegment `json:"iata:transportMeans:transportSegment,omitempty"` // Associated transport segment
	TypicalCO2Coefficient  *Value `json:"iata:transportMeans:typicalCO2Coefficient,omitempty"`  // Required for some CO2 calculations
	TypicalFuelConsumption *Value `json:"iata:transportMeans:typicalFuelConsumption,omitempty"` // Typical fuel comsumption (e.g. 20 000L / 1 000nm)
	VehicleModel           string `json:"iata:transportMeans:vehicleModel,omitempty"`           // Model or make of the vehicle (e.g. A330-300)
	VehicleRegistration    string `json:"iata:transportMeans:vehicleRegistration,omitempty"`    // Vehicle identification - e.g. aircraft registration number
	VehicleSize            string `json:"iata:transportMeans:vehicleSize,omitempty"`            // Size of the vehicle - free text
	VehicleType            string `json:"iata:transportMeans:vehicleType,omitempty"`            // Vehicle or container type. Refer UNECE28, e.g. 4.00.0 - Aircraft, type unknown.For Air refer to IATA Standard Schedules Information Manua in section ATA/IATA Aircraft Types
}

func (o TransportMeans) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
