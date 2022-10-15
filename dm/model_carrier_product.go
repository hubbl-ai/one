package dm

type CarrierProduct struct {
	// BookingOption *BookingOption `json:"iata:carrierProduct:bookingOption,omitempty"` // Reference to the BookingOption where the carrier product is used
	ProductCode        string `json:"iata:carrierProduct:productCode,omitempty"`        // Carrier's product code
	ProductDescription string `json:"iata:carrierProduct:productDescription,omitempty"` // Carrier's product description
	*LogisticsObject
}

func (o CarrierProduct) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
