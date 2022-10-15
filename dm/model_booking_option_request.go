package dm

// NB: Ontology says ExpectedCommodities, data model says ExpectedCommodity

type BookingOptionRequest struct {
	// BookingSegment *BookingSegment `json:"iata:bookingOptionRequest:bookingSegment,omitempty"` // The Booking Segment linked to the Booking Option Request
	Parties                []*Party      `json:"iata:bookingOptionRequest:parties,omitempty"`                // Parties involved if known
	RatingsPreference      *Ratings      `json:"iata:bookingOptionRequest:ratingsPreference,omitempty"`      // Ratings preferences of the request
	RoutingPreferences     *Routing      `json:"iata:bookingOptionRequest:routingPreferences,omitempty"`     // Routing details that are part of the request, these details will be used to determine if the offer is a perfect match
	ShipmentDetails        *Shipment     `json:"iata:bookingOptionRequest:shipmentDetails,omitempty"`        // Details of the shipement that is to be shipped
	TimePreferences        *BookingTimes `json:"iata:bookingOptionRequest:timePreferences,omitempty"`        // Schedule preferences of the request
	UnitsPreference        []*Value      `json:"iata:bookingOptionRequest:unitsPreference,omitempty"`        // Unit preferences of the request (e.g. kg or cm)
	Allotment              string        `json:"iata:bookingOptionRequest:allotment,omitempty"`              // Reference to the Allotment as per the contracts between forwarders and carriers
	ExpectedCommodities    string        `json:"iata:bookingOptionRequest:expectedCommodities,omitempty"`    // Expected commodity for quote request purposes only. To be defined by MCD
	RequestType            string        `json:"iata:bookingOptionRequest:requestType,omitempty"`            // Identification of the request type: Quote or Booking (to be confirmed)
	RequestedHandling      string        `json:"iata:bookingOptionRequest:requestedHandling,omitempty"`      // Requested handling information for quote request purposes only
	ShipmentSecurityStatus string        `json:"iata:bookingOptionRequest:shipmentSecurityStatus,omitempty"` // Indicate the security state of the shipment, screened or not
	*LogisticsObject
}

func (o BookingOptionRequest) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
