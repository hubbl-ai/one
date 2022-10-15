package dm

// NB: LocationName is called Name in the data model

type Location struct {
	// Location type - e.g. Airport, Freight terminal, Rail
	// station, Seaport, etc
	LocationType string `json:"iata:location:locationType,omitempty"`
	LocationName string `json:"iata:location:locationName,omitempty"` // Full name of the location
	// Location code of airport, freight terminal, se
	// station. UN/LOCODE city code (5 letter) or IATA airport
	// code (3 letter)
	Code        string       `json:"iata:location:code,omitempty"`
	Address     *Address     `json:"iata:location:address,omitempty"`     // Address details
	Geolocation *Geolocation `json:"iata:location:geolocation,omitempty"` // Geolocation details
}

func (o Location) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
