package dm

type Ranges struct {
	Amount                float64 `json:"iata:ranges:amount,omitempty"`          // Amount
	MaximumQuantity       float64 `json:"iata:ranges:maximumQuantity,omitempty"` // Maximum quantity
	MinimumQuantity       float64 `json:"iata:ranges:minimumQuantity,omitempty"` // Minimum quantity
	RateClassCode         string  `json:"iata:ranges:rateClassCode,omitempty"`   // Rate class code e.g. Q. Refer to CXML Code List 1.4 Rate Class Codes
	RatingType            string  `json:"iata:ranges:ratingType,omitempty"`      // rating type - Refer to CXML Code List 1.44 ULD Charge Codes
	SpecificCommodityCode string  `json:"iata:ranges:scr,omitempty"`             // Specific commodity rates linked to commodity
	UnitBasis             string  `json:"iata:ranges:unitBasis,omitempty"`       // Specific commodity code linked to commodity
}

func (o Ranges) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
