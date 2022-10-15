package dm

import "time"

type ScheduledLegs struct {
	ArrivalLocation   *Location `json:"iata:scheduledLegs:arrivalLocation,omitempty"`   // Arrival location of the leg
	DepartureLocation *Location `json:"iata:scheduledLegs:departureLocation,omitempty"` // Departure Location of the leg
	ArrivalDate       time.Time `json:"iata:scheduledLegs:arrivalDate,omitempty"`       // Arrival date and time of the leg
	DepartureDate     time.Time `json:"iata:scheduledLegs:departureDate,omitempty"`     // Departure date and time of the leg
	SequenceNumber    uint32    `json:"iata:scheduledLegs:sequenceNumber,omitempty"`    // Sequence number of the leg
	TransportID       string    `json:"iata:scheduledLegs:transportId,omitempty"`       // Transport Id of the leg. E.g. Flight number, truck route identifier, etc.
}

func (o ScheduledLegs) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
