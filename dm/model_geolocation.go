package dm

type Geolocation struct {
	Elevation       *Value  `json:"iata:Geolocation:elevation,omitempty"`
	GeolocationUnit string  `json:"iata:Geolocation:geolocationUnit,omitempty"`
	Latitude        float64 `json:"iata:Geolocation:latitude,omitempty"`
	Longitude       float64 `json:"iata:Geolocation:longitude,omitempty"`
}

func (o Geolocation) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
