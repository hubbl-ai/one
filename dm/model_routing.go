package dm

import "time"

type Routing struct {
	ScheduledLegs           []*ScheduledLegs `json:"iata:routing:scheduledLegs,omitempty"`           // Scheduled Legs class to be used to identify legs. Can be used with Booking Option Request as an indicator of preferred Routing or with Booking Option when a carrier proposes a specific Routing.
	AircraftPossibilityCode string           `json:"iata:routing:aircraftPossibilityCode,omitempty"` // Aircraft possibility code
	LatestArrivalDateTime   time.Time        `json:"iata:routing:latestArrivalDateTime,omitempty"`   // Latest Arrival date time (requested or proposed)
	MaxConnections          uint32           `json:"iata:routing:maxConnections,omitempty"`          // Maximum number of connections of the transport movement (requested or proposed)
	OnlineInd               bool             `json:"iata:routing:onlineInd,omitempty"`               // Indicates interlining (requested or proposed)
	RoadFeederServiceInd    bool             `json:"iata:routing:rfsInd,omitempty"`                  // Indicates if RFS (Road Feeder Services) is included (requested or proposed)
}

func (o Routing) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
