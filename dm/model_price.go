package dm

import "time"

type Price struct {
	// BookingOption *BookingOption `json:"iata:price:bookingOption,omitempty"` // Reference to the Booking or Offer
	Ratings           []*Ratings `json:"iata:price:ratings,omitempty"`           // Rating used for pricing
	CarrierChargeCode *string    `json:"iata:price:carrierChargeCode,omitempty"` // Charge code for carrier, eg. CA, CB, etc
	GrandTotal        float64    `json:"iata:price:grandTotal,omitempty"`        // Total price
	ValidTo           time.Time  `json:"iata:price:validTo,omitempty"`           // Terms of validity
}

func (o Price) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
