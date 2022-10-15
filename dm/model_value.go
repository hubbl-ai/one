package dm

type Value struct {
	Value string `json:"iata:value:value,omitempty"`
	Unit  string `json:"iata:value:unit,omitempty"`
}

func (o Value) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
