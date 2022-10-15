package dm

type Address struct {
	Country         *Country `json:"iata:address:country,omitempty"`
	AddressCode     []string `json:"iata:address:addressCode,omitempty"`     // Address identifier using special coding systems e.g. US CBP FIRMS code
	AddressCodeType string   `json:"iata:address:addressCodeType,omitempty"` // Type of address code e.g. US CBP FIRMS
	CityCode        string   `json:"iata:address:cityCode,omitempty"`        // UN/LOCODE city code (5 letter) or IATA city code (3 letter)
	CityName        string   `json:"iata:address:cityName,omitempty"`        // If no CityCode provided, full name of the city
	POBox           string   `json:"iata:address:poBox,omitempty"`           // Post Office box number / code
	PostalCode      string   `json:"iata:address:postalCode,omitempty"`      // Postal / ZIP code
	RegionCode      string   `json:"iata:address:regionCode,omitempty"`      // Region/ State / Department. Refer ISO 3166-2
	RegionName      string   `json:"iata:address:regionName,omitempty"`      // If no StateCode provided, full name of the region / state / province / canton
	Street          []string `json:"iata:address:street,omitempty"`          // Street address including street name, street number, building number, apartment etc
}

func (o Address) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
