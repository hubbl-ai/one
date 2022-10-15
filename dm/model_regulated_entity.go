package dm

// NB: Branch is deprecated. Should Entity be to a CompanyBranch?

type RegulatedEntity struct {
	// Entity *Branch `json:"iata:regulatedEntity:entity,omitempty"` // Branch/Company
}

func (o RegulatedEntity) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
