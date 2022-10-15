package dm

type OtherIdentifier struct {
	Identifier string `json:"iata:otherIdentifier:identifier,omitempty"`          // Item identifier
	Type       string `json:"iata:otherIdentifier:otherIdentifierType,omitempty"` // Identifier type or description
}

func (o OtherIdentifier) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
