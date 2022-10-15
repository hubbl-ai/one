package dm

type CompanyBranch struct {
	BranchName                       string    `json:"iata:companybranch:branchName,omitempty"`                       // Company branch name
	IATACargoAgentLocationIdentifier string    `json:"iata:companybranch:iataCargoAgentLocationIdentifier,omitempty"` // IATA CASS cargo agent 4 digit branch number / location identifier
	Company                          *Company  `json:"iata:companybranch:company,omitempty"`                          // Refers to the mother company of the branch
	ContactPersons                   []*Person `json:"iata:companybranch:contactPersons,omitempty"`                   // Contact person details
	Location                         *Location `json:"iata:companybranch:location,omitempty"`                         // Location and address details
	// Other identifiers (e.g. LEI (Legal Entity Identifier), TIN
	// (Trader Identification Number), PIMA address, Account
	// number, VAT/Tax id, Legal Registration id, DUNS number,
	// etc)
	OtherIdentifiers []*OtherIdentifier `json:"iata:companybranch:otherIdentifiers,omitempty"`
}

func (o CompanyBranch) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
