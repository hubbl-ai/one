package dm

// Dimension details
type Dimensions struct {
	Height *Value `json:"iata:dimension:height,omitempty"`
	Length *Value `json:"iata:dimension:length,omitempty"`
	Width  *Value `json:"iata:dimension:width,omitempty"`
	Volume *Value `json:"iata:dimension:volume,omitempty"`
}

func (o Dimensions) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
