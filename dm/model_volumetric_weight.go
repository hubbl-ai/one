package dm

type VolumetricWeight struct {
	ChargeableWeight *Value `json:"iata:volumetricWeight:chargeableWeight,omitempty"`
	ConversionFactor *Value `json:"iata:volumetricWeight:conversionFactor,omitempty"`
}

func (o VolumetricWeight) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
