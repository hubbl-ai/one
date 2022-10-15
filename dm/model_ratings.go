package dm

type Ratings struct {
	Ranges                  []*Ranges `json:"iata:ratings:ranges,omitempty"`                  // Reference to the ranges
	BillingChargeIdentifier *string   `json:"iata:ratings:billingChargeIdentifier,omitempty"` // Billing charge identifiers to be used for CASS. Refer to CargoXML Code List 1.33
	ChargeCode              string    `json:"iata:ratings:chargeCode,omitempty"`              // Charge code, refer to CargoXML Code List 1.1
	ChargeDescription       string    `json:"iata:ratings:chargeDescription,omitempty"`       // Description of the charge e.g. Airfreight, fuel, etc.
	ChargePaymentType       string    `json:"iata:ratings:chargePaymentType,omitempty"`       // Indicates if charge is prepaid or collect (P, C)
	ChargeType              string    `json:"iata:ratings:chargeType,omitempty"`              // Type of charge that should match the code expressed in either chargeCode, otherChargeCode or billingChargeIndentifier data properties.
	Entitlement             string    `json:"iata:ratings:entitlement,omitempty"`             // Entitlement code to define if charges are Due carrier (C) or Due agent (A). Refer to CXML Code List 1.3
	OtherChargeCode         string    `json:"iata:ratings:otherChargeCode,omitempty"`         // Refer to CargoXML Code List 1.2 for Other Charges
	PriceSpecification      string    `json:"iata:ratings:priceSpecification,omitempty"`      // Specification of the price e.g. Street, Group, Spot, etc.
	PriceSpecificationRef   *string   `json:"iata:ratings:priceSpecificationRef,omitempty"`   // Reference of price specifications
	Quantity                string    `json:"iata:ratings:quantity,omitempty"`                // Used if there is an applicable quantity to the rate (Usually a Time or a Number)
	RatingsType             string    `json:"iata:ratings:ratingsType,omitempty"`             // Used to identify if the Ratings are Face, Published or Actual ratings. Expected values are F, A, C.
	RateCombinationPoint    string    `json:"iata:ratings:rcp,omitempty"`                     // IATA 3-letter code of the rate combination point
	SubTotal                float64   `json:"iata:ratings:subTotal,omitempty"`                // Subtotal of the charge
}

func (o Ratings) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
