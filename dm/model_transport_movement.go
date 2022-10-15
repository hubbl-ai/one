package dm

type TransportMovement struct {
	// CO2CalculationMethod    *CO2CalcMethod         `json:"iata:transportMovement:co2CalculationMethod,omitempty"`    // Method of calculation of the CO2 emissions
	ArrivalLocation         *Location              `json:"iata:transportMovement:arrivalLocation,omitempty"`         // Arrival location details
	Co2Emissions            []*CO2Emissions        `json:"iata:transportMovement:co2Emissions,omitempty"`            // Amount of CO2 emitted (e.g. 34 kg/km)
	DepartureLocation       *Location              `json:"iata:transportMovement:departureLocation,omitempty"`       // Departure location details
	DistanceCalculated      *Value                 `json:"iata:transportMovement:distanceCalculated,omitempty"`      // Distance calculated if distance measured is not available
	DistanceMeasured        *Value                 `json:"iata:transportMovement:distanceMeasured,omitempty"`        // Distance measured
	ExternalReferences      []*ExternalReference   `json:"iata:transportMovement:externalReferences,omitempty"`      // Reference to document or logistics object (URI)
	FuelAmountCalculated    *Value                 `json:"iata:transportMovement:fuelAmountCalculated,omitempty"`    // calculated fuel consumption, if measured not available
	FuelAmountMeasured      *Value                 `json:"iata:transportMovement:fuelAmountMeasured,omitempty"`      // actual measured fuel consumption
	MovementTimes           []*MovementTimes       `json:"iata:transportMovement:movementTimes,omitempty"`           // Reference to all Movement Times such as Departure, Arrival, etc.
	Payload                 *Value                 `json:"iata:transportMovement:payload,omitempty"`                 // Actual payload for the transport
	TransportMeans          *TransportMeans        `json:"iata:transportMovement:transportMeans,omitempty"`          // Transport means details
	TransportMeansOperators []*Person              `json:"iata:transportMovement:transportMeansOperators,omitempty"` // Name of the person operating the transport means (e.g. aircraft captain, truck driver)
	TransportedPieces       []*Piece               `json:"iata:transportMovement:transportedPieces,omitempty"`       // Pieces assigned to the transport segment
	TransportedULDs         []*ULD                 `json:"iata:transportMovement:transportedUlds,omitempty"`         // ULDs assigned to the transport segment
	FuelType                string                 `json:"iata:transportMovement:fuelType,omitempty"`                // e.g. Kerosene, Diesel, SAF, Electricity [renewable], Electricity [non-renewable]
	ModeCode                string                 `json:"iata:transportMovement:modeCode,omitempty"`                // Mode of transport code, refer to UNECE R
	ModeQualifier           TransportModeQualifier `json:"iata:transportMovement:modeQualifier,omitempty"`           // Pre-Carriage, Main-Carriage or On-Carriage
	Seal                    string                 `json:"iata:transportMovement:seal,omitempty"`                    // Seal identifier
	TransportIdentifier     string                 `json:"iata:transportMovement:transportIdentifier,omitempty"`     // Airline flight number, or rail /  truck / maritime line id
	UnplannedStop           string                 `json:"iata:transportMovement:unplannedStop,omitempty"`           // Free text field to be used for additional details regarding unplanned stops such as technical stops.
}

func (o TransportMovement) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
