package dm

import "time"

type Event struct {
	EventTypeIndicator EventType `json:"iata:event:eventTypeIndicator"`
	DateTime           time.Time `json:"iata:event:dateTime"` // Date and time of the event
	// Movement or milestone code. Can refer to CXML Code List
	// 1.18, e.g. DEP, ARR, FOH, RCS but not restricted to it.
	EventCode         string    `json:"iata:event:eventCode,omitempty"`
	EventName         string    `json:"iata:event:eventName,omitempty"`         // If no EventCode provided, event name - e.g. Security clearance
	Location          *Location `json:"iata:event:location,omitempty"`          // Location of event
	PerformedBy       *Company  `json:"iata:event:performedBy,omitempty"`       // Party performing the event
	PerformedByPerson *Person   `json:"iata:event:performedByPerson,omitempty"` // rdfs:comment "Document originator details and contacts
	LinkedObject      string    `json:"iata:event:linkedObject,omitempty"`      // Refers to the URI of the linked object
}

func (o Event) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
