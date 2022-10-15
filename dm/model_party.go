package dm

type Party struct {
	PartyRole    PartyRole `json:"iata:party:partyRole,omitempty"`
	PartyDetails *Company  `json:"iata:party:partyDetails,omitempty"`
	// Reference to other identifiers for parties. In the context
	// of the AWB, otherIdentifier object can be used with types
	// "PrimaryID" (internal ID), "Additional ID" (Standard
	// ID) or "AccountID" (Account numbers).
	//
	// Note that this is not present in the data model, but present
	// in the ontology.
	OtherIdentifier []*OtherIdentifier `json:"iata:party:otherIdentifier,omitempty"`
}

func (o Party) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
