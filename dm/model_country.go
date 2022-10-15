package dm

type Country struct {
	CountryCode  string `json:"iata:country:countryCode,omitempty"`
	CountryValue string `json:"iata:country:countryValue,omitempty"`
}

func (o Country) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
