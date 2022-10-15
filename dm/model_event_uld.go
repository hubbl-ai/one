package dm

type EventULD struct {
	*Event
	LoadingPosition string `json:"iata:eventUld:loadingPosition,omitempty"` // Position of the shipment in the aircraft - e.g. lower or main deck
}

func (o EventULD) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
