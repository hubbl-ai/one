package dm

type Carrier struct {
	*Company
	CarrierName      string `json:"iata:carrier:carrierName,omitempty"`      // Official carrier name
	CarrierShortName string `json:"iata:carrier:carrierShortName,omitempty"` // Carrier short name if any
	AirlineCode      string `json:"iata:carrier:airlineCode,omitempty"`      // IATA two-character airline code
	AirlinePrefix    string `json:"iata:carrier:airlinePrefix,omitempty"`    // IATA three-numeric airline prefix number
}

func (o Carrier) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
