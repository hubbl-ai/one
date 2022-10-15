package dm

type DGDeclaration struct {
	AircraftLimitationInformation string `json:"iata:dgDeclaration:aircraftLimitationInformation,omitempty"` // Contains the Special Handling Code related to the prescribed limitation. Hardcoded to PASSENGER AND CARGO AIRCRAFT or CARGO AIRCRAFT ONLY. This field is mandatory for air (Air)
	ComplianceDeclarationText     string `json:"iata:dgDeclaration:complianceDeclarationText,omitempty"`     // Contains the warning message complying with the regulations text note. This field is mandatory for air (Air)
	ExclusiveUseIndicator         bool   `json:"iata:dgDeclaration:exclusiveUseIndicator,omitempty"`         // Indicates an exclusive use shipment
	HandlingInformation           string `json:"iata:dgDeclaration:handlingInformation,omitempty"`           // Free text. This may include items such as Control temperature for substances stabilized by temperature control, name and telephone number of a responsible person for infectious substances.
	LogisticsObject
}

func (o DGDeclaration) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
