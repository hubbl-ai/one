package dm

type Insurance struct {
	CoveringParty     *Company  `json:"iata:insurance:coveringParty,omitempty"`     // Party covering the insurance
	InsuranceAmount   *Value    `json:"iata:insurance:insuranceAmount,omitempty"`   // Insured amount - amount covered by the insurance policy
	InsuranceShipment *Shipment `json:"iata:insurance:insuranceShipment,omitempty"` // Reference to the shipment insured
	NVDIndicator      bool      `json:"iata:insurance:nvdIndicator,omitempty"`      // When no value is declared for Insurance this field should be completed with the value TRUE otherwise FALSE
}

func (o Insurance) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
