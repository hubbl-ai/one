package dm

type CO2Emissions struct {
	CalculatedEmissions *Value             `json:"iata:co2Emissions:calculatedEmissions,omitempty"` // CO2 emissions calculated
	TransportMovement   *TransportMovement `json:"iata:co2Emissions:transportMovement,omitempty"`   // Transport Movement linked to the CO2 Emissions object
	*LogisticsObject
}

func (o CO2Emissions) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
