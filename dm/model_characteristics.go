package dm

// NB: Product is not in the data model

type Characteristics struct {
	Product             *Product `json:"iata:characteristics:product,omitempty"`             // Reference to the product
	CharacteristicsType string   `json:"iata:characteristics:characteristicsType,omitempty"` // Product characteristics code - e.g. CLR - Color. Not restricted to a list.
	Value               string   `json:"iata:characteristics:value,omitempty"`               // Product characteristics value / attribute - e.g. Blue...
	*LogisticsObject
}

func (o Characteristics) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
