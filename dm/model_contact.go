package dm

type Contact struct {
	ContactType  ContactType `json:"iata:contact:contactType,omitempty"`
	ContactValue string      `json:"iata:contact:contactValue,omitempty"`
}

func (o Contact) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
