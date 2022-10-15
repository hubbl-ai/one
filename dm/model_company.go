package dm

type Company struct {
	CompanyName        string           `json:"iata:company:companyName,omitempty"`        // Name of company or organization
	IATACargoAgentCode string           `json:"iata:company:iataCargoAgentCode,omitempty"` // IATA accredited cargo agent 7 digit number
	Branch             []*CompanyBranch `json:"iata:company:branch,omitempty"`             // Company branches. Ontology says 1:n, model says 1:1
}

func (o Company) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
