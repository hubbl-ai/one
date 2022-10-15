package dm

import "time"

type BookingTimes struct {
	BookingOption          *BookingOption        `json:"iata:bookingtimes:bookingOption,omitempty"`          // Reference to the BookingOption where the booking times are used
	BookingOptionRequest   *BookingOptionRequest `json:"iata:bookingtimes:bookingOptionRequest,omitempty"`   // Reference to the BookingOptionRequest where the booking times are used
	EarliestAcceptanceTime time.Time             `json:"iata:bookingtimes:earliestAcceptanceTime,omitempty"` // Earliest acceptance date time (requested or proposed)
	LatestAcceptanceTime   time.Time             `json:"iata:bookingtimes:latestAcceptanceTime,omitempty"`   // Latest Acceptance time as per CargoIQ definition (requested, proposed or actual)
	TimeOfAvailability     time.Time             `json:"iata:bookingtimes:timeOfAvailability,omitempty"`     // Time of availability of the shipment as per CargoIQ definition
	TotalTransitTime       time.Duration         `json:"iata:bookingtimes:totalTransitTime,omitempty"`       // Total transit time as per CargoIQ definition, expressed as a duration
	*LogisticsObject
}

func (o BookingTimes) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
